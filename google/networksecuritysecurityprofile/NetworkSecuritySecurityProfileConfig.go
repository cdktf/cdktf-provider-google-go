// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecuritysecurityprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkSecuritySecurityProfileConfig struct {
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
	// The name of the security profile resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile#name NetworkSecuritySecurityProfile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of security profile. Possible values: ["THREAT_PREVENTION", "CUSTOM_MIRRORING", "CUSTOM_INTERCEPT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile#type NetworkSecuritySecurityProfile#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// custom_intercept_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile#custom_intercept_profile NetworkSecuritySecurityProfile#custom_intercept_profile}
	CustomInterceptProfile *NetworkSecuritySecurityProfileCustomInterceptProfile `field:"optional" json:"customInterceptProfile" yaml:"customInterceptProfile"`
	// custom_mirroring_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile#custom_mirroring_profile NetworkSecuritySecurityProfile#custom_mirroring_profile}
	CustomMirroringProfile *NetworkSecuritySecurityProfileCustomMirroringProfile `field:"optional" json:"customMirroringProfile" yaml:"customMirroringProfile"`
	// An optional description of the security profile. The Max length is 512 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile#description NetworkSecuritySecurityProfile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile#id NetworkSecuritySecurityProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A map of key/value label pairs to assign to the resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile#labels NetworkSecuritySecurityProfile#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location of the security profile. The default value is 'global'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile#location NetworkSecuritySecurityProfile#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The name of the parent this security profile belongs to. Format: organizations/{organization_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile#parent NetworkSecuritySecurityProfile#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// threat_prevention_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile#threat_prevention_profile NetworkSecuritySecurityProfile#threat_prevention_profile}
	ThreatPreventionProfile *NetworkSecuritySecurityProfileThreatPreventionProfile `field:"optional" json:"threatPreventionProfile" yaml:"threatPreventionProfile"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile#timeouts NetworkSecuritySecurityProfile#timeouts}
	Timeouts *NetworkSecuritySecurityProfileTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

