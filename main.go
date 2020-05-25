package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"kvdb/kochetkov/consts"
	"kvdb/kochetkov/database"
)

func main() {
	// load db map[string]string
	base := database.NewDatabase()
	if err := base.Load(); err != nil {
		fmt.Println(err)
	}
	// load db map[string]int
	baseInt := database.NewDatabaseInt()
	if err := baseInt.Load(); err != nil {
		fmt.Println(err)
	}
	// Command line handler
	fmt.Println(consts.Text)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		command, key, value, valueInt := Handler(text)
		switch command {
		case consts.Create:
			fmt.Println(base.Create(key, value))
		case consts.Read:
			fmt.Println(base.Read(key))
		case consts.Update:
			fmt.Println(base.Update(key, value))
		case consts.Delete:
			fmt.Println(base.Delete(key))
		case consts.Exist:
			fmt.Println(base.IsExist(key))
		case consts.Quit:
			if err := base.Save(); err != nil {
				fmt.Println(err)
			}
			if err := baseInt.Save(); err != nil {
				fmt.Println(err)
			}
			return
		case consts.Createint:
			fmt.Println(baseInt.Create(key, valueInt))
		case consts.Updateint:
			fmt.Println(baseInt.Update(key, valueInt))
		case consts.Readint:
			fmt.Println(baseInt.Read(key))
		case consts.Deleteint:
			fmt.Println(baseInt.Delete(key))
		case consts.Existint:
			fmt.Println(baseInt.IsExist(key))
		case consts.Sum:
			fmt.Println(baseInt.Sum())
		case consts.Avg:
			fmt.Println(baseInt.Avg())
		case consts.Gt:
			for _, v := range strings.Split(baseInt.GtVal(valueInt), " ") {
				fmt.Println(v)
			}
		case consts.Lt:
			for _, v := range strings.Split(baseInt.LtVal(valueInt), " ") {
				fmt.Println(v)
			}
		case consts.Eq:
			for _, v := range strings.Split(baseInt.EqVal(valueInt), " ") {
				fmt.Println(v)
			}
		case consts.Count:
			fmt.Println(baseInt.Count())
		case consts.Med:
			fmt.Println(baseInt.Med())
		default:
			fmt.Println(consts.Invalid)
		}
	}
}
