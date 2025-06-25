// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iapsettings


type IapSettingsAccessSettingsOauthSettings struct {
	// Domain hint to send as hd=?
	//
	// parameter in OAuth request flow.
	// Enables redirect to primary IDP by skipping Google's login screen.
	// (https://developers.google.com/identity/protocols/OpenIDConnect#hd-param)
	// Note: IAP does not verify that the id token's hd claim matches this value
	// since access behavior is managed by IAM policies.
	// * loginHint setting is not a replacement for access control. Always enforce an appropriate access policy if you want to restrict access to users outside your domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/iap_settings#login_hint IapSettings#login_hint}
	LoginHint *string `field:"optional" json:"loginHint" yaml:"loginHint"`
	// List of client ids allowed to use IAP programmatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/iap_settings#programmatic_clients IapSettings#programmatic_clients}
	ProgrammaticClients *[]*string `field:"optional" json:"programmaticClients" yaml:"programmaticClients"`
}

