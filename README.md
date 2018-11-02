# r53dyndns

Dynamic DNS service backed by AWS Route 53.

[![CircleCI](https://circleci.com/gh/gregarmer/r53dyndns/tree/master.svg?style=svg)](https://circleci.com/gh/gregarmer/r53dyndns/tree/master)

## Overview

This app will update a DNS record on Route 53 with your current IP.

## Installation

### Debian

```bash
$ curl -s https://packagecloud.io/install/repositories/gregarmer/packages/script.deb.sh | sudo bash
```

### CentOS (or RHEL / other rpm based distro's)

```bash
$ curl -s https://packagecloud.io/install/repositories/gregarmer/packages/script.rpm.sh | sudo bash
```

## Getting Started

1. Follow the installation steps above.
2. Configure ~/.r53dyndns or r53dyndns.cfg - something like this:

    ```json
    {
      "aws_access_key": "UBGKJEBGKE56783JHVFW",
      "aws_secret_key": "webgrwebgjwbewegfkeg"
    }
    ```

3. Run `r53dyndns -d your.domain.com`

## Usage

```
$ r53dyndns -h
Usage of r53dyndns:
  -c string  path to the config file (default "~/.r53dyndns")
  -d string  domain to update
  -v         be verbose
```

`-c config.cfg` is optional. If you don't specify a config it'll default to
`~/.r53dyndns`.

The `-v` parameter will make r53dyndns be verbose about what is actually
happening.
