package main

import (
	"log"
	"net"
)

type server struct {
	rooms    map[string]*room
	commands chan commands
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan commands),
	}
}
func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listRooms(cmd.client, cmd.args)
		default:
		}
	}
}
func (s *server) newClient(conn net.Conn) {
	log.Printf("new client connection %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		nick:     "anon",
		commands: s.commands,
	}
	c.readInput()
}

func (s *server) nick(c *client, args []string) {

}

func (s *server) join(c *client, args []string) {

}

func (s *server) msg(c *client, args []string) {

}

func (s *server) listRooms(c *client, args []string) {

}

func (s *server) quit(c *client, args []string) {

}
