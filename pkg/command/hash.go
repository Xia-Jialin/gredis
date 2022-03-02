package command

import (
	"log"

	"github.com/Xia-Jialin/gredis/pkg/gredis"
)

func hset(c gredis.Client) {
	if len(c.Args) != 4 {
		c.WriteError(newWrongNumOfArgsError(string(c.Args[0])).Error())
		return
	}

	b := c.HExists(c.Args[1], c.Args[2])
	_, err := c.HSet(c.Args[1], c.Args[2], c.Args[3])
	if err != nil {
		log.Println(err.Error())
		c.WriteInt(0)
		return
	}
	//如果字段是哈希表中的一个新建字段，并且值设置成功，返回 1 。 如果哈希表中域字段已经存在且旧值已被新值覆盖，返回 0 。
	if b {
		c.WriteInt(0)
		return
	}
	c.WriteInt(1)
}

func hget(c gredis.Client) {
	if len(c.Args) != 3 {
		c.WriteError(newWrongNumOfArgsError(string(c.Args[0])).Error())
		return
	}
	val := c.HGet(c.Args[1], c.Args[2])
	if val == nil {
		c.WriteNull()
		return
	}
	c.WriteBulk(val)
}
