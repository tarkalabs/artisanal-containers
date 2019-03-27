// +build linux
package container

import (
	"os"
	"os/exec"
	"syscall"

	log "github.com/sirupsen/logrus"
)

type ArtisanalContainer struct {
	Command string
	Args    []string
	Uid     int
	Gid     int
}

func (ac *ArtisanalContainer) ExecuteCommand() error {
	path, err := exec.LookPath(ac.Command)
	if err != nil {
		log.Errorf("Could not find binary %s in path", ac.Command)
	}
	err = syscall.Sethostname([]byte("artisanal-container"))
	if err != nil {
		log.Errorf("unable to set host name %s", err.Error())
		return err
	}
	log.Info("Running ...", path, ac.Args)
	args := append([]string{path}, ac.Args...)
	return syscall.Exec(path, args, os.Environ())
}

func (ac *ArtisanalContainer) ForkYourself() error {
	args := append([]string{"/proc/self/exe", "exec", "--fork=true", ac.Command}, ac.Args...)
	cmd := exec.Cmd{
		Path: "/proc/self/exe",
		Args: args,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUSER | syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS,
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
func (ac *ArtisanalContainer) Start(isFork bool) error {
	if isFork {
		return ac.ExecuteCommand()
	}
	return ac.ForkYourself()
}
