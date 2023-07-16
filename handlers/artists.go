package ConcertAPI

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// function that being called when page is reloaded, or search result is clicked
func GetArtists(w http.ResponseWriter, r *http.Request) {
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
	myDate, err := time.Parse("02-01-2006 15:04", cache.Artists[pers].FirstAlbum+" 04:35")
	if err != nil {
		log.Println("Error during time formatting. Error:", err)
	}
	return Data{
		BandId:     cache.Artists[pers].ID,
		Image:         cache.Artists[pers].Image,
		Name:          cache.Artists[pers].Name,
		Members:       cache.Artists[pers].Members,
		CreationDate:  cache.Artists[pers].CreationDate,
		FirstAlbum:    myDate.Format("02/01/2006"),
		LocationsLink: cache.Artists[pers].Locations,
		ConcertDates:  cache.Artists[pers].ConcertDates,
		Relations:     cache.Artists[pers].Relations,

		Locations:      cache.Locations.Index[pers].Locations,
		LocationsDates: cache.Locations.Index[pers].Dates,

		Dates:          cache.Dates.Index[pers].Dates,
		RelationStruct: cache.Relation.Index[pers].DatesLocations,

		JSONLen: len(cache.Artists),
	}
}
