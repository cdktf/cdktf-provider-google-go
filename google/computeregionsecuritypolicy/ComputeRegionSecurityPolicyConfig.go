// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionsecuritypolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionSecurityPolicyConfig struct {
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
	// Name of the resource.
	//
	// Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035.
	// Specifically, the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy#name ComputeRegionSecurityPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// ddos_protection_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy#ddos_protection_config ComputeRegionSecurityPolicy#ddos_protection_config}
	DdosProtectionConfig *ComputeRegionSecurityPolicyDdosProtectionConfig `field:"optional" json:"ddosProtectionConfig" yaml:"ddosProtectionConfig"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy#description ComputeRegionSecurityPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy#id ComputeRegionSecurityPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy#project ComputeRegionSecurityPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The Region in which the created Region Security Policy should reside.
	//
	// If it is not provided, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy#region ComputeRegionSecurityPolicy#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy#rules ComputeRegionSecurityPolicy#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy#timeouts ComputeRegionSecurityPolicy#timeouts}
	Timeouts *ComputeRegionSecurityPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The type indicates the intended use of the security policy.
	//
	// - CLOUD_ARMOR: Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers.
	// - CLOUD_ARMOR_EDGE: Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache.
	// - CLOUD_ARMOR_NETWORK: Cloud Armor network policies can be configured to filter packets targeting network load balancing resources such as backend services, target pools, target instances, and instances with external IPs. They filter requests before the request is served from the application.
	// This field can be set only at resource creation time. Possible values: ["CLOUD_ARMOR", "CLOUD_ARMOR_EDGE", "CLOUD_ARMOR_NETWORK"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy#type ComputeRegionSecurityPolicy#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// user_defined_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy#user_defined_fields ComputeRegionSecurityPolicy#user_defined_fields}
	UserDefinedFields interface{} `field:"optional" json:"userDefinedFields" yaml:"userDefinedFields"`
}

