package postgresql

import (
	"testing"

	"github.com/Yscream/go-form-reg/pkg/models"
)

//tips for named test cases :
//func always gotta start with "Test" in the func name

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

var userWithCorrectFields = models.User{ID: 1, Name: "Big", LastName: "Bob", Email: "bigbog123@gmail.com", Password: "bigbob123"}
var userWithSameEmail = models.User{ID: 2, Name: "Putin", LastName: "Huilo", Email: "bigbog123@gmail.com", Password: "bigbob123"}
var userWithSameID = models.User{ID: 1, Name: "John", LastName: "Armstrong", Email: "johnbog1234@gmail.com", Password: "bigbob123"}
var userWithEmptyFields = models.User{Name: "", LastName: "", Email: "", Password: ""}

var credentialsWithCorrectFields = models.Credentials{ID: userWithCorrectFields.ID, Salt: "7oGQ7CwmdjEXV7NU", Hash: "ztS5F8G5IfOH3mSu"}
var credentialsFalse = models.Credentials{Salt: "", Hash: ""}
var accessTokenWithCorrectFields = models.AccessToken{ID: userWithCorrectFields.ID, Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}
var accessTokenFalse = models.AccessToken{Token: ""}

func Test_InsertUserWithCorrectFields_Success(t *testing.T) {
	db := getDB(t, testURL)

	err := db.InsertUser(&userWithCorrectFields)
	if err != nil {
		t.Errorf("cannot insert user %s", err.Error())
	}
	t.Run("getUser", func(t *testing.T) {
		getuser, err := db.GetUser(userWithCorrectFields.Email)
		if userWithCorrectFields.ID == getuser.ID &&
			userWithCorrectFields.Name == getuser.Name &&
			userWithCorrectFields.LastName == getuser.Lname &&
			userWithCorrectFields.Email == getuser.Email {
			t.Logf("db:( ID: %d Name: %s, Lname: %s Email: %s) coincides with userCorrectFields:( ID: %d Name: %s, Lname: %s Email: %s)",
				getuser.ID, getuser.Name, getuser.Lname, getuser.Email, userWithCorrectFields.ID, userWithCorrectFields.Name, userWithCorrectFields.LastName, userWithCorrectFields.Email)
		}
		if err != nil {
			t.Errorf("cannot take user, %s", err.Error())
		}
	})
}

func Test_InsertCredentials_Success(t *testing.T) {
	db := getDB(t, testURL)

	err := db.InsertCredentials(&credentialsWithCorrectFields)
	if err != nil {
		t.Errorf("cannot insert credentials, %s", err.Error())
	}
	t.Run("getCredetials", func(t *testing.T) {
		getCred, err := db.GetCredentials(credentialsWithCorrectFields.ID)
		if credentialsWithCorrectFields.ID == getCred.ID &&
			credentialsWithCorrectFields.Salt == getCred.Salt &&
			credentialsWithCorrectFields.Hash == getCred.Hash {
			t.Logf("db:( ID: %d, Salt: %s, Hash: %s) match with credentialsWithCorrectFields:(ID: %d, Salt: %s, Hash: %s)",
				getCred.ID, getCred.Salt, getCred.Hash, credentialsWithCorrectFields.ID, credentialsWithCorrectFields.Salt, credentialsWithCorrectFields.Hash)
		}
		if err != nil {
			t.Errorf("cannot take credentials %s", err.Error())
		}
	})
}

func Test_InsertToken_Success(t *testing.T) {
	db := getDB(t, testURL)

	err := db.InsertToken(&accessTokenWithCorrectFields)
	if err != nil {
		t.Errorf("cannot insert credentials, %s", err.Error())
	}
	t.Run("selectToken", func(t *testing.T) {
		getToken, err := db.SelectToken(accessTokenWithCorrectFields.ID)
		if accessTokenWithCorrectFields.Token == getToken {
			t.Logf("db:(Token: %s) match with accessTokenWithCorrectFields:(Token: %s)", getToken, accessTokenFalse.Token)
		}
		if err != nil {
			t.Errorf("cannot take token %s", err.Error())
		}
	})
}

func Test_GetId_Success(t *testing.T) {
	db := getDB(t, testURL)

	testU := models.User{ID: 3, Name: "Michel", LastName: "Oddman", Email: "michelodd121@gmail.com", Password: "michelodmichel1"}
	t.Run("insertUser", func(t *testing.T) {
		err := db.InsertUser(&testU)
		if err != nil {
			t.Errorf("cannot insert user %s", err.Error())
		}
	})
	id, err := db.GetId(testU.Email)
	if id == testU.ID {
		t.Logf("db: (ID: %d) match with testU: (ID: %d)", id, testU.ID)
	}
	if err != nil {
		t.Errorf("cannot get id %s", err.Error())
	}
}

func Test_GetEmail_Success(t *testing.T) {
	db := getDB(t, testURL)

	testU := models.User{ID: 4, Name: "Mike", LastName: "Deal", Email: "mikysq121@gmail.com", Password: "micksssas"}
	t.Run("insertUser", func(t *testing.T) {
		err := db.InsertUser(&testU)
		if err != nil {
			t.Errorf("cannot insert user %s", err.Error())
		}
	})
	email, err := db.GetEmail(testU.Email)
	if email == testU.Email {
		t.Logf("db: (Email: %s) match with testU: (Email: %d)", email, testU.ID)
	}
	if err != nil {
		t.Errorf("cannot get email %s", err.Error())
	}
}

func Test_GetCredentials_Success(t *testing.T) {
	db := getDB(t, testURL)

	testC := models.Credentials{ID: 2, Salt: "7oGQ7sd2CwmdjEXV7sds1NU", Hash: "xx2431ztS5F8G5sda21IfOH3mSu"}
	t.Run("insertCredentials", func(t *testing.T) {
		err := db.InsertCredentials(&testC)
		if err != nil {
			t.Errorf("cannot insert user %s", err.Error())
		}
	})
	getCred, err := db.GetCredentials(testC.ID)
	if testC.ID == getCred.ID &&
		testC.Salt == getCred.Salt &&
		testC.Hash == getCred.Hash {
		t.Logf("db:( ID: %d, Salt: %s, Hash: %s) match with credentialsWithCorrectFields:(ID: %d, Salt: %s, Hash: %s)",
			getCred.ID, getCred.Salt, getCred.Hash, testC.ID, testC.Salt, testC.Hash)
	}
	if err != nil {
		t.Errorf("cannot take credentials %s", err.Error())
	}
}

// func Test_GetToken_Success(t *testing.T) {
// 	db := getDB(t, testURL)

// }

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
