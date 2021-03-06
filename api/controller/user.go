package controller

import (
	"dependency_injection_tut/model"
	"dependency_injection_tut/service"
	"dependency_injection_tut/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{
	userServie *service.UserService
}

type CreateUserCredentials struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}
type LoginUserCredentials struct{
	Email string `json:"email"`
	Password string `json:"password"`
}
type UserReturnType struct{
	ID uint `json:"id"`
	NAME string `json:"name"`
	EMAIL string	`json:"email"`
}

func NewUserController(us *service.UserService) *UserController{
	return &UserController{userServie: us}
}

func (uc *UserController) ReturnUser(u *model.User) *UserReturnType {
	rtrnUser := &UserReturnType{
		ID:u.ID,
		EMAIL:u.Email,
		NAME:u.Name,
	}
return rtrnUser
}

func(uc *UserController) Create(c *gin.Context){
	user := &model.User{}
	if err:= c.Bind(user);err!=nil{
		log.Println("Binding context to user faile")
		return 
	}
	
	user,err := uc.userServie.Create(user)
	if err!=nil{
		log.Println("Binding context to user failed")
	}
	returnUser:= uc.ReturnUser(user)


	c.JSON(200,gin.H{
		"data":returnUser,
	})

}


func (uc *UserController) Login(c *gin.Context){
	var payload LoginUserCredentials
	if err:=c.Bind(&payload);err!=nil{
		log.Println("Binding context to user faile")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		user:= &model.User{
			Email: payload.Email,
		}


	user,err := uc.userServie.CheckUserExists(user)	
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "User doesnot exists"})
		return
	}
	isPasswordMatched := utils.CheckPasswordHash(payload.Password,string(user.Password) )
	if !isPasswordMatched{
		c.JSON(http.StatusForbidden, gin.H{"error": "Password Incorrect"})
		return
	}

	token,err:= utils.GenerateToken(user)
	if err!=nil{
		c.JSON(http.StatusForbidden, gin.H{"error": "Couldnot create token"})
		return
	}

	c.JSON(200,gin.H{
		"data":gin.H{
			"token":token,
			"user":uc.ReturnUser(user),
		},
	})
	
}

func(uc *UserController) GetAll(c *gin.Context){
	fmt.Println("Get all ")
}