package server

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) HandleWs(ws *websocket.Conn) {
	log.Println("New incoming connection from client:", ws.RemoteAddr())

	// TODO: Change to mutex
	s.conns[ws] = true

	s.ReadLoop(ws)
}

func (s *Server) Broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				log.Println("error broadcasting", err)
			}
		}(ws)
	}

}

func (s *Server) Disconnect(ws *websocket.Conn) {
	ws.Close()
	delete(s.conns, ws)
}

func (s *Server) ReadLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				s.Disconnect(ws)
				break
			}
			log.Println("read err", err)
			continue
		}
		msg := buf[:n]
		fmt.Println(string(msg))

		s.Broadcast(msg)
	}
}
