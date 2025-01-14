// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeenvironmentkeyvaluemapsentries

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeEnvironmentKeyvaluemapsEntriesConfig struct {
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
	// The Apigee environment keyvalumaps Id associated with the Apigee environment, in the format 'organizations/{{org_name}}/environments/{{env_name}}/keyvaluemaps/{{keyvaluemap_name}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/apigee_environment_keyvaluemaps_entries#env_keyvaluemap_id ApigeeEnvironmentKeyvaluemapsEntries#env_keyvaluemap_id}
	EnvKeyvaluemapId *string `field:"required" json:"envKeyvaluemapId" yaml:"envKeyvaluemapId"`
	// Required. Resource URI that can be used to identify the scope of the key value map entries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/apigee_environment_keyvaluemaps_entries#name ApigeeEnvironmentKeyvaluemapsEntries#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Required. Data or payload that is being retrieved and associated with the unique key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/apigee_environment_keyvaluemaps_entries#value ApigeeEnvironmentKeyvaluemapsEntries#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/apigee_environment_keyvaluemaps_entries#id ApigeeEnvironmentKeyvaluemapsEntries#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/apigee_environment_keyvaluemaps_entries#timeouts ApigeeEnvironmentKeyvaluemapsEntries#timeouts}
	Timeouts *ApigeeEnvironmentKeyvaluemapsEntriesTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

