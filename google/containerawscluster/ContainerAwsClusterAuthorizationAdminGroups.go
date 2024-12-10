// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerawscluster


type ContainerAwsClusterAuthorizationAdminGroups struct {
	// The name of the group, e.g. `my-group@domain.com`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/container_aws_cluster#group ContainerAwsCluster#group}
	Group *string `field:"required" json:"group" yaml:"group"`
}

