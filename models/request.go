package models

type DeveloperRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DeveloperResponseDB struct {
	Id int64 `db:"user_id"`
}

type DBIDResponse struct {
	ID int64 `db:"user_id"`
}

type DeveloperProfile struct {
	Id            int64  `db:"userid,omitempty"`
	FirstName     string `db:"firstname,omitempty"`
	LastName      string `db:"lastname,omitempty"`
	UserName      string `db:"username,omitempty"`
	EmailAddress  string `db:"emailaddress,omitempty"`
	Mobile_number string `db:"mobile_number,omitempty"`
	Active        string `db:"active,omitempty"`
	DateAdded     string `db:"dateadded,omitempty"`
	Title         string `db:"title,omitempty"`
	Password      string `db:"userpassword,omitempty"`
	Image         string `db:"image,omitempty"`
}

type DeveloperRegister struct {
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Email_address string `json:"email_address"`
	Mobile_number string `json:"mobile_number"`
	Title         string `json:"title"`
	Password      string `json:"password"`
}

type GetDeveloperProfile struct {
	EmailAddress string `json:"dev_email_address"`
}

type UpdateProfile struct {
	ID         int64  `json:"developer_id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email_address"`
	Mobile     string `json:"mobile_number"`
	Title      string `json:"title"`
}

type UpdateImage struct {
	ID       int64  `json:"developer_id"`
	ImageURL string `json:"imageUrl"`
}

type UpdateImageDB struct {
	ID       int64  `json:"developer_id"`
	ImageURL string `json:"imageUrl"`
	Email    string `json:"email_address"`
}
