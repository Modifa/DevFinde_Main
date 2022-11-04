package models

type Education struct {
	ID                  int64  `json:"developer_id"`
	Intsitution         string `json:"intsitution_"`
	Start_date          string `json:"start_date_"`
	EndDate             string `json:"end_date_"`
	Qualification_name  string `json:"qualification_name_"`
	Qualification_type_ int64  `json:"qualification_type"`
}

type EducationRequest struct {
	ID                  int64  `json:"developer_id"`
	UserName            string `json:"username"`
	Intsitution         string `json:"intsitution_"`
	Start_date          string `json:"start_date_"`
	EndDate             string `json:"end_date_"`
	Qualification_name  string `json:"qualification_name_"`
	Qualification_type_ int64  `json:"qualification_type"`
}
