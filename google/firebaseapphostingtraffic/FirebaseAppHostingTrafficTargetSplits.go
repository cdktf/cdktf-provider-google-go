// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaseapphostingtraffic


type FirebaseAppHostingTrafficTargetSplits struct {
	// The build that traffic is being routed to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/firebase_app_hosting_traffic#build FirebaseAppHostingTraffic#build}
	BuildAttribute *string `field:"required" json:"buildAttribute" yaml:"buildAttribute"`
	// The percentage of traffic to send to the build. Currently must be 100 or 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/firebase_app_hosting_traffic#percent FirebaseAppHostingTraffic#percent}
	Percent *float64 `field:"required" json:"percent" yaml:"percent"`
}

