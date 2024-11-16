// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolNodeConfigHostMaintenancePolicy struct {
	// .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/container_node_pool#maintenance_interval ContainerNodePool#maintenance_interval}
	MaintenanceInterval *string `field:"required" json:"maintenanceInterval" yaml:"maintenanceInterval"`
}

