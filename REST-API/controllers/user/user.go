package user

import (
	"errors"
	"fmt"
	"net/http"
	"restful/controllers"
	userModel "restful/models/user"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type profileReq struct {
	Name        string           `json:"name"`
	Gender      userModel.Gender `json:"gender"`
	Email       string           `json:"email"`
	PhoneNumber string           `json:"phone"`
}

type userResp struct {
	ID      uint        `json:"id"`
	Profile profileResp `json:"info,omitempty"`
}

type profileResp struct {
	Name        string           `json:"name,omitempty"`
	Gender      userModel.Gender `json:"gender,omitempty"`
	Email       string           `json:"email,omitempty"`
	PhoneNumber string           `json:"phone,omitempty"`
}

func (r profileReq) verify() error {
	if r.Name == "" {
		return errors.New("Invalid Param: name")
	}

	if r.Gender.Verify() != nil {
		return errors.New("Invalid Param: gender")
	}

	if r.Email == "" {
		return errors.New("Invalid Param: email")
	}

	if r.PhoneNumber == "" {
		return errors.New("Invalid Param: phone")
	}

	return nil
}

const (
	JWTAuthSecret   string = "Super@ppCid"
	JWTAuthIssuer   string = "SuperAppKid@gmail.com"
	JWTAuthAudience string = "com.superappkid.www"
)

func Auth(c *gin.Context) {
	authStr := c.Request.Header.Get("Authorization")
	authParts := strings.Split(authStr, " ")

	if len(authParts) > 0 && len(authParts) != 2 && authParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, controllers.TextResponse("invalid authorization"))
		return
	}

	token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTAuthSecret), nil
	},
		jwt.WithIssuer(JWTAuthIssuer),
		jwt.WithAudience(JWTAuthAudience),
		jwt.WithValidMethods([]string{
			jwt.SigningMethodHS256.Name,
			jwt.SigningMethodHS384.Name,
			jwt.SigningMethodHS512.Name,
		}),
	)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, controllers.ErrorResponse(err))
		return
	}

	sub, err := token.Claims.GetSubject()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, controllers.ErrorResponse(err))
		return
	}

	intID, err := strconv.Atoi(sub)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	if intID <= 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("invalid user id"))
		return
	}

	c.Set("id", intID)
	c.Next()
}

func GetAll(c *gin.Context) {
	users, err := userModel.FetchAll()
	if err != nil {
		c.JSON(http.StatusNotFound, controllers.ErrorResponse(err))
		return
	}

	respData := make([]userResp, 0, len(users))
	for _, user := range users {
		respData = append(respData, transUser2Resp(user))
	}

	c.JSON(http.StatusOK, controllers.DataResponse(respData))
}

func GetProfileByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusNotFound, controllers.TextResponse(fmt.Sprintf("user: %s not exist", idStr)))
		return
	}

	user, err := userModel.FetchByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, controllers.ErrorResponse(err))
		return
	}

	resp := transUser2Resp(user)
	c.JSON(http.StatusOK, controllers.DataResponse(resp))
}

func CreateTokenByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusNotFound, controllers.TextResponse(fmt.Sprintf("user: %s not exist", idStr)))
		return
	}

	if _, err := userModel.FetchByID(id); err != nil {
		c.JSON(http.StatusNotFound, controllers.ErrorResponse(err))
		return
	}

	now := time.Now()
	claims := &jwt.RegisteredClaims{
		Issuer:    JWTAuthIssuer,
		Audience:  []string{JWTAuthAudience},
		Subject:   strconv.Itoa(id),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Minute)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(JWTAuthSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, controllers.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, controllers.DataResponse(
		gin.H{
			"token": tokenStr,
		},
	))
}

func Create(c *gin.Context) {
	req := profileReq{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, controllers.ErrorResponse(err))
		return
	}

	if err := req.verify(); err != nil {
		c.JSON(http.StatusNotFound, controllers.ErrorResponse(err))
		return
	}

	newUser := &userModel.User{
		Profile: userModel.Profile{
			Name:        req.Name,
			Gender:      req.Gender,
			Email:       req.Email,
			PhoneNumber: req.PhoneNumber,
		},
	}

	if err := newUser.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, controllers.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, controllers.DataResponse(
		gin.H{
			"id": newUser.ID,
		},
	))
}

func Delete(c *gin.Context) {
	id := c.GetInt("id")

	user := userModel.User{ID: uint(id)}
	if err := user.Delete(); err != nil {
		c.JSON(http.StatusBadRequest, controllers.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, controllers.TextResponse("success"))
}

func GetSelfProfile(c *gin.Context) {
	id := c.GetInt("id")

	user, err := userModel.FetchByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, controllers.ErrorResponse(err))
		return
	}

	resp := transUser2Resp(user)
	c.JSON(http.StatusOK, controllers.DataResponse(resp))
}

func UpdateSelfProfile(c *gin.Context) {
	id := c.GetInt("id")

	req := profileReq{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, controllers.ErrorResponse(err))
		return
	}

	if err := req.verify(); err != nil {
		c.JSON(http.StatusBadRequest, controllers.ErrorResponse(err))
		return
	}

	profile := &userModel.Profile{
		Name:        req.Name,
		Gender:      req.Gender,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		UserID:      uint(id),
	}

	if err := profile.UpdateAll(); err != nil {
		c.JSON(http.StatusInternalServerError, controllers.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, controllers.TextResponse("success"))
}

func PatchSelfProfile(c *gin.Context) {
	id := c.GetInt("id")

	patchMap := make(map[string]interface{})

	if err := c.ShouldBindJSON(&patchMap); err != nil {
		c.JSON(http.StatusBadRequest, controllers.ErrorResponse(err))
		return
	}

	profile := &userModel.Profile{UserID: uint(id)}

	if name, ok := patchMap["name"].(string); ok {
		profile.Name = name
	}

	if genderStr, ok := patchMap["gender"].(string); ok {
		gender := userModel.Gender(genderStr)
		if err := gender.Verify(); err != nil {
			c.JSON(http.StatusBadRequest, controllers.ErrorResponse(err))
			return
		}
		profile.Gender = gender
	}

	if email, ok := patchMap["email"].(string); ok {
		profile.Email = email
	}

	if phoneNumber, ok := patchMap["phone"].(string); ok {
		profile.PhoneNumber = phoneNumber
	}

	if err := profile.UpdateNonZero(); err != nil {
		c.JSON(http.StatusInternalServerError, controllers.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, controllers.TextResponse("success"))
}

func transUser2Resp(user *userModel.User) userResp {
	return userResp{
		ID: user.ID,
		Profile: profileResp{
			Name:        user.Profile.Name,
			Gender:      user.Profile.Gender,
			Email:       user.Profile.Email,
			PhoneNumber: user.Profile.PhoneNumber,
		},
	}
}
