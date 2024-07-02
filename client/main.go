package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"gsm/client/notifications"
	"gsm/client/swagger"
	"gsm/config"
	"log"

	_ "gsm/api"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	fmt.Println("client starting...")

	engine := gin.Default()

	conn, err := newClient()
	if err != nil {
		log.Println(err)
		return
	}

	//task.MakeConnection(conn)
	notifications.MakeConnection(conn)

	group := engine.Group("api")
	swagger.AddRoutes(group)
	//task.AddRoutes(group)
	notifications.AddRoutes(group)

	engine.Run(":8080")
}

func generateTLSCreds() (credentials.TransportCredentials, error) {
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}

	return credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	}), nil
}

func newClient() (*grpc.ClientConn, error) {

	tlsCreds, err := generateTLSCreds()
	if err != nil {
		log.Fatal(err)
	}

	//conn, err := grpc.NewClient(config.SERVER_GRPC_ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(config.SERVER_GRPC_ADDRESS, grpc.WithTransportCredentials(tlsCreds))

	return conn, err
}
