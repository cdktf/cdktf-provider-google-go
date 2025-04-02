// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chronicleretrohunt


type ChronicleRetrohuntProcessInterval struct {
	// Exclusive end of the interval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/chronicle_retrohunt#end_time ChronicleRetrohunt#end_time}
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// Inclusive start of the interval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/chronicle_retrohunt#start_time ChronicleRetrohunt#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

