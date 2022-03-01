package command

import "github.com/Xia-Jialin/gredis/pkg/gredis"

func del(c gredis.Client) {
	if len(c.Args) != 2 {
		c.WriteError(newWrongNumOfArgsError(string(c.Args[0])).Error())
		return
	}
	err := c.Remove(c.Args[1])
	if err != nil {
		c.WriteInt(0)
	} else {
		c.WriteInt(1)
	}
}

func ttl(c gredis.Client) {
	if len(c.Args) != 2 {
		c.WriteError(newWrongNumOfArgsError(string(c.Args[0])).Error())
		return
	}
	c.WriteInt64(c.TTL(c.Args[1]))
}
