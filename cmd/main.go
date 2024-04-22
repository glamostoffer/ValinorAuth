package main

import (
	"fmt"
	"github.com/glamostoffer/ValinorProtos/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	fmt.Println(cfg)
}
