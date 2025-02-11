// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitypostureposturedeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityposturePostureDeploymentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The location of the resource, eg. global'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/securityposture_posture_deployment#location SecurityposturePostureDeployment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The parent of the resource, an organization. Format should be 'organizations/{organization_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/securityposture_posture_deployment#parent SecurityposturePostureDeployment#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// ID of the posture deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/securityposture_posture_deployment#posture_deployment_id SecurityposturePostureDeployment#posture_deployment_id}
	PostureDeploymentId *string `field:"required" json:"postureDeploymentId" yaml:"postureDeploymentId"`
	// Relative name of the posture which needs to be deployed. It should be in the format:   organizations/{organization_id}/locations/{location}/postures/{posture_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/securityposture_posture_deployment#posture_id SecurityposturePostureDeployment#posture_id}
	PostureId *string `field:"required" json:"postureId" yaml:"postureId"`
	// Revision_id the posture which needs to be deployed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/securityposture_posture_deployment#posture_revision_id SecurityposturePostureDeployment#posture_revision_id}
	PostureRevisionId *string `field:"required" json:"postureRevisionId" yaml:"postureRevisionId"`
	// The resource on which the posture should be deployed. This can be in one of the following formats: projects/{project_number}, folders/{folder_number}, organizations/{organization_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/securityposture_posture_deployment#target_resource SecurityposturePostureDeployment#target_resource}
	TargetResource *string `field:"required" json:"targetResource" yaml:"targetResource"`
	// Description of the posture deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/securityposture_posture_deployment#description SecurityposturePostureDeployment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/securityposture_posture_deployment#id SecurityposturePostureDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/securityposture_posture_deployment#timeouts SecurityposturePostureDeployment#timeouts}
	Timeouts *SecurityposturePostureDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

