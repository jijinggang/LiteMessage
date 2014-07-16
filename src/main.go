package main

import (
//"fmt"
)

func main() {
	//fmt.Println("ok.")
	msgs, _ := Parse("../testdata/msg.lm")
	//	fmt.Println(msgs)
	Gen("../bin/cs", msgs, "")
}

type Message struct {
	Name   string
	Msgid  int
	Fields []MessageField
}

type MessageField struct {
	Name string
	Type string
	//Default string
}
