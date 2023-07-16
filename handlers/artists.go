package ConcertAPI

import (
	"encoding/json"
	"log"
	"net/http"
)

// function that being called when page is reloaded, or search result is clicked
func GetBands(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var dataArr []Data
		for i := 0; i < 52; i++ {
			dataArr = append(dataArr, getData(i))
		}
		b, err1 := json.Marshal(dataArr)
		if err1 != nil {
			log.Println("Error during json marshlling. Error:", err1)
		}
		w.Write(b)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("This function does not support " + r.Method + " method."))
	}
}

func getData(pers int) Data {
	return Data{
		BandId:     inputs.Artists[pers].ID,
		Image:         inputs.Artists[pers].Image,
		Name:          inputs.Artists[pers].Name,
		Members:       inputs.Artists[pers].Members,
		CreationDate:  inputs.Artists[pers].CreationDate,
		FirstAlbum:    inputs.Artists[pers].FirstAlbum,
		LocationsLink: inputs.Artists[pers].Locations,
		ConcertDates:  inputs.Artists[pers].ConcertDates,
		Relations:     inputs.Artists[pers].Relations,
		Locations:      inputs.Locations.Index[pers].Locations,
		LocationsDates: inputs.Locations.Index[pers].Dates,
		Dates:          inputs.Dates.Index[pers].Dates,
		RelationStruct: inputs.Relation.Index[pers].DatesLocations,
		JSONLen: len(inputs.Artists),
	}
}
