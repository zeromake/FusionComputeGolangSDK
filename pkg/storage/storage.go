package storage

import (
	"context"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/common"
	"strings"
)

const (
	siteMask     = "<site_uri>"
	datastoreUrl = "<site_uri>/datastores"
)

type Interface interface {
	ListDataStore(ctx context.Context) ([]Datastore, error)
}

func NewManager(client client.FusionComputeClient, siteUri string) Interface {
	return &manager{client: client, siteUri: siteUri}
}

type manager struct {
	client  client.FusionComputeClient
	siteUri string
}

func (m *manager) ListDataStore(ctx context.Context) ([]Datastore, error) {
	var adapters []Datastore
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	var listAdapterResponse ListDataStoreResponse
	resp, err := api.R().
		SetContext(ctx).
		SetResult(&listAdapterResponse).
		Get(strings.Replace(datastoreUrl, siteMask, m.siteUri, -1))
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, common.FormatHttpError(resp)
	} else {
		adapters = listAdapterResponse.Datastores
	}
	return adapters, nil
}
