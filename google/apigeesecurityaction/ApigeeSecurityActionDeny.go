// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeesecurityaction


type ApigeeSecurityActionDeny struct {
	// The HTTP response code if the Action = DENY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apigee_security_action#response_code ApigeeSecurityAction#response_code}
	ResponseCode *float64 `field:"optional" json:"responseCode" yaml:"responseCode"`
}

