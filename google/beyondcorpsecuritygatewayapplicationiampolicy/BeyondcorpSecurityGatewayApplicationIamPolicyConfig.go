// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package beyondcorpsecuritygatewayapplicationiampolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BeyondcorpSecurityGatewayApplicationIamPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/beyondcorp_security_gateway_application_iam_policy#application_id BeyondcorpSecurityGatewayApplicationIamPolicy#application_id}.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/beyondcorp_security_gateway_application_iam_policy#policy_data BeyondcorpSecurityGatewayApplicationIamPolicy#policy_data}.
	PolicyData *string `field:"required" json:"policyData" yaml:"policyData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/beyondcorp_security_gateway_application_iam_policy#security_gateway_id BeyondcorpSecurityGatewayApplicationIamPolicy#security_gateway_id}.
	SecurityGatewayId *string `field:"required" json:"securityGatewayId" yaml:"securityGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/beyondcorp_security_gateway_application_iam_policy#id BeyondcorpSecurityGatewayApplicationIamPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/beyondcorp_security_gateway_application_iam_policy#project BeyondcorpSecurityGatewayApplicationIamPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

