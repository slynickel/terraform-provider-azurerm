package client

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-05-01/network"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/common"
)

type Client struct {
	LoadBalancersClient      *network.LoadBalancersClient
	LoadBalancingRulesClient *network.LoadBalancerLoadBalancingRulesClient
}

func NewClient(o *common.ClientOptions) *Client {
	loadBalancersClient := network.NewLoadBalancersClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&loadBalancersClient.Client, o.ResourceManagerAuthorizer)

	loadBalancingRulesClient := network.NewLoadBalancerLoadBalancingRulesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&loadBalancingRulesClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		LoadBalancersClient:      &loadBalancersClient,
		LoadBalancingRulesClient: &loadBalancingRulesClient,
	}
}
