// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagooglecomputeinstanceguestattributes

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleComputeInstanceGuestAttributesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/data-sources/compute_instance_guest_attributes#name DataGoogleComputeInstanceGuestAttributes#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/data-sources/compute_instance_guest_attributes#id DataGoogleComputeInstanceGuestAttributes#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/data-sources/compute_instance_guest_attributes#project DataGoogleComputeInstanceGuestAttributes#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/data-sources/compute_instance_guest_attributes#query_path DataGoogleComputeInstanceGuestAttributes#query_path}.
	QueryPath *string `field:"optional" json:"queryPath" yaml:"queryPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/data-sources/compute_instance_guest_attributes#region DataGoogleComputeInstanceGuestAttributes#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/data-sources/compute_instance_guest_attributes#variable_key DataGoogleComputeInstanceGuestAttributes#variable_key}.
	VariableKey *string `field:"optional" json:"variableKey" yaml:"variableKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/data-sources/compute_instance_guest_attributes#zone DataGoogleComputeInstanceGuestAttributes#zone}.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

