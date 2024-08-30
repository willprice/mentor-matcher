package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
    "github.com/urfave/cli/v2"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    app := &cli.App{
        Name:  "mentor-matcher",
        Usage: "Run the mentor matcher server",
        Flags: []cli.Flag{
            &cli.IntFlag{
                Name:  "port",
                Value: 8080,
                Usage: "Port to listen on",
            },
        },

        Action: func(ctx *cli.Context) error {
            http.HandleFunc("/", handler)
            addr := fmt.Sprintf(":%v", ctx.Int("port"))
            log.Printf("Starting server on %v", addr)
            log.Fatal(http.ListenAndServe(addr, nil))
            return nil
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
