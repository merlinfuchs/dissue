# Dissue

Synchronize Github issues with a Discord forum channel

This is work in progress!

## Quick Start

### Install

```shell
go install github.com/merlinfuchs/dissue
```

### Configure

Create a file called `config.yml` in the working directory with the following format:

```yml
discord:
  token: ""

github:
  token: ""
  webhook_secret: ""

api:
  host: "localhost"
  port: 8080

db:
  path: "data"
```

### Github Webhook

Create a Github Webhook that points to `/webhook` of the API.

### Start

```shell
dissue
```

### Usage

In Discord run the following command to enable synchronization:

```
/dissue enable channel: #some-channel repository: github.com/some/repo
```
