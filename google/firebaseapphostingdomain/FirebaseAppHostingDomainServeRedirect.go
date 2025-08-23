// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaseapphostingdomain


type FirebaseAppHostingDomainServeRedirect struct {
	// The URI of the redirect's intended destination.
	//
	// This URI will be
	// prepended to the original request path. URI without a scheme are
	// assumed to be HTTPS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/firebase_app_hosting_domain#uri FirebaseAppHostingDomain#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// The status code to use in a redirect response.
	//
	// Must be a valid HTTP 3XX
	// status code. Defaults to 302 if not present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/firebase_app_hosting_domain#status FirebaseAppHostingDomain#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

