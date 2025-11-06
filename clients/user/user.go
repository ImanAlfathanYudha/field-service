package clients

import (
	"context"
	"field-service/clients/config"
	"field-service/common/util"
	config2 "field-service/config"
	"field-service/constants"
	"fmt"
	"net/http"
	"time"
)

type UserClient struct {
	client config.IClientConfig
}

type IUserClient interface {
	GetUserByToken(ctx context.Context) (*UserData, error)
}

func NewUserClient(client config.IClientConfig) IUserClient {
	return &UserClient{client: client}
}

func (u *UserClient) GetUserByToken(ctx context.Context) (*UserData, error) {
	unixTime := time.Now().Unix()
	generateApiKey := fmt.Sprintf("%s:%s:%d", config2.Config.AppName, u.client.SignatureKey(), unixTime)
	apiKey := util.GenerateSHA256(generateApiKey)
	token := ctx.Value(constants.Token).(string)

	var response UserResponse
	request := u.client.Client().
		Get(fmt.Sprintf("%s/api/v1/auth/user", u.client.BaseURL())).
		Set("Authorization", fmt.Sprintf("%s", token)).
		Set("x-api-key", apiKey).
		Set("x-service-name", config2.Config.AppName).
		Set("x-request-at", fmt.Sprintf("%d", unixTime))
	resp, _, errs := request.EndStruct(&response)

	if len(errs) > 0 {
		return nil, errs[0]
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("user response: %s", response.Message)
	}
	return &response.Data, nil
}
