// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodeConfigTaint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.84.0/docs/resources/container_cluster#effect ContainerCluster#effect}.
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.84.0/docs/resources/container_cluster#key ContainerCluster#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.84.0/docs/resources/container_cluster#value ContainerCluster#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

