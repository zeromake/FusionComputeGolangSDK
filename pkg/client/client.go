package client

import (
	"context"
	"errors"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/common"
	"github.com/go-resty/resty/v2"
)

type Session string

type FusionComputeClient interface {
	Connect(ctx context.Context) error
	DisConnect(ctx context.Context) error
	SetSession(token string)
	GetSession() Session
	GetHost() string
	GetUser() string
	GetPassword() string
	GetApiClient() (*resty.Client, error)
	GetUserType() string
	SetUserType(userType string)
}

func NewFusionComputeClient(host string, user string, password string) FusionComputeClient {
	return &fusionComputeClient{
		user:     user,
		password: password,
		host:     host,
		userType: "2",
	}
}

type fusionComputeClient struct {
	session  Session
	user     string
	password string
	host     string
	userType string
}

func (f *fusionComputeClient) GetUserType() string {
	return f.userType
}

func (f *fusionComputeClient) SetUserType(userType string) {
	f.userType = userType
}

func (f *fusionComputeClient) SetSession(token string) {
	f.session = Session(token)
}

func (f *fusionComputeClient) GetSession() Session {
	return f.session
}

func (f *fusionComputeClient) Connect(ctx context.Context) error {
	a := NewAuth(f)
	err := a.Login(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (f *fusionComputeClient) DisConnect(ctx context.Context) error {
	a := NewAuth(f)
	err := a.Logout(ctx)
	if err != nil {
		return err
	}
	return nil
}
func (f *fusionComputeClient) GetHost() string {
	return f.host
}
func (f *fusionComputeClient) GetUser() string {
	return f.user
}
func (f *fusionComputeClient) GetPassword() string {
	return f.password
}

func (f *fusionComputeClient) GetApiClient() (*resty.Client, error) {
	r := common.NewHttpClient()
	if f.GetSession() == "" {
		return nil, errors.New("no session exists,please login and try it again")
	}
	f.setDefaultHeader(r)
	r.SetHeader(XAuthToken, string(f.GetSession())).
		SetHostURL(f.host)
	return r, nil
}

func (f *fusionComputeClient) setDefaultHeader(client *resty.Client) {
	client.SetHeaders(map[string]string{
		"Accept":          "application/json;version=v8.0;charset=UTF-8;",
		"Accept-Language": "zh_CN:1.0",
	})
}
