package computereservation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeReservationConfig struct {
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
	// Name of the resource.
	//
	// Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_reservation#name ComputeReservation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// specific_reservation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_reservation#specific_reservation ComputeReservation#specific_reservation}
	SpecificReservation *ComputeReservationSpecificReservation `field:"required" json:"specificReservation" yaml:"specificReservation"`
	// The zone where the reservation is made.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_reservation#zone ComputeReservation#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_reservation#description ComputeReservation#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_reservation#id ComputeReservation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_reservation#project ComputeReservation#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// share_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_reservation#share_settings ComputeReservation#share_settings}
	ShareSettings *ComputeReservationShareSettings `field:"optional" json:"shareSettings" yaml:"shareSettings"`
	// When set to true, only VMs that target this reservation by name can consume this reservation.
	//
	// Otherwise, it can be consumed by VMs with
	// affinity for any reservation. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_reservation#specific_reservation_required ComputeReservation#specific_reservation_required}
	SpecificReservationRequired interface{} `field:"optional" json:"specificReservationRequired" yaml:"specificReservationRequired"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_reservation#timeouts ComputeReservation#timeouts}
	Timeouts *ComputeReservationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

