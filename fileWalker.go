package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

//ProcessPath processing path
func ProcessPath(root string, rootWg *sync.WaitGroup) {

	//if root wait group is existing, then Done it when func will end
	if rootWg != nil {
		defer rootWg.Done()
	}

	//slice of processed folders
	mapFolders := make(map[string]bool)

	//local wait group
	var wg sync.WaitGroup

	filepath.Walk(root, func(dir string, fi os.FileInfo, err error) error {
		filePath := dir //filepath.Join(dir, fi.Name())

		newLog := new(ProcessingPathLog)
		newLog.Path = filePath

		if _, exists := mapFolders[filePath]; !exists {
			mapFolders[filePath] = true

			if err == nil {
				//don't check root directory itself
				//if _, tmpErr := os.Stat(filePath); !os.IsNotExist(tmpErr) {
				if root != dir {
					wg.Add(1)

					//if curent object is a directory, then process it async
					if fi.IsDir() {
						go ProcessPath(filePath, &wg)
					} else {
						//current object is a file, process it async
						go processFile(filePath, &wg)
					}
				}
				newLog.IsSuccess = true

			} else {
				newLog.IsSuccess = false
				newLog.Err = err
			}

			strProcesingResult := newLog.Path + ": "
			if newLog.IsSuccess {
				strProcesingResult += "Success"
			} else {
				strProcesingResult += "Failed - "
				if newLog.Err != nil {
					strProcesingResult += newLog.Err.Error()
				}
			}

			fmt.Println(strProcesingResult)
		}
		return nil
	})

	wg.Wait()
}

type ProcessingPathLog struct {
	Path      string
	IsSuccess bool
	Err       error
}
