// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareenginecluster


type VmwareengineClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/vmwareengine_cluster#create VmwareengineCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/vmwareengine_cluster#delete VmwareengineCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/vmwareengine_cluster#update VmwareengineCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

