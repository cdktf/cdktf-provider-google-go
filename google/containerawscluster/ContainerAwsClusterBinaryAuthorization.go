// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerawscluster


type ContainerAwsClusterBinaryAuthorization struct {
	// Mode of operation for Binary Authorization policy evaluation. Possible values: DISABLED, PROJECT_SINGLETON_POLICY_ENFORCE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/container_aws_cluster#evaluation_mode ContainerAwsCluster#evaluation_mode}
	EvaluationMode *string `field:"optional" json:"evaluationMode" yaml:"evaluationMode"`
}

