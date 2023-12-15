# Translate

## Usage

### Install

```shell
go get github.com/eiixy/go-translate
```

### Example
```go
package main

import (
	"fmt"
	"github.com/eiixy/go-translate"
)

func main() {
	tr := translate.NewClient()
	texts, err := tr.Translates([]string{"hello", "world"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(texts)

	text, err := tr.Translate("hello world")
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println(text)
}

```