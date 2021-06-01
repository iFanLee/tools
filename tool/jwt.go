package tool

import "github.com/dgrijalva/jwt-go"

/**
 * @Author: Lee
 * @Date: 2021/6/1 15:22
 * @Desc:
 */
const secret = "test"
type Claims struct {
	Id int `json:"id"`
	UpdateTime string `json:"updateTime"`
	jwt.StandardClaims
}
//生成jwt token
func GenerateToken(id int,uuid int,mobile string, updateTime string) (string, error) {
	//nowTime := time.Now()
	//expireTime := nowTime.Add(30 * time.Hour)
	claims := Claims{
		id,
		updateTime,
		jwt.StandardClaims {
			//ExpiresAt : expireTime.Unix(),
			Issuer : "test",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(secret))
	return token, err
}
//解密jwt token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}