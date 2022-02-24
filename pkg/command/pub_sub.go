package command

import (
	"strings"

	"github.com/Xia-Jialin/gredis/pkg/gredis"
)

func subscribe(c gredis.Client) {
	if len(c.Args) < 2 {
		c.WriteError(newWrongNumOfArgsError(string(c.Args[0])).Error())
		return
	}

	command := strings.ToLower(string(c.Args[0]))
	for i := 1; i < len(c.Args); i++ {
		if command == "subscribe" {
			c.Subscribe(c.Conn, string(c.Args[i]))
		}
	}
}

func psubscribe(c gredis.Client) {
	if len(c.Args) < 2 {
		c.WriteError(newWrongNumOfArgsError(string(c.Args[0])).Error())
		return
	}

	command := strings.ToLower(string(c.Args[0]))
	for i := 1; i < len(c.Args); i++ {
		if command == "psubscribe" {
			c.Psubscribe(c.Conn, string(c.Args[i]))
		}
	}
}

func publish(c gredis.Client) {
	if len(c.Args) != 3 {
		c.WriteError(newWrongNumOfArgsError(string(c.Args[0])).Error())
		return
	}
	c.WriteInt(c.Publish(string(c.Args[1]), string(c.Args[2])))
}
