// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagoogleiamworkloadidentitypool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleIamWorkloadIdentityPoolConfig struct {
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
	// The ID to use for the pool, which becomes the final component of the resource name.
	//
	// This
	// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
	// 'gcp-' is reserved for use by Google, and may not be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/data-sources/iam_workload_identity_pool#workload_identity_pool_id DataGoogleIamWorkloadIdentityPool#workload_identity_pool_id}
	WorkloadIdentityPoolId *string `field:"required" json:"workloadIdentityPoolId" yaml:"workloadIdentityPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/data-sources/iam_workload_identity_pool#id DataGoogleIamWorkloadIdentityPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/data-sources/iam_workload_identity_pool#project DataGoogleIamWorkloadIdentityPool#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

