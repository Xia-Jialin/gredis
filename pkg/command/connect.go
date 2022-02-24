package command

import "github.com/Xia-Jialin/gredis/pkg/gredis"

func quit(c gredis.Client) {
	c.WriteString("OK")
	c.Conn.Close()
}

func ping(c gredis.Client) {
	c.WriteString("PONG")
	c.Conn.Close()
}
