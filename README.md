## Simple env to struct library in Go

Supported data types:
1. bool
2. int, int8, int16, int32, int64
3. uint, uint8, uint16, uint32, uint64
4. float32, float64
6. string

Supported struct tags: `env`, `default`

`env` is a tag which describes environment variable name

`default` is a tag which describes default value for given environment variable

Installation

```bash
go get github.com/MrWebUzb/goenv
```

Run tests

```bash
make test
```

or

```bash
go test ./...
```

Usage example

```go
package main

import (
	"fmt"
	"log"

	"github.com/MrWebUzb/goenv"
)

// Config structure
type Config struct {
	Debug      bool    `env:"DEBUG" default:"true"`
	DBName     string  `env:"DB_NAME" default:"postgres"`
	Port       int64   `env:"PORT" default:"8080"`
	FeePercent float32 `env:"FEE_PERCENT" default:"3.3"`
}

func main() {
	env, err := goenv.New()
	if err != nil {
		log.Fatalf("could not create env configuration: %v\n", err)
		return
	}

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		log.Fatalf("could not parse env to struct %v\n", err)
		return
	}

	fmt.Printf("Name: Debug, Type: %T, Value: %v\n", cfg.Debug, cfg.Debug)
	fmt.Printf("Name: DBName, Type: %T, Value: %v\n", cfg.DBName, cfg.DBName)
	fmt.Printf("Name: Port, Type: %T, Value: %v\n", cfg.Port, cfg.Port)
	fmt.Printf("Name: FeePercent, Type: %T, Value: %v\n", cfg.FeePercent, cfg.FeePercent)
}
```