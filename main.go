package main

import (
	"github.com/zserge/lorca"
	"log"
	"os"
	"os/signal"
)

func main() {
	ui, err := lorca.New("https://www.baidu.com", "", 800, 600)
	if err != nil {
		return
	}
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	log.Println("exiting...")
}
