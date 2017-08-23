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

var setters = map[reflect.Kind]func(parser.PathArg, string){
    reflect.Bool: boolWriter,
    reflect.String: stringWriter,
}
