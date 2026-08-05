package config

import (
	"charm.land/log/v2"
)

func StartupHooks() {
	log.Info("running startup hooks")
	Spawn("swaybg", "-m", "fill", "-i", "./assets/bg.jpg")
	Spawn("ashell")
	Spawn("wlr-randr", "--output", "DP-4", "--pos", "0,600",
		"--output", "DP-5", "--pos", "1920,0", "--output", "HDMI-A-2", "--pos", "1920,600")
}
