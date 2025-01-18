package pkg

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type Config struct {
	// TODO
}

func readConfig(path string) (*Config, error) {
	if path == "" {
		return nil, errors.New("path is empty")
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "can't open file")
	}
	defer file.Close()
	cfg := &Config{}
	return cfg, nil
}

func WorkingWithErrors(){
	fmt.Println("Reading config file")
	config, err := readConfig("apple/config.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		// os.Exit(1)
	}
	fmt.Println(config)
}