package main

import (
	"fmt"
	"time"
)

const (
	RFC3339 = "2006-01-02T15:04:05Z07:00"
)

func main() {
	t, _ := time.Parse(RFC3339, "2019-01-01 15:22:22")
	fmt.Println(t)
}
