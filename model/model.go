package model

type Artist struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
}

// type Concert struct {
// 	ID       int    `json:"id"`
// 	Date     string `json:"date"`
// 	ArtistID int    `json:"ArtistID"`
// }

// type Venue struct {
// 	ID      int    `json:"id"`
// 	City    string `json:"city"`
// 	Country string `json:"country"`
// }
