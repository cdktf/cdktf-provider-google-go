// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigtableschemabundle

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigtableSchemaBundleConfig struct {
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
	// proto_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/bigtable_schema_bundle#proto_schema BigtableSchemaBundle#proto_schema}
	ProtoSchema *BigtableSchemaBundleProtoSchema `field:"required" json:"protoSchema" yaml:"protoSchema"`
	// The unique name of the schema bundle in the form '[_a-zA-Z0-9][-_.a-zA-Z0-9]*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/bigtable_schema_bundle#schema_bundle_id BigtableSchemaBundle#schema_bundle_id}
	SchemaBundleId *string `field:"required" json:"schemaBundleId" yaml:"schemaBundleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/bigtable_schema_bundle#id BigtableSchemaBundle#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If true, allow backwards incompatible changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/bigtable_schema_bundle#ignore_warnings BigtableSchemaBundle#ignore_warnings}
	IgnoreWarnings interface{} `field:"optional" json:"ignoreWarnings" yaml:"ignoreWarnings"`
	// The name of the instance to create the schema bundle within.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/bigtable_schema_bundle#instance BigtableSchemaBundle#instance}
	Instance *string `field:"optional" json:"instance" yaml:"instance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/bigtable_schema_bundle#project BigtableSchemaBundle#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The name of the table to create the schema bundle within.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/bigtable_schema_bundle#table BigtableSchemaBundle#table}
	Table *string `field:"optional" json:"table" yaml:"table"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/bigtable_schema_bundle#timeouts BigtableSchemaBundle#timeouts}
	Timeouts *BigtableSchemaBundleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

