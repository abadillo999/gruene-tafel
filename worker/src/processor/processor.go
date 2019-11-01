package processor

import (
	"fmt"
	"config"
	"model"
	"time"
)

type Processor struct {
	_config* config.Config
}

func NewProcessor (config *config.ENVConfig) *Processor {
		fmt.Printf("processor\n")
		processor := Processor{}
		return &processor
}

func (processor *Processor) CheckRequest(task *model.Task) bool {
	createId(task)
	return true
}

func (processor *Processor) Process(task *model.Task) string {
	return ""
}

func createId(task *model.Task) {
	timestamp := time.Now()
	formatedTime := timestamp.Format(time.RFC3339)
	task.Id = task.ClientId + "|" + formatedTime
}