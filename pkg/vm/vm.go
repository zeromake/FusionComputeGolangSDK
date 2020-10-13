package vm

import (
	"encoding/json"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/common"
	"strings"
)

const (
	siteMask = "<site_uri>"
	vmUrl    = "<site_uri>/vms"
)

type Manager interface {
	ListVm(isTemplate bool) ([]Vm, error)
}

func NewManager(client client.FusionComputeClient, siteUri string) Manager {
	return &manager{client: client, siteUri: siteUri}
}

type manager struct {
	client  client.FusionComputeClient
	siteUri string
}

func (m *manager) ListVm(isTemplate bool) ([]Vm, error) {
	var vms []Vm
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	request := api.R()
	if isTemplate {
		request.SetQueryParam("isTemplate", "true")
	}
	resp, err := request.Get(strings.Replace(vmUrl, siteMask, m.siteUri, -1))
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		var listVmResponse ListVmResponse
		err := json.Unmarshal(resp.Body(), &listVmResponse)
		if err != nil {
			return nil, err
		}
		vms = listVmResponse.Vms

	} else {
		return nil, common.FormatHttpError(resp)
	}
	return vms, nil
}
