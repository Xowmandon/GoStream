package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"encoding/json"

	"github.com/Xowmandon/GoStream/TRN"
	"github.com/Xowmandon/GoStream/TRN/ApexPlayer"
	//"golang.org/x/tools/go/analysis/passes/nilfunc"
	//"github.com/Xowmandon/GoStream/Twitch"
	//"fmt"
)


func main() {

	userName := "harrumaki"

	userURL := "https://public-api.tracker.gg/v2/apex/standard/profile/origin/" + userName
	connection := TRN.NewTRNconnection(userURL)

	data, err := connection.Get(userURL)

	body, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Fatal(err)
	}

    ioutil.WriteFile( userName + ".json", body, 0600)

	var user ApexProfile.ApexProfile

	error := json.Unmarshal(body, &user)

	if error != nil {
		log.Fatal(error)
	}

    fmt.Println(user.Data.Segments[0].GlobalStats.GlobalSeasonData.Season8Wins)

	//fmt.Println(user)

    //fmt.Println( reflect.TypeOf(body) )
    fmt.Println("done")
}


