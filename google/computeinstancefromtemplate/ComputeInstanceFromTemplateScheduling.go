// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancefromtemplate


type ComputeInstanceFromTemplateScheduling struct {
	// Specifies if the instance should be restarted if it was terminated by Compute Engine (not a user).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/compute_instance_from_template#automatic_restart ComputeInstanceFromTemplate#automatic_restart}
	AutomaticRestart interface{} `field:"optional" json:"automaticRestart" yaml:"automaticRestart"`
	// Specifies the availability domain, which this instance should be scheduled on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/compute_instance_from_template#availability_domain ComputeInstanceFromTemplate#availability_domain}
	AvailabilityDomain *float64 `field:"optional" json:"availabilityDomain" yaml:"availabilityDomain"`
	// Specifies the action GCE should take when SPOT VM is preempted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/compute_instance_from_template#instance_termination_action ComputeInstanceFromTemplate#instance_termination_action}
	InstanceTerminationAction *string `field:"optional" json:"instanceTerminationAction" yaml:"instanceTerminationAction"`
	// local_ssd_recovery_timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/compute_instance_from_template#local_ssd_recovery_timeout ComputeInstanceFromTemplate#local_ssd_recovery_timeout}
	LocalSsdRecoveryTimeout *ComputeInstanceFromTemplateSchedulingLocalSsdRecoveryTimeout `field:"optional" json:"localSsdRecoveryTimeout" yaml:"localSsdRecoveryTimeout"`
	// max_run_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/compute_instance_from_template#max_run_duration ComputeInstanceFromTemplate#max_run_duration}
	MaxRunDuration *ComputeInstanceFromTemplateSchedulingMaxRunDuration `field:"optional" json:"maxRunDuration" yaml:"maxRunDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/compute_instance_from_template#min_node_cpus ComputeInstanceFromTemplate#min_node_cpus}.
	MinNodeCpus *float64 `field:"optional" json:"minNodeCpus" yaml:"minNodeCpus"`
	// node_affinities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/compute_instance_from_template#node_affinities ComputeInstanceFromTemplate#node_affinities}
	NodeAffinities interface{} `field:"optional" json:"nodeAffinities" yaml:"nodeAffinities"`
	// Describes maintenance behavior for the instance. One of MIGRATE or TERMINATE,.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/compute_instance_from_template#on_host_maintenance ComputeInstanceFromTemplate#on_host_maintenance}
	OnHostMaintenance *string `field:"optional" json:"onHostMaintenance" yaml:"onHostMaintenance"`
	// on_instance_stop_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/compute_instance_from_template#on_instance_stop_action ComputeInstanceFromTemplate#on_instance_stop_action}
	OnInstanceStopAction *ComputeInstanceFromTemplateSchedulingOnInstanceStopAction `field:"optional" json:"onInstanceStopAction" yaml:"onInstanceStopAction"`
	// Whether the instance is preemptible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/compute_instance_from_template#preemptible ComputeInstanceFromTemplate#preemptible}
	Preemptible interface{} `field:"optional" json:"preemptible" yaml:"preemptible"`
	// Whether the instance is spot. If this is set as SPOT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/compute_instance_from_template#provisioning_model ComputeInstanceFromTemplate#provisioning_model}
	ProvisioningModel *string `field:"optional" json:"provisioningModel" yaml:"provisioningModel"`
	// Specifies the timestamp, when the instance will be terminated, in RFC3339 text format.
	//
	// If specified, the instance termination action
	// will be performed at the termination time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/compute_instance_from_template#termination_time ComputeInstanceFromTemplate#termination_time}
	TerminationTime *string `field:"optional" json:"terminationTime" yaml:"terminationTime"`
}

