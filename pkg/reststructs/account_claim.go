package reststructs

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

type AccountClaim struct {
	AccountID uuid.UUID `json:"account_id"`
	jwt.StandardClaims
}
