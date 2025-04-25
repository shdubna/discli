## discli auth logout

Log out of a registry

```
discli auth logout [SERVER] [flags]
```

### Examples

```
  # Log out of reg.example.com
  discli auth logout reg.example.com
```

### Options

```
  -h, --help   help for logout
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

