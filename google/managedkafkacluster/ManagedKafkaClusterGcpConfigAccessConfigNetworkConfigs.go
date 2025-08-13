// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedkafkacluster


type ManagedKafkaClusterGcpConfigAccessConfigNetworkConfigs struct {
	// Name of the VPC subnet from which the cluster is accessible.
	//
	// Both broker and bootstrap server IP addresses and DNS entries are automatically created in the subnet. There can only be one subnet per network, and the subnet must be located in the same region as the cluster. The project may differ. The name of the subnet must be in the format 'projects/PROJECT_ID/regions/REGION/subnetworks/SUBNET'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/managed_kafka_cluster#subnet ManagedKafkaCluster#subnet}
	Subnet *string `field:"required" json:"subnet" yaml:"subnet"`
}

