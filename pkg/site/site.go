package site

import (
	"context"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/common"
)

const (
	siteUrl = "/service/sites"
)

type Interface interface {
	ListSite(ctx context.Context) ([]Site, error)
	GetSite(ctx context.Context, siteUri string) (*Site, error)
}

func NewManager(client client.FusionComputeClient) Interface {
	return &manager{client: client}
}

type manager struct {
	client client.FusionComputeClient
}

func (m *manager) GetSite(ctx context.Context, siteUri string) (*Site, error) {
	var site Site
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().
		SetContext(ctx).
		SetResult(&site).
		Get(siteUri)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, common.FormatHttpError(resp)
	}

	return &site, nil
}

func (m *manager) ListSite(ctx context.Context) ([]Site, error) {
	var listSiteResponse ListSiteResponse
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().
		SetContext(ctx).
		SetResult(&listSiteResponse).
		Get(siteUrl)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, common.FormatHttpError(resp)
	}
	return listSiteResponse.Sites, nil
}
