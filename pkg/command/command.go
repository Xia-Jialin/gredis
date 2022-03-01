package command

import (
	"fmt"

	"github.com/Xia-Jialin/gredis/pkg/gredis"
)

var CommandTable map[string]func(c gredis.Client)

func init() {
	CommandTable = make(map[string]func(c gredis.Client))
	// Key
	CommandTable["del"] = del
	CommandTable["ttl"] = ttl
	// sting
	CommandTable["set"] = set
	CommandTable["setnx"] = setnx
	CommandTable["setex"] = setex
	CommandTable["get"] = get
	CommandTable["getset"] = getSet

	// connect
	CommandTable["quit"] = quit
	CommandTable["ping"] = ping

	// PubSub
	CommandTable["subscribe"] = subscribe
	CommandTable["psubscribe"] = psubscribe
	CommandTable["publish"] = publish
}

func newWrongNumOfArgsError(cmd string) error {
	return fmt.Errorf("wrong number of arguments for '%s' command", cmd)
}
