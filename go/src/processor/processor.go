package processor

import (
	"fmt"
	"config.Config"
)

type processor struct {
	config config.Config
}

func newProcessor(config *config.Config) {
		fmt.Printf("processor\n")
}
