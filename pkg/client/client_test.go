package client

import (
	"context"
	"log"
	"testing"
)

func TestFusionComputeClient_Connect(t *testing.T) {
	ctx := context.Background()
	c := NewFusionComputeClient("https://100.199.16.208:7443", "kubeoperator", "Calong@2015")
	err := c.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = c.DisConnect(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
