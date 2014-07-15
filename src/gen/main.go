package main

import (
	"fmt"
	//	"io"
	"io/ioutil"
	"regexp"
)

func main() {
	//fmt.Println("ok.")
	parse("../../testdata/msg.lm")
}

func parse(filePath string) {
	content := ReadStringFile(filePath)
	content = removeComment(content)
	parseMessages(content)
}
func ReadStringFile(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return ""
	}
	return string(content)
}

func removeComment(str string) string {
	re, _ := regexp.Compile(`([\w\s]*)(\/\/.*)`)
	strNew := re.ReplaceAllString(str, "$1")
	return strNew
}

func parseMessages(str string) []Message {
	const PATTERN_MSG = `\bmessage\b\s+([a-zA-Z][a-zA-Z0-9_]*)\s*(=\s*(\d+)\s*)?\{([\w\s\[\]\=\"]*)\}`
	re, _ := regexp.Compile(PATTERN_MSG)
	results := re.FindAllString(str, -1)
	for i, v := range results {
		fmt.Printf("%d\n%s", i, v)
	}
	return nil
}
func parseMessage() {

}

type Message struct {
	Name   string
	Msgid  int
	Fields []MessageField
}

type MessageField struct {
	Name    string
	Type    string
	Default string
}
