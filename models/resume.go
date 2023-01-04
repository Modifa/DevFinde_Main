package models

//Link
type AddResumeLink struct {
	ID        int64  `json:"developer_id"`
	ResumeURL string `json:"resume_link"`
}

type AddResumeLinkDB struct {
	ID       int64  `json:"_developer_id"`
	LinkType string `json:"_link_type"`
	Link     string `json:"Link_"`
	Username string `json:"username"`
}
type AddResumeLinkDBPOST struct {
	ID       int64  `json:"_developer_id"`
	LinkType int64  `json:"_link_type"`
	Link     string `json:"Link_"`
}
type AddResumeDescDB struct {
	ID        int64  `json:"developer_id"`
	ResumeURL string `json:"resume_link"`
	Email     string `json:"email_address"`
}

type UpdateResumeLinkDB struct {
	ID       int64  `json:"_developer_id"`
	LinkId   int64  `json:"link_id"`
	Link     string `json:"Link_"`
	Username string `db:"username"`
}
type UpdateResumeLinkPOST struct {
	LinkId int64  `json:"link_id"`
	ID     int64  `json:"_developer_id"`
	Link   string `json:"Link_"`
}
type DeleteResumeLink struct {
	ID       int64  `json:"_developer_id"`
	LinkId   int64  `json:"link_id"`
	Username string `db:"username"`
}

type DeleteResumeLinkPOST struct {
	LinkId int64 `json:"link_id"`
	ID     int64 `json:"_developer_id"`
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

type ResumedescReq struct {
	ID      int64  `json:"_developer_id"`
	AboutMe string `json:"short_desc"`
}

//
type ResumedescRes struct {
	ResID       int64  `db:"id"`
	Description string `db:"description"`
	Username    string `db:"username"`
	Dateadded   string `db:"dateadded"`
}

type ResumedescRedis struct {
	ResID       int64  `db:"id"`
	Description string `db:"description"`
	Dateadded   string `db:"dateadded"`
}

type ResumedescRedisUP struct {
	Username     string `db:"username"`
	Developer_ID int64  `json:"_developer_id"`
	Description  string `json:"short_desc"`
}

//
type ResumeDesc struct {
	Developer_ID int64  `json:"_developer_id"`
	Description  string `json:"short_desc"`
}
