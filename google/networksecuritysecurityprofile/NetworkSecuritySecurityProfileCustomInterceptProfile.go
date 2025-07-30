// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecuritysecurityprofile


type NetworkSecuritySecurityProfileCustomInterceptProfile struct {
	// The Intercept Endpoint Group to which matching traffic should be intercepted. Format: projects/{project_id}/locations/global/interceptEndpointGroups/{endpoint_group_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/network_security_security_profile#intercept_endpoint_group NetworkSecuritySecurityProfile#intercept_endpoint_group}
	InterceptEndpointGroup *string `field:"required" json:"interceptEndpointGroup" yaml:"interceptEndpointGroup"`
}

