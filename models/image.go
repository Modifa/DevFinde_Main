package models

type UpdateImage struct {
	ID       int64  `json:"developer_id"`
	ImageURL string `json:"imageUrl"`
}

type UpdateImageDB struct {
	ID       int64  `json:"developer_id"`
	ImageURL string `json:"imageUrl"`
	Email    string `json:"email_address"`
}
