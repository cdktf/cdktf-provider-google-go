// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagoogledatacatalogtaxonomyiampolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleDataCatalogTaxonomyIamPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/data-sources/data_catalog_taxonomy_iam_policy#taxonomy DataGoogleDataCatalogTaxonomyIamPolicy#taxonomy}.
	Taxonomy *string `field:"required" json:"taxonomy" yaml:"taxonomy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/data-sources/data_catalog_taxonomy_iam_policy#id DataGoogleDataCatalogTaxonomyIamPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/data-sources/data_catalog_taxonomy_iam_policy#project DataGoogleDataCatalogTaxonomyIamPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/data-sources/data_catalog_taxonomy_iam_policy#region DataGoogleDataCatalogTaxonomyIamPolicy#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

