---
title: Check Volume
menu:
  product_searchlight_6.0.0-alpha.0:
    identifier: hyperalert-check-volume
    name: Check Volume
    parent: hyperalert-cli
product_name: searchlight
section_menu_id: reference
menu_name: product_searchlight_6.0.0-alpha.0
---
## hyperalert check_volume

Check kubernetes volume

### Synopsis

Check kubernetes volume

```
hyperalert check_volume [flags]
```

### Options

```
  -c, --critical float      Critical level value (usage percentage) (default 95)
  -h, --help                help for check_volume
  -H, --host string         Icinga host name
      --nodeStat            Checking Node disk size
  -s, --secretName string   Kubernetes secret name
  -N, --volumeName string   Volume name
  -w, --warning float       Warning level value (usage percentage) (default 80)
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --context string                   Use the context in kubeconfig
      --kubeconfig string                Path to kubeconfig file with authorization information (the master location is set by the master flag).
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
      --stderrthreshold severity         logs at or above this threshold go to stderr
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO

* [hyperalert](/docs/reference/hyperalert/hyperalert.md)	 - AppsCode Icinga2 plugin


