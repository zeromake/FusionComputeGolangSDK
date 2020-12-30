package network

import (
	"context"
	"fmt"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/common"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/vm"
	"path"
)

const (
	dvSwitchUrl  = "%s/dvswitchs"
	vmScopeUrl   = "%s/vms?scope=%s"
	portGroupUrl = "%s/portgroups"
)

type Manager interface {
	ListDVSwitch(ctx context.Context) ([]DVSwitch, error)
	ListPortGroupBySwitch(ctx context.Context, dvSwitchIdUri string) ([]PortGroup, error)
	ListPortGroupInUseIp(ctx context.Context, portGroupUrn string) ([]string, error)
	ListPortGroup(ctx context.Context) ([]PortGroup, error)
}

func NewManager(client client.FusionComputeClient, siteUri string) Manager {
	return &manager{client: client, siteUri: siteUri}
}

type manager struct {
	client  client.FusionComputeClient
	siteUri string
}

func (m *manager) ListPortGroup(ctx context.Context) ([]PortGroup, error) {
	var listPortGroupResponse ListPortGroupResponse
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().
		SetContext(ctx).
		SetResult(&listPortGroupResponse).
		Get(fmt.Sprintf(portGroupUrl, m.siteUri))
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, common.FormatHttpError(resp)
	}
	return listPortGroupResponse.PortGroups, nil
}

func (m *manager) ListPortGroupBySwitch(ctx context.Context, dvSwitchIdUri string) ([]PortGroup, error) {
	var listPortGroupResponse ListPortGroupResponse
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().
		SetContext(ctx).
		SetResult(&listPortGroupResponse).
		Get(path.Join(dvSwitchIdUri, "portgroups"))
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, common.FormatHttpError(resp)
	}
	return listPortGroupResponse.PortGroups, nil
}

func (m *manager) ListDVSwitch(ctx context.Context) ([]DVSwitch, error) {
	var listDVSwitchResponse ListDVSwitchResponse
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().
		SetContext(ctx).
		SetResult(&listDVSwitchResponse).
		Get(fmt.Sprintf(dvSwitchUrl, m.siteUri))
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, common.FormatHttpError(resp)
	}
	return listDVSwitchResponse.DVSwitchs, nil
}

func (m *manager) ListPortGroupInUseIp(ctx context.Context, portGroupUrn string) ([]string, error) {
	var results []string
	var listVmResponse vm.ListVmResponse
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().
		SetContext(ctx).
		SetResult(&listVmResponse).
		Get(fmt.Sprintf(vmScopeUrl, m.siteUri, portGroupUrn))
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, common.FormatHttpError(resp)
	}
	for _, v := range listVmResponse.Vms {
		for _, nic := range v.VmConfig.Nics {
			if nic.Ip != "0.0.0.0" {
				results = append(results, nic.Ip)
			}
		}
	}
	return results, nil
}
