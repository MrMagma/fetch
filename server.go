package fetch

import (
    "net/http"
    "time"
)

type Server struct {
    http.Server
    Router Router
    ReadTimeout time.Duration
    WriteTimeout time.Duration
    MaxHeaderBytes int
}

func (s Server) Start(addr string) error {
    s.Server = http.Server{
        Handler: &s.Router,
        Addr: addr,
        ReadTimeout: s.ReadTimeout,
        WriteTimeout: s.WriteTimeout,
        MaxHeaderBytes: s.MaxHeaderBytes,
    }
    return s.ListenAndServe()
}
