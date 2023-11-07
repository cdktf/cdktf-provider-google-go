// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeature


type GkeHubFeatureSpecFleetobservability struct {
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.5.0/docs/resources/gke_hub_feature#logging_config GkeHubFeature#logging_config}
	LoggingConfig *GkeHubFeatureSpecFleetobservabilityLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
}

