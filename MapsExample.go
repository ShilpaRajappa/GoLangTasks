package main

import (
	"fmt"
	"sort"
)

func main() {

	nodeMap := map[string][]string{
		"SR":{"7750","7450","7950"},
		"SAR":{"7705","7250"},
	}
fmt.Println(nodeMap)
//append

nodeMap["MG"] = append(nodeMap["MG"], "7705")

fmt.Println("after append ",nodeMap)

//find

srStrArray := nodeMap["SR"]
fmt.Println("find in SR array",srStrArray)
res1 := sort.SearchStrings(srStrArray,"7250")
res2 := sort.SearchStrings(srStrArray,"7750")
fmt.Println("Find Result 7250:",res1)
fmt.Println("Find Result 7750:",res2)

//delete
delete(nodeMap,"MG")

fmt.Println("After delete MG",nodeMap)

//Use case 2



type EventDay struct {
	day int
	month int
	year int
}

type OrganisingTeam struct {

	teamMemebers  []string
	primaryContact string
}

type NSPEvent struct {
	eventName string 
	primaryContact string
	eventDay EventDay
	organisingTeam OrganisingTeam
}

nspEvent := NSPEvent { eventName:"Event1", primaryContact: "Contact1",
eventDay: EventDay{24,1,2024}, 
organisingTeam: OrganisingTeam{[]string{"Mem1","Mem2"},"Mem1"}}

fmt.Println("struct : ",nspEvent)
}