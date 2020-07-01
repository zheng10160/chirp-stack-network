package api

import (
	"github.com/pkg/errors"

	"github.com/jon177/lky-network-server/internal/config"
	"github.com/jon177/lky-network-server/internal/api/external"

)

func Setup(conf config.Config) error {
	if err := external.Setup(conf); err != nil {
		return errors.Wrap(err, "setup external api error")
	}

	return nil
}
