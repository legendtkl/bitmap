# bitmap
bitmap is implemented by golang

# example
```go
package main

import (
    "github.com/legendtkl/bitmap"
    "fmt"
)

func main() {
    bitMap, _ := bitmap.NewBitMap(100)
    err := bitMap.SetOne(101)   // out of size
    if err != nil {
        fmt.Println(err)
    }

    bitMap.SetOne(99)
    x, err = bitMap.GetPosition(99)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(x)
    }
}
```
