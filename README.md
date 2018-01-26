# Golang Sample App
This is a golang sample app that I use for my own projects

## Features
- Supports multiple binaries
- Uses travis-ci as the primary CI
- Uses Codecov for aggregating test coverage
- Uses dep for dependency management
- Has a Dockerfile to ship the binary in a container
- Has a docker-compose file in external which gives you a postgres, redis and rabbitmq
- Uses viper for config

## Installation instructions
```
$ git clone git@github.com:sudhanshuraheja/golang-sample.git
$ make build_fresh
```