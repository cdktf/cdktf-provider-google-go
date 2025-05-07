// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chroniclerule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ChronicleRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/chronicle_rule#instance ChronicleRule#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// The location of the resource.
	//
	// This is the geographical region where the Chronicle instance resides, such as "us" or "europe-west2".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/chronicle_rule#location ChronicleRule#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Policy to determine if the rule should be deleted forcefully.
	//
	// If deletion_policy = "FORCE", any retrohunts and any detections associated with the rule
	// will also be deleted. If deletion_policy = "DEFAULT", the call will only succeed if the
	// rule has no associated retrohunts, including completed retrohunts, and no
	// associated detections. Regardless of this field's value, the rule
	// deployment associated with this rule will also be deleted.
	// Possible values: DEFAULT, FORCE
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/chronicle_rule#deletion_policy ChronicleRule#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// The etag for this rule.
	//
	// If this is provided on update, the request will succeed if and only if it
	// matches the server-computed value, and will fail with an ABORTED error
	// otherwise.
	// Populated in BASIC view and FULL view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/chronicle_rule#etag ChronicleRule#etag}
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/chronicle_rule#id ChronicleRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/chronicle_rule#project ChronicleRule#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Rule Id is the ID of the Rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/chronicle_rule#rule_id ChronicleRule#rule_id}
	RuleId *string `field:"optional" json:"ruleId" yaml:"ruleId"`
	// Resource name of the DataAccessScope bound to this rule.
	//
	// Populated in BASIC view and FULL view.
	// If reference lists are used in the rule, validations will be performed
	// against this scope to ensure that the reference lists are compatible with
	// both the user's and the rule's scopes.
	// The scope should be in the format:
	// "projects/{project}/locations/{location}/instances/{instance}/dataAccessScopes/{scope}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/chronicle_rule#scope ChronicleRule#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The YARA-L content of the rule. Populated in FULL view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/chronicle_rule#text ChronicleRule#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/chronicle_rule#timeouts ChronicleRule#timeouts}
	Timeouts *ChronicleRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

