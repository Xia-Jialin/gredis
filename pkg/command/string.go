package command

import (
	"log"
	"strconv"

	"github.com/Xia-Jialin/gredis/pkg/gredis"
)

func set(c gredis.Client) {
	if len(c.Args) != 3 {
		c.WriteError(newWrongNumOfArgsError(string(c.Args[0])).Error())
		return
	}
	err := c.Set(c.Args[1], c.Args[2])
	if err != nil {
		c.WriteError(err.Error())
		return
	}
	c.WriteString("OK")
}

func setnx(c gredis.Client) {
	if len(c.Args) != 3 {
		c.WriteError(newWrongNumOfArgsError(string(c.Args[0])).Error())
		return
	}

	ok, err := c.SetNx(c.Args[1], c.Args[2])
	if err != nil {
		log.Println(err.Error())
		c.WriteInt(0)
		return
	}
	if !ok {
		c.WriteInt(0)
		return
	}
	c.WriteInt(1)
}

func setex(c gredis.Client) {
	if len(c.Args) != 4 {
		c.WriteError(newWrongNumOfArgsError(string(c.Args[0])).Error())
		return
	}
	duration, err := strconv.Atoi(string(c.Args[2]))
	if err != nil {
		c.WriteError("ERR value is not an integer or out of range")
		return
	}
	err = c.SetEx(c.Args[1], c.Args[3], int64(duration))
	if err != nil {
		log.Println(err.Error())
		return
	}
	c.WriteString("OK")
}

func get(c gredis.Client) {
	if len(c.Args) != 2 {
		c.WriteError(newWrongNumOfArgsError(string(c.Args[0])).Error())
		return
	}

	val := make([]byte, 1)
	err := c.Get(c.Args[1], &val)
	if err != nil {
		c.WriteNull()
		log.Println(err)
		return
	}
	c.WriteBulk(val)
}

func getSet(c gredis.Client) {
	if len(c.Args) != 3 {
		c.WriteError(newWrongNumOfArgsError(string(c.Args[0])).Error())
		return
	}

	val := make([]byte, 1)
	err := c.GetSet(c.Args[1], c.Args[2], &val)
	if err != nil {
		log.Println(err.Error())
		c.WriteNull()
		return
	}
	c.WriteBulk(val)
}

func append_str(c gredis.Client) {
	if len(c.Args) != 3 {
		c.WriteError(newWrongNumOfArgsError(string(c.Args[0])).Error())
		return
	}

	err := c.Append(c.Args[1], string(c.Args[2]))
	if err != nil {
		c.WriteInt(0)
		log.Println(err.Error())
		return
	}
	val := make([]byte, 1)
	err = c.Get(c.Args[1], &val)
	if err != nil {
		c.WriteInt(0)
		log.Println(err)
		return
	}
	c.WriteInt(len(val))
}
