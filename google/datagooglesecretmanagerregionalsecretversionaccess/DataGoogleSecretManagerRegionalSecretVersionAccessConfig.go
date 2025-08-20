// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagooglesecretmanagerregionalsecretversionaccess

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleSecretManagerRegionalSecretVersionAccessConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/data-sources/secret_manager_regional_secret_version_access#secret DataGoogleSecretManagerRegionalSecretVersionAccess#secret}.
	Secret *string `field:"required" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/data-sources/secret_manager_regional_secret_version_access#id DataGoogleSecretManagerRegionalSecretVersionAccess#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/data-sources/secret_manager_regional_secret_version_access#is_secret_data_base64 DataGoogleSecretManagerRegionalSecretVersionAccess#is_secret_data_base64}.
	IsSecretDataBase64 interface{} `field:"optional" json:"isSecretDataBase64" yaml:"isSecretDataBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/data-sources/secret_manager_regional_secret_version_access#location DataGoogleSecretManagerRegionalSecretVersionAccess#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/data-sources/secret_manager_regional_secret_version_access#project DataGoogleSecretManagerRegionalSecretVersionAccess#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/data-sources/secret_manager_regional_secret_version_access#version DataGoogleSecretManagerRegionalSecretVersionAccess#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

