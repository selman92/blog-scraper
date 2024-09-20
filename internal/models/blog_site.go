package models

type BlogSite struct {
	ID            int    `json:"id"`
	URL           string `json:"url"`
	TitleSelector string `json:"title_selector"`
	TimeSelector  string `json:"time_selector"`
}
