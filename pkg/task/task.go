package task

import (
	"context"
	"path"

	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/common"
)

type Manager interface {
	Get(ctx context.Context, taskUri string) (*Task, error)
}

func NewManager(client client.FusionComputeClient, siteUri string) Manager {
	return &manager{client: client, siteUri: siteUri}
}

type manager struct {
	client  client.FusionComputeClient
	siteUri string
}

func (m *manager) Get(ctx context.Context, taskUri string) (*Task, error) {
	var task Task
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().
		SetContext(ctx).
		SetResult(&task).
		Get(path.Join(taskUri))
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, common.FormatHttpError(resp)
	}
	return &task, nil
}
