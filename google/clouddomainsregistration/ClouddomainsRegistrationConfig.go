// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddomainsregistration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClouddomainsRegistrationConfig struct {
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
	// contact_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/clouddomains_registration#contact_settings ClouddomainsRegistration#contact_settings}
	ContactSettings *ClouddomainsRegistrationContactSettings `field:"required" json:"contactSettings" yaml:"contactSettings"`
	// Required. The domain name. Unicode domain names must be expressed in Punycode format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/clouddomains_registration#domain_name ClouddomainsRegistration#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/clouddomains_registration#location ClouddomainsRegistration#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// yearly_price block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/clouddomains_registration#yearly_price ClouddomainsRegistration#yearly_price}
	YearlyPrice *ClouddomainsRegistrationYearlyPrice `field:"required" json:"yearlyPrice" yaml:"yearlyPrice"`
	// The list of contact notices that the caller acknowledges. Possible value is PUBLIC_CONTACT_DATA_ACKNOWLEDGEMENT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/clouddomains_registration#contact_notices ClouddomainsRegistration#contact_notices}
	ContactNotices *[]*string `field:"optional" json:"contactNotices" yaml:"contactNotices"`
	// dns_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/clouddomains_registration#dns_settings ClouddomainsRegistration#dns_settings}
	DnsSettings *ClouddomainsRegistrationDnsSettings `field:"optional" json:"dnsSettings" yaml:"dnsSettings"`
	// The list of domain notices that you acknowledge. Possible value is HSTS_PRELOADED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/clouddomains_registration#domain_notices ClouddomainsRegistration#domain_notices}
	DomainNotices *[]*string `field:"optional" json:"domainNotices" yaml:"domainNotices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/clouddomains_registration#id ClouddomainsRegistration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of labels associated with the Registration.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/clouddomains_registration#labels ClouddomainsRegistration#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// management_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/clouddomains_registration#management_settings ClouddomainsRegistration#management_settings}
	ManagementSettings *ClouddomainsRegistrationManagementSettings `field:"optional" json:"managementSettings" yaml:"managementSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/clouddomains_registration#project ClouddomainsRegistration#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/clouddomains_registration#timeouts ClouddomainsRegistration#timeouts}
	Timeouts *ClouddomainsRegistrationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

