---
title: Check Event
menu:
  product_searchlight_6.0.0-alpha.0:
    identifier: hyperalert-check-event
    name: Check Event
    parent: hyperalert-cli
product_name: searchlight
section_menu_id: reference
menu_name: product_searchlight_6.0.0-alpha.0
---
## hyperalert check_event

Check kubernetes events for all namespaces

### Synopsis

Check kubernetes events for all namespaces

```
hyperalert check_event [flags]
```

### Options

```
  -c, --checkInterval duration           Icinga check_interval in duration. [Format: 30s, 5m]
  -s, --clockSkew duration               Add skew with check_interval in duration. [Default: 30s] (default 30s)
  -h, --help                             help for check_event
  -H, --host string                      Icinga host name
      --involvedObjectKind string        Involved object kind used to select events
      --involvedObjectName string        Involved object name used to select events
      --involvedObjectNamespace string   Involved object namespace used to select events
      --involvedObjectUID string         Involved object uid used to select events
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


