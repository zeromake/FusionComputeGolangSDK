package cluster

import (
	"context"
	"fmt"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/client"
	"github.com/KubeOperator/FusionComputeGolangSDK/pkg/site"
	"log"
	"testing"
)

func TestManager_List(t *testing.T) {
	ctx := context.Background()
	c := client.NewFusionComputeClient("https://100.199.16.208:7443", "fit2cloud", "Huawei@1234")
	err := c.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer c.DisConnect(ctx)

	sm := site.NewManager(c)
	ss, err := sm.ListSite(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range ss {
		cm := NewManager(c, s.Uri)
		cs, err := cm.ListCluster(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(cs)
	}
}
