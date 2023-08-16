package bigquerycapacitycommitment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryCapacityCommitmentConfig struct {
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
	// Capacity commitment plan. Valid values are at https://cloud.google.com/bigquery/docs/reference/reservations/rpc/google.cloud.bigquery.reservation.v1#commitmentplan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_capacity_commitment#plan BigqueryCapacityCommitment#plan}
	Plan *string `field:"required" json:"plan" yaml:"plan"`
	// Number of slots in this commitment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_capacity_commitment#slot_count BigqueryCapacityCommitment#slot_count}
	SlotCount *float64 `field:"required" json:"slotCount" yaml:"slotCount"`
	// The optional capacity commitment ID.
	//
	// Capacity commitment name will be generated automatically if this field is
	// empty. This field must only contain lower case alphanumeric characters or dashes. The first and last character
	// cannot be a dash. Max length is 64 characters. NOTE: this ID won't be kept if the capacity commitment is split
	// or merged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_capacity_commitment#capacity_commitment_id BigqueryCapacityCommitment#capacity_commitment_id}
	CapacityCommitmentId *string `field:"optional" json:"capacityCommitmentId" yaml:"capacityCommitmentId"`
	// The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_capacity_commitment#edition BigqueryCapacityCommitment#edition}
	Edition *string `field:"optional" json:"edition" yaml:"edition"`
	// If true, fail the request if another project in the organization has a capacity commitment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_capacity_commitment#enforce_single_admin_project_per_org BigqueryCapacityCommitment#enforce_single_admin_project_per_org}
	EnforceSingleAdminProjectPerOrg *string `field:"optional" json:"enforceSingleAdminProjectPerOrg" yaml:"enforceSingleAdminProjectPerOrg"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_capacity_commitment#id BigqueryCapacityCommitment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is US.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_capacity_commitment#location BigqueryCapacityCommitment#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_capacity_commitment#project BigqueryCapacityCommitment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The plan this capacity commitment is converted to after commitmentEndTime passes.
	//
	// Once the plan is changed, committed period is extended according to commitment plan. Only applicable some commitment plans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_capacity_commitment#renewal_plan BigqueryCapacityCommitment#renewal_plan}
	RenewalPlan *string `field:"optional" json:"renewalPlan" yaml:"renewalPlan"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_capacity_commitment#timeouts BigqueryCapacityCommitment#timeouts}
	Timeouts *BigqueryCapacityCommitmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

