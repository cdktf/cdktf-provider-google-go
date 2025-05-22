// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinterconnect


type ComputeInterconnectMacsec struct {
	// pre_shared_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/compute_interconnect#pre_shared_keys ComputeInterconnect#pre_shared_keys}
	PreSharedKeys interface{} `field:"required" json:"preSharedKeys" yaml:"preSharedKeys"`
	// If set to true, the Interconnect connection is configured with a should-secure MACsec security policy, that allows the Google router to fallback to cleartext traffic if the MKA session cannot be established.
	//
	// By default, the Interconnect
	// connection is configured with a must-secure security policy that drops all traffic
	// if the MKA session cannot be established with your router.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/compute_interconnect#fail_open ComputeInterconnect#fail_open}
	FailOpen interface{} `field:"optional" json:"failOpen" yaml:"failOpen"`
}

