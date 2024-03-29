package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyService struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func CareteJwt() {
	mySigningKey := []byte("fuckyou")
	c := MyService{
		Username: "NongChen",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,      //生效时间 当前时间减去60秒
			ExpiresAt: time.Now().Unix() + 60*60*2, //失效时间 当前时间加上60秒
			Issuer:    "NongChen",                  //签发人
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c) //生成token
	str, err := t.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("%s", err)
	}
	fmt.Println("加密字符串", str)

}
