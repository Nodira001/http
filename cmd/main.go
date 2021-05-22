package main

import (
	"log"
	"net"
	"os"
	"strconv"

	"github.com/Nodira001/http/pkg/server"
)

func main() {
	host := "0.0.0.0"
	port := "9999"
	if err := execute(host, port); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func BodyWriter(conn net.Conn, body string) (err error) {
	_, err = conn.Write([]byte(
		"HTTP/1.1 200 OK \r\n" +
			"Content-Length: " + strconv.Itoa(len(body)) + "\r\n" +
			"Content-Type: text/html\r\n" +
			"Connection: close\r\n" + "\r\n" + body,
	))
	return err
}

func execute(host string, port string) (err error) {
	svr := server.NewServer(net.JoinHostPort(host, port))

	svr.Register("/api/users", func(req *server.Request) {
		// log.Println(req.PathParams)
		name := req.QueryParams["name"][0]
		age := req.QueryParams["age"][0]
		log.Println(req.QueryParams)
		body := "Имя: " + name + ", возраст: " + age
		err := BodyWriter(req.Conn, body)
		if err != nil {
			log.Println(err)
		}
	})

	return svr.Start()
}
