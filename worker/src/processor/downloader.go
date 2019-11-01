package processor

import (
	"fmt"
	"config"
)

type Downloader struct {
	config *config.Config
}

func NewDownloader(config *config.Config) {
	fmt.Printf("processor\n")
}