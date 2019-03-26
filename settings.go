package main

import "path/filepath"

var settings Settings

type Settings struct {
	SearchPath      string
	DestinationPath string
}

func init() {
	settings = Settings{}
}

func (this *Settings) GetDestinationPath(path string) (string, error) {
	var (
		destinationPath string
		err             error
		relPath         string
	)

	relPath, err = filepath.Rel(this.SearchPath, path)
	if err == nil {
		destinationPath = filepath.Join(this.DestinationPath, relPath)

		//updating extention
		destinationPath = destinationPath[0:len(destinationPath)-2] + "go"
	}

	return destinationPath, err
}
