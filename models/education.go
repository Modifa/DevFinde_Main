package models

type Education struct {
	ID                  int64  `db:"developer_id"`
	Intsitution         string `db:"intsitution_"`
	Start_date          string `db:"start_date_"`
	EndDate             string `db:"end_date_"`
	Qualification_name  string `db:"qualification_name_"`
	Qualification_type_ int64  `db:"qualification_type"`
}

type EducationRequest struct {
	ID                  int64  `json:"developer_id"`
	UserName            string `json:"username"`
	Intsitution         string `json:"intsitution_"`
	Start_date          string `json:"start_date_"`
	EndDate             string `json:"end_date_"`
	Qualification_name  string `json:"qualification_name_"`
	Qualification_type_ string `json:"qualification_type"`
}
