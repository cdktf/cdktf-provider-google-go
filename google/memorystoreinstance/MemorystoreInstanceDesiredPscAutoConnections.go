// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memorystoreinstance


type MemorystoreInstanceDesiredPscAutoConnections struct {
	// Required. The consumer network where the IP address resides, in the form of projects/{project_id}/global/networks/{network_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/memorystore_instance#network MemorystoreInstance#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Required. The consumer project_id where the forwarding rule is created from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/memorystore_instance#project_id MemorystoreInstance#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

