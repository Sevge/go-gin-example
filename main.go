package main

import (
    "fmt"
    "net/http"

    "github.com/sevge/go-gin-example/routers"
    "github.com/sevge/go-gin-example/pkg/setting"
)

func main() {
    router := routers.InitRouter()

    s := &http.Server{
        Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
        Handler:        router,
        ReadTimeout:    setting.ReadTimeout,
        WriteTimeout:   setting.WriteTimeout,
        MaxHeaderBytes: 1 << 20,
    }

    s.ListenAndServe()
}

