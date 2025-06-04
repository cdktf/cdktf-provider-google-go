// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataproccluster


type DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyInstanceSelectionListStruct struct {
	// Full machine-type names, e.g. "n1-standard-16".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/dataproc_cluster#machine_types DataprocCluster#machine_types}
	MachineTypes *[]*string `field:"optional" json:"machineTypes" yaml:"machineTypes"`
	// Preference of this instance selection.
	//
	// Lower number means higher preference. Dataproc will first try to create a VM based on the machine-type with priority rank and fallback to next rank based on availability. Machine types and instance selections with the same priority have the same preference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/dataproc_cluster#rank DataprocCluster#rank}
	Rank *float64 `field:"optional" json:"rank" yaml:"rank"`
}

