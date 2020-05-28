package users

import "encoding/json"

// Public User template
type Public struct {
	Id int64 `json:"id"`
	// FirstName   string `json:"firstname"`
	// LastName    string `json:"lastname"`
	// Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	// Password    string `json:"password"` //->> if you want to display value to json and vice versa
}

// Private User template
type Private struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"` //->> if you want to display value to json and vice versa
}

// Marshall Formating users result
func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

// Marshall Formating user result
func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return Public{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	userJSON, _ := json.Marshal(user)
	var privateUser Private
	json.Unmarshal(userJSON, &privateUser)
	return privateUser
}
