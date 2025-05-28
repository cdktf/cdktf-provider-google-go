// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeesecurityprofilev2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeSecurityProfileV2Config struct {
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
	// The Apigee Organization associated with the Apigee Security Profile V2, in the format 'organizations/{{org_name}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/apigee_security_profile_v2#org_id ApigeeSecurityProfileV2#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// profile_assessment_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/apigee_security_profile_v2#profile_assessment_configs ApigeeSecurityProfileV2#profile_assessment_configs}
	ProfileAssessmentConfigs interface{} `field:"required" json:"profileAssessmentConfigs" yaml:"profileAssessmentConfigs"`
	// Resource ID of the security profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/apigee_security_profile_v2#profile_id ApigeeSecurityProfileV2#profile_id}
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
	// Description of the security profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/apigee_security_profile_v2#description ApigeeSecurityProfileV2#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/apigee_security_profile_v2#id ApigeeSecurityProfileV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/apigee_security_profile_v2#timeouts ApigeeSecurityProfileV2#timeouts}
	Timeouts *ApigeeSecurityProfileV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

