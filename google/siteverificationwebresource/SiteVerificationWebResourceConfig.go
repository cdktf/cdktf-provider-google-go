// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siteverificationwebresource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SiteVerificationWebResourceConfig struct {
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
	// site block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/site_verification_web_resource#site SiteVerificationWebResource#site}
	Site *SiteVerificationWebResourceSite `field:"required" json:"site" yaml:"site"`
	// The verification method for the Site Verification system to use to verify this site or domain.
	//
	// Possible values: ["ANALYTICS", "DNS_CNAME", "DNS_TXT", "FILE", "META", "TAG_MANAGER"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/site_verification_web_resource#verification_method SiteVerificationWebResource#verification_method}
	VerificationMethod *string `field:"required" json:"verificationMethod" yaml:"verificationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/site_verification_web_resource#id SiteVerificationWebResource#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/site_verification_web_resource#timeouts SiteVerificationWebResource#timeouts}
	Timeouts *SiteVerificationWebResourceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

