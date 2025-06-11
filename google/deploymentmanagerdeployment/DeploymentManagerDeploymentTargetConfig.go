// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deploymentmanagerdeployment


type DeploymentManagerDeploymentTargetConfig struct {
	// The full YAML contents of your configuration file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/deployment_manager_deployment#content DeploymentManagerDeployment#content}
	Content *string `field:"required" json:"content" yaml:"content"`
}

