package fetch

import (
    "net/http"
    "fmt"
)

func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    response := Response{w, false}
    request := Request{req}
    r.serveConstHTTP(response, request)
    if !response.Written {
        r.serveVarHTTP(response, request)
    }
}

/*
Tries to find a constant listener to handle a request before it is passed on to variable listeners.
*/
func (r *Router) serveConstHTTP(response Response, request Request) {
    handlers := r.findConstantHandlers(request.URL.Path)
    for _, h := range handlers {
        h.handle(&h.data, response, request)
        if response.Written {
            break
        }
    }
}

/*
Finds all constant listeners that could handle the given path and returns them as a slice.
*/
func (r *Router) findConstantHandlers(path string) (listeners []*constantListener) {
    for i := 0; i < len(r.constant); i++ {
        if r.constant[i].path == path {
            listeners = append(listeners, &r.constant[i])
        }
    }
    return
}

/*
Finds and tries any variable handler that matches the given request path until either one writes a response or there are no more listeners left.
*/
func (r *Router) serveVarHTTP(response Response, request Request) {
    for _, h := range r.variable {
        s := h.pattern.FindStringSubmatch(request.URL.Path)
        if len(s) == len(h.args) + 1 {
            h.tryRun(s, response, request)
        }
        if response.Written {
            break
        }
    }
}
