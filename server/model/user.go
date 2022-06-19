package model

import "golang.org/x/crypto/bcrypt"

type UserInfo struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type TryLoginUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
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

func CheckNameAndPassword(tryLoginUser TryLoginUser, findUser *User) (err error) {
	// dbに接続
	db := DbConnect()
	// 名前から名前とパスワードを取得
	err = db.Where("name = ?", tryLoginUser.Name).First(findUser).Error
	if err != nil {
		return err
	}
	defer db.Close()
	// パスワードが一致するか確認
	err = bcrypt.CompareHashAndPassword(findUser.Password, []byte(tryLoginUser.Password))
	if err != nil {
		return err
	}
	return nil
}
