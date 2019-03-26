// +build linux
package main

import (
	"os"
	"os/exec"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

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

func main() {
	ac := &ArtisanalContainer{Command: "sh", Uid: os.Getuid(), Gid: os.Getgid()}
	err := ac.Start()
	if err != nil {
		log.Fatal(err)
	}
}
