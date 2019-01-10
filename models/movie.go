package models

type Movie struct {
	Id    int    `json:id`
	Title string `json:title`
	Genre string `json:genre`
	Year  string `json:year`
}
