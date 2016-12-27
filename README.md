# suddendeath

suddeandeath is

```
＿人人人人人人＿
＞　突然の死　＜
￣Y^Y^Y^Y^Y^Y￣
```

## Command

``` sh
$ go get github.com/kechako/suddendeath/cmd/suddendeath
$ suddendeath 突然の死
＿人人人人人人＿
＞　突然の死　＜
￣Y^Y^Y^Y^Y^Y￣
```

## Package

``` sh
$ go get github.com/kechako/suddendeath
```

``` go
packamge main

import (
    "fmt"

    "github.com/kechako/suddendeath"
)

func main() {
    fmt.Println(suddendeath.Generate("突然の死"))
    // ＿人人人人人人＿
    // ＞　突然の死　＜
    // ￣Y^Y^Y^Y^Y^Y￣
}
```
