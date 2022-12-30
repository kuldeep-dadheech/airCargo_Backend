package helpers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	uuid "github.com/gofrs/uuid"
)

type TokenDetails struct {
	AccessToken string;
	RefreshToken string
	AccessUuid string
	RefreshUuid string
	AtExpires int64
	RtExpires int64
}

type RedisInstance struct {
	client *redis.Client
}

type AccessTokenPayload struct {
	Exp time.Time `json:"exp"`
	UserId int `json:"user_id"`
	AccessUuid string `json:"access_uuid"`
	Authorised bool `json:"authorised"`
}

type settingValues struct {
	value string
	expires time.Time
}


func CreateToken(user_id int)(*TokenDetails, error){

	// used for genrating access token and refresh token 

	td := &TokenDetails{}
	var err error
	at_uuid, er := uuid.NewV4()
	if er != nil {
		return nil, er
	}
	rt_uuid, e := uuid.NewV4()
	if e != nil {
		return nil ,e
	}
	td.AccessUuid = at_uuid.String()
	td.RefreshUuid = rt_uuid.String()
	td.AtExpires = time.Now().Add(time.Minute * 3600).Unix()
  	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	
	os.Setenv("ACCESS_SECRET", "HEY_BABY!")
	os.Setenv("REFRESH_SECRET" , "YEAH_BABY!")

	at_claims := jwt.MapClaims{}
	at_claims["authorised"] = true;
	at_claims["user_id"] = user_id
	at_claims["exp"] = td.AtExpires
	at_claims["access_uuid"] = td.AccessUuid

	rt_claims := jwt.MapClaims{}
	rt_claims["authorised"] = true;
	rt_claims["refresh_uuid"] = td.RefreshUuid
	rt_claims["user_id"] = user_id
	rt_claims["exp"] = td.RtExpires

	at_tk := jwt.NewWithClaims(jwt.SigningMethodHS256, at_claims)
	rt_tk := jwt.NewWithClaims(jwt.SigningMethodHS256, rt_claims)

	td.AccessToken , err = at_tk.SignedString([]byte("HEY_BABY!"))
	if err != nil {
		return nil, err
	}

	td.RefreshToken , err = rt_tk.SignedString([]byte(os.Getenv("YEAH_BABY!")))
	
	if err != nil {
		return nil , err
	}
	return td, nil
}

func ExtractToken(r *http.Request) string {
  bearToken := r.Header.Get("Authorization")
  fmt.Println(bearToken, "bear token!")
  //normally Authorization the_token_xxx
  strArr := strings.Split(bearToken, " ")
  fmt.Println(strArr, "string arr in extract token!!!")
  if len(strArr) == 2 {
     return strArr[1]
  }
  return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
  tokenString := ExtractToken(r)
  fmt.Println(tokenString, "tokenstringngng11")
  keyfunc := func(token *jwt.Token) (interface{}, error) {
     //Make sure that the token method conform to "SigningMethodHMAC"
     if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
     }
     return []byte("HEY_BABY!"), nil
  }
  unquoteStr,er := strconv.Unquote(tokenString)
  if er!= nil {
	return nil, er
  }
  token, err := jwt.Parse(unquoteStr,keyfunc)
  fmt.Println(token ,"tokennenenenne")
  fmt.Println(err, "errorororor")
  if err != nil {
     return nil, err
  }
  return token, nil
}


func TokenValid(r *http.Request) error {
  token, err := VerifyToken(r)
  fmt.Println(token, "token herere~!")
  if err != nil {
     return err
  }
  if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
     return err
  }
  return nil
}


func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
  token, err := VerifyToken(r)
  if err != nil {
     return nil, err
  }
  claims, ok := token.Claims.(jwt.MapClaims)
  if ok && token.Valid {
     accessUuid, ok := claims["access_uuid"].(string)
     if !ok {
        return nil, err
     }
     userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
     if err != nil {
        return nil, err
     }
     return &AccessDetails{
        AccessUuid: accessUuid,
        UserId:   userId,
     }, nil
  }
  return nil, err
}
type AccessDetails struct {
    AccessUuid string
    UserId   uint64
}

func TokenMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
	 fmt.Println("in middleware 1")
	 fmt.Println(c.Request, "request herere!!")
	 fmt.Println(c.Request.Header, "headre herere")
     err := TokenValid(c.Request)
	 fmt.Println("in middleware 2")
     if err != nil {
        c.JSON(http.StatusUnauthorized, err.Error())
        c.Abort()
        return
     }
     c.Next()
  }
}