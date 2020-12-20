# amka-go

[![GoDoc](https://godoc.org/github.com/zoispag/amka-go?status.svg)](https://godoc.org/github.com/zoispag/amka-go)
![CI](https://github.com/zoispag/amka-go/workflows/CI/badge.svg)

A validator for greek social security number (AMKA)

## Usage

```go
import "github.com/zoispag/amka-go"

func main() {
	// An invalid AMKA
    isValid, _ = amka.Validate("09095986680"); // false
    
    // A valid AMKA
    isValid, err = amka.Validate("09095986684"); // true
}
```
