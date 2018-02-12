# Sample

Golang-Sample is a sample app that I use to initialise go-lang projects.

## Getting Started

Golang-Samples lets you do the following

* Supports multiple binaries
* Uses travis-ci as the primary CI
* Uses Codecov for aggregating test coverage
* Uses dep for dependency management
* Has a Dockerfile to ship the binary in a container
* Has a docker-compose file in external which gives you a postgres, redis and rabbitmq
* Supports 12factorapps

### Prerequisites

Here are the things that you would require before you get started

1. [Install git](https://www.atlassian.com/git/tutorials/install-git)
1. [Install golang](https://golang.org/doc/install)
1. [Install docker](https://docs.docker.com/install/#supported-platforms), we use it both for deployment and development

### Installing

Create a new_project and pull files into it

```bash
mkdir new_project
cd new_project
curl -L https://github.com/sudhanshuraheja/golang-sample/archive/master.tar.gz | tar -xzv --strip 1
make build_fresh
```

Setup postgres

```bash
$ cd external
$ docker-compose up
$ tanker migrate
repeat above till you see *Sadly, found no new migrations to run*
```

## Running the tests

If you would like to run the automated tests for the complete package, run this

```bash
make coverage
open ./coverage.html
```

### And coding style tests

We use the default golang coding conventions. Run the following to test for those

```bash
make fmt
make vet
make lint
```

## Deployment

There is no functionality in this repo and hence cannot be deployed

## Built With

* [DEP](https://github.com/golang/dep) - For dependency management
* [CLI](github.com/urfave/cli) - For accessing the binary on CLI
* [VIPER](github.com/spf13/viper) - For configuration management
* [LOGRUS](github.com/sirupsen/logrus) - For logging
* [NEGRONI](github.com/urfave/negroni) - HTTP Middleware
* [MUX](github.com/gorilla/mux) - For routing each request to the correct place
* [PQ](github.com/lib/pq) - SQL driver for postgres
* [SQLX](github.com/jmoiron/sqlx) - For connecting to postgres
* [MIGRATE](github.com/mattes/migrate) - For migrating postgres
* [TESTIFY](github.com/stretchr/testify) - For asserting tests
* [GO-SQLMOCK](github.com/DATA-DOG/go-sqlmock) - For mocking postgres

## Contributing

Please read [CONTRIBUTING.md](https://github.com/sudhanshuraheja/tanker/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](https://semver.org/spec/v2.0.0.html) for versioning based on the recommendation from [Dave Chaney](https://dave.cheney.net/2016/06/24/gophers-please-tag-your-releases). For the versions available, see the [tags on this repository](https://github.com/sudhanshuraheja/tanker/tags).

## Authors

* [Sudhanshu Raheja](https://github.com/sudhanshuraheja)

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](https://github.com/sudhanshuraheja/tanker/blob/master/LICENSE) file for details

## Acknowledgments

* Hat tip to every gopher out there who took the time to write articles so that newbies like me could learn
