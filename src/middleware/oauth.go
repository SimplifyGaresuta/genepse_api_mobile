package middleware

import (
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
)

var user common.User

func GetUser(providerName string, urlQuery string) (common.User, error) {
	provider, err := gomniauth.Provider(providerName)
	if err != nil {
		return nil, err
	}

	creds, err := provider.CompleteAuth(objx.MustFromURLQuery(urlQuery))
	if err != nil {
		return nil, err
	}

	user, err = provider.GetUser(creds)
	if err != nil {
		return nil, err
	}
	return user, nil
}
