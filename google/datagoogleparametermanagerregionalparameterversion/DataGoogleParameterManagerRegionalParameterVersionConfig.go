// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagoogleparametermanagerregionalparameterversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleParameterManagerRegionalParameterVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/data-sources/parameter_manager_regional_parameter_version#parameter DataGoogleParameterManagerRegionalParameterVersion#parameter}.
	Parameter *string `field:"required" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/data-sources/parameter_manager_regional_parameter_version#parameter_version_id DataGoogleParameterManagerRegionalParameterVersion#parameter_version_id}.
	ParameterVersionId *string `field:"required" json:"parameterVersionId" yaml:"parameterVersionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/data-sources/parameter_manager_regional_parameter_version#id DataGoogleParameterManagerRegionalParameterVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/data-sources/parameter_manager_regional_parameter_version#location DataGoogleParameterManagerRegionalParameterVersion#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/data-sources/parameter_manager_regional_parameter_version#project DataGoogleParameterManagerRegionalParameterVersion#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

