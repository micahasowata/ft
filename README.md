## ft

Get a file category based on its extension. It was built for a small project I'm working on which should explain why the list does not contain every possibility and its function is simple.

### How to use

Import the module...

```
$ go get github.com/spobly/ft@latest
```

...use it

```go

package main

import (
	"fmt"
    "errors"
	"github.com/spobly/ft/v2"
)

func main() {
	c, err := ft.GetCategory(".png")

	if err != nil {
       if errors.Is(err, ft.ErrNotFound) {
		// Handle not unsupported extension
	   }
	}

	fmt.Println(c)
}
```

### Contributing

This list does and cannot contain every file type and category so if you don't find what you are looking for, simply edit `ft.go` and create a PR.

I would merge it, you can be sure about that.

Happy hacking
