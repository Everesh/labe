package main

import (
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"charm.land/log/v2"
	"codeberg.org/everesh/labe/internal/wm"
	"hazelnut.eclair.cafe/wlcl"
)

func main() {
	log.SetLevel(getLogLevel())

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
	log.Info("graceful shutdown")
}

func run(ctx context.Context) error {
	conn, err := wlcl.Connect(ctx, "")
	if err != nil {
		return fmt.Errorf("failed to connect to a wayland session: %w", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Error("failed to close connection to a wayland session", "err", err)
		}
	}()

	wm, err := wm.New(ctx, conn)
	if err != nil {
		return fmt.Errorf("failed to initialize the window manager: %w", err)
	}

	for !wm.Done {
		if err := conn.Dispatch(ctx); err != nil {
			if errors.Is(err, context.Canceled) {
				return nil
			}
			return fmt.Errorf("dispatch failed: %w", err)
		}
	}

	if wm.Err != nil {
		return fmt.Errorf("window manager exited due to an error: %w", err)
	}

	return nil
}

func getLogLevel() log.Level {
	if level, ok := os.LookupEnv("LABE_LOGLEVEL"); ok {
		switch strings.ToUpper(level) {
		case "DEBUG":
			return log.DebugLevel
		case "INFO":
			return log.InfoLevel
		case "WARN":
			return log.WarnLevel
		case "ERROR":
			return log.ErrorLevel
		case "FATAL":
			return log.FatalLevel
		case "NONE":
			return log.Level(math.MaxInt)
		default:
			log.Warn("unknown log level, defaulting to Info")
			return log.InfoLevel
		}
	}
	return log.InfoLevel
}
