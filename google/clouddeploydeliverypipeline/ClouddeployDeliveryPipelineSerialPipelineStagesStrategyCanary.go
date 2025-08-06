// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeploydeliverypipeline


type ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanary struct {
	// canary_deployment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/clouddeploy_delivery_pipeline#canary_deployment ClouddeployDeliveryPipeline#canary_deployment}
	CanaryDeployment *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment `field:"optional" json:"canaryDeployment" yaml:"canaryDeployment"`
	// custom_canary_deployment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/clouddeploy_delivery_pipeline#custom_canary_deployment ClouddeployDeliveryPipeline#custom_canary_deployment}
	CustomCanaryDeployment *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeployment `field:"optional" json:"customCanaryDeployment" yaml:"customCanaryDeployment"`
	// runtime_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/clouddeploy_delivery_pipeline#runtime_config ClouddeployDeliveryPipeline#runtime_config}
	RuntimeConfig *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfig `field:"optional" json:"runtimeConfig" yaml:"runtimeConfig"`
}

