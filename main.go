package main

import (
	"log"
	"os"
	"os/signal"

	//"daltondiaz/async-jobs/http"
	"daltondiaz/async-jobs/pkg"
)

func main() {
	pkg.Run()
    //http.Start()
    listen()
}

func listen() {
    sig := make(chan os.Signal)
    signal.Notify(sig, os.Interrupt, os.Kill)
    <-sig
    log.Println("Finished job")
}
