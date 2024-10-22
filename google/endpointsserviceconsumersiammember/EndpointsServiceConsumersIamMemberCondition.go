// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package endpointsserviceconsumersiammember


type EndpointsServiceConsumersIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/endpoints_service_consumers_iam_member#expression EndpointsServiceConsumersIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/endpoints_service_consumers_iam_member#title EndpointsServiceConsumersIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/endpoints_service_consumers_iam_member#description EndpointsServiceConsumersIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

