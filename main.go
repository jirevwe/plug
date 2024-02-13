package plug

import (
	"context"
	"log"
	"time"
)

func Main() {
	ctx, cancel := New(context.Background())
	defer cancel()

	l, err := ctx.LoadModuleByID(ID)
	if err != nil {
		log.Fatal(err)
	}

	logger := l.(*Logger)
	logger.Info("running")

	ids := GetModules()

	counter := time.NewTicker(time.Second)
	for {
		select {
		case hmm := <-counter.C:
			for _, id := range ids {
				mod, err := ctx.LoadModuleByID(id)
				if err != nil {
					log.Fatal(err)
				}

				err = mod.(Emitter).Emit(hmm)
				if err != nil {
					log.Fatal(err)
				}
			}
			break
		}
	}
}
