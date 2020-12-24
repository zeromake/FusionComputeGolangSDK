package client

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/common"
)

const (
	XAuthUser     = "X-Auth-User"
	XAuthKey      = "X-Auth-Key"
	XAuthUserType = "X-Auth-UserType"
	XAuthToken    = "X-Auth-Token"

	authUri = "/service/session"
)

type Auth interface {
	Login(ctx context.Context) error
	Logout(ctx context.Context) error
}

func NewAuth(client FusionComputeClient) Auth {
	return &auth{client: client}
}

type auth struct {
	client FusionComputeClient
}

func (a *auth) Login(ctx context.Context) error {
	host := a.client.GetHost()
	r := common.NewHttpClient()
	r.SetHostURL(host).
		SetHeader(XAuthUser, a.client.GetUser()).
		SetHeader(XAuthKey, encodePassword(a.client.GetPassword())).
		SetHeader(XAuthUserType, a.client.GetUserType())
	resp, err := r.R().SetContext(ctx).Post(authUri)
	if err != nil {
		return err
	}
	if resp.IsSuccess() {
		var loginResponse LoginResponse
		_ = json.Unmarshal(resp.Body(), &loginResponse)
		token := resp.Header().Get(XAuthToken)
		a.client.SetSession(token)
	} else {
		return common.FormatHttpError(resp)
	}
	return nil
}

func (a *auth) Logout(ctx context.Context) error {
	host := a.client.GetHost()
	r := common.NewHttpClient()
	r.SetHostURL(host).
		SetHeader(XAuthToken, string(a.client.GetSession()))
	resp, err := r.R().SetContext(ctx).Delete(authUri)
	if err != nil {
		return err
	}
	if resp.IsSuccess() {
		a.client.SetSession("")
	} else {
		return common.FormatHttpError(resp)
	}
	return nil
}

func encodePassword(pass string) string {
	bs := sha256.Sum256([]byte(pass))
	return hex.EncodeToString(bs[:])
}
