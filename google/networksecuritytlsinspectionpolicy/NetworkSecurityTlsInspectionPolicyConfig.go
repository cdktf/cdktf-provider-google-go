// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecuritytlsinspectionpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkSecurityTlsInspectionPolicyConfig struct {
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
	// A CA pool resource used to issue interception certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/network_security_tls_inspection_policy#ca_pool NetworkSecurityTlsInspectionPolicy#ca_pool}
	CaPool *string `field:"required" json:"caPool" yaml:"caPool"`
	// Short name of the TlsInspectionPolicy resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/network_security_tls_inspection_policy#name NetworkSecurityTlsInspectionPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of custom TLS cipher suites selected.
	//
	// This field is valid only if the selected tls_feature_profile is CUSTOM. The compute.SslPoliciesService.ListAvailableFeatures method returns the set of features that can be specified in this list. Note that Secure Web Proxy does not yet honor this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/network_security_tls_inspection_policy#custom_tls_features NetworkSecurityTlsInspectionPolicy#custom_tls_features}
	CustomTlsFeatures *[]*string `field:"optional" json:"customTlsFeatures" yaml:"customTlsFeatures"`
	// Free-text description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/network_security_tls_inspection_policy#description NetworkSecurityTlsInspectionPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trustConfig.
	//
	// These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trustConfig will be accepted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/network_security_tls_inspection_policy#exclude_public_ca_set NetworkSecurityTlsInspectionPolicy#exclude_public_ca_set}
	ExcludePublicCaSet interface{} `field:"optional" json:"excludePublicCaSet" yaml:"excludePublicCaSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/network_security_tls_inspection_policy#id NetworkSecurityTlsInspectionPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The location of the tls inspection policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/network_security_tls_inspection_policy#location NetworkSecurityTlsInspectionPolicy#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Minimum TLS version that the firewall should use when negotiating connections with both clients and servers.
	//
	// If this is not set, then the default value is to allow the broadest set of clients and servers (TLS 1.0 or higher). Setting this to more restrictive values may improve security, but may also prevent the firewall from connecting to some clients or servers. Note that Secure Web Proxy does not yet honor this field. Possible values: ["TLS_VERSION_UNSPECIFIED", "TLS_1_0", "TLS_1_1", "TLS_1_2", "TLS_1_3"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/network_security_tls_inspection_policy#min_tls_version NetworkSecurityTlsInspectionPolicy#min_tls_version}
	MinTlsVersion *string `field:"optional" json:"minTlsVersion" yaml:"minTlsVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/network_security_tls_inspection_policy#project NetworkSecurityTlsInspectionPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/network_security_tls_inspection_policy#timeouts NetworkSecurityTlsInspectionPolicy#timeouts}
	Timeouts *NetworkSecurityTlsInspectionPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The selected Profile.
	//
	// If this is not set, then the default value is to allow the broadest set of clients and servers (\"PROFILE_COMPATIBLE\"). Setting this to more restrictive values may improve security, but may also prevent the TLS inspection proxy from connecting to some clients or servers. Note that Secure Web Proxy does not yet honor this field. Possible values: ["PROFILE_UNSPECIFIED", "PROFILE_COMPATIBLE", "PROFILE_MODERN", "PROFILE_RESTRICTED", "PROFILE_CUSTOM"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/network_security_tls_inspection_policy#tls_feature_profile NetworkSecurityTlsInspectionPolicy#tls_feature_profile}
	TlsFeatureProfile *string `field:"optional" json:"tlsFeatureProfile" yaml:"tlsFeatureProfile"`
	// A TrustConfig resource used when making a connection to the TLS server.
	//
	// This is a relative resource path following the form \"projects/{project}/locations/{location}/trustConfigs/{trust_config}\". This is necessary to intercept TLS connections to servers with certificates signed by a private CA or self-signed certificates. Trust config and the TLS inspection policy must be in the same region. Note that Secure Web Proxy does not yet honor this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/network_security_tls_inspection_policy#trust_config NetworkSecurityTlsInspectionPolicy#trust_config}
	TrustConfig *string `field:"optional" json:"trustConfig" yaml:"trustConfig"`
}

