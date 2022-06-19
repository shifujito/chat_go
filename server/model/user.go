package model

type UserInfo struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func FindAllUser(users *[]UserInfo) {
	modelUsers := []User{}
	db := DbConnect()
	db.Find(&modelUsers)
	for _, val := range modelUsers {
		user := UserInfo{Id: val.ID, Name: val.Name}
		*users = append(*users, user)
	}
	defer db.Close()
}
