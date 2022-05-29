# Guvnor

Tools for managing Cosmos on-chain governance proposals.

Currently `guvnor` supports converting a chains' governance proposals into an RSS feed.

## Setup

### Golang

If you are new to Golang, please follow the setup instructions [here](https://golang.org/doc/install).

### Environment

Before running the `guvnor` service, please ensure that you have the following environment variables set:

|Var|Description|
|---|-----------|
|`GUVNOR_PORT`|The port that the service should run on (e.g.: `3000`)|
|`GUVNOR_CONFIG`|The full path to the `config.toml` file.|
|`GUVNOR_DOMAIN`|The domain that the service is being run on.|
|`REDIS_URL`|The Redis URL (e.g.: `localhost:6379`).|
|`REDIS_PASSWORD`|The Redis password (leave blank if no password is set).|
|`PROPOSAL_FEED_AUTHOR`|The name of the author, for the feed.|
|`PROPOSAL_FEED_EMAIL`|The email address of the author, for the feed.|

## Config

### Setup

To setup the config, please run:

```console
make setup-config
```

This will create a `~/.guvnor/config` directory and copy the example config file into. It will also set the `GUVNOR_CONFIG` environment variable.

You can easily override the config location by changing `GUVNOR_CONFIG`.

### Chains

Remove the placeholder content and update it with those that are relevant to the chains that you wish to generate feeds for. The structure of the config file is very simple, and you add a given chain like so:

```
[[chain]]
chain_id = "<chain_id>"
api = "<api_host>"
```

## Run

You can run the `guvnor` service on any cloud or bare metal provider. A Heroku `Procfile` (please see [here](https://devcenter.heroku.com/articles/getting-started-with-go) for how to launch this on Heroku) as also been included.

Please ensure that you have a Redis instance available, as the `guvnor` service makes use of Redis to cache requests (for up to 6 hours).

### Install

To install the binary, run:

```console
make clean install
```

### Start

To then to start the service:

```console
make run-guvnor-service
```

Once running, you may access the feed at `<hostname>/<chain_id>/proposals/rss`.

## Development

### Linter

To run the linter:

```console
make lint
```

### Tests

To run the tests and see test coverage:

```console
make tests
```

### Cache

As mentioned above, Redis is used to cache frequently requested objects. To start a local Redis instance, run:

```console
make docker-redis
```
