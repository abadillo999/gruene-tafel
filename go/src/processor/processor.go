package processor

import (
	"fmt"
	"config"
)

type Processor struct {
	_config* config.Config
}

func NewProcessor(config *config.Config) {
		fmt.Printf("processor\n")
}
