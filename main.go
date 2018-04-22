package main

import (
	"fmt"
	"log"

	"github.com/thingful/zenroom-go"
)

func main() {
	script := `
	octet = require'octet'
	ecdh = require 'ecdh'
	msg = octet.new(#DATA)
	msg:string(DATA)

	ed25519 =ecdh.new('ec25519')
	pk, sk = ed25519:keygen()

	sess = ed25519:session(pk, sk)

	zmsg = ed25519:encrypt(sess, msg)

	decipher = ed25519:decrypt(sess, zmsg)
	print(decipher:string())
	`
	res, err := zenroom.Exec(script, "", "")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
