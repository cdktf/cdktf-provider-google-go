// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chroniclewatchlist

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ChronicleWatchlistConfig struct {
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
	// Required. Display name of the watchlist. Note that it must be at least one character and less than 63 characters (https://google.aip.dev/148).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_watchlist#display_name ChronicleWatchlist#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// entity_population_mechanism block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_watchlist#entity_population_mechanism ChronicleWatchlist#entity_population_mechanism}
	EntityPopulationMechanism *ChronicleWatchlistEntityPopulationMechanism `field:"required" json:"entityPopulationMechanism" yaml:"entityPopulationMechanism"`
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_watchlist#instance ChronicleWatchlist#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// The location of the resource.
	//
	// This is the geographical region where the Chronicle instance resides, such as "us" or "europe-west2".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_watchlist#location ChronicleWatchlist#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Optional. Description of the watchlist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_watchlist#description ChronicleWatchlist#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_watchlist#id ChronicleWatchlist#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. Weight applied to the risk score for entities in this watchlist. The default is 1.0 if it is not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_watchlist#multiplying_factor ChronicleWatchlist#multiplying_factor}
	MultiplyingFactor *float64 `field:"optional" json:"multiplyingFactor" yaml:"multiplyingFactor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_watchlist#project ChronicleWatchlist#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_watchlist#timeouts ChronicleWatchlist#timeouts}
	Timeouts *ChronicleWatchlistTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Optional.
	//
	// The ID to use for the watchlist,
	// which will become the final component of the watchlist's resource name.
	// This value should be 4-63 characters, and valid characters
	// are /a-z-/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_watchlist#watchlist_id ChronicleWatchlist#watchlist_id}
	WatchlistId *string `field:"optional" json:"watchlistId" yaml:"watchlistId"`
	// watchlist_user_preferences block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_watchlist#watchlist_user_preferences ChronicleWatchlist#watchlist_user_preferences}
	WatchlistUserPreferences *ChronicleWatchlistWatchlistUserPreferences `field:"optional" json:"watchlistUserPreferences" yaml:"watchlistUserPreferences"`
}

