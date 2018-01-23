package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudhanshuraheja/golang-sample/config"
)

func TestConfigValues(t *testing.T) {
	config.Init()
	assert.Equal(t, "sample", config.Name())
	assert.Equal(t, "0.0.1", config.Version())
	assert.Equal(t, "debug", config.LogLevel())
	assert.Equal(t, "3000", config.Port())
	assert.Equal(t, false, config.EnableStaticFileServer())
	assert.Equal(t, true, config.EnableGzipCompression())
	assert.Equal(t, false, config.EnableDelayMiddleware())

	assert.Equal(t, "dbname=sample user=sample password='sample' host=postgres port=5432 sslmode=disable", config.Database().ConnectionString())
	assert.Equal(t, "postgres://sample:sample@postgres:5432/sample?sslmode=disable", config.Database().ConnectionURL())
	assert.Equal(t, 10, config.Database().MaxPoolSize())
}
