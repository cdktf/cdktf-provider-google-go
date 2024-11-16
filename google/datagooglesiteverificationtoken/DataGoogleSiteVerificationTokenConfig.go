// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagooglesiteverificationtoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleSiteVerificationTokenConfig struct {
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
	// The site identifier.
	//
	// If the type is set to SITE, the identifier is a URL. If the type is
	// set to INET_DOMAIN, the identifier is a domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/data-sources/site_verification_token#identifier DataGoogleSiteVerificationToken#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// The type of resource to be verified, either a domain or a web site. Possible values: ["INET_DOMAIN", "SITE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/data-sources/site_verification_token#type DataGoogleSiteVerificationToken#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The verification method for the Site Verification system to use to verify this site or domain.
	//
	// Possible values: ["ANALYTICS", "DNS_CNAME", "DNS_TXT", "FILE", "META", "TAG_MANAGER"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/data-sources/site_verification_token#verification_method DataGoogleSiteVerificationToken#verification_method}
	VerificationMethod *string `field:"required" json:"verificationMethod" yaml:"verificationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/data-sources/site_verification_token#id DataGoogleSiteVerificationToken#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/data-sources/site_verification_token#timeouts DataGoogleSiteVerificationToken#timeouts}
	Timeouts *DataGoogleSiteVerificationTokenTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

