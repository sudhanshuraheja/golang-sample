package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudhanshuraheja/golang-sample/pkg/config"
)

func TestConfigValues(t *testing.T) {
	conf := config.NewConfig()
	assert.Equal(t, "debug", conf.LogLevel())
	assert.Equal(t, "dbname=sample user=sample password='sample' host=postgres port=5432 sslmode=disable", conf.Database().ConnectionString())
	assert.Equal(t, "postgres://sample:sample@postgres:5432/sample?sslmode=disable", conf.Database().ConnectionURL())
}
