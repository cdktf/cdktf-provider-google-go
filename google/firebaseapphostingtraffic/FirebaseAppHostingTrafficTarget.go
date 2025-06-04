// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaseapphostingtraffic


type FirebaseAppHostingTrafficTarget struct {
	// splits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/firebase_app_hosting_traffic#splits FirebaseAppHostingTraffic#splits}
	Splits interface{} `field:"required" json:"splits" yaml:"splits"`
}

