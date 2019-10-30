package processor

import (
	"fmt"
	"config"
)

type Processor struct {
	_config* config.Config
}

func NewProcessor (config *config.ENVConfig) *Processor {
		fmt.Printf("processor\n")
		processor := Processor{}
		return &processor
}
