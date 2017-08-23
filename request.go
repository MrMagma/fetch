package fetch

import (
    "net/http"
)

type Request struct {
    *http.Request
}

func (r *Request) Send() {

}

func (r *Request) Resolve() {

}
