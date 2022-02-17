package handlers

import (
	"net/http"

	"github.com/ATM/pkg/common"
	"github.com/ATM/pkg/reststructs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h *ATMHandler) LoginWithPin(c *gin.Context) {
	var login reststructs.Login
	err := common.ParseRequest(c.Request, &login)
	if err != nil {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"error_msg": "invalid payload is provided"})
		return
	}
	if len(login.Pin) != 4 {
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"error_msg": "invalid pin is provided"})
		return
	}
	account, err := common.GetAccountByAPIkey(login.APIKey)
	if err != nil {
		c.Abort()
		c.JSON(http.StatusNotFound, gin.H{"error_msg": "account is not found"})
		return
	}
	if account.Pin != login.Pin {
		c.Abort()
		c.JSON(http.StatusForbidden, gin.H{"error_msg": "pin number is not correct"})
		return
	}
	claim := reststructs.AccountClaim{
		AccountID: account.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "nameOfWebsiteHere",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte("secureSecretText"))

	if err != nil {
		c.Abort()
		c.JSON(http.StatusInternalServerError, gin.H{"error_msg": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, reststructs.LoginToken{Token: signedToken})
}
