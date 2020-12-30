package cluster

import (
	"context"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/common"
	"strings"
)

const (
	siteMask   = "<site_uri>"
	clusterUrl = "<site_uri>/clusters"
)

type Manager interface {
	ListCluster(ctx context.Context) ([]Cluster, error)
}

func NewManager(client client.FusionComputeClient, siteUri string) Manager {
	return &manager{client: client, siteUri: siteUri}
}

type manager struct {
	client  client.FusionComputeClient
	siteUri string
}

func (m *manager) ListCluster(ctx context.Context) ([]Cluster, error) {
	var listClusterResponse ListClusterResponse
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().
		SetContext(ctx).
		SetResult(&listClusterResponse).
		Get(strings.Replace(clusterUrl, siteMask, m.siteUri, -1))
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, common.FormatHttpError(resp)
	}
	return listClusterResponse.Clusters, nil

}
