package processor

import (
	"fmt"
	"config"
)

type Executor struct {
    _config* config.Config
}

func NewExecutor(config *config.Config) {
	fmt.Printf("processor\n")
}