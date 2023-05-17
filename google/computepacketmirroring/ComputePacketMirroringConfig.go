package computepacketmirroring

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputePacketMirroringConfig struct {
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
	// collector_ilb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_packet_mirroring#collector_ilb ComputePacketMirroring#collector_ilb}
	CollectorIlb *ComputePacketMirroringCollectorIlb `field:"required" json:"collectorIlb" yaml:"collectorIlb"`
	// mirrored_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_packet_mirroring#mirrored_resources ComputePacketMirroring#mirrored_resources}
	MirroredResources *ComputePacketMirroringMirroredResources `field:"required" json:"mirroredResources" yaml:"mirroredResources"`
	// The name of the packet mirroring rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_packet_mirroring#name ComputePacketMirroring#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_packet_mirroring#network ComputePacketMirroring#network}
	Network *ComputePacketMirroringNetwork `field:"required" json:"network" yaml:"network"`
	// A human-readable description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_packet_mirroring#description ComputePacketMirroring#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_packet_mirroring#filter ComputePacketMirroring#filter}
	Filter *ComputePacketMirroringFilter `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_packet_mirroring#id ComputePacketMirroring#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Since only one rule can be active at a time, priority is used to break ties in the case of two rules that apply to the same instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_packet_mirroring#priority ComputePacketMirroring#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_packet_mirroring#project ComputePacketMirroring#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_packet_mirroring#region ComputePacketMirroring#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_packet_mirroring#timeouts ComputePacketMirroring#timeouts}
	Timeouts *ComputePacketMirroringTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

