// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerawscluster


type ContainerAwsClusterFleet struct {
	// The number of the Fleet host project where this cluster will be registered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.15.0/docs/resources/container_aws_cluster#project ContainerAwsCluster#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

