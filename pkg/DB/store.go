package DB

import (
	"encoding/base64"

	"github.com/Yscream/go-form-reg/pkg/encryption"
	"github.com/Yscream/go-form-reg/pkg/models"
	_ "github.com/lib/pq"
)

func SaveData(user *models.User) {
	db := GetConnection()

	defer db.Close()

	salt := encryption.GenerateRandomString([]byte(user.Password))
	saltToString := base64.StdEncoding.EncodeToString(salt)
	combination := saltToString + user.Password
	hash, _ := encryption.HashPassword(combination)

	db.Exec("INSERT INTO users_data (fname,lname,email)  VALUES($1, $2, $3)", user.Name, user.LastName, user.Email)
	db.Exec("INSERT INTO credentials (users_id, salt, hash)  VALUES($1, $2, $3)", GetId(user.Email), saltToString, hash)
}
