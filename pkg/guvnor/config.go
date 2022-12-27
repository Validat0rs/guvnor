package guvnor

import (
	"os"

	"github.com/BurntSushi/toml"

	"github.com/Validat0rs/guvnor/pkg/guvnor/types"
)

type IReader interface {
	readFile() ([]byte, error)
}

type Reader struct {
	fileName string
}

func (g *Guvnor) ParseConfig(reader IReader) (*types.Config, error) {
	var config types.Config

	body, err := reader.readFile()
	if err != nil {
		return nil, err
	}

	if _, err := toml.Decode(string(body), &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (r *Reader) readFile() ([]byte, error) {
	file, err := os.ReadFile(r.fileName)

	if err != nil {
		return nil, err
	}

	return file, err
}
