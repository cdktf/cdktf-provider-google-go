// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaseapphostingdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirebaseAppHostingDomainConfig struct {
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
	// The ID of the Backend that this Domain is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/firebase_app_hosting_domain#backend FirebaseAppHostingDomain#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Id of the domain to create. Must be a valid domain name, such as "foo.com".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/firebase_app_hosting_domain#domain_id FirebaseAppHostingDomain#domain_id}
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The location of the Backend that this Domain is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/firebase_app_hosting_domain#location FirebaseAppHostingDomain#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/firebase_app_hosting_domain#id FirebaseAppHostingDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/firebase_app_hosting_domain#project FirebaseAppHostingDomain#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// serve block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/firebase_app_hosting_domain#serve FirebaseAppHostingDomain#serve}
	Serve *FirebaseAppHostingDomainServe `field:"optional" json:"serve" yaml:"serve"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/firebase_app_hosting_domain#timeouts FirebaseAppHostingDomain#timeouts}
	Timeouts *FirebaseAppHostingDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

