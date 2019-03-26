// +build linux
package container

import (
	"os"
	"os/exec"
	"syscall"
)

type ArtisanalContainer struct {
	Command string
	Args    []string
	Uid     int
	Gid     int
}

func (ac *ArtisanalContainer) Start() error {
	cmd := exec.Command(ac.Command, ac.Args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUSER | syscall.CLONE_NEWPID,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      ac.Uid,
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      ac.Gid,
				Size:        1,
			},
		},
	}
	return cmd.Run()
}
