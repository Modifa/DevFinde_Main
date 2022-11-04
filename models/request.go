package models

type DBIDResponse struct {
	ID int64 `db:"user_id"`
}

type DBIDRequest struct {
	ID int64 `json:"user_id"`
}
