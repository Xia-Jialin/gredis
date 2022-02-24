package gredis

import (
	"github.com/roseduan/rosedb"
	"github.com/tidwall/redcon"
)

type Client struct {
	redcon.Conn
	redcon.Command
	*rosedb.RoseDB
	*redcon.PubSub
}
