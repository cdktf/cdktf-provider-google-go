// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamworkloadidentitypool


type IamWorkloadIdentityPoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/iam_workload_identity_pool#create IamWorkloadIdentityPool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/iam_workload_identity_pool#delete IamWorkloadIdentityPool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/iam_workload_identity_pool#update IamWorkloadIdentityPool#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

