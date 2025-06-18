// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmsekmconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KmsEkmConnectionConfig struct {
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
	// The location for the EkmConnection. A full list of valid locations can be found by running 'gcloud kms locations list'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/kms_ekm_connection#location KmsEkmConnection#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name for the EkmConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/kms_ekm_connection#name KmsEkmConnection#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// service_resolvers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/kms_ekm_connection#service_resolvers KmsEkmConnection#service_resolvers}
	ServiceResolvers interface{} `field:"required" json:"serviceResolvers" yaml:"serviceResolvers"`
	// Optional.
	//
	// Identifies the EKM Crypto Space that this EkmConnection maps to. Note: This field is required if KeyManagementMode is CLOUD_KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/kms_ekm_connection#crypto_space_path KmsEkmConnection#crypto_space_path}
	CryptoSpacePath *string `field:"optional" json:"cryptoSpacePath" yaml:"cryptoSpacePath"`
	// Optional. Etag of the currently stored EkmConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/kms_ekm_connection#etag KmsEkmConnection#etag}
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/kms_ekm_connection#id KmsEkmConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional.
	//
	// Describes who can perform control plane operations on the EKM. If unset, this defaults to MANUAL Default value: "MANUAL" Possible values: ["MANUAL", "CLOUD_KMS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/kms_ekm_connection#key_management_mode KmsEkmConnection#key_management_mode}
	KeyManagementMode *string `field:"optional" json:"keyManagementMode" yaml:"keyManagementMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/kms_ekm_connection#project KmsEkmConnection#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/kms_ekm_connection#timeouts KmsEkmConnection#timeouts}
	Timeouts *KmsEkmConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

