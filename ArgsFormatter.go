// 🤖 Github Repository: https://github.com/Arion-Kun/GoLaunch

package GoLaunch

import (
	"container/list"
	"os"
	"strings"
)

var _ArgsCache []map[string][]string

func GetSanitizedArgs() []map[string][]string {

	if _ArgsCache != nil {
		return _ArgsCache
	}

	argMap := parseArgs()
	mapBuffer := make([]map[string][]string, argMap.Len())

	for a := 0; a < argMap.Len(); a++ {

		argMapFront := argMap.Front()
		for str, strArr := range argMapFront.Value.(map[string][]string) {
			mapBuffer = append(mapBuffer, map[string][]string{str: strArr})
			argMap.MoveToBack(argMapFront)
		}
	}
	_ArgsCache = mapBuffer
	return mapBuffer
}

func TryGetValue(key string) (bool, []string) {
	if Contains(key) {
		return true, Get(key)
	}
	return false, nil
}

func Contains(key string) bool {
	args := GetSanitizedArgs()

	//Iterate over the arguments and check if the key is present, s2 is the key, and the value is discarded.
	for _, arg := range args {
		for s2 := range arg {
			if key == s2 {
				return true
			}
		}
	}
	return false
}
func Get(key string) []string {
	args := GetSanitizedArgs()

	//Iterate over the arguments and check if the key is present, s2 is the key, and the value is discarded.
	for _, arg := range args {
		for s2 := range arg {
			if key == s2 {
				return arg[s2]
			}
		}
	}

	return nil
}

// Rather create one than a bunch of new ones right?
const stringEmpty = ""

// private static list* parseArgs() ?
var parseArgs = func() *list.List {
	//os.Args = append(os.Args, "-a", "test", "one", "two", "three")
	//os.Args = append(os.Args, "--b", "four", "five", "six")

	var preparedList = list.New()

	keyBuffer := stringEmpty
	var valueBuffer []string
	for index, i := range os.Args {

		isVariableDeclaration := strings.HasPrefix(i, "-")

		if isVariableDeclaration { // Variable = -something or --something

			//i = strings.Replace(i, "-", stringEmpty, 2)

			if i == stringEmpty {
				continue
			}

			if keyBuffer != stringEmpty { // Needs saving
				preparedList.PushFront(map[string][]string{keyBuffer: valueBuffer}) // preparedList.Add( (string, string[]) )
			}
			keyBuffer = i
			valueBuffer = nil
		} else {
			if keyBuffer == stringEmpty || i == stringEmpty {
				continue
			} // if it doesn't start with a variable like - or --

			valueBuffer = append(valueBuffer, i) // Append strings till we find a new variableDeclaration

			if len(os.Args) == index+1 { // Last element
				preparedList.PushFront(map[string][]string{keyBuffer: valueBuffer})
			}
		}

	}

	return preparedList
}
