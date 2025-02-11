// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecuritysecurityprofilegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkSecuritySecurityProfileGroupConfig struct {
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
	// The name of the security profile group resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile_group#name NetworkSecuritySecurityProfileGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Reference to a SecurityProfile with the CustomIntercept configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile_group#custom_intercept_profile NetworkSecuritySecurityProfileGroup#custom_intercept_profile}
	CustomInterceptProfile *string `field:"optional" json:"customInterceptProfile" yaml:"customInterceptProfile"`
	// Reference to a SecurityProfile with the custom mirroring configuration for the SecurityProfileGroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile_group#custom_mirroring_profile NetworkSecuritySecurityProfileGroup#custom_mirroring_profile}
	CustomMirroringProfile *string `field:"optional" json:"customMirroringProfile" yaml:"customMirroringProfile"`
	// An optional description of the profile. The Max length is 512 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile_group#description NetworkSecuritySecurityProfileGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile_group#id NetworkSecuritySecurityProfileGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A map of key/value label pairs to assign to the resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile_group#labels NetworkSecuritySecurityProfileGroup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location of the security profile group. The default value is 'global'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile_group#location NetworkSecuritySecurityProfileGroup#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The name of the parent this security profile group belongs to. Format: organizations/{organization_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile_group#parent NetworkSecuritySecurityProfileGroup#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// Reference to a SecurityProfile with the threat prevention configuration for the SecurityProfileGroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile_group#threat_prevention_profile NetworkSecuritySecurityProfileGroup#threat_prevention_profile}
	ThreatPreventionProfile *string `field:"optional" json:"threatPreventionProfile" yaml:"threatPreventionProfile"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/network_security_security_profile_group#timeouts NetworkSecuritySecurityProfileGroup#timeouts}
	Timeouts *NetworkSecuritySecurityProfileGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

