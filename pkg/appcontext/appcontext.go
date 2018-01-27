package appcontext

import (
	"github.com/sudhanshuraheja/golang-sample/pkg/config"
)

type AppContext struct{
	config *config.Config
}

func NewAppContext(config *config.Config) {
	return &AppContext{
		config: config
	}
}

func (a *AppContext) GetConfig() *config.Config {
	return a.config
}