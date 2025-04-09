// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securesourcemanagerinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecureSourceManagerInstanceConfig struct {
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
	// The name for the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/secure_source_manager_instance#instance_id SecureSourceManagerInstance#instance_id}
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The location for the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/secure_source_manager_instance#location SecureSourceManagerInstance#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/secure_source_manager_instance#id SecureSourceManagerInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Customer-managed encryption key name, in the format projects/* /locations/* /keyRings/* /cryptoKeys/*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/secure_source_manager_instance#kms_key SecureSourceManagerInstance#kms_key}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/secure_source_manager_instance#labels SecureSourceManagerInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// private_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/secure_source_manager_instance#private_config SecureSourceManagerInstance#private_config}
	PrivateConfig *SecureSourceManagerInstancePrivateConfig `field:"optional" json:"privateConfig" yaml:"privateConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/secure_source_manager_instance#project SecureSourceManagerInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/secure_source_manager_instance#timeouts SecureSourceManagerInstance#timeouts}
	Timeouts *SecureSourceManagerInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// workforce_identity_federation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/secure_source_manager_instance#workforce_identity_federation_config SecureSourceManagerInstance#workforce_identity_federation_config}
	WorkforceIdentityFederationConfig *SecureSourceManagerInstanceWorkforceIdentityFederationConfig `field:"optional" json:"workforceIdentityFederationConfig" yaml:"workforceIdentityFederationConfig"`
}

