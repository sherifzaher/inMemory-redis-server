package server

import (
	"bufio"
	"fmt"
	"github.com/sherifzaher/inMemory-redis-server/pkg/internal/db"
	"io"
	"log"
	"net"
	"strings"
)

type Server struct {
	address string
	db      *db.DB
}

func New(address string) *Server {
	return &Server{
		address: address,
		db:      db.New(),
	}
}

func (server *Server) Start() error {
	listener, err := net.Listen("tcp", server.address)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go server.handleConnection(conn)
	}

	return nil
}

func (server *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	// instructions
	io.WriteString(conn, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")

	// read & write
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		// logic
		if len(fs) < 1 {
			continue
		}
		switch fs[0] {
		case "GET":
			key := fs[1]
			v := server.db.Get(key)
			fmt.Fprintf(conn, "%s\r\n", v)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "EXPECTED VALUE\r\n")
				continue
			}
			key := fs[1]
			value := fs[2]
			server.db.Set(key, value)
		case "DEL":
			key := fs[1]
			server.db.Del(key)
		default:
			fmt.Fprintln(conn, "INVALID COMMAND "+fs[0]+"\r\n")
			continue
		}
	}
}
