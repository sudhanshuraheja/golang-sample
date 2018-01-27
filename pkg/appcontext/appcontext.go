package appcontext

import (
	"github.com/sudhanshuraheja/golang-sample/pkg/config"
)

type AppContext struct {
	config *config.Config
}

func New(config *config.Config) *AppContext {
	return &AppContext{
		config: config,
	}
}

func (a *AppContext) GetConfig() *config.Config {
	return a.config
}
