package common

import (
	"encoding/json"
	"net/http"

	"github.com/ATM/pkg/reststructs"
)

func ParseTransactRequest(request *http.Request) (*reststructs.Transact, error) {
	transact := &reststructs.Transact{}
	err := json.NewDecoder(request.Body).Decode(transact)
	if err != nil {
		return nil, err
	}
	return transact, nil
}
