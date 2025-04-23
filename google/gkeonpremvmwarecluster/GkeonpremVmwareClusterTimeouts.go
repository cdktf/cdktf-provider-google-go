// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster


type GkeonpremVmwareClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/gkeonprem_vmware_cluster#create GkeonpremVmwareCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/gkeonprem_vmware_cluster#delete GkeonpremVmwareCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/gkeonprem_vmware_cluster#update GkeonpremVmwareCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

