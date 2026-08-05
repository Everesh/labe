package config

import (
	"charm.land/log/v2"
)

func StartupHooks() {
	log.Info("running startup hooks")
	Spawn("swaybg", "-m", "fill", "-i", "./assets/bg.jpg")
	Spawn("ashell")
}
