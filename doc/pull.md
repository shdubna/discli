## discli pull

Pull remote images by reference and store their contents locally

```
discli pull IMAGE TARBALL [flags]
```

### Options

```
      --annotate-ref        Preserves image reference used to pull as an annotation when used with --format=oci
  -c, --cache_path string   Path to cache image layers
      --format string       Format in which to save images ("tarball", "legacy", or "oci") (default "tarball")
  -h, --help                help for pull
```

### Options inherited from parent commands

```
      --allow-nondistributable-artifacts   Allow pushing non-distributable (foreign) layers
      --insecure                           Allow image references to be fetched without TLS
      --platform platform                  Specifies the platform in the form os/arch[/variant][:osversion], or all for all available platforms (e.g. linux/amd64). (default linux/amd64)
  -v, --verbose                            Enable debug logs
```

