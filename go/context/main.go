package main

import (
	C "context"
	"log"
	"sync"

	"gitlab.com/rwxrob/cmdtab"
)

func main() {
	wg := new(sync.WaitGroup)
	ctx, cancel := C.WithCancel(C.Background())
	go first(ctx, wg)
	go second(ctx, wg)
	go third(ctx, wg)
	cmdtab.WaitForInterrupt(cancel)
	wg.Wait()
}

func first(ctx C.Context, wg *sync.WaitGroup) {
	log.Println("starting first")
	wg.Add(1)
	<-ctx.Done()
	log.Println("stopping first")
	wg.Done()
}

func second(ctx C.Context, wg *sync.WaitGroup) {
	log.Println("starting second")
	wg.Add(1)
	<-ctx.Done()
	log.Println("stopping second")
	wg.Done()
}

func third(ctx C.Context, wg *sync.WaitGroup) {
	log.Println("starting third")
	wg.Add(1)
	<-ctx.Done()
	log.Println("stopping third")
	wg.Done()
}
