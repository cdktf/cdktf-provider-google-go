// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagoogleiamtestablepermissions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleIamTestablePermissionsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/data-sources/iam_testable_permissions#full_resource_name DataGoogleIamTestablePermissions#full_resource_name}.
	FullResourceName *string `field:"required" json:"fullResourceName" yaml:"fullResourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/data-sources/iam_testable_permissions#custom_support_level DataGoogleIamTestablePermissions#custom_support_level}.
	CustomSupportLevel *string `field:"optional" json:"customSupportLevel" yaml:"customSupportLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/data-sources/iam_testable_permissions#id DataGoogleIamTestablePermissions#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/data-sources/iam_testable_permissions#stages DataGoogleIamTestablePermissions#stages}.
	Stages *[]*string `field:"optional" json:"stages" yaml:"stages"`
}

