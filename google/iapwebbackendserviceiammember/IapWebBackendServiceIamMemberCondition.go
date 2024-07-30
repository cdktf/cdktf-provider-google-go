// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iapwebbackendserviceiammember


type IapWebBackendServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/iap_web_backend_service_iam_member#expression IapWebBackendServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/iap_web_backend_service_iam_member#title IapWebBackendServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/iap_web_backend_service_iam_member#description IapWebBackendServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

