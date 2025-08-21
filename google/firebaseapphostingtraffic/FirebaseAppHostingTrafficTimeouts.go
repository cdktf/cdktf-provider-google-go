// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaseapphostingtraffic


type FirebaseAppHostingTrafficTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/firebase_app_hosting_traffic#create FirebaseAppHostingTraffic#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/firebase_app_hosting_traffic#delete FirebaseAppHostingTraffic#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/firebase_app_hosting_traffic#update FirebaseAppHostingTraffic#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

