## discli ls

discli the tags in a repo

```
discli ls REPO [flags]
```

### Options

```
      --full-ref           (Optional) if true, print the full image reference
  -h, --help               help for ls
  -O, --omit-digest-tags   (Optional), if true, omit digest tags (e.g., ':sha256-...')
```

### Options inherited from parent commands

```
      --allow-nondistributable-artifacts   Allow pushing non-distributable (foreign) layers
      --insecure                           Allow image references to be fetched without TLS
      --platform platform                  Specifies the platform in the form os/arch[/variant][:osversion], or all for all available platforms (e.g. linux/amd64). (default linux/amd64)
  -v, --verbose                            Enable debug logs
```

