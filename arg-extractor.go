package fetch

import (
    "reflect"
    "strconv"
    "fetch/parser"
    "fmt"
)

/*
Attempts to run a variable listener, extracting its arguments from the request path, parsing them, and passing them to the data structure. Will fail if it is unable to pass a variable to the data structure.
*/
func (l *variableListener) tryRun(matches []string, response Response, request Request) {
    names := l.pattern.SubexpNames()
    for i, value := range matches[1:len(matches)] {
        arg := l.args.GetArg(names[i + 1])
        fmt.Println(arg)
        if arg.CanSet() {
            setters[arg.Kind()](arg, value)
        } else {
            fmt.Println("Unable to set field, " + names[i + 1])
            return
        }
    }
    l.handle(&l.data, response, request)
}

func boolWriter(f parser.PathArg, s string) {
    v, err := strconv.ParseBool(s)
    if err == nil {
        fmt.Println(v)
        f.SetBool(v)
    } else {
        fmt.Println(err)
    }
}

func stringWriter(f parser.PathArg, s string) {
    f.SetString(s)
}

func intWriter(f parser.PathArg, s string) {
    v, err := strconv.ParseInt(s, 10, f.Type().Bits())
    if err == nil {
        f.SetInt(v)
    } else {
        fmt.Println(err)
    }
}

func uintWriter(f parser.PathArg, s string) {
    v, err := strconv.ParseUint(s, 10, f.Type().Bits())
    if err == nil {
        f.SetUint(v)
    } else {
        fmt.Println(err)
    }
}

func floatWriter(f parser.PathArg, s string) {
    v, err := strconv.ParseFloat(s, f.Type().Bits())
    if err == nil {
        f.SetFloat(v)
    } else {
        fmt.Println(err)
    }
}

var setters = map[reflect.Kind]func(parser.PathArg, string){
    reflect.Bool: boolWriter,
    reflect.String: stringWriter,
    reflect.Int: intWriter,
    reflect.Int8: intWriter,
    reflect.Int16: intWriter,
    reflect.Int32: intWriter,
    reflect.Int64: intWriter,
    reflect.Uint: uintWriter,
    reflect.Uint8: uintWriter,
    reflect.Uint16: uintWriter,
    reflect.Uint32: uintWriter,
    reflect.Uint64: uintWriter,
    reflect.Float32: floatWriter,
    reflect.Float64: floatWriter,
}
