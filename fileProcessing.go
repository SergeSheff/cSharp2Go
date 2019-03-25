package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)

var classRegexp *regexp.Regexp
var classPropertyRegexp *regexp.Regexp

var CURLY_BRACES_OPEN_REGEXP *regexp.Regexp
var CURLY_BRACES_CLOSE_REGEXP *regexp.Regexp

func processFile(filePath string, rootWg *sync.WaitGroup) {
	defer rootWg.Done()

	file, err := os.Open(filePath)
	if err == nil {
		defer file.Close()

		var classData *ClassData

		curlyBrasesCount := 0

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			str := scanner.Text()

			//if we found a class, start count braces and remember fields
			if classData != nil {

				//if we found all opened and closed braces (class code is finished), then create a go file
				if classData.startCalculation && curlyBrasesCount == 0 {
					newFilePath := filePath + ".go"
					f, fileError := os.Create(newFilePath)
					if fileError == nil {
						defer f.Close()

						_, fileError = f.WriteString(classData.GetGoType())
						if fileError == nil {
							fmt.Printf("%s was created\n", newFilePath)
						} else {
							fmt.Println(fileError)
						}
					} else {
						fmt.Println(fileError)
					}
					classData = nil
				} else {

					matches := classPropertyRegexp.FindAllStringSubmatch(str, -1)
					if len(matches) > 0 && len(matches[0]) > 3 {

						propertyName := applyAccessor(matches[0][1], matches[0][3])
						propertyType, typeError := getGolangType(matches[0][2])

						if typeError == nil {
							classData.fields[propertyName] = propertyType
						} else {
							fmt.Println(typeError.Error())
						}
					}

					//calculate braces count
					curlyBrasesCount += calculateMatches(str, CURLY_BRACES_OPEN_REGEXP)
					if curlyBrasesCount > 0 {
						classData.startCalculation = true
					}
					curlyBrasesCount -= calculateMatches(str, CURLY_BRACES_CLOSE_REGEXP)
				}

			} else {

				matches := classRegexp.FindAllStringSubmatch(str, -1)
				if len(matches) > 0 && len(matches[0]) > 3 {

					classData = &ClassData{
						className: matches[0][3],
					}

					classData.InitFields()

					classData.className = applyAccessor(matches[0][1], classData.className)

					curlyBrasesCount += calculateMatches(str, CURLY_BRACES_OPEN_REGEXP)
					if curlyBrasesCount > 0 {
						classData.startCalculation = true
					}
					curlyBrasesCount -= calculateMatches(str, CURLY_BRACES_CLOSE_REGEXP)
				}
			}
		}
	}
}

func calculateMatches(str string, regExp *regexp.Regexp) int {
	return len(regExp.FindAllStringSubmatch(str, -1))
}

func init() {
	classRegexp = regexp.MustCompile(`(?P<accessor>[a-z]+)\s+(partial\s+)?class\s+(?P<class>[^{]+)`)
	classPropertyRegexp = regexp.MustCompile(`([a-z]+)\s([\w<>\.]+)\s(\w+)\s`)

	CURLY_BRACES_OPEN_REGEXP = regexp.MustCompile(`(\{)`)
	CURLY_BRACES_CLOSE_REGEXP = regexp.MustCompile(`(\})`)
}

func applyAccessor(accessor string, str string) string {

	accessor = strings.ToLower(accessor)

	var goAccessor = string((str)[0])
	if accessor == "public" || accessor == "internal" {
		goAccessor = strings.ToUpper(goAccessor)
	} else {
		goAccessor = strings.ToLower(goAccessor)
	}

	return goAccessor + (str)[1:len(str)]
}
