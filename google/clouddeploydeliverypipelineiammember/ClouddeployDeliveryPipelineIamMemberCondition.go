// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeploydeliverypipelineiammember


type ClouddeployDeliveryPipelineIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/clouddeploy_delivery_pipeline_iam_member#expression ClouddeployDeliveryPipelineIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/clouddeploy_delivery_pipeline_iam_member#title ClouddeployDeliveryPipelineIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/clouddeploy_delivery_pipeline_iam_member#description ClouddeployDeliveryPipelineIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

