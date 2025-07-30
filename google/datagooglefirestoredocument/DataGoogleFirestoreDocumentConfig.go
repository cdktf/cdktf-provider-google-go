// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagooglefirestoredocument

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleFirestoreDocumentConfig struct {
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
	// The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/data-sources/firestore_document#collection DataGoogleFirestoreDocument#collection}
	Collection *string `field:"required" json:"collection" yaml:"collection"`
	// The Firestore database id. Defaults to '"(default)"'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/data-sources/firestore_document#database DataGoogleFirestoreDocument#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// The client-assigned document ID to use for this document during creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/data-sources/firestore_document#document_id DataGoogleFirestoreDocument#document_id}
	DocumentId *string `field:"required" json:"documentId" yaml:"documentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/data-sources/firestore_document#id DataGoogleFirestoreDocument#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/data-sources/firestore_document#project DataGoogleFirestoreDocument#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

