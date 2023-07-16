package ConcertAPI

import (
	"sync"
)

var apiLink = "https://groupietrackers.herokuapp.com/api"
var inputs Inputs

func Parse() {
	//parse api and save everthing into the struct
	var wg sync.WaitGroup

	SendRequest(apiLink, &inputs)
	wg.Add(1)
	go func() {
		go SendRequest(inputs.BandsUrl, &inputs.Artists)
		go SendRequest(inputs.LocUrl, &inputs.Locations)
		go SendRequest(inputs.DatesUrl, &inputs.Dates)
		go SendRequest(inputs.RelationUrl, &inputs.Relation)
		wg.Done()
	}()
	wg.Wait()
}
