// This is awful code :)
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
    "github.com/mediocregopher/radix"
)

var (
	static http.Handler
    cache *radix.Pool
)

func main() {
	static = http.FileServer(http.Dir("."))

	http.HandleFunc("/version", handleVersion)
	http.HandleFunc("/args", handleArgs)
    http.HandleFunc("/panic", handlePanic)
    http.HandleFunc("/counter", handleCounter)
    http.Handle("/", static)


	log.Fatal(http.ListenAndServe(":3001", nil))
}

func handleVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "0.2")
}

func handleArgs(w http.ResponseWriter, r *http.Request) {
	for _, arg := range os.Args[1:] {
		fmt.Fprintln(w, arg)
	}
}

func handlePanic(w http.ResponseWriter, r *http.Request) {
    panic("Eek!")
}

func handleCounter(w http.ResponseWriter, r *http.Request) {
    var err error
    var counter int

    if cache == nil {
        cache, err = radix.NewPool("tcp", "cache:6379", 10)

        if err != nil {
            panic(err)
        }
    }

    err = cache.Do(radix.Cmd(&counter, "INCR", "theCounter"))

    if err != nil {
        panic(err)
    }

    fmt.Fprintln(w, counter)
}

