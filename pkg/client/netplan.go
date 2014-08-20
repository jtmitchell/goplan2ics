package client

import (
	"net"
	"fmt"
	"os"
	"os/user"
)

type Client struct {
	Name string
	Conn net.Conn
}

func NewClient(address string) *Client {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c := &Client{clientId(), conn}
	c.Send("=" + c.Name)
	c.Send("t0")
	
	return c
}

func (c *Client) Send(data string) {
	fmt.Fprintf(c.Conn,"%s\r\n", data)
}

func clientId() (string) {
	uid := "1"
	gid := "1"
	pid := os.Getpid()

	user, err := user.Current()
	if err == nil {
		uid = user.Uid
		gid = user.Gid
	}
	
	return fmt.Sprintf("%s<uid=%s,gid=%s,pid=%d>","goplan2ics", uid, gid, pid)
}	