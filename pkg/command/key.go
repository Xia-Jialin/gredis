package command

import (
	"strconv"

	"github.com/Xia-Jialin/gredis/pkg/gredis"
	"github.com/roseduan/rosedb"
)

func del(c gredis.Client) {
	if len(c.Args) != 2 {
		c.WriteError(newWrongNumOfArgsError(string(c.Args[0])).Error())
		return
	}

	var err error
	switch getKeyType(c) {
	case rosedb.String:
		err = c.Remove(c.Args[1])
	case rosedb.List:
		err = c.LClear(c.Args[1])
	case rosedb.Hash:
		err = c.HClear(c.Args[1])
	case rosedb.Set:
		err = c.SClear(c.Args[1])
	case rosedb.ZSet:
		err = c.ZClear(c.Args[1])
	default:
		c.WriteInt(0)
		return
	}

	if err != nil {
		c.WriteInt(0)
	} else {
		c.WriteInt(1)
	}
}

func exists(c gredis.Client) {
	if len(c.Args) != 2 {
		c.WriteError(newWrongNumOfArgsError(string(c.Args[0])).Error())
		return
	}
	if getKeyType(c) > rosedb.ZSet {
		c.WriteInt(0)
		return
	}
	c.WriteInt(1)
}

func expire(c gredis.Client) {
	if len(c.Args) != 3 {
		c.WriteError(newWrongNumOfArgsError(string(c.Args[0])).Error())
		return
	}
	key := c.Args[1]
	duration, err := strconv.Atoi(string(c.Args[2]))
	if err != nil {
		c.WriteError("ERR value is not an integer or out of range")
		return
	}

	switch getKeyType(c) {
	case rosedb.String:
		c.Expire(key, int64(duration))
	case rosedb.List:
		c.LExpire(key, int64(duration))
	case rosedb.Hash:
		c.HExpire(key, int64(duration))
	case rosedb.Set:
		c.SExpire(key, int64(duration))
	case rosedb.ZSet:
		c.ZExpire(key, int64(duration))
	default:
		c.WriteInt64(0)
		return
	}

	c.WriteInt(1)
}

func ttl(c gredis.Client) {
	if len(c.Args) != 2 {
		c.WriteError(newWrongNumOfArgsError(string(c.Args[0])).Error())
		return
	}

	var liveTime int64
	switch getKeyType(c) {
	case rosedb.String:
		liveTime = c.TTL(c.Args[1])
	case rosedb.List:
		liveTime = c.LTTL(c.Args[1])
	case rosedb.Hash:
		liveTime = c.HTTL(c.Args[1])
	case rosedb.Set:
		liveTime = c.STTL(c.Args[1])
	case rosedb.ZSet:
		liveTime = c.ZTTL(c.Args[1])
	default:
		c.WriteInt64(-2)
		return
	}
	c.WriteInt64(liveTime)
}

func pttl(c gredis.Client) {
	if len(c.Args) != 2 {
		c.WriteError(newWrongNumOfArgsError(string(c.Args[0])).Error())
		return
	}

	var liveTime int64
	switch getKeyType(c) {
	case rosedb.String:
		liveTime = c.TTL(c.Args[1])
	case rosedb.List:
		liveTime = c.LTTL(c.Args[1])
	case rosedb.Hash:
		liveTime = c.HTTL(c.Args[1])
	case rosedb.Set:
		liveTime = c.STTL(c.Args[1])
	case rosedb.ZSet:
		liveTime = c.ZTTL(c.Args[1])
	default:
		c.WriteInt64(-2)
		return
	}
	c.WriteInt64(liveTime * 1000)
}

func getKeyType(c gredis.Client) rosedb.DataType {
	if c.StrExists(c.Args[1]) {
		return rosedb.String
	}
	if c.HKeyExists(c.Args[1]) {
		return rosedb.Hash
	}
	if c.LKeyExists(c.Args[1]) {
		return rosedb.List
	}
	if c.SKeyExists(c.Args[1]) {
		return rosedb.Set
	}
	if c.ZKeyExists(c.Args[1]) {
		return rosedb.ZSet
	}
	return rosedb.ZSet + 1
}
