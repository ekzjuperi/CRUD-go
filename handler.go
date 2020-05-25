package main

import (
	"strconv"
	"strings"

	"kvdb/kochetkov/consts"
)

// Handler function command string
func Handler(text string) (command string, key string, value string, valueInt int) {
	command, key, value = "", "", ""
	valueInt = 0
	arr := strings.Split(text, " ")
	switch len(arr) {
	case 1:
		command = arr[0]
	case 2:
		command = arr[0]
		key = arr[1]
	case 3:
		command = arr[0]
		key = arr[1]
		value = arr[2]
	}
	if len(arr) == 2 && (command == consts.Eq || command == consts.Gt || command == consts.Lt) {
		valueInt, _ = strconv.Atoi(arr[1])
		return command, key, "", valueInt
	}
	if len(arr) == 3 && (command == consts.Createint || command == consts.Updateint || command == consts.Lt) {
		valueInt, _ = strconv.Atoi(arr[2])
		return command, key, "", valueInt
	}
	return command, key, value, 0
}
