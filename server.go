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
c.nick = args [1];
c.msg(fmt.Sprintf("hello,%s", c.nicks))
}

func (s *server) join(c *client, args []string) {
	roomName := args[1]
	r,ok := s.rooms[roomName]
	if !ok{
		r = &room{
			name: roomName,
			members: make (map[net.Addr]*client),
		}
		s.rooms [roomName] = r;
	}
	r.members[c.conn.RemoteAddr()] = c;
	s.quitThisRoom(c)
	c.room = r 
	r.broadcast (c,fmt.Sprintf (c.nick "%s has joined us"))
	c.msg (fmt.Sprintf ("welcome to %s", r.name))
}

func (s *server) msg(c *client, args []string) {
	if c.room == nil {
		c.err(error.new ("need to join the room firstly!"))
		return
	}
	c.room.broadcast(c, c.nick+": "strings.join(args[1:len(args)], " "))
}

func (s *server) listRooms(c *client, args []string) {
	var rooms []string
	for name := range s.rooms{
		rooms = append (rooms,name)
	}
	c.msg(fmt.Sprintf("rooms are available to join:%s",strings.join(rooms,",")))
}

func (s *server) quit(c *client, args []string) {
	log.Printf("client has disconnected:%s", c.conn.RemoteAddr().String())
	s.quitThisRoom(c)
	c.msg ("bye")
	c.conn.Close()
}

func (s *server) quitThisRoom (c *client){
	if c.room != nil{
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.broadcast (c,fmt.Sprintf(c.nick, "succesfully leave this room"))
	}
}
