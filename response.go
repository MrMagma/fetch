package fetch

import (
    "net/http"
)

type Response struct {
    http.ResponseWriter
    Written bool
}

func (r Response) Write(b []byte) (int, error) {
    r.Written = true
    return r.ResponseWriter.Write(b)
}
