// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeploydeploypolicy


type ClouddeployDeployPolicySelectors struct {
	// delivery_pipeline block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/clouddeploy_deploy_policy#delivery_pipeline ClouddeployDeployPolicy#delivery_pipeline}
	DeliveryPipeline *ClouddeployDeployPolicySelectorsDeliveryPipeline `field:"optional" json:"deliveryPipeline" yaml:"deliveryPipeline"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/clouddeploy_deploy_policy#target ClouddeployDeployPolicy#target}
	Target *ClouddeployDeployPolicySelectorsTarget `field:"optional" json:"target" yaml:"target"`
}

