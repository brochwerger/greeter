package main

import (
	"fmt"
	"greeter/stress"
	"log"
	"net/http"
	"strconv"
	"time"
)

func stressCpu(w http.ResponseWriter, r *http.Request) {
    n := 20
	params := r.URL.Query()
	for k, v := range params {
		switch k {
		// TODO: add error handling code
		case "N":
			n, _ = strconv.Atoi(v[0])
        }
    }

    t0 := time.Now()
    fib := stress.RecursiveFibonacci(n)
    elapsed := time.Since(t0)
    fmt.Fprintf(w, "fibonacci(%d) == %d (%s)\n", n, fib, elapsed)
}

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World, I am a golang based server\n")
    })

    http.HandleFunc("/stress", stressCpu)

    log.Fatal(http.ListenAndServe(":8080", nil))

}
