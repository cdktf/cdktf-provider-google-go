// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chroniclewatchlist


type ChronicleWatchlistEntityPopulationMechanism struct {
	// manual block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/chronicle_watchlist#manual ChronicleWatchlist#manual}
	Manual *ChronicleWatchlistEntityPopulationMechanismManual `field:"optional" json:"manual" yaml:"manual"`
}

