package main

import (
	"fmt"
	"os"
)

func ReadFile(name string) (*os.File, error) {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %q", err)
	}
	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("error reading stat: %q", err)
	}
	if info.Size() == 0 {
		file, err := os.Create(name)
		if err != nil {
			return nil, fmt.Errorf("error creating file: %q", err)
		}
		return file, nil
	}
	return file, nil
}
