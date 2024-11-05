// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package biglaketable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BiglakeTableConfig struct {
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
	// Output only. The name of the Table. Format: projects/{project_id_or_number}/locations/{locationId}/catalogs/{catalogId}/databases/{databaseId}/tables/{tableId}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/biglake_table#name BiglakeTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The id of the parent database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/biglake_table#database BiglakeTable#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// hive_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/biglake_table#hive_options BiglakeTable#hive_options}
	HiveOptions *BiglakeTableHiveOptions `field:"optional" json:"hiveOptions" yaml:"hiveOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/biglake_table#id BiglakeTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/biglake_table#timeouts BiglakeTable#timeouts}
	Timeouts *BiglakeTableTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The database type. Possible values: ["HIVE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/biglake_table#type BiglakeTable#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

