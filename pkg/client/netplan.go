package client

import (
	"net"
	"fmt"
	"os"
	"os/user"
)

type Client struct {
	conn net.Conn
}

func (c *Client) connect(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c.conn = conn
	
	c.send("=" + c.clientId())
	c.send("t0")
}

func (c *Client) send(data string) {
	fmt.Fprintf(c.conn,"%s\r\n", data)
}

func (c *Client) clientId() (string) {
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