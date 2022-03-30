package postgresql

import (
	"testing"

	"github.com/Yscream/go-form-reg/pkg/models"
)

//tips for named test cases :
//func always gotta start with "Tesn" in the func name

// MethodName_StateUnderTest_ExpectedBehavior
// example: isAdult_AgeLessThan18_False

// MethodName_ExpectedBehavior_StateUnderTest
// example: isAdult_False_AgeLessThan18

// testFeatureBeingTested
// example: testIsNotAnAdultIfAgeLessThan18

// FeatureToBeTested
// example: IsNotAnAdultIfAgeLessThan18

// Should_ExpectedBehavior_When_StateUnderTest
// example: Should_ThrowException_When_AgeLessThan18

// When_StateUnderTest_Expect_ExpectedBehavior
// example: When_AgeLessThan18_Expect_isAdultAsFalse

// Given_Preconditions_When_StateUnderTest_Then_ExpectedBehavior — Behavior-Driven Development (BDD)
// example: Given_UserIsAuthenticated_When_InvalidAccountNumberIsUsedToWithdrawMoney_Then_TransactionsWillFail
var userTrue = models.User{Name: "Big", LastName: "Bob", Email: "bigbog123@gmail.com", Password: "bigbob123"}
var userFalse = models.User{Name: "", LastName: "", Email: "", Password: ""}
var credentialsTrue = models.Credentials{Salt: "7oGQ7CwmdjEXV7NU", Hash: "ztS5F8G5IfOH3mSu"}
var credentialsFalse = models.Credentials{Salt: "", Hash: ""}
var accessTokenTrue = models.AccessToken{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}
var accessTokenFalse = models.AccessToken{Token: ""}

// func TestXxx(t *testing.T) {
// 	db := getDB(t, testURL)
// 	name, lname, err := db.GetUser(userFalse.Email)
// }
func Test_InsertUser_CorrectFields_True(t *testing.T) {
	db := getDB(t, testURL)

	err := db.InsertUser(&models.User{
		Name: "Big", LastName: "Bob", Email: "bigbog123@gmail.com", Password: "bigbob123",
	})
	if err != nil {
		t.Error("cannot insert user", err.Error())
	}
}

func Test_InsertUser_IncorrectFields_False(t *testing.T) {
	db := getDB(t, testURL)

	// if len(userFalse.Name) == 0 || len(userFalse.LastName) == 0 || len(userFalse.Email) == 0 || len(userFalse.Password) == 0 {
	// 	t.Errorf("cannot insert user")
	// }
	err := db.InsertUser(&models.User{
		Name: "", LastName: "", Email: "", Password: "",
	})
	if err != nil {
		t.Error("cannot insert user", err.Error())
	}
}

// func Test_InsertPassword_CorrectField_True(t *testing.T) {
// 	db := getDB(t, testURL)

// 	err := db.InsertPassword(userTrue.ID, credentialsTrue.Salt, credentialsTrue.Hash)
// 	if err != nil {
// 		t.Error("cannot insertPassword", err.Error())
// 	}
// }

// func Test_InsertPassword_IncorrectField_False(t *testing.T) {
// 	db := getDB(t, testURL)

// 	if len(credentialsFalse.Salt) < 0 || len(credentialsFalse.Hash) < 0 {
// 		t.Errorf("cannot insert user")
// 	}
// 	err := db.InsertPassword(userFalse.ID, credentialsTrue.Salt, credentialsTrue.Hash)
// 	if err != nil {
// 		t.Error("cannot insertPassword", err.Error())
// 	}
// }

// func Test_GetId_FieldWithIdIsNotEmpty_True(t *testing.T) {
// 	db := getDB(t, testURL)

// 	testId, err := db.GetId(userTrue.Email)
// 	if err != nil {
// 		t.Error("cannot take id", testId, err.Error())
// 	}
// }

// func Test_GetId_FieldWithIdIsEmpty_False(t *testing.T) {
// 	db := getDB(t, testURL)
// 	testId, err := db.GetId(userFalse.Email)
// 	if err != nil {
// 		t.Error("cannot take id", testId, err.Error())
// 	}
// }

// func Test_GetEmail_FieldWithEmailIsNotEmpty_True(t *testing.T) {
// 	db := getDB(t, testURL)
// 	testEmail, err := db.GetEmail(userTrue.Email)
// 	if err != nil {
// 		t.Error("cannot take email", testEmail, err.Error())
// 	}
// }

