// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerawscluster


type ContainerAwsClusterAuthorizationAdminUsers struct {
	// The name of the user, e.g. `my-gcp-id@gmail.com`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/container_aws_cluster#username ContainerAwsCluster#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

