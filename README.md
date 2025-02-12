[![Go Report Card](https://goreportcard.com/badge/github.com/sgaunet/envtemplate)](https://goreportcard.com/report/github.com/sgaunet/envtemplate)
[![GitHub release](https://img.shields.io/github/release/sgaunet/envtemplate.svg)](https://github.com/sgaunet/envtemplate/releases/latest)
![GitHub Downloads](https://img.shields.io/github/downloads/sgaunet/envtemplate/total)
[![Test Coverage](https://api.codeclimate.com/v1/badges/214a2cb2c610e725f513/test_coverage)](https://codeclimate.com/github/sgaunet/envtemplate/test_coverage)
[![Maintainability](https://api.codeclimate.com/v1/badges/214a2cb2c610e725f513/maintainability)](https://codeclimate.com/github/sgaunet/envtemplate/maintainability)
[![GoDoc](https://godoc.org/github.com/sgaunet/envtemplate?status.svg)](https://godoc.org/github.com/sgaunet/envtemplate)
[![License](https://img.shields.io/github/license/sgaunet/envtemplate.svg)](LICENSE)

[Forked from https://github.com/orls/envtemplate](https://github.com/orls/envtemplate)

Why forking the project ?

* Want a different CLI (add some flags)
* Author does not seem to be active
* Want binary for multiple architectures and a docker image to copy the binary from in multi stage build


# envtemplate

A super-lightweight tool for templating config files from environment variables – and nothing else.

## Usage

`$ envtemplate -i my-template-file > my-output-file`

or

```bash
$ cat test/input1.txt | envtemplate
Hello sylvain! Your home dir is /home/sylvain.
```

Usage:

```bash
envtemplate -h
  -h    Print help
  -i string
        File to encrypt/decrypt
  -v    Get version
```

Templating is done by the [go template package](https://golang.org/pkg/text/template/), where the only configured variables are the process's environment variables.

For example, a simple template might look like:

```
Hello {{ .USER }}! {{ if .HOME }}Your home dir is {{ .HOME }}.{{ else }}You don't appear to have a home dir set.{{ end }}
```

There is an indent function:

```
$ cat test/input2.txt 
{{ indent 4 .MULTILINE }}
$ export MULTILINE="ligne1
ligne2
ligne3"
$ go run . -i test/input2.txt 
    ligne1
    ligne2
    ligne3
```

## Demo

![demo](doc/demo.gif)

## Install

### homebrew

```bash
brew tap sgaunet/homebrew-tools
brew install sgaunet/tools/envtemplate
```

### from source

Assuming a working go installation, just  `git clone`, `cd` and `go build`. A binary named `envtemplate` should appear.

### prebuilt binary

A binary is attached to each [github release](https://github.com/sgaunet/envtemplate/releases). If you're happy to trust that, just fetch it with curl/similar (being sure to follow redirects):

`curl -L https://github.com/sgaunet/envtemplate/releases/download/0.1.0/envtemplate > /usr/bin/envtemplate && chmod +x /usr/bin/nvtemplate`

## Why?

This was borne out of frustration with using regular shell techniques – heredocs, `sed`, and similar – in various docker image-building and container-runtime configuration arrangements; for many config file formats (hi, nginx!) it starts to become unwieldy to manage conditional blocks, escaping, etc.

It is a kind of hybrid of [gotpl](https://github.com/tsg/gotpl) and [envtpl](https://github.com/andreasjansson/envtpl). In the target environment of docker container management, it's useful to:

- have small, easily-installable binary tools
    - ...ruling out `envtpl`; the extra docker image bloat of a python+pip install is.... far from zero
- provide variables directly as env vars
    - ...ruling out `gotpl`, which takes a yml file, needing extra pre-processing if env vars are your only means of configuring.

It shares some spiritual affinity to [confd](https://github.com/kelseyhightower/confd), in the way that a butter knife shares some spiritual affinity to a swiss-army knife.

If you want config pulled from remote datastores, or from yml files, or already have python in the relevant envs and like jinja syntax, then those projects may be better fits.
