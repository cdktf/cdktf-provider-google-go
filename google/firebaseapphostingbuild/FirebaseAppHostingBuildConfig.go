// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaseapphostingbuild

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirebaseAppHostingBuildConfig struct {
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
	// The ID of the Backend that this Build applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/firebase_app_hosting_build#backend FirebaseAppHostingBuild#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// The user-specified ID of the build being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/firebase_app_hosting_build#build_id FirebaseAppHostingBuild#build_id}
	BuildId *string `field:"required" json:"buildId" yaml:"buildId"`
	// The location of the Backend that this Build applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/firebase_app_hosting_build#location FirebaseAppHostingBuild#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/firebase_app_hosting_build#source FirebaseAppHostingBuild#source}
	Source *FirebaseAppHostingBuildSource `field:"required" json:"source" yaml:"source"`
	// Unstructured key value map that may be set by external tools to store and arbitrary metadata.
	//
	// They are not queryable and should be
	// preserved when modifying objects.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/firebase_app_hosting_build#annotations FirebaseAppHostingBuild#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Human-readable name. 63 character limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/firebase_app_hosting_build#display_name FirebaseAppHostingBuild#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/firebase_app_hosting_build#id FirebaseAppHostingBuild#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Unstructured key value map that can be used to organize and categorize objects.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/firebase_app_hosting_build#labels FirebaseAppHostingBuild#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/firebase_app_hosting_build#project FirebaseAppHostingBuild#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/firebase_app_hosting_build#timeouts FirebaseAppHostingBuild#timeouts}
	Timeouts *FirebaseAppHostingBuildTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

