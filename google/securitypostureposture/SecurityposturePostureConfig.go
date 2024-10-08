// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitypostureposture

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityposturePostureConfig struct {
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
	// Location of the resource, eg: global.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/securityposture_posture#location SecurityposturePosture#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The parent of the resource, an organization. Format should be 'organizations/{organization_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/securityposture_posture#parent SecurityposturePosture#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// policy_sets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/securityposture_posture#policy_sets SecurityposturePosture#policy_sets}
	PolicySets interface{} `field:"required" json:"policySets" yaml:"policySets"`
	// Id of the posture. It is an immutable field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/securityposture_posture#posture_id SecurityposturePosture#posture_id}
	PostureId *string `field:"required" json:"postureId" yaml:"postureId"`
	// State of the posture.
	//
	// Update to state field should not be triggered along with
	// with other field updates. Possible values: ["DEPRECATED", "DRAFT", "ACTIVE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/securityposture_posture#state SecurityposturePosture#state}
	State *string `field:"required" json:"state" yaml:"state"`
	// Description of the posture.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/securityposture_posture#description SecurityposturePosture#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/securityposture_posture#id SecurityposturePosture#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/securityposture_posture#timeouts SecurityposturePosture#timeouts}
	Timeouts *SecurityposturePostureTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

