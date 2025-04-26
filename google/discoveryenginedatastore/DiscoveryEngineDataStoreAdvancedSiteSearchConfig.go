// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginedatastore


type DiscoveryEngineDataStoreAdvancedSiteSearchConfig struct {
	// If set true, automatic refresh is disabled for the DataStore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/discovery_engine_data_store#disable_automatic_refresh DiscoveryEngineDataStore#disable_automatic_refresh}
	DisableAutomaticRefresh interface{} `field:"optional" json:"disableAutomaticRefresh" yaml:"disableAutomaticRefresh"`
	// If set true, initial indexing is disabled for the DataStore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/discovery_engine_data_store#disable_initial_index DiscoveryEngineDataStore#disable_initial_index}
	DisableInitialIndex interface{} `field:"optional" json:"disableInitialIndex" yaml:"disableInitialIndex"`
}

