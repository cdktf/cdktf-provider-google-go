// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chroniclewatchlist


type ChronicleWatchlistWatchlistUserPreferences struct {
	// Optional. Whether the watchlist is pinned on the dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/chronicle_watchlist#pinned ChronicleWatchlist#pinned}
	Pinned interface{} `field:"optional" json:"pinned" yaml:"pinned"`
}

