// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataproccluster


type DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyProvisioningModelMix struct {
	// The base capacity that will always use Standard VMs to avoid risk of more preemption than the minimum capacity you need.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dataproc_cluster#standard_capacity_base DataprocCluster#standard_capacity_base}
	StandardCapacityBase *float64 `field:"optional" json:"standardCapacityBase" yaml:"standardCapacityBase"`
	// The percentage of target capacity that should use Standard VM. The remaining percentage will use Spot VMs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dataproc_cluster#standard_capacity_percent_above_base DataprocCluster#standard_capacity_percent_above_base}
	StandardCapacityPercentAboveBase *float64 `field:"optional" json:"standardCapacityPercentAboveBase" yaml:"standardCapacityPercentAboveBase"`
}

