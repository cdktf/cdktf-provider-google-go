// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectivityinternalrange


type NetworkConnectivityInternalRangeAllocationOptions struct {
	// Optional.
	//
	// Sets the strategy used to automatically find a free range of a size given by prefixLength. Can be set only when trying to create a reservation that automatically finds the free range to reserve. Possible values: ["RANDOM", "FIRST_AVAILABLE", "RANDOM_FIRST_N_AVAILABLE", "FIRST_SMALLEST_FITTING"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/network_connectivity_internal_range#allocation_strategy NetworkConnectivityInternalRange#allocation_strategy}
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Must be set when allocation_strategy is RANDOM_FIRST_N_AVAILABLE, otherwise must remain unset.
	//
	// Defines the size of the set of free ranges from which RANDOM_FIRST_N_AVAILABLE strategy randomy selects one,
	// in other words it sets the N in the RANDOM_FIRST_N_AVAILABLE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/network_connectivity_internal_range#first_available_ranges_lookup_size NetworkConnectivityInternalRange#first_available_ranges_lookup_size}
	FirstAvailableRangesLookupSize *float64 `field:"optional" json:"firstAvailableRangesLookupSize" yaml:"firstAvailableRangesLookupSize"`
}

