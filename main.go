package main

import (
	"golang.org/x/sync/errgroup"
	"log"

	"github.com/monstrum/grpc-iris-demo/cmd"
)

var g errgroup.Group

// go run *.go
func main() {
	cmds := []cmd.Command{
		cmd.CreateGrpcServer(),
	}

	for _, c := range cmds {
		g.Go(c.Execute)
	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
