package run

import (
	"os"
	"os/exec"
	"syscall"
)

type service struct {
	cmd     *exec.Cmd
	running bool
}

func (s *service) Start() error {
	if s.running {
		return nil
	}

	s.cmd = exec.Command("go", "run", ".")
	s.cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	s.cmd.Stdout = os.Stdout
	s.cmd.Stderr = os.Stderr
	s.running = true

	return s.cmd.Start()
}

func (s *service) Stop() error {
	if !s.running {
		return nil
	}

	s.running = false
	return syscall.Kill(-s.cmd.Process.Pid, syscall.SIGTERM)
}

func (s *service) Wait() error {
	return s.cmd.Wait()
}

func newService() *service {
	service := new(service)
	return service
}
