package parser

import (
    "reflect"
    "sort"
    "fmt"
)

var validKinds = [14]reflect.Kind{
    reflect.Bool,
    reflect.Int,
    reflect.Int8,
    reflect.Int16,
    reflect.Int32,
    reflect.Int64,
    reflect.Uint,
    reflect.Uint8,
    reflect.Uint16,
    reflect.Uint32,
    reflect.Uint64,
    reflect.Float32,
    reflect.Float64,
    reflect.String,
}

type PathArg struct {
    reflect.Value
    Name string
}

type ArgList []PathArg

/*
Finds the index of an argument with a given name in a list of arguments using a binary search algorithm. Returns -1 if no argument is found with the given name.
*/
func (args ArgList) GetArgIndex(name string) int {
    min, max := 0, len(args) - 1
    var avg int
    for max > min {
        avg = (min + max) / 2
        if name == args[avg].Name {
            return avg
        } else if stringLt(args[avg].Name, name) {
            min = avg + 1
        } else {
            max = avg - 1
        }
    }
    return -1
}

/*
Finds and returns an argument with the given name.
*/
func (args ArgList) GetArg(name string) (f PathArg) {
    index := args.GetArgIndex(name)
    if index >= 0 {
        f = args[index]
    }
    return
}

func (args ArgList) Len() int {
    return len(args)
}

func (args ArgList) Swap(i, j int) {
    args[i], args[j] = args[j], args[i]
}

/*
Utility function which tells whether one string is "less than" another. Used to sort a list of arguments and to quickly find an argument with a given name in that list.
*/
func stringLt(s1, s2 string) bool {
    for i, _ := range s1 {
        if i > len(s2) {
            return false
        } else if s1[i] < s2[i] {
            return true
        }
    }
    return false
}

/*
Used to sort the elements of an ArgList alphabetically.
*/
func (args ArgList) Less(i, j int) bool {
    return stringLt(args[i].Name, args[j].Name)
}

/*
Checks whether or not a struct's field's type is valid by using binary search to determine whether or not it is in a list of valid types.
*/
func kindIsValid(kind reflect.Kind) bool {
    min, max := 0, len(validKinds) - 1
    var avg int
    for max >= min {
        avg = (min + max) / 2
        if validKinds[avg] > kind {
            max = avg - 1
        } else if validKinds[avg] < kind {
            min = avg + 1
        } else {
            return true
        }
    }
    return false
}

/*
Extracts data from an arbitrary structure provided by the user, whose public fields are only of basic types (int, uint, float32, string, int16, etc.), into an array containing the names and types of those fields. Panics if a non-struct value is passed in or one of the struct's public fields is invalid.
*/
func ParseArgs(data interface{}) (args ArgList) {
    dataValue := reflect.ValueOf(data).Elem()
    fmt.Println(dataValue)
    dataType := dataValue.Type()
    fmt.Println(dataType)
    if dataType.Kind() != reflect.Struct {
        panic("Struct type required")
    }
    n := dataValue.NumField()
    var field reflect.Value
    for i := 0; i < n; i++ {
        field = dataValue.Field(i)
        if kindIsValid(field.Kind()) {
            args = append(args, PathArg{field, dataType.Field(i).Name})
        } else {
            panic("Did not recieve a valid type!")
        }
    }
    sort.Sort(args)
    return
}

