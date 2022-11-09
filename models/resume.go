package models

//Link
type AddResumeLink struct {
	ID        int64  `json:"developer_id"`
	ResumeURL string `json:"resume_link"`
}

type AddResumeLinkDB struct {
	ID        int64  `json:"developer_id"`
	ResumeURL string `json:"resume_link"`
	Email     string `json:"email_address"`
}

//Resume
type AddResume struct {
	ID        int64  `json:"developer_id"`
	AboutMe   string `json:"description"`
	ResumeURL string `json:"resume_link"`
}

type AddResumeDB struct {
	ID        int64  `json:"developer_id"`
	AboutMe   string `json:"description"`
	ResumeURL string `json:"resume_link"`
	Email     string `json:"email_address"`
}

type ResumeResponse struct {
	AboutMe   string `db:"description"`
	ResumeURL string `db:"resume_link"`
	Username  string `db:"username"`
}

type ResumeRequestRedis struct {
	Username string `json:"username"`
}
