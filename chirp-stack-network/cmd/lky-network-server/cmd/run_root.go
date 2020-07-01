package cmd

import (
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"context"
	"github.com/jon177/lky-network-server/internal/api"
	"github.com/jon177/lky-network-server/internal/config"
	"github.com/pkg/errors"
	"github.com/jon177/lky-network-server/internal/storage"
)

func run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	task := []func() error{
		setupStorage,
		setupApi,
	}

	for _, t := range task {
		if err := t(); err != nil {
			log.Fatal(err)
		}
	}

	sigChan := make(chan os.Signal)
	exitChan := make(chan struct{})
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	log.WithField("signal", <-sigChan).Info("signal received")

	go func() {
		log.Warning("stopping lky-network-server")
		exitChan <- struct{}{}
	}()

	select {
	case <-exitChan:
	case s := <-sigChan:
		log.WithField("signal", s).Info("signal received, stopping immediately")
	}
	return nil
}

func setupStorage() error {
	if err:= storage.Setup(config.C); err != nil {
		return  errors.Wrap(err, "setup storage error")
	}

	return nil
}

func setupApi() error {
	if err := api.Setup(config.C); err != nil {
		return errors.Wrap(err, "setup api error")
	}

	return nil
}
