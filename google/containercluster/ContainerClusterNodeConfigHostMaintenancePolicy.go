// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodeConfigHostMaintenancePolicy struct {
	// .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/container_cluster#maintenance_interval ContainerCluster#maintenance_interval}
	MaintenanceInterval *string `field:"required" json:"maintenanceInterval" yaml:"maintenanceInterval"`
}

