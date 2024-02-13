package main

import (
	"context"
	"github.com/jirevwe/plug"
	"log"
	"time"
)

func main() {
	ctx, cancel := plug.New(context.Background())
	defer cancel()

	l, err := ctx.LoadModuleByID(plug.ID)
	if err != nil {
		log.Fatal(err)
	}

	logger := l.(*plug.Logger)
	logger.Info("running")

	modules := plug.GetModules()

	counter := time.NewTicker(time.Second)
	for {
		select {
		case hmm := <-counter.C:
			for _, id := range modules {
				mod, err := ctx.LoadModuleByID(id)
				if err != nil {
					log.Fatal(err)
				}

				err = mod.(plug.Emitter).Emit(hmm)
				if err != nil {
					log.Fatal(err)
				}
			}
			break
		}
	}
}
