// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chronicledataaccesslabel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ChronicleDataAccessLabelConfig struct {
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
	// The ID to use for the data access label, which will become the label's
	// display name and the final component of the label's resource name. The
	// maximum number of characters should be 63. Regex pattern is as per AIP:
	// https://google.aip.dev/122#resource-id-segments
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/chronicle_data_access_label#data_access_label_id ChronicleDataAccessLabel#data_access_label_id}
	DataAccessLabelId *string `field:"required" json:"dataAccessLabelId" yaml:"dataAccessLabelId"`
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/chronicle_data_access_label#instance ChronicleDataAccessLabel#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// The location of the resource.
	//
	// This is the geographical region where the Chronicle instance resides, such as "us" or "europe-west2".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/chronicle_data_access_label#location ChronicleDataAccessLabel#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// A UDM query over event data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/chronicle_data_access_label#udm_query ChronicleDataAccessLabel#udm_query}
	UdmQuery *string `field:"required" json:"udmQuery" yaml:"udmQuery"`
	// Optional. A description of the data access label for a human reader.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/chronicle_data_access_label#description ChronicleDataAccessLabel#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/chronicle_data_access_label#id ChronicleDataAccessLabel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/chronicle_data_access_label#project ChronicleDataAccessLabel#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/chronicle_data_access_label#timeouts ChronicleDataAccessLabel#timeouts}
	Timeouts *ChronicleDataAccessLabelTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

