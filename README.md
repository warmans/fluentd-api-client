# fluentd-api-client
Golang client for fluentd monitoring API (http://docs.fluentd.org/articles/monitoring)

The monitoring API provides information on active plugins and their configuration.


### Usage

```
package main

import (
  "fmt"
  "monitoring"
)

func main() 
  host := monitoring.NewHost("127.0.0.1:24220")
  host.Update()
  fmt.Print(host.Plugins)
}
```
