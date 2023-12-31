package ConcertAPI

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func FindBand(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

		var dataArr []Data
		var data Data

		var currIndex int
		dataArrIndexCounter := 0

		//convert everything to lower case to ease search algorithm
		searchingFor := strings.ToLower(r.FormValue("search"))

		for pers, art := range inputs.Artists {
			foundBy := ""
			//search for artists by the group name
			if strings.Contains(strings.ToLower(art.Name), searchingFor) {
				data = FetchData(pers)
				dataArr = append(dataArr, data)
				currIndex++
				foundBy += "group name"
				//search for creation dates
			} else if strings.Contains(strconv.Itoa(art.CreationDate), searchingFor) {
				if len(dataArr) >= 1 {
					if dataArr[currIndex-1].Name != art.Name {
						data = FetchData(pers)
						foundBy += "creation date"
						dataArr = append(dataArr, data)
						currIndex++
					} else {
						if !strings.Contains(foundBy, "creation date") {
							foundBy += ", creation date"
						}
					}
				} else {
					data = FetchData(pers)
					foundBy += "creation date"
					dataArr = append(dataArr, data)
					currIndex++
				}
			} else {
				myDate, _ := time.Parse("02-01-2006 15:04", art.FirstAlbum+" 04:35")
				if strings.Contains(myDate.Format("02/01/2006"), searchingFor) || strings.Contains(art.FirstAlbum, searchingFor) {
					if len(dataArr) >= 1 {
						if dataArr[currIndex-1].Name != art.Name {
							data = FetchData(pers)
							foundBy += "first album"
							dataArr = append(dataArr, data)
							currIndex++
						} else {
							if !strings.Contains(foundBy, "first album") {
								foundBy += ", first album"
							}
						}
					} else {
						data = FetchData(pers)
						foundBy += "by first album"
						dataArr = append(dataArr, data)
						currIndex++
					}
				}
			}
			//search for members
			for _, member := range art.Members {
				if strings.Contains(strings.ToLower(member), searchingFor) {
					if len(dataArr) >= 1 {
						if dataArr[currIndex-1].Name != art.Name {
							data = FetchData(pers)
							foundBy += "member name"
							dataArr = append(dataArr, data)
							currIndex++
						} else {
							if !strings.Contains(foundBy, "member name") {
								foundBy += ", member name"
							} else {
								break
							}
						}
					} else {
						data = FetchData(pers)
						foundBy += "member name"
						dataArr = append(dataArr, data)
						currIndex++
					}
				}
			}

			for _, location := range inputs.Locations.Index[art.ID-1].Locations {
				location = (strings.ToLower(location))
				locationDefault := location
				location = strings.Replace(location, "-", " ", -1)
				location = strings.Replace(location, "_", " ", -1)
				if strings.Contains(location, searchingFor) || strings.Contains(locationDefault, searchingFor) {
					if len(dataArr) >= 1 {
						if dataArr[currIndex-1].Name != art.Name {
							data = FetchData(pers)
							foundBy += "location"
							dataArr = append(dataArr, data)
							currIndex++
						} else {
							if !strings.Contains(foundBy, "location") {
								foundBy += ", location"
							} else {
								break
							}
						}
					} else {
						data = FetchData(pers)
						dataArr = append(dataArr, data)
						foundBy += "location"
						currIndex++
					}
				}
			}
			if foundBy != "" {
				data.FoundBy = append(data.FoundBy, foundBy)
				dataArr[dataArrIndexCounter].FoundBy = data.FoundBy
				dataArrIndexCounter++
			}
		}
		b, err := json.Marshal(dataArr)
		if err != nil {
			log.Println("Error during json marshlling. Error:", err)
		}

		w.Write(b)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("This function does not support " + r.Method + " method."))
	}
}
