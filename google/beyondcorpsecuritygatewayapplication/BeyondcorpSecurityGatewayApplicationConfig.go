// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package beyondcorpsecuritygatewayapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BeyondcorpSecurityGatewayApplicationConfig struct {
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
	// User-settable Application resource ID.
	//
	// * Must start with a letter.
	// * Must contain between 4-63 characters from '/a-z-/'.
	// * Must end with a number or letter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/beyondcorp_security_gateway_application#application_id BeyondcorpSecurityGatewayApplication#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// endpoint_matchers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/beyondcorp_security_gateway_application#endpoint_matchers BeyondcorpSecurityGatewayApplication#endpoint_matchers}
	EndpointMatchers interface{} `field:"required" json:"endpointMatchers" yaml:"endpointMatchers"`
	// ID of the Security Gateway resource this belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/beyondcorp_security_gateway_application#security_gateway_id BeyondcorpSecurityGatewayApplication#security_gateway_id}
	SecurityGatewayId *string `field:"required" json:"securityGatewayId" yaml:"securityGatewayId"`
	// Optional. An arbitrary user-provided name for the Application resource. Cannot exceed 64 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/beyondcorp_security_gateway_application#display_name BeyondcorpSecurityGatewayApplication#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/beyondcorp_security_gateway_application#id BeyondcorpSecurityGatewayApplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/beyondcorp_security_gateway_application#project BeyondcorpSecurityGatewayApplication#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/beyondcorp_security_gateway_application#timeouts BeyondcorpSecurityGatewayApplication#timeouts}
	Timeouts *BeyondcorpSecurityGatewayApplicationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// upstreams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/beyondcorp_security_gateway_application#upstreams BeyondcorpSecurityGatewayApplication#upstreams}
	Upstreams interface{} `field:"optional" json:"upstreams" yaml:"upstreams"`
}

