package models

type Fact struct {
	ID int `json:id`
	Title string `json:title`
	Categories []string `json:categories`
	Text string `json:text`

}