package handler

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/tosashimanto/cd-center/model"
	"github.com/tosashimanto/cd-center/util"
	"net/http"
	"time"
)

// 1. トークン生成
func GetToken(c echo.Context) error {

	tokenPost := new(model.TokenJSONPost)
	if err := c.Bind(tokenPost); err != nil {
		return err
	}
	fmt.Println("Token=", tokenPost.Token)

	// Create token(JWT)
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	//claims["name"] = username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	fmt.Println("token=" + t)
	c.Response().Header().Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", t))

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

// 2. 一覧取得
func Constructions(c echo.Context) error {
	var constructionArray [5]model.Construction
	constructionArray[0] = model.Construction{
		ConstructionId:    1,
		ExecutionNum:      "100000000012345",
		PropertyName:      "PropertyName その1",
		Address:           "東京都XX-X1",
		ConstManagerPhone: "000-0000-0000",
	}
	constructionArray[1] = model.Construction{
		ConstructionId:    2,
		ExecutionNum:      "100000000012346",
		PropertyName:      "PropertyName その2",
		Address:           "東京都XX-X2",
		ConstManagerPhone: "000-0000-0002",
	}
	constructionArray[2] = model.Construction{
		ConstructionId:    3,
		ExecutionNum:      "100000000012347",
		PropertyName:      "PropertyName その4",
		Address:           "東京都XX-X4",
		ConstManagerPhone: "000-0000-0003",
	}
	constructionArray[3] = model.Construction{
		ConstructionId:    4,
		ExecutionNum:      "100000000012348",
		PropertyName:      "PropertyName その5",
		Address:           "東京都XX-X5",
		ConstManagerPhone: "000-0000-0004",
	}
	constructionArray[4] = model.Construction{
		ConstructionId:    5,
		ExecutionNum:      "100000000012349",
		PropertyName:      "PropertyName その6",
		Address:           "東京都XX-X6",
		ConstManagerPhone: "000-0000-0006",
	}
	constructions := &model.Constructions{
		ConstructionArray: []model.Construction{
			constructionArray[0],
			constructionArray[1],
			constructionArray[2],
			constructionArray[3],
			constructionArray[4],
		},
	}
	jsonString, _ := json.Marshal(constructions)
	util.JSONFormatOut(jsonString)

	return c.JSON(http.StatusOK, constructions)
}
