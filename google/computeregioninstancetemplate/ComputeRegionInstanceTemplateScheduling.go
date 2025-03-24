// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancetemplate


type ComputeRegionInstanceTemplateScheduling struct {
	// Specifies whether the instance should be automatically restarted if it is terminated by Compute Engine (not terminated by a user).
	//
	// This defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_region_instance_template#automatic_restart ComputeRegionInstanceTemplate#automatic_restart}
	AutomaticRestart interface{} `field:"optional" json:"automaticRestart" yaml:"automaticRestart"`
	// Specifies the availability domain, which this instance should be scheduled on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_region_instance_template#availability_domain ComputeRegionInstanceTemplate#availability_domain}
	AvailabilityDomain *float64 `field:"optional" json:"availabilityDomain" yaml:"availabilityDomain"`
	// Specifies the action GCE should take when SPOT VM is preempted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_region_instance_template#instance_termination_action ComputeRegionInstanceTemplate#instance_termination_action}
	InstanceTerminationAction *string `field:"optional" json:"instanceTerminationAction" yaml:"instanceTerminationAction"`
	// local_ssd_recovery_timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_region_instance_template#local_ssd_recovery_timeout ComputeRegionInstanceTemplate#local_ssd_recovery_timeout}
	LocalSsdRecoveryTimeout interface{} `field:"optional" json:"localSsdRecoveryTimeout" yaml:"localSsdRecoveryTimeout"`
	// max_run_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_region_instance_template#max_run_duration ComputeRegionInstanceTemplate#max_run_duration}
	MaxRunDuration *ComputeRegionInstanceTemplateSchedulingMaxRunDuration `field:"optional" json:"maxRunDuration" yaml:"maxRunDuration"`
	// Minimum number of cpus for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_region_instance_template#min_node_cpus ComputeRegionInstanceTemplate#min_node_cpus}
	MinNodeCpus *float64 `field:"optional" json:"minNodeCpus" yaml:"minNodeCpus"`
	// node_affinities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_region_instance_template#node_affinities ComputeRegionInstanceTemplate#node_affinities}
	NodeAffinities interface{} `field:"optional" json:"nodeAffinities" yaml:"nodeAffinities"`
	// Defines the maintenance behavior for this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_region_instance_template#on_host_maintenance ComputeRegionInstanceTemplate#on_host_maintenance}
	OnHostMaintenance *string `field:"optional" json:"onHostMaintenance" yaml:"onHostMaintenance"`
	// on_instance_stop_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_region_instance_template#on_instance_stop_action ComputeRegionInstanceTemplate#on_instance_stop_action}
	OnInstanceStopAction *ComputeRegionInstanceTemplateSchedulingOnInstanceStopAction `field:"optional" json:"onInstanceStopAction" yaml:"onInstanceStopAction"`
	// Allows instance to be preempted. This defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_region_instance_template#preemptible ComputeRegionInstanceTemplate#preemptible}
	Preemptible interface{} `field:"optional" json:"preemptible" yaml:"preemptible"`
	// Whether the instance is spot. If this is set as SPOT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_region_instance_template#provisioning_model ComputeRegionInstanceTemplate#provisioning_model}
	ProvisioningModel *string `field:"optional" json:"provisioningModel" yaml:"provisioningModel"`
	// Specifies the timestamp, when the instance will be terminated, in RFC3339 text format.
	//
	// If specified, the instance termination action
	// will be performed at the termination time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_region_instance_template#termination_time ComputeRegionInstanceTemplate#termination_time}
	TerminationTime *string `field:"optional" json:"terminationTime" yaml:"terminationTime"`
}

