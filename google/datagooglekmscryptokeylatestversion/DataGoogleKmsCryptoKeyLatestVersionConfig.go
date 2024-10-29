// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagooglekmscryptokeylatestversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleKmsCryptoKeyLatestVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/data-sources/kms_crypto_key_latest_version#crypto_key DataGoogleKmsCryptoKeyLatestVersion#crypto_key}.
	CryptoKey *string `field:"required" json:"cryptoKey" yaml:"cryptoKey"`
	// The filter argument is used to add a filter query parameter that limits which type of cryptoKeyVersion is retrieved as the latest by the data source: ?filter={{filter}}. When no value is provided there is no filtering.
	//
	// Example filter values if filtering on state.
	//
	// * "state:ENABLED" will retrieve the latest cryptoKeyVersion that has the state "ENABLED".
	//
	// [See the documentation about using filters](https://cloud.google.com/kms/docs/sorting-and-filtering)
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/data-sources/kms_crypto_key_latest_version#filter DataGoogleKmsCryptoKeyLatestVersion#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/data-sources/kms_crypto_key_latest_version#id DataGoogleKmsCryptoKeyLatestVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

