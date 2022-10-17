package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"github.com/kataras/iris/v12"
	_ "github.com/kataras/iris/v12"
	service "github.com/monstrum/grpc-iris-demo/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	address = "127.0.0.1:3443" // Not working
)

func main() {
	app := iris.New()
	// cred := insecure.NewCredentials()
	// cred, err := credentials.NewClientTLSFromFile("./server.crt", "localhost")
	cred := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
		Renegotiation:      tls.RenegotiateNever,
	})

	interceptors := []grpc.DialOption{
		grpc.WithTransportCredentials(cred),
		grpc.WithBlock(),
	}

	app.Get("/", func(ctx iris.Context) {
		conn, err := grpc.Dial(
			address,
			interceptors...,
		)
		if err != nil {
			_, _ = ctx.Writef("conn err %s", err)
			return
		}
		defer conn.Close()

		timeout, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		c := service.NewApiClient(conn)
		req := service.CreateProductReq{
			Product: &service.ProductReq{
				Name:        "Demo",
				Description: "Descr",
			},
		}

		r, err := c.Create(timeout, &req)
		if err != nil {
			_, _ = ctx.Writef("c.Create err %s", err)
			return
		}
		r.GetProduct()
		_, _ = ctx.Writef("client %s", r.Product)
	})

	app.Get("/test", func(ctx iris.Context) {
		jsonStr := []byte(`{"product":{"name":"HelloWorld", "description":"Hello Description"}}`)
		req, err := http.NewRequest("POST", "https://localhost:3443/product.Api/Create", bytes.NewBuffer(jsonStr))
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		r, err := client.Do(req)
		if err != nil {
			_, _ = ctx.Writef("conn err %s", err)
			return
		}
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			_, _ = ctx.Writef("conn err %s", err)
			return
		}
		_, _ = ctx.Writef("body %s", body)
	})

	// http://localhost:18080
	// http://localhost:18080/test
	app.Listen(":18080")
}
