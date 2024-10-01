// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package beyondcorpappconnector


type BeyondcorpAppConnectorPrincipalInfoServiceAccount struct {
	// Email address of the service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/beyondcorp_app_connector#email BeyondcorpAppConnector#email}
	Email *string `field:"required" json:"email" yaml:"email"`
}

