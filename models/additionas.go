package models

type LinksRequest struct {
	Id          int64  `json:"_developer_id"`
	Link        string `json:"_link"`
	Description string `json:"_description"`
}

type LinksRequestDB struct {
	Id          int64  `json:"_developer_id"`
	Link        string `json:"_link"`
	Description string `json:"_description"`
	Username    string `json:"username"`
}

type LinksRequestReponse struct {
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
	Description string `db:"description"`
	Title       string `db:"title"`
	Company     string `db:"company"`
	Start_Date  string `db:"start_date"`
	End_Date    string `db:"end_date"`
}

type IDRequest struct {
	Id int64 `json:"_developer_id"`
}
