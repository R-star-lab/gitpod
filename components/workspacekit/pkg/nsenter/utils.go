// +build linux

package nsenter

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"golang.org/x/sys/unix"
)

type Namespace int

const (
	// NamespaceMount refers to the mount namespace
	NamespaceMount = iota
	// NamespaceNet refers to the network namespace
	NamespaceNet
	// NamespaceNet refers to the network namespace
	NamespacePID
)

// Run executes a workspacekit handler in a namespace.
// Use preflight to check libseccomp.NotifIDValid().
func Run(pid int, args []string, preflight func() error, enterNamespace ...Namespace) error {
	nss := []struct {
		Env    string
		Source string
		Flags  int
		NS     Namespace
	}{
		{"_LIBNSENTER_ROOTFD", fmt.Sprintf("/proc/%d/root", pid), unix.O_PATH, -1},
		{"_LIBNSENTER_CWDFD", fmt.Sprintf("/proc/%d/cwd", pid), unix.O_PATH, -1},
		{"_LIBNSENTER_MNTNSFD", fmt.Sprintf("/proc/%d/ns/mnt", pid), os.O_RDONLY, NamespaceMount},
		{"_LIBNSENTER_MNTNSFD", fmt.Sprintf("/proc/%d/ns/net", pid), os.O_RDONLY, NamespaceNet},
		{"_LIBNSENTER_MNTNSFD", fmt.Sprintf("/proc/%d/ns/pid", pid), os.O_RDONLY, NamespacePID},
	}

	stdioFdCount := 3
	cmd := exec.Command("/proc/self/exe", append([]string{"handler"}, args...)...)
	cmd.Env = append(cmd.Env, "_LIBNSENTER_INIT=1")
	for _, ns := range nss {
		var enter bool
		if ns.NS == -1 {
			enter = true
		} else {
			for _, s := range enterNamespace {
				if ns.NS == s {
					enter = true
					break
				}
			}
		}
		if !enter {
			continue
		}

		f, err := os.OpenFile(ns.Source, ns.Flags, 0)
		if err != nil {
			return fmt.Errorf("cannot open %s: %w", ns.Source, err)
		}
		defer f.Close()
		cmd.Env = append(cmd.Env, fmt.Sprint(stdioFdCount+len(cmd.ExtraFiles)))
		cmd.ExtraFiles = append(cmd.ExtraFiles, f)
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("cannot start handler: %w: %s", err, string(out))
	}
	return nil
}

// Mount executes mount in the mount namespace of PID
func Mount(pid int, source, target string, flags int, data string) error {
	args := []string{"mount",
		"--source", source,
		"--target", target,
		"--flags", strconv.Itoa(flags),
	}
	if data != "" {
		args = append(args, "--data", data)
	}
	return Run(pid, args, nil, NamespaceMount)
}
