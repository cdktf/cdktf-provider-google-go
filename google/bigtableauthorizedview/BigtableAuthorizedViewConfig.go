// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigtableauthorizedview

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigtableAuthorizedViewConfig struct {
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
	// The name of the Bigtable instance in which the authorized view belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/bigtable_authorized_view#instance_name BigtableAuthorizedView#instance_name}
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// The name of the authorized view.
	//
	// Must be 1-50 characters and must only contain hyphens, underscores, periods, letters and numbers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/bigtable_authorized_view#name BigtableAuthorizedView#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the Bigtable table in which the authorized view belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/bigtable_authorized_view#table_name BigtableAuthorizedView#table_name}
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// A field to make the authorized view protected against data loss i.e. when set to PROTECTED, deleting the authorized view, the table containing the authorized view, and the instance containing the authorized view would be prohibited. If not provided, currently deletion protection will be set to UNPROTECTED as it is the API default value. Note this field configs the deletion protection provided by the API in the backend, and should not be confused with Terraform-side deletion protection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/bigtable_authorized_view#deletion_protection BigtableAuthorizedView#deletion_protection}
	DeletionProtection *string `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/bigtable_authorized_view#id BigtableAuthorizedView#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the project in which the resource belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/bigtable_authorized_view#project BigtableAuthorizedView#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// subset_view block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/bigtable_authorized_view#subset_view BigtableAuthorizedView#subset_view}
	SubsetView *BigtableAuthorizedViewSubsetView `field:"optional" json:"subsetView" yaml:"subsetView"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/bigtable_authorized_view#timeouts BigtableAuthorizedView#timeouts}
	Timeouts *BigtableAuthorizedViewTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

