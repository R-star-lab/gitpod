// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package seccomp

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/gitpod-io/gitpod/common-go/log"
	"github.com/gitpod-io/gitpod/workspacekit/pkg/nsenter"
	"github.com/gitpod-io/gitpod/workspacekit/pkg/readarg"
	daemonapi "github.com/gitpod-io/gitpod/ws-daemon/api"
	libseccomp "github.com/seccomp/libseccomp-golang"
	"golang.org/x/sys/unix"
)

func handleMount(req *libseccomp.ScmpNotifReq, stagingDir string, daemon daemonapi.InWorkspaceServiceClient) (val uint64, errno int32, cont bool) {
	log := log.WithField("syscall", "mount")

	memFile, err := readarg.OpenMem(req.Pid)
	if err != nil {
		return returnErrno(unix.EPERM)
	}
	defer memFile.Close()

	source, err := readarg.ReadString(memFile, int64(req.Data.Args[0]))
	if err != nil {
		log.WithField("pid", req.Pid).WithField("arg", 0).WithError(err).Error("cannot read argument")
		return returnErrno(unix.EFAULT)
	}
	dest, err := readarg.ReadString(memFile, int64(req.Data.Args[1]))
	if err != nil {
		log.WithField("pid", req.Pid).WithField("arg", 1).WithError(err).Error("cannot read argument")
		return returnErrno(unix.EFAULT)
	}
	filesystem, err := readarg.ReadString(memFile, int64(req.Data.Args[2]))
	if err != nil {
		log.WithField("pid", req.Pid).WithField("arg", 2).WithError(err).Error("cannot read argument")
		return returnErrno(unix.EFAULT)
	}
	// flags, err := readarg.ReadUintptr(memFile, int64(req.Data.Args[3]))
	// if err != nil {
	// 	log.WithField("pid", req.Pid).WithField("arg", 3).WithError(err).Error("cannot read argument")
	// 	return returnErrno(unix.EFAULT)
	// }
	// data, err := readarg.ReadString(memFile, int64(req.Data.Args[4]))
	// if err != nil {
	// 	log.WithField("pid", req.Pid).WithField("arg", 4).WithError(err).Error("cannot read argument")
	// 	return returnErrno(unix.EFAULT)
	// }

	log.WithFields(map[string]interface{}{
		"source": source,
		"dest":   dest,
		// "data":   data,
	}).Info("handling mount syscall")
	if filesystem == "proc" {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		resp, err := daemon.MountProc(ctx, &daemonapi.MountProcRequest{
			Pid: int64(req.Pid),
		})
		if err != nil {
			log.WithField("pid", req.Pid).WithField("dest", dest).WithError(err).Error("cannot mount proc")
			return returnErrno(unix.EFAULT)
		}

		err = MoveMountIntoRing2(int(req.Pid), resp.Location, stagingDir, dest)
		if err != nil {
			log.WithField("pid", req.Pid).WithField("dest", dest).WithField("loc", resp.Location).WithError(err).Error("cannot move proc")
			return returnErrno(unix.EFAULT)
		}
	} else {

	}

	return returnSuccess()
}

// MoveMountIntoMountNS moves a mount from source, via staging to dest.
// dest is a path as seen from ring2. staging is expected to be visible inside ring2 as `/.staging`
func MoveMountIntoRing2(pid int, source, staging, dest string) error {
	var (
		err error

		id     = fmt.Sprint(rand.Uint32())
		staged = filepath.Join(staging, id)
	)
	err = os.MkdirAll(staged, 0755)
	if err != nil {
		return err
	}
	err = unix.Mount(source, staged, "", unix.MS_MOVE, "")
	if err != nil {
		return err
	}

	return nsenter.Mount(pid, filepath.Join("/.staging", id), dest, unix.MS_MOVE, "")
}
