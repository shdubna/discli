## discli auth

Log in or access credentials

```
discli auth [flags]
```

### Options

```
  -h, --help   help for auth
```

### Options inherited from parent commands

```
      --allow-nondistributable-artifacts   Allow pushing non-distributable (foreign) layers
      --insecure                           Allow image references to be fetched without TLS
      --platform platform                  Specifies the platform in the form os/arch[/variant][:osversion], or all for all available platforms (e.g. linux/amd64). (default linux/amd64)
  -v, --verbose                            Enable debug logs
```

### SEE ALSO

* [discli auth get](auth_get.md)	     - Implements a credential helper
* [discli auth login](auth_login.md)	 - Log in to a registry
* [discli auth logout](auth_logout.md)	 - Log out of a registry
* [discli auth token](auth_token.md)	 - Retrieves a token for a remote repo

