// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package cmd

import (
	"io/ioutil"

	"github.com/gitpod-io/gitpod/common-go/log"
	"github.com/spf13/cobra"
	"golang.org/x/sys/unix"
)

var handlerMountOpts struct {
	Target string
	Source string
	FSType string
	Flags  uint
	Data   string
}

// handlerCmd represents the base command for all syscall handler
var handlerMountCmd = &cobra.Command{
	Use:   "mount",
	Short: "In-namespace mount handler",
	Run: func(cmd *cobra.Command, args []string) {
		if handlerMountOpts.FSType == "proc" || handlerMountOpts.FSType == "sys" {
			log.Fatal("cannot mount proc or sys")
		}

		// TODO(cw): is mounting onto symlinks a problem?
		// TODO(cw): should we create the mount target?

		ls, _ := ioutil.ReadDir(handlerMountOpts.Source)
		log.WithFields(map[string]interface{}{
			"source": handlerMountOpts.Source,
			"target": handlerMountOpts.Target,
			"flags":  handlerMountOpts.Flags,
			"fstype": handlerMountOpts.FSType,
			"data":   handlerMountOpts.Data,
			"len ls": len(ls),
		}).Info("mounting")

		// DEBUGGING ONLY
		handlerMountOpts.Flags = unix.MS_MOVE

		err := unix.Mount(handlerMountOpts.Source, handlerMountOpts.Target, handlerMountOpts.FSType, uintptr(handlerMountOpts.Flags), handlerMountOpts.Data)
		if err != nil {
			log.WithError(err).Fatal("cannot mount")
		}
	},
}

func init() {
	handlerMountCmd.Flags().StringVar(&handlerMountOpts.Target, "target", "", "target mount point")
	handlerMountCmd.Flags().StringVar(&handlerMountOpts.Source, "source", "", "mount source")
	handlerMountCmd.Flags().StringVar(&handlerMountOpts.FSType, "fs", "", "filesystem type")
	handlerMountCmd.Flags().UintVar(&handlerMountOpts.Flags, "flags", 0, "mount flags")
	handlerMountCmd.Flags().StringVar(&handlerMountOpts.Data, "data", "", "mount data")

	handlerCmd.AddCommand(handlerMountCmd)
}
