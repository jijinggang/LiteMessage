package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
)

func Parse(filePath string) ([]Message, error) {
	content, err := ReadStringFile(filePath)
	if err != nil {
		return nil, err
	}
	content = removeComment(content)
	msgs, err := parseMessages(content)
	return msgs, err
}
func ReadStringFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func removeComment(str string) string {
	re := regexp.MustCompile(`([\w\s]*)(\/\/.*)`)
	strNew := re.ReplaceAllString(str, "$1")
	return strNew
}

func parseMessages(str string) (msgs []Message, err error) {
	const PATTERN_MSG = `\bmessage\b\s+([a-zA-Z][a-zA-Z0-9_]*)\s*(?:=\s*(\d+)\s*)?\{([\w\s\[\]\=\"]*)\}`
	re := regexp.MustCompile(PATTERN_MSG)
	results := re.FindAllStringSubmatch(str, -1)
	msgs = []Message{}
	for _, v := range results {
		msg := Message{}
		//fmt.Println(i, v[1], "|", v[2], "|", v[3])
		msg.Name = v[1]
		msg.Msgid, _ = strconv.Atoi(v[2])
		msg.Fields, err = parseMessageField(v[3])
		if err != nil {
			return
		}
		msgs = append(msgs, msg)
	}
	return
}
func parseMessageField(str string) ([]MessageField, error) {
	const PATTERN_FIELD = `\s*((?:\w+)(?:\[\])?)\s+(\w+)`
	re := regexp.MustCompile(PATTERN_FIELD)
	results := re.FindAllStringSubmatch(str, -1)
	fields := []MessageField{}
	for _, v := range results {
		field := MessageField{}
		//fmt.Println(i, v[1], "|", v[2])
		field.Type = v[1]
		field.Name = v[2]
		//field.Default = v[3]
		fields = append(fields, field)
	}
	return fields, nil
}
