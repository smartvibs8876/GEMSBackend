package controllers

import (
	_ "crypto/aes"
	_ "crypto/cipher"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rest-go-demo/database"
	"rest-go-demo/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_secret_key")

type LoginRequestBody struct {
	Password string
	Email    string
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type ResponseForGetUserDetailsWithToken struct {
	Userid    int    `json:"userid"`
	Fname     string `json:"fname"`
	Lname     string `json:"lname"`
	Email     string `json:"email"`
	Mobile_no string `json:"mobile_no"`
	Address   string `json:"address"`
}

// func decryptPassword(encryptedPassword string) string { //[]byte("ABCDEFGHIJKLMOPQ")
// 	key := []byte("ABCDEFGHIJKLMOPQ")
// 	ciphertext := []byte(encryptedPassword)
// 	c, _ := aes.NewCipher(key)
// 	gcm, _ := cipher.NewGCM(c)

// 	nonceSize := gcm.NonceSize()
// 	if len(ciphertext) < nonceSize {
// 		panic("ciphertext size is less than nonceSize")
// 	}

// 	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
// 	plaintext, _ := gcm.Open(nil, nonce, ciphertext, nil)
// 	return string(plaintext)
// }
func Registration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestBody, _ := ioutil.ReadAll(r.Body)
	newUser := entity.Users{}
	json.Unmarshal(requestBody, &newUser)
	//newUser.Password = decryptPassword(newUser.Password)
	if newUser.Email == "" || newUser.Password == "" || newUser.F_name == "" || newUser.L_name == "" || newUser.Address == "" || newUser.Mo_no == "" {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(false)
		return
	}
	user := entity.Users{}
	database.Connector.Where("email = ?", newUser.Email).First(&user)
	newUser.Password, _ = HashPassword(newUser.Password)
	if user.Email != newUser.Email {
		if err := database.Connector.Create(newUser).Error; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(false)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(true)
		return
	} else {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(false)
		return
	}
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user entity.Users
	json.Unmarshal(requestBody, &user)
	userEnteredPassword := user.Password
	//user.Password = decryptPassword(user.Password)
	user.Password, _ = HashPassword(user.Password)
	database.Connector.Where("email = ?", user.Email).First(&user)
	if user.User_id != 0 {
		match := CheckPasswordHash(userEnteredPassword, user.Password)
		if match == true {
			expirationTime := time.Now().Add(24 * time.Hour)
			claims := &Claims{
				Email: user.Email,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(jwtKey)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(false)
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(tokenString)
			return
		} else {
			//fmt.Println("Incorrect Password")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(false)
			return
		}
	} else {
		//fmt.Println("Incorrect Email")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(false)
		return
	}
}

func GetUserDetailsWithToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := Authorization(w, r)
	if (user == entity.Users{}) {
		return
	}
	var response ResponseForGetUserDetailsWithToken
	response.Userid = user.User_id
	response.Address = user.Address
	response.Email = user.Email
	response.Fname = user.F_name
	response.Lname = user.L_name
	response.Mobile_no = user.Mo_no
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
