package fetch

import (
    "fetch/parser"
    "reflect"
    "regexp"
)

type constantListener struct {
    method string
    path string
    data interface{}
    handle func(*interface{}, Response, Request)
}

type variableListener struct {
    method string
    pattern *regexp.Regexp
    args parser.ArgList
    data interface{}
    handle func(*interface{}, Response, Request)
}

type Router struct {
    constant []constantListener
    variable []variableListener
}

type Handler struct {
    Method string
    Pattern string
    Data interface{}
    Handle func(*interface{}, Response, Request)
}

/*
Registers a handler with a router, checking to make sure that the handler is valid (arguments consistent). Panics if the handler is invalid.
*/
func (r *Router) Handle(h Handler) {
    patternArgs, pattern := parser.ParsePattern(h.Pattern)
    args := parser.ParseArgs(h.Data)
    if !argsConsistent(patternArgs, args) {
        panic("Arguments are not consistent")
    }
    if len(args) == 0 {
        r.constant = append(r.constant, constantListener{
            method: h.Method,
            path: pattern,
            handle: h.Handle,
            data: h.Data,
        })
    } else {
        r.variable = append(r.variable, variableListener{
            method: h.Method,
            pattern: regexp.MustCompile("^" + pattern + "$"),
            args: args,
            handle: h.Handle,
            data: h.Data,
        })
    }
}

/*
Verifies that there are no inconsistencies between the arguments extracted from a handler's pattern and its data interface. Namely, each argument extracted from the data interface should be captured once and only once in the pattern and the pattern should not capture any arguments with names that were not extracted from the data interface.
*/
func argsConsistent(pArgs parser.PatternArgs, args parser.ArgList) bool {
    if len(args) != len(pArgs) {
        return false
    }
    if len(args) == 0 {
        return true
    }
    var argIndex int
    for i := 0; i < len(pArgs); i++ {
        if i < len(pArgs) - 1 && pArgs[i].Name == pArgs[i + 1].Name {
            return false
        }
        argIndex = args.GetArgIndex(pArgs[i].Name)
        if argIndex < 0 {
            return false
        }
        if pArgs[i].Type == parser.Multi && args[argIndex].Kind() != reflect.String {
            return false
        }
    }
    return true
}

