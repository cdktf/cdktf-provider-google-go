// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appengineservicenetworksettings


type AppEngineServiceNetworkSettingsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/app_engine_service_network_settings#create AppEngineServiceNetworkSettings#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/app_engine_service_network_settings#delete AppEngineServiceNetworkSettings#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/app_engine_service_network_settings#update AppEngineServiceNetworkSettings#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

