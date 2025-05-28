// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamworkforcepoolproviderkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamWorkforcePoolProviderKeyConfig struct {
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
	// key_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/iam_workforce_pool_provider_key#key_data IamWorkforcePoolProviderKey#key_data}
	KeyData *IamWorkforcePoolProviderKeyKeyData `field:"required" json:"keyData" yaml:"keyData"`
	// The ID to use for the key, which becomes the final component of the resource name.
	//
	// This value must be 4-32 characters, and may contain the characters [a-z0-9-].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/iam_workforce_pool_provider_key#key_id IamWorkforcePoolProviderKey#key_id}
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/iam_workforce_pool_provider_key#location IamWorkforcePoolProviderKey#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID of the provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/iam_workforce_pool_provider_key#provider_id IamWorkforcePoolProviderKey#provider_id}
	ProviderId *string `field:"required" json:"providerId" yaml:"providerId"`
	// The purpose of the key. Possible values: ["ENCRYPTION"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/iam_workforce_pool_provider_key#use IamWorkforcePoolProviderKey#use}
	Use *string `field:"required" json:"use" yaml:"use"`
	// The ID of the workforce pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/iam_workforce_pool_provider_key#workforce_pool_id IamWorkforcePoolProviderKey#workforce_pool_id}
	WorkforcePoolId *string `field:"required" json:"workforcePoolId" yaml:"workforcePoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/iam_workforce_pool_provider_key#id IamWorkforcePoolProviderKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/iam_workforce_pool_provider_key#timeouts IamWorkforcePoolProviderKey#timeouts}
	Timeouts *IamWorkforcePoolProviderKeyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

