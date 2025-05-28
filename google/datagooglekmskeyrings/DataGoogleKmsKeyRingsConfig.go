// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagooglekmskeyrings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleKmsKeyRingsConfig struct {
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
	// The canonical id for the location. For example: "us-east1".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/data-sources/kms_key_rings#location DataGoogleKmsKeyRings#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The filter argument is used to add a filter query parameter that limits which keys are retrieved by the data source: ?filter={{filter}}. Example values:.
	//
	// * "name:my-key-" will retrieve key rings that contain "my-key-" anywhere in their name. Note: names take the form projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}.
	// * "name=projects/my-project/locations/global/keyRings/my-key-ring" will only retrieve a key ring with that exact name.
	//
	// [See the documentation about using filters](https://cloud.google.com/kms/docs/sorting-and-filtering)
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/data-sources/kms_key_rings#filter DataGoogleKmsKeyRings#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/data-sources/kms_key_rings#id DataGoogleKmsKeyRings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Project ID of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/data-sources/kms_key_rings#project DataGoogleKmsKeyRings#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

