package controllers

import (
	"bytes"
	"encoding/json"
	"gin-mongo-api/configs"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var loginCollection *mongo.Collection = configs.GetCollection(configs.DB, "login")
var validateLogin = validator.New()

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var login models.Login
		// defer cancel()

		// validate the request body
		if err := c.BindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{
				StatusCode: http.StatusBadRequest,
				Message:    "error",
				Result:     map[string]interface{}{"data": err.Error()}})
			return
		}

		// use the validator library to validate required fields
		if validationErr := validateLogin.Struct(&login); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.LoginResponse{
				StatusCode: http.StatusBadRequest,
				Message:    "error",
				Result:     map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		// url_api := configs.EnvMaterai()
		reqData, _ := json.Marshal(login)

		strData := bytes.NewBuffer(reqData)

		resp, err := http.Post("https://emeteraidevent.scm.perurica.co.id/api/users/login", "application/json", strData)

		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer resp.Body.Close()
		//Read the response body
		respBody, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}

		jsonString := string(respBody)
		log.Printf(jsonString)

		dataJson := map[string]interface{}{}
		if err := json.Unmarshal([]byte(respBody), &dataJson); err != nil {
			log.Fatal(err)
		}

		// newLogin := models.Login{
		// 	User:     login.User,
		// 	Password: login.Password,
		// }

		// result, err := loginCollection.InsertOne(ctx, newLogin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LoginResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Result:     map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, dataJson)

		// c.JSON(http.StatusCreated, responses.LoginResponse{
		// 	StatusCode: http.StatusCreated,
		// 	Message:    "success",
		// 	Result:     map[string]interface{}{"data": sb}})

	}
}
