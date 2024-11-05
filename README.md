# Korean Romanizer in Golang

Simply ported from [korean-romanizer](https://github.com/osori/korean-romanizer)

## Usage

### Installation
``` bash
go get github.com/lokks307/korean-romanizer-go
```

### Basic Usage
```go
import (
    "fmt"
    korom "github.com/lokks307/korean-romanizer-go"
)

func main() {
    fmt.Println(korom.Romanize("안녕하세요"))   // annyeonghaseyo
}
```