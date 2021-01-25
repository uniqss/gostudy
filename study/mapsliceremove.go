package main

import (
	"log"
	"strconv"
)

func main() {

	var userAgentMap = make(map[string][]string)
	for i := 0; i < 10; i++ {
		is := "agent" + strconv.Itoa(i)
		var value []string
		for c := 1; c < 4; c++ {
			cs := "uid" + strconv.Itoa(int(c))
			value = append(value, cs)
		}
		userAgentMap[is] = value
	}

	log.Println(userAgentMap)

	oldAgentID := "agent" + "3"
	userID := "uid2"

	userList := userAgentMap[oldAgentID]
	for i, id := range userList {
		if id == userID {
			userList = append(userList[:i], userList[i+1:]...)
			break
		}
	}

	userAgentMap[oldAgentID] = userList

	log.Println(userAgentMap)
}
