## Introduction

This is a repo for capturing various pre-boarding activities I've done to ramp-up before joining the GCP Serverless Runtime Team.

View the [project](https://github.com/users/smsohan/projects/1) to understand the context for this repo.

## Code Structure

```bash

|- api # contains the HTTP handler code. This is the function for GCF
|- go-buildpack # an opinionated and minimal buildpack for go
|- main.go # Starts and listens to :8080 serving a simple JSON response
|- app.yaml # the config for deploying the app using GAE
|- package.toml # a minimal CNB package declaration
|- run_pack.sh # a single command to build the image using CNB

```

## Prerequisites

- [x] Go 1.16 or newer
- [x] Cloud Native Buildpacks
- [x] Docker

## Build Status

[![Go](https://github.com/smsohan/gcp-rampup/actions/workflows/go.yml/badge.svg)](https://github.com/smsohan/gcp-rampup/actions/workflows/go.yml)

