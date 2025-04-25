## discli auth get

Implements a credential helper

```
discli auth get [REGISTRY_ADDR] [flags]
```

### Examples

```
  # Read configured credentials for reg.example.com
  $ echo "reg.example.com" | crane auth get
  {"username":"AzureDiamond","password":"hunter2"}
  # or
  $ crane auth get reg.example.com
  {"username":"AzureDiamond","password":"hunter2"}
```

### Options

```
  -h, --help   help for get
```

### Options inherited from parent commands

```
      --allow-nondistributable-artifacts   Allow pushing non-distributable (foreign) layers
      --insecure                           Allow image references to be fetched without TLS
      --platform platform                  Specifies the platform in the form os/arch[/variant][:osversion], or all for all available platforms (e.g. linux/amd64). (default linux/amd64)
  -v, --verbose                            Enable debug logs
```

### SEE ALSO

* [discli auth](auth.md)	 - Log in or access credentials

