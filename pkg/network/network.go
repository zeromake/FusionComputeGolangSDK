package network

import (
	"context"
	"encoding/json"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/common"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/vm"
	"path"
	"strings"
)

const (
	siteMask     = "<site_uri>"
	dvSwitchUrl  = "<site_uri>/dvswitchs"
	vmScopeUrl   = "<site_uri>/vms?scope=<resource_urn>"
	portGroupUrl = "<site_uri>/portgroups"
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
	var portGroups []PortGroup
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().SetContext(ctx).Get(strings.Replace(portGroupUrl, siteMask, m.siteUri, -1))
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		var listPortGroupResponse ListPortGroupResponse
		err := json.Unmarshal(resp.Body(), &listPortGroupResponse)
		if err != nil {
			return nil, err
		}
		portGroups = listPortGroupResponse.PortGroups
	} else {
		return nil, common.FormatHttpError(resp)
	}
	return portGroups, nil
}

func (m *manager) ListPortGroupBySwitch(ctx context.Context, dvSwitchIdUri string) ([]PortGroup, error) {

	var portGroups []PortGroup
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().SetContext(ctx).Get(path.Join(dvSwitchIdUri, "portgroups"))
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		var listPortGroupResponse ListPortGroupResponse
		err := json.Unmarshal(resp.Body(), &listPortGroupResponse)
		if err != nil {
			return nil, err
		}
		portGroups = listPortGroupResponse.PortGroups
	} else {
		return nil, common.FormatHttpError(resp)
	}
	return portGroups, nil
}

func (m *manager) ListDVSwitch(ctx context.Context) ([]DVSwitch, error) {
	var dvSwitchs []DVSwitch
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().SetContext(ctx).Get(strings.Replace(dvSwitchUrl, siteMask, m.siteUri, -1))
	if err != nil {
		return nil, err
	}
	if resp.IsSuccess() {
		var listDVSwitchResponse ListDVSwitchResponse
		err := json.Unmarshal(resp.Body(), &listDVSwitchResponse)
		if err != nil {
			return nil, err
		}
		dvSwitchs = listDVSwitchResponse.DVSwitchs
	} else {
		return nil, common.FormatHttpError(resp)
	}
	return dvSwitchs, nil
}

func (m *manager) ListPortGroupInUseIp(ctx context.Context, portGroupUrn string) ([]string, error) {
	var results []string
	api, err := m.client.GetApiClient()
	if err != nil {
		return nil, err
	}
	resp, err := api.R().SetContext(ctx).Get(strings.Replace(strings.Replace(vmScopeUrl, siteMask, m.siteUri, -1), "<resource_urn>", portGroupUrn, -1))
	if err != nil {
		return nil, err
	}
	var vms []vm.Vm
	if resp.IsSuccess() {
		var listVmResponse vm.ListVmResponse
		err := json.Unmarshal(resp.Body(), &listVmResponse)
		if err != nil {
			return nil, err
		}
		vms = listVmResponse.Vms
	}
	for _, v := range vms {
		for _, nic := range v.VmConfig.Nics {
			if nic.Ip != "0.0.0.0" {
				results = append(results, nic.Ip)
			}
		}
	}
	return results, nil
}
