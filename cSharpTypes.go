package main

import (
	"errors"
	"fmt"
)

var cSharpTypes map[string]string

//import null "gopkg.in/guregu/null.v3"

func init() {
	cSharpTypes = map[string]string{
		"decimal":         "float64",
		"int":             "int",
		"string":          "string",
		"long":            "int64",
		"short":           "int16",
		"byte":            "int16",
		"bool":            "bool",
		"System.DateTime": "int",
		"System.Guid":     "string",

		"Nullable<decimal>":         "null.Float",
		"Nullable<int>":             "null.Int",
		"Nullable<string>":          "null.String",
		"Nullable<long>":            "null.Int",
		"Nullable<short>":           "null.Int",
		"Nullable<byte>":            "null.Int",
		"Nullable<bool>":            "null.Bool",
		"Nullable<System.DateTime>": "null.Int",
		"Nullable<System.Guid>":     "null.String",
	}
}

func getGolangType(cSharpType string) (string, error) {

	var (
		goType string
		err    error
		exists bool
	)

	goType, exists = cSharpTypes[cSharpType]
	if !exists {
		err = errors.New(fmt.Sprintf("%s doesn't exists", cSharpType))
	}

	return goType, err
}
