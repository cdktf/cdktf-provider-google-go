package computeregionautoscaler

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionAutoscalerConfig struct {
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
	// autoscaling_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_autoscaler#autoscaling_policy ComputeRegionAutoscaler#autoscaling_policy}
	AutoscalingPolicy *ComputeRegionAutoscalerAutoscalingPolicy `field:"required" json:"autoscalingPolicy" yaml:"autoscalingPolicy"`
	// Name of the resource.
	//
	// The name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_autoscaler#name ComputeRegionAutoscaler#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// URL of the managed instance group that this autoscaler will scale.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_autoscaler#target ComputeRegionAutoscaler#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_autoscaler#description ComputeRegionAutoscaler#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_autoscaler#id ComputeRegionAutoscaler#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_autoscaler#project ComputeRegionAutoscaler#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// URL of the region where the instance group resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_autoscaler#region ComputeRegionAutoscaler#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_autoscaler#timeouts ComputeRegionAutoscaler#timeouts}
	Timeouts *ComputeRegionAutoscalerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

