package ConcertAPI

type Inputs struct {
	BandsUrl   string `json:"artists"`
	LocUrl string `json:"locations"`
	DatesUrl     string `json:"dates"`
	RelationUrl  string `json:"relation"`

	Artists   []Artist
	Locations Locations
	Dates     Dates
	Relation  Relation
}

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type Dates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}
type Relation struct {
	Index []struct {
		ID             int `json:"id"`
		DatesLocations map[string][]string
	} `json:"index"`
}

type Data struct {
	BandId     	  int
	Image         string
	Name          string
	Members       []string
	CreationDate  int
	FirstAlbum    string
	LocationsLink string
	ConcertDates  string
	Relations     string
	Locations      []string
	LocationsDates string
	RelationStruct map[string][]string
	Dates []string
	ErrorCode int
	Error     string
	FoundBy []string
	JSONLen int
}

type Geodata struct {
	Index []struct {
		CountryCoords map[string][]string
	} `json:"index"`
}
