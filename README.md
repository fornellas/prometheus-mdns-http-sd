[![Latest Release](https://img.shields.io/github/v/release/fornellas/prometheus-mdns-http-sd)](https://github.com/fornellas/prometheus-mdns-http-sd/releases)
[![Push](https://github.com/fornellas/prometheus-mdns-http-sd/actions/workflows/push.yaml/badge.svg)](https://github.com/fornellas/prometheus-mdns-http-sd/actions/workflows/push.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/fornellas/prometheus-mdns-http-sd)](https://goreportcard.com/report/github.com/fornellas/prometheus-mdns-http-sd)
[![Go Reference](https://pkg.go.dev/badge/github.com/fornellas/prometheus-mdns-http-sd.svg)](https://pkg.go.dev/github.com/fornellas/prometheus-mdns-http-sd)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Buy me a beer: donate](https://img.shields.io/badge/Donate-Buy%20me%20a%20beer-yellow)](https://www.paypal.com/donate?hosted_button_id=AX26JVRT2GS2Q)

# Prometheus mDNS HTTP Service Discovery

This is an implementation of [Prometheus HTTP Service Discovery](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#http_sd_config) for discovering targets via mDNS.

## Install

Pick the [latest release](https://github.com/fornellas/prometheus-mdns-http-sd/releases) with:

```bash
GOARCH=$(case $(uname -m) in i[23456]86) echo 386;; x86_64) echo amd64;; armv6l|armv7l) echo arm;; aarch64) echo arm64;; *) echo Unknown machine $(uname -m) 1>&2 ; exit 1 ;; esac) && wget -O- https://github.com/fornellas/prometheus-mdns-http-sd/releases/latest/download/prometheus-mdns-http-sd.linux.$GOARCH.gz | gunzip > prometheus-mdns-http-sd && chmod 755 prometheus-mdns-http-sd
./prometheus-mdns-http-sd --help
```

## Development

[Docker](https://www.docker.com/) is used to create a reproducible development environment on any machine:

```bash
git clone git@github.com:fornellas/prometheus-mdns-http-sd.git
cd prometheus-mdns-http-sd/
./builld.sh
```

Typically you'll want to stick to `./builld.sh rrb`, as it enables you to edit files as preferred, and the build will automatically be triggered on any file changes.