// func Test_GetSaltAndHash_FieldWithCredetialsIsNotEmpty_True(t *testing.T) {
// 	db := getDB(t, testURL)
// 	getSalt, getHash, err := db.GetSaltAndHash(userTrue.ID)
// 	if err != nil {
// 		t.Error("cannot take salt and id", getSalt, getHash, err)
// 	}
// }

// func Test_GetSaltAndHash_FieldWithCredetialsIsEmpty_False(t *testing.T) {
// 	db := getDB(t, testURL)
// 	getSalt, getHash, err := db.GetSaltAndHash(userFalse.ID)
// 	if len(getSalt) < 0 || len(getHash) < 0 {
// 		t.Error("cannot insert user", err.Error())
// 	}
// 	// if err != nil {
// 	// 	t.Error("cannot take salt and id", getSalt, getHash, err)
// 	// }
// }

/////////////////////////////////////////////////////////////////////////////////////////
// func InsertUser_IncorrectFields_False(t *testing.T) {
// 	db := getDB(t, testURL)
// 	user := models.User{Name: "", LastName: "", Email: "", Password: ""}
// }

// func TestReq(t *testing.T) {
// 	db := getDB(t, testURL)

// testCases1 := map[string]models.User{
// 	"Correct data": {Name: "putin", LastName: "dayn", Email: "putinloh@gmail.com", Password: "1111111"},
// 	"Invalid data": {Name: "", LastName: "", Email: "", Password: ""},
// }

// testCases2 := map[string]struct {
// 	Id   int
// 	Salt string
// 	Hash string
// }{
// 	"Correct data": {Id: 1, Salt: "7oGQ7CwmdjEXV7NU", Hash: "ztS5F8G5IfOH3mSu"},
// 	"Invalid data": {Id: 2, Salt: "", Hash: ""},
// }

// testCases3 := map[string]struct {
// 	Id    int
// 	Token string
// }{
// 	"Correct data": {Id: 1, Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"},
// 	"Invalid data": {Id: 2, Token: ""},
// }
// tc := models.User{Name: "Big", LastName: "Bob", Email: "bigbog123@gmail.com", Password: "bigbob123"}
// // for _, tc := range testCases1 {

// if len(tc.Name) > 0 || len(tc.LastName) > 0 || len(tc.Email) > 0 || len(tc.Password) > 0 {
// 	db.InsertUser(&tc)
// 	return
// }
// t.Error("cannot insert user")

// t.Run("testUser", func(t *testing.T) {
// 	testName, testLname, err := db.GetUser(tc.Email)
// 	if err != nil {
// 		t.Error("cannot take user:", testName, testLname, err)
// 	}
// })

// t.Run("testId", func(t *testing.T) {
// 	testId, err := db.GetId(tc.Email)
// 	if err != nil {
// 		t.Error("cannot take id", testId, err)
// 	}
// })

// t.Run("testEmail", func(t *testing.T) {
// 	testEmail := db.GetEmail(tc.Email)
// 	if testEmail != nil {
// 		t.Error("cannot take email", testEmail)
// 	}
// })
// }

// for _, tc := range testCases2 {
// 	t.Run("insertPassword ", func(t *testing.T) {
// 		err := db.InsertPassword(tc.Id, tc.Salt, tc.Hash)
// 		if err != nil {
// 			t.Error("cannot insertPassword", err)
// 		}
// 	})

// 	t.Run("testSaltandHash ", func(t *testing.T) {
// 		getSalt, getHash, err := db.GetSaltAndHash(tc.Id)
// 		if err != nil {
// 			t.Error("cannot take salt and id", getSalt, getHash, err)
// 		}
// 	})
// }

// for _, tc := range testCases3 {
// 	t.Run("insertToken", func(t *testing.T) {
// 		err := db.InsertToken(tc.Id, tc.Token)
// 		if err != nil {
// 			t.Error("cannot take id:", err)
// 		}
// 	})

// 	t.Run("selectToken", func(t *testing.T) {
// 		_, err := db.SelectToken(tc.Id)
// 		if err != nil {
// 			t.Error(err)
// 		}
// 	})

// 	t.Run("deleteToken", func(t *testing.T) {
// 		if len(tc.Token) > 0 {
// 			db.DeleteToken(tc.Token)
// 			return
// 		}
// 		t.Error("cannot delete token")
// 	})
// }
// }
