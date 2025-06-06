## discli auth login

Log in to a registry

```
discli auth login [OPTIONS] [SERVER] [flags]
```

### Examples

```
  # Log in to reg.example.com
  discli auth login reg.example.com -u AzureDiamond -p hunter2
```

### Options

```
  -h, --help              help for login
  -p, --password string   Password
      --password-stdin    Take the password from stdin
  -u, --username string   Username
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

