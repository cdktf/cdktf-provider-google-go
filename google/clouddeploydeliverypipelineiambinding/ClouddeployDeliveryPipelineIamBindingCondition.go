// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeploydeliverypipelineiambinding


type ClouddeployDeliveryPipelineIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/clouddeploy_delivery_pipeline_iam_binding#expression ClouddeployDeliveryPipelineIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/clouddeploy_delivery_pipeline_iam_binding#title ClouddeployDeliveryPipelineIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/clouddeploy_delivery_pipeline_iam_binding#description ClouddeployDeliveryPipelineIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

