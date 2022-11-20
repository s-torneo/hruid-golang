# hruid-golang

Golang version of this repository: https://github.com/nebbles/hruid-python.

## Usage

Install
```
go get github.com/s-torneo/hruid-golang
```

Use in Golang code
```
package main

import (
	"fmt"

	"github.com/s-torneo/hruid-golang"
)

func main() {
    // generate uid with default entities values
    uid := hruid.Generate()
    fmt.Println(uid)
}
```