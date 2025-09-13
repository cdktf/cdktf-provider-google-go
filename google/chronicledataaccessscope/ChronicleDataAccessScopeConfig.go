// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chronicledataaccessscope

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ChronicleDataAccessScopeConfig struct {
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
	// Required.
	//
	// The user provided scope id which will become the last part of the name
	// of the scope resource.
	// Needs to be compliant with https://google.aip.dev/122
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/chronicle_data_access_scope#data_access_scope_id ChronicleDataAccessScope#data_access_scope_id}
	DataAccessScopeId *string `field:"required" json:"dataAccessScopeId" yaml:"dataAccessScopeId"`
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/chronicle_data_access_scope#instance ChronicleDataAccessScope#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// The location of the resource.
	//
	// This is the geographical region where the Chronicle instance resides, such as "us" or "europe-west2".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/chronicle_data_access_scope#location ChronicleDataAccessScope#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Optional.
	//
	// Whether or not the scope allows all labels, allow_all and
	// allowed_data_access_labels are mutually exclusive and one of them must be
	// present. denied_data_access_labels can still be used along with allow_all.
	// When combined with denied_data_access_labels, access will be granted to all
	// data that doesn't have labels mentioned in denied_data_access_labels. E.g.:
	// A customer with scope with denied labels A and B and allow_all will be able
	// to see all data except data labeled with A and data labeled with B and data
	// with labels A and B.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/chronicle_data_access_scope#allow_all ChronicleDataAccessScope#allow_all}
	AllowAll interface{} `field:"optional" json:"allowAll" yaml:"allowAll"`
	// allowed_data_access_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/chronicle_data_access_scope#allowed_data_access_labels ChronicleDataAccessScope#allowed_data_access_labels}
	AllowedDataAccessLabels interface{} `field:"optional" json:"allowedDataAccessLabels" yaml:"allowedDataAccessLabels"`
	// denied_data_access_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/chronicle_data_access_scope#denied_data_access_labels ChronicleDataAccessScope#denied_data_access_labels}
	DeniedDataAccessLabels interface{} `field:"optional" json:"deniedDataAccessLabels" yaml:"deniedDataAccessLabels"`
	// Optional. A description of the data access scope for a human reader.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/chronicle_data_access_scope#description ChronicleDataAccessScope#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/chronicle_data_access_scope#id ChronicleDataAccessScope#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/chronicle_data_access_scope#project ChronicleDataAccessScope#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/chronicle_data_access_scope#timeouts ChronicleDataAccessScope#timeouts}
	Timeouts *ChronicleDataAccessScopeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

