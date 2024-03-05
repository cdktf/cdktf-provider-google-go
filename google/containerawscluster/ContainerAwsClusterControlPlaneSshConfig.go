// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerawscluster


type ContainerAwsClusterControlPlaneSshConfig struct {
	// The name of the EC2 key pair used to login into cluster machines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/container_aws_cluster#ec2_key_pair ContainerAwsCluster#ec2_key_pair}
	Ec2KeyPair *string `field:"required" json:"ec2KeyPair" yaml:"ec2KeyPair"`
}

