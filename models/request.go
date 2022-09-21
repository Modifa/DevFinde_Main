package models

type DeveloperRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DeveloperResponseDB struct {
	Id int64 `db:"user_id"`
}

type DeveloperProfile struct {
	Id            int64  `db:"userid"`
	FirstName     string `db:"firstname"`
	LastName      string `db:"lastname"`
	UserName      string `db:"username"`
	EmailAddress  string `db:"emailaddress"`
	Mobile_number string `json:"mobile_number"`
	Active        string `db:"active"`
	DateAdded     string `db:"dateadded"`
	Password      string `db:"userpassword"`
}

type DeveloperRegister struct {
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Email_address string `json:"email_address"`
	Mobile_number string `json:"mobile_number"`
	Password      string `json:"password"`
}

type GetDeveloperProfile struct {
	EmailAddress string `json:"dev_email_address"`
}
