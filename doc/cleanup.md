## discli cleanup

Remove tags from repos specified by regex

```
discli cleanup REGISTRY [flags]
```

### Options

```
  -d, --dry-run             (Optional) Do everything except remove.
  -h, --help                help for cleanup
  -k, --keep-tags int       (Optional) Number of last SemVer sorted tags that will not be cleared. (default 3)
  -r, --repo-regex string   (Optional) The regulat expression for the repository selection. (default ".*")
  -t, --tag-regex string    (Optional) The regulat expression for the tag selection. (default ".*")
```

### Options inherited from parent commands

```
      --allow-nondistributable-artifacts   Allow pushing non-distributable (foreign) layers
      --insecure                           Allow image references to be fetched without TLS
      --platform platform                  Specifies the platform in the form os/arch[/variant][:osversion], or all for all available platforms (e.g. linux/amd64). (default linux/amd64)
  -v, --verbose                            Enable debug logs
```

