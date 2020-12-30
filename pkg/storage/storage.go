package storage

import (
	"context"
	"fmt"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/common"
)

const (
	//siteMask     = "<site_uri>"
	datastoreUrl = "%s/datastores"
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
		Get(fmt.Sprintf(datastoreUrl, m.siteUri))
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
