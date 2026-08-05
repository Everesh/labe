package config

import (
	"os/exec"

	"charm.land/log/v2"
)

func Spawn(name string, args ...string) {
	log.Info("launching", "cmd", name, "args", args)
	cmd := exec.Command(name, args...)
	if err := cmd.Start(); err != nil {
		log.Error("failed to launch command", "cmd", name, "err", err)
		return
	}
	go func() { _, _ = cmd.Process.Wait() }()
}
