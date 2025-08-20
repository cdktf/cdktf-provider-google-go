// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagementconnectivitytest


type NetworkManagementConnectivityTestDestination struct {
	// A Cloud SQL instance URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_management_connectivity_test#cloud_sql_instance NetworkManagementConnectivityTest#cloud_sql_instance}
	CloudSqlInstance *string `field:"optional" json:"cloudSqlInstance" yaml:"cloudSqlInstance"`
	// Forwarding rule URI. Forwarding rules are frontends for load balancers, PSC endpoints, and Protocol Forwarding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_management_connectivity_test#forwarding_rule NetworkManagementConnectivityTest#forwarding_rule}
	ForwardingRule *string `field:"optional" json:"forwardingRule" yaml:"forwardingRule"`
	// A DNS endpoint of Google Kubernetes Engine cluster control plane.
	//
	// Requires gke_master_cluster to be set, can't be used simultaneoulsly with
	// ip_address or network. Applicable only to destination endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_management_connectivity_test#fqdn NetworkManagementConnectivityTest#fqdn}
	Fqdn *string `field:"optional" json:"fqdn" yaml:"fqdn"`
	// A cluster URI for Google Kubernetes Engine cluster control plane.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_management_connectivity_test#gke_master_cluster NetworkManagementConnectivityTest#gke_master_cluster}
	GkeMasterCluster *string `field:"optional" json:"gkeMasterCluster" yaml:"gkeMasterCluster"`
	// A Compute Engine instance URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_management_connectivity_test#instance NetworkManagementConnectivityTest#instance}
	Instance *string `field:"optional" json:"instance" yaml:"instance"`
	// The IP address of the endpoint, which can be an external or internal IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_management_connectivity_test#ip_address NetworkManagementConnectivityTest#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// A VPC network URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_management_connectivity_test#network NetworkManagementConnectivityTest#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The IP protocol port of the endpoint. Only applicable when protocol is TCP or UDP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_management_connectivity_test#port NetworkManagementConnectivityTest#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Project ID where the endpoint is located.
	//
	// The project ID can be derived from the URI if you provide a endpoint or
	// network URI.
	// The following are two cases where you may need to provide the project ID:
	// 1. Only the IP address is specified, and the IP address is within a Google
	// Cloud project.
	// 2. When you are using Shared VPC and the IP address that you provide is
	// from the service project. In this case, the network that the IP address
	// resides in is defined in the host project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_management_connectivity_test#project_id NetworkManagementConnectivityTest#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// A Redis Cluster URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_management_connectivity_test#redis_cluster NetworkManagementConnectivityTest#redis_cluster}
	RedisCluster *string `field:"optional" json:"redisCluster" yaml:"redisCluster"`
	// A Redis Instance URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/network_management_connectivity_test#redis_instance NetworkManagementConnectivityTest#redis_instance}
	RedisInstance *string `field:"optional" json:"redisInstance" yaml:"redisInstance"`
}

