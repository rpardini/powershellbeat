package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/rpardini/powershellbeat/beater"
)

func main() {
	err := beat.Run("powershellbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
