// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chroniclereferencelist

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ChronicleReferenceListConfig struct {
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
	// Required. A user-provided description of the reference list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/chronicle_reference_list#description ChronicleReferenceList#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// entries block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/chronicle_reference_list#entries ChronicleReferenceList#entries}
	Entries interface{} `field:"required" json:"entries" yaml:"entries"`
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/chronicle_reference_list#instance ChronicleReferenceList#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// The location of the resource.
	//
	// This is the geographical region where the Chronicle instance resides, such as "us" or "europe-west2".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/chronicle_reference_list#location ChronicleReferenceList#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Required.
	//
	// The ID to use for the reference list. This is also the display name for
	// the reference list. It must satisfy the following requirements:
	// - Starts with letter.
	// - Contains only letters, numbers and underscore.
	// - Has length < 256.
	// - Must be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/chronicle_reference_list#reference_list_id ChronicleReferenceList#reference_list_id}
	ReferenceListId *string `field:"required" json:"referenceListId" yaml:"referenceListId"`
	// Possible values: REFERENCE_LIST_SYNTAX_TYPE_PLAIN_TEXT_STRING REFERENCE_LIST_SYNTAX_TYPE_REGEX REFERENCE_LIST_SYNTAX_TYPE_CIDR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/chronicle_reference_list#syntax_type ChronicleReferenceList#syntax_type}
	SyntaxType *string `field:"required" json:"syntaxType" yaml:"syntaxType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/chronicle_reference_list#id ChronicleReferenceList#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/chronicle_reference_list#project ChronicleReferenceList#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/chronicle_reference_list#timeouts ChronicleReferenceList#timeouts}
	Timeouts *ChronicleReferenceListTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

