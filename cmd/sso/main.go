package main

import (
	"Andrew-Savin-msk/sso/internal/config"
	"fmt"
)

func main() {
	cfg := config.Load()

	fmt.Println(cfg)
}
