// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chronicleretrohunt

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ChronicleRetrohuntConfig struct {
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
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/chronicle_retrohunt#instance ChronicleRetrohunt#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// The location of the resource.
	//
	// This is the geographical region where the Chronicle instance resides, such as "us" or "europe-west2".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/chronicle_retrohunt#location ChronicleRetrohunt#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// process_interval block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/chronicle_retrohunt#process_interval ChronicleRetrohunt#process_interval}
	ProcessInterval *ChronicleRetrohuntProcessInterval `field:"required" json:"processInterval" yaml:"processInterval"`
	// The Rule ID of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/chronicle_retrohunt#rule ChronicleRetrohunt#rule}
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/chronicle_retrohunt#id ChronicleRetrohunt#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/chronicle_retrohunt#project ChronicleRetrohunt#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The retrohunt ID of the Retrohunt.
	//
	// A retrohunt is an execution of a Rule over a time range in the past.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/chronicle_retrohunt#retrohunt ChronicleRetrohunt#retrohunt}
	Retrohunt *string `field:"optional" json:"retrohunt" yaml:"retrohunt"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/chronicle_retrohunt#timeouts ChronicleRetrohunt#timeouts}
	Timeouts *ChronicleRetrohuntTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

