//go:build !tui
// +build !tui

package main

import "fmt"

func main() {
	fmt.Println("TUI build tag is disabled. Use: go run -tags tui ./cmd/tui [host:port|port]")
}
