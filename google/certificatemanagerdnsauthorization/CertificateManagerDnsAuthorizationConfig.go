// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certificatemanagerdnsauthorization

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CertificateManagerDnsAuthorizationConfig struct {
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
	// A domain which is being authorized.
	//
	// A DnsAuthorization resource covers a
	// single domain and its wildcard, e.g. authorization for "example.com" can
	// be used to issue certificates for "example.com" and "*.example.com".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/certificate_manager_dns_authorization#domain CertificateManagerDnsAuthorization#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// Name of the resource;
	//
	// provided by the client when the resource is created.
	// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
	// and all following characters must be a dash, underscore, letter or digit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/certificate_manager_dns_authorization#name CertificateManagerDnsAuthorization#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A human-readable description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/certificate_manager_dns_authorization#description CertificateManagerDnsAuthorization#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/certificate_manager_dns_authorization#id CertificateManagerDnsAuthorization#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of label tags associated with the DNS Authorization resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/certificate_manager_dns_authorization#labels CertificateManagerDnsAuthorization#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The Certificate Manager location. If not specified, "global" is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/certificate_manager_dns_authorization#location CertificateManagerDnsAuthorization#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/certificate_manager_dns_authorization#project CertificateManagerDnsAuthorization#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/certificate_manager_dns_authorization#timeouts CertificateManagerDnsAuthorization#timeouts}
	Timeouts *CertificateManagerDnsAuthorizationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// type of DNS authorization.
	//
	// If unset during the resource creation, FIXED_RECORD will
	// be used for global resources, and PER_PROJECT_RECORD will be used for other locations.
	//
	// FIXED_RECORD DNS authorization uses DNS-01 validation method
	//
	// PER_PROJECT_RECORD DNS authorization allows for independent management
	// of Google-managed certificates with DNS authorization across multiple
	// projects. Possible values: ["FIXED_RECORD", "PER_PROJECT_RECORD"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/certificate_manager_dns_authorization#type CertificateManagerDnsAuthorization#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

