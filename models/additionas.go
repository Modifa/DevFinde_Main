package models

type LinksRequest struct {
	Id       int64  `json:"_developer_id"`
	Link     string `json:"_link"`
	LinkType int64  `json:"_link_type"`
}

type LinksRequestDB struct {
	Id       int64  `json:"_developer_id"`
	Link     string `json:"_link"`
	LinkType string `json:"_link_type"`
	Username string `json:"username"`
}

type LinksRequestReponse struct {
	Id          int64  `db:"id"`
	LinkType    string `db:"link_type"`
	LinkTypeId  int64  `db:"link_type_id"`
	Description string `db:"description"`
	Link        string `db:"link"`
}

type ExperienceRequest struct {
	Id          int64  `json:"_developer_id"`
	Description string `json:"description"`
	Title       string `json:"title_name"`
	Company     string `json:"company"`
	Start_date  string `json:"start_date"`
	End_date    string `json:"end_date"`
	Username    string `json:"username"`
}

type ExperienceRequestDB struct {
	Id          int64  `json:"_developer_id"`
	Description string `json:"description"`
	Title       string `json:"title_name"`
	Company     string `json:"company"`
	Start_date  string `json:"start_date"`
	End_date    string `json:"end_date"`
	Username    string `json:"username"`
}

type ExperienceResponseDB struct {
	Id          int64  `db:"id"`
	Description string `db:"description"`
	Title       string `db:"title"`
	Company     string `db:"company"`
	Start_Date  string `db:"start_date"`
	End_Date    string `db:"end_date"`
}

type IDRequest struct {
	Id int64 `json:"_developer_id"`
}
