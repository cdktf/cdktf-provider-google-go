// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deploymentmanagerdeployment


type DeploymentManagerDeploymentLabels struct {
	// Key for label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.5.0/docs/resources/deployment_manager_deployment#key DeploymentManagerDeployment#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Value of label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.5.0/docs/resources/deployment_manager_deployment#value DeploymentManagerDeployment#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

