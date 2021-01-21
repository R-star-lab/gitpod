// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package cmd

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/bmatcuk/doublestar/v2"
	"github.com/gitpod-io/gitpod/gitpod-cli/pkg/theialib"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Save and restore files",
}

var cacheSaveCmd = &cobra.Command{
	Use:   "save <name> <globs...>",
	Short: "Saves files and folders into the cache",
	Long: `All files matching the given globs are added to a tar file and uploaded to the cache storage.
	
See https://github.com/bmatcuk/doublestar#patterns for valid glob patterns.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fail := func(msg string) {
			fmt.Fprintf(os.Stderr, "Error: %s\n", msg)
			os.Exit(-1)
		}

		name := args[0]
		if err := validateName(name); err != nil {
			fail(err.Error())
		}
		globs := args[1:]
		if verbose {
			fmt.Printf("Saving cache with name '%s' using %d glob pattern(s).\n", name, len(globs))
		}

		piper, pipew := io.Pipe()
		go func() {
			defer pipew.Close()
			if err := createTarball(pipew, globs); err != nil {
				fail(err.Error())
			}
		}()

		service, err := theialib.NewServiceFromEnv()
		if err != nil {
			fail(err.Error())
		}

		req := theialib.GetContentBlobUploadUrlRequest{Name: name, MediaType: "application/x-gitpod-cache"}
		resp, err := service.GetContentBlobUploadUrl(req)
		if err != nil {
			fail(err.Error())
		}
		if verbose {
			fmt.Printf("Upload URL: '%s'\n", resp.Url)
		}

		client := &http.Client{}
		httpreq, err := http.NewRequest(http.MethodPut, resp.Url, piper)
		if err != nil {
			fail(err.Error())
		}
		httpresp, err := client.Do(httpreq)
		if err != nil {
			fail(err.Error())
		}
		body, err := ioutil.ReadAll(httpresp.Body)
		if err != nil {
			fail(err.Error())
		}
		if verbose {
			fmt.Printf("Upload response: '%s'\n", body)
		}
		fmt.Printf("Successfully saved '%s'.\n", name)
	},
}

var cacheRestoreCmd = &cobra.Command{
	Use:   "restore <name> [target-dir]",
	Short: "Restores stored file from the cache",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		fail := func(msg string) {
			fmt.Fprintf(os.Stderr, "Error: %s\n", msg)
			os.Exit(-1)
		}

		name := args[0]
		if err := validateName(name); err != nil {
			fail(err.Error())
		}
		var targetDir = "/"
		if len(args) > 1 {
			targetDir = args[1]
		}
		if verbose {
			fmt.Printf("Restoring cache with name '%s' into directory '%s'.\n", name, targetDir)
		}

		service, err := theialib.NewServiceFromEnv()
		if err != nil {
			fail(err.Error())
		}

		req := theialib.GetContentBlobDownloadUrlRequest{Name: name}
		resp, err := service.GetContentBlobDownloadUrl(req)
		if err != nil {
			fail(err.Error())
		}

		httpresp, err := http.Get(resp.Url)
		if err != nil {
			fail(err.Error())
		}

		gzipReader, err := gzip.NewReader(httpresp.Body)
		if err != nil {
			fail(err.Error())
		}
		defer gzipReader.Close()

		tarReader := tar.NewReader(gzipReader)
		for {
			header, err := tarReader.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				fail(err.Error())
			}
			filePath := filepath.Join(targetDir, header.Name)
			perm := os.FileMode(header.Mode)
			if err := writeFile(tarReader, filePath, perm); err != nil {
				fail(err.Error())
			}
		}
		fmt.Printf("Successfully restored '%s'.\n", name)
	},
}

var verbose bool
var restoreSkipExisting bool
var restoreIgnoreWriteError bool

func init() {
	cacheCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	cacheCmd.AddCommand(cacheSaveCmd)
	cacheRestoreCmd.Flags().BoolVar(&restoreSkipExisting, "skip-existing", false, "do not override existing files, skip them instead")
	cacheRestoreCmd.Flags().BoolVar(&restoreIgnoreWriteError, "ignore-write-error", false, "do not abort if a file cannot be restored (e.g. due to missing write permissions), skip them instead")
	cacheCmd.AddCommand(cacheRestoreCmd)
	rootCmd.AddCommand(cacheCmd)
}

func validateName(name string) error {
	pattern := `[a-zA-Z0-9._\-]+`
	b, err := regexp.MatchString(pattern, name)
	if err != nil {
		return err
	}
	if !b {
		return fmt.Errorf("name needs to match pattern '%s'", pattern)
	}
	return nil
}

func createTarball(writer io.Writer, globs []string) error {
	gzipWriter := gzip.NewWriter(writer)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	for _, glob := range globs {
		if verbose {
			fmt.Printf("Processing glob pattern '%s'...\n", glob)
		}
		filePaths, err := doublestar.Glob(glob)
		if err != nil {
			return err
		}
		for _, filePath := range filePaths {
			filePath, err := filepath.Abs(filePath)
			if err != nil {
				return err
			}
			stat, err := os.Stat(filePath)
			if err != nil {
				return err
			}
			if stat.Mode().IsRegular() { // ignore dirs
				if err := addFileToTarball(tarWriter, filePath); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func addFileToTarball(tarWriter *tar.Writer, filePath string) error {
	if verbose {
		fmt.Printf("Adding file '%s'...", filePath)
		defer fmt.Println()
	}
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}
	header := &tar.Header{
		Name:    filePath,
		Size:    stat.Size(),
		Mode:    int64(stat.Mode()),
		ModTime: stat.ModTime(),
	}
	if err = tarWriter.WriteHeader(header); err != nil {
		return err
	}
	if _, err = io.Copy(tarWriter, file); err != nil {
		return err
	}
	if verbose {
		fmt.Print(" done.")
	}
	return nil
}

func writeFile(reader io.Reader, filePath string, perm os.FileMode) error {
	if verbose {
		fmt.Printf("Restoring file '%s'...", filePath)
		defer fmt.Println()
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		if restoreIgnoreWriteError {
			if verbose {
				fmt.Printf(" failed. (%s)", err.Error())
			}
			return nil
		}
		return err
	}

	flag := os.O_WRONLY | os.O_CREATE | os.O_TRUNC // create new or truncate existing file
	if restoreSkipExisting {
		flag = os.O_WRONLY | os.O_CREATE | os.O_EXCL // create new file or fail if file exists
	}
	file, err := os.OpenFile(filePath, flag, perm)
	if os.IsExist(err) {
		if verbose {
			fmt.Print(" skipped.")
		}
		return nil
	} else if err != nil {
		if restoreIgnoreWriteError {
			if verbose {
				fmt.Printf(" failed. (%s)", err.Error())
			}
			return nil
		}
		return err
	}
	defer file.Close()
	if _, err = io.Copy(file, reader); err != nil {
		if restoreIgnoreWriteError {
			if verbose {
				fmt.Printf(" failed. (%s)", err.Error())
			}
			return nil
		}
		return err
	}
	if verbose {
		fmt.Print(" done.")
	}
	return nil
}
