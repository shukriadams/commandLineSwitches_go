package commandLineArgs

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type CommandLineArgs struct {
	Arguments          map[string]string
	ConvertKeysToLower bool

	// Prints all arguments out to console
	Print func()

	GetValue func(name string) (string, error)
	HasValue func(name string) bool
}

// constructor
func New() *CommandLineArgs {
	var args = os.Args[1:]

	arguments := make(map[string]string)
	argsLength := len(args)
	for i := 0; i < argsLength; i++ {
		item := args[i]

		if !strings.HasPrefix(item, "-") && !strings.HasPrefix(item, "--") {
			continue
		}

		if strings.HasPrefix(item, "--") {
			item = item[2:]
		} else if strings.HasPrefix(item, "-") {
			item = item[1:]
			// ignore switches prefixed with - if their lengths are greater than 1
			if len(item) > 1 {
				continue
			}
		}

		value := ""
		nextItem := ""
		nextItemIsValue := false
		if i < argsLength-1 {
			nextItem = args[i+1]
			nextItemIsValue = !strings.HasPrefix(nextItem, "-") && !strings.HasPrefix(nextItem, "--")
		}

		if nextItemIsValue {
			value = nextItem
		}

		arguments[item] = value
	}

	var this = CommandLineArgs{
		Arguments: arguments,
	}

	this.Print = func() {
		for key, value := range this.Arguments {
			fmt.Println(key, value)
		}
	}

	this.GetValue = func(name string) (string, error) {

		value, ok := this.Arguments[name]
		if ok {
			return value, nil
		}

		return "", errors.New("No item with that name")
	}

	this.HasValue = func(name string) bool {
		_, ok := this.Arguments[name]

		if ok {
			return true
		}
		return false
	}

	return &this
}
