package main

import (
	"fmt"
)

// Sercive control command
const (
	RestartCmd = iota
	StopCmd
	StartCmd
	EnableCmd
	DisableCmd
)

// Sercive status
const (
	ServiceRunning = iota
	ServiceStopped
	ServiceDone
	ServiceFailed
)

type ControlSystem struct {
	Cmd        chan int
	ExitStatus chan error
	Stop       chan bool
}

func (c *ControlSystem) Init() {
	c.Cmd = make(chan int)
	c.ExitStatus = make(chan error)
	c.Stop = make(chan bool)
}

func (s *Service) Supervise() {
	if s.isSupervised {
		return
	}
	s.isSupervised = true
	fmt.Println("Supervising:", s.Info.Name)
	s.Ctrl.Init()
	for {
		select {
		case status := <-s.Ctrl.ExitStatus:
			if status != nil {
				s.Status = ServiceFailed
				fmt.Println(s.Info.Name+":", status)
				if s.Info.AutoRestart {
					go s.Run()
				}
			} else {
				s.Status = ServiceStopped
				fmt.Println(s.Info.Name + ": Exit success")
			}
		case cmd := <-s.Ctrl.Cmd:
			s.HandleCmd(cmd)
		case stop := <-s.Ctrl.Stop:
			if stop {
				s.Stop()
			}
		}
	}
	s.isSupervised = false
}
