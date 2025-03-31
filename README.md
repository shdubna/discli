# Docker registry client
[![Go Report Card](https://goreportcard.com/badge/github.com/shdubna/discli)](https://goreportcard.com/report/github.com/shdubna/discli)
[![GitHub CodeQL](https://github.com/shdubna/discli/workflows/CodeQL/badge.svg)](https://github.com/shdubna/discli/actions?query=workflow%3CodeQL)
[![GitHub Release](https://github.com/shdubna/discli/workflows/Release/badge.svg)](https://github.com/shdubna/discli/actions?query=workflow%3ARelease)
[![GitHub license](https://img.shields.io/github/license/shdubna/discli.svg)](https://github.com/shdubna/discli/blob/main/LICENSE)
[![GitHub tag](https://img.shields.io/github/v/tag/shdubna/discli?label=latest)](https://github.com/shdubna/discli/releases)

Client for managing container images on [Docker registry](https://github.com/distribution/distribution) 
based on [go-containerregistry](https://github.com/google/go-containerregistry) 
and [crane](https://github.com/google/go-containerregistry/tree/main/cmd/crane).

## Installation

1. Get the [latest release](https://github.com/shdubna/discli/releases/latest) version.

   ```sh
   VERSION=$(curl -s "https://api.github.com/repos/shdubna/discli/releases/latest" | jq -r '.tag_name')
   ```

   or set a specific version:

   ```sh
   VERSION=vX.Y.Z   # Version number with a leading v
   ```

2. Download the release.

   ```sh
   OS=Linux       # or Darwin, Windows
   ARCH=x86_64    # or arm64, x86_64, armv6
   curl -sL "https://github.com/shdubna/discli/releases/download/${VERSION}/discli_${OS}_${ARCH}.tar.gz" > discli.tar.gz
   ```

3. Unpack it in the PATH.

   ```sh
   tar -zxvf discli.tar.gz -C /usr/local/bin/ discli
   ```

### Options

```
      --allow-nondistributable-artifacts   Allow pushing non-distributable (foreign) layers
  -h, --help                               help for discli
      --insecure                           Allow image references to be fetched without TLS
      --platform platform                  Specifies the platform in the form os/arch[/variant][:osversion] (e.g. linux/amd64). (default all)
  -v, --verbose                            Enable debug logs
```

## Commands
- [discli auth](doc/auth.md) - Log in or access credentials
- [discli catalog](doc/catalog.md) - List the repos in a registry
- [discli cleanup](doc/cleanup.md) - Cleanup registry from outdated images
- [discli copy](doc/copy.md) - Efficiently copy a remote image from src to dst while retaining the digest value
- [discli delete](doc/delete.md) - Delete an image reference from its registry
- [discli ls](doc/ls.md) - List the tags in a repo
- [discli pull](doc/pull.md) - Pull remote images by reference and store their contents locally
- [discli push](doc/push.md) - Push local image contents to a remote registry
- discli version - Print the version