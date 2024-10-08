// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firestoreindex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirestoreIndexConfig struct {
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
	// The collection being indexed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/firestore_index#collection FirestoreIndex#collection}
	Collection *string `field:"required" json:"collection" yaml:"collection"`
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/firestore_index#fields FirestoreIndex#fields}
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// The API scope at which a query is run. Default value: "ANY_API" Possible values: ["ANY_API", "DATASTORE_MODE_API"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/firestore_index#api_scope FirestoreIndex#api_scope}
	ApiScope *string `field:"optional" json:"apiScope" yaml:"apiScope"`
	// The Firestore database id. Defaults to '"(default)"'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/firestore_index#database FirestoreIndex#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/firestore_index#id FirestoreIndex#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/firestore_index#project FirestoreIndex#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The scope at which a query is run. Default value: "COLLECTION" Possible values: ["COLLECTION", "COLLECTION_GROUP", "COLLECTION_RECURSIVE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/firestore_index#query_scope FirestoreIndex#query_scope}
	QueryScope *string `field:"optional" json:"queryScope" yaml:"queryScope"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/firestore_index#timeouts FirestoreIndex#timeouts}
	Timeouts *FirestoreIndexTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

