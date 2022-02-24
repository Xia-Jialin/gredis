package main

import (
	"log"
	"strings"

	"github.com/Xia-Jialin/gredis/pkg/command"
	"github.com/Xia-Jialin/gredis/pkg/gredis"
	"github.com/roseduan/rosedb"
	"github.com/tidwall/redcon"
)

var addr = ":6379"

func main() {
	config := rosedb.DefaultConfig()
	db, err := rosedb.Open(config)
	if err != nil {
		log.Fatal(err)
	}

	var ps redcon.PubSub
	go log.Printf("started server at %s", addr)
	err = redcon.ListenAndServe(addr,
		func(conn redcon.Conn, cmd redcon.Command) {
			if len(cmd.Args) < 1 {
				conn.WriteError("ERR unknown command '" + string(cmd.Args[0]) + "'")
				return
			}
			comd, ok := command.CommandTable[strings.ToLower(string(cmd.Args[0]))]
			if !ok {
				conn.WriteError("ERR unknown command '" + string(cmd.Args[0]) + "'")
				return
			}
			comd(gredis.Client{Conn: conn, Command: cmd, RoseDB: db, PubSub: &ps})
		},
		func(conn redcon.Conn) bool {
			// Use this function to accept or deny the connection.
			// log.Printf("accept: %s", conn.RemoteAddr())
			return true
		},
		func(conn redcon.Conn, err error) {
			// This is called when the connection has been closed
			// log.Printf("closed: %s, err: %v", conn.RemoteAddr(), err)
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
