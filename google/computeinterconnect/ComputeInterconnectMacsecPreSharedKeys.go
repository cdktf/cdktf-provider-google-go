// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinterconnect


type ComputeInterconnectMacsecPreSharedKeys struct {
	// A name for this pre-shared key.
	//
	// The name must be 1-63 characters long, and
	//  comply with RFC1035. Specifically, the name must be 1-63 characters long and match
	//  the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character
	//  must be a lowercase letter, and all following characters must be a dash, lowercase
	//  letter, or digit, except the last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/compute_interconnect#name ComputeInterconnect#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// If set to true, the Interconnect connection is configured with a should-secure MACsec security policy, that allows the Google router to fallback to cleartext traffic if the MKA session cannot be established.
	//
	// By default, the Interconnect
	// connection is configured with a must-secure security policy that drops all traffic
	// if the MKA session cannot be established with your router.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/compute_interconnect#fail_open ComputeInterconnect#fail_open}
	FailOpen interface{} `field:"optional" json:"failOpen" yaml:"failOpen"`
	// A RFC3339 timestamp on or after which the key is valid.
	//
	// startTime can be in the
	// future. If the keychain has a single key, startTime can be omitted. If the keychain
	// has multiple keys, startTime is mandatory for each key. The start times of keys must
	// be in increasing order. The start times of two consecutive keys must be at least 6
	// hours apart.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/compute_interconnect#start_time ComputeInterconnect#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

