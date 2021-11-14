package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Service struct {
	Info         ServiceInfo
	Cmd          *exec.Cmd
	Ctrl         ControlSystem
	Status       int
	isSupervised bool
}
type ServiceInfo struct {
	Name              string
	Desc              string
	Exec              string
	Type              string
	Target            string
	AutoRestart       bool
	Required_services []string
	Required_targets  []string
}

func (s *Service) Run() {
	args := strings.Fields(s.Info.Exec)
	if len(args) == 1 {
		s.Cmd = exec.Command(args[0])
	} else {
		s.Cmd = exec.Command(args[0], args[1:len(args)]...)
	}
	logfile, err := os.Create(LOG_PATH + "/" + s.Info.Name + ".log")
	if err != nil {
		fmt.Println(err)
	}
	defer logfile.Close()
	s.Cmd.Stdout = logfile
	s.Cmd.Stderr = logfile
	go s.Supervise()
	s.Status = ServiceRunning
	err = s.Cmd.Run()
	s.Ctrl.ExitStatus <- err
}

func (s *Service) Stop() {
	s.Cmd.Process.Kill()
}

func (s *Service) HandleCmd(cmd int) {
	switch cmd {
	case RestartCmd:
		s.Stop()
		go s.Run()
	case StopCmd:
	case StartCmd:
	case EnableCmd:
	case DisableCmd:
	}
}
