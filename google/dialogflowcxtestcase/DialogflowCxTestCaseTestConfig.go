// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxtestcase


type DialogflowCxTestCaseTestConfig struct {
	// Flow name to start the test case with.
	//
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	// Only one of flow and page should be set to indicate the starting point of the test case. If neither is set, the test case will start with start page on the default start flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/dialogflow_cx_test_case#flow DialogflowCxTestCase#flow}
	Flow *string `field:"optional" json:"flow" yaml:"flow"`
	// The page to start the test case with.
	//
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>.
	// Only one of flow and page should be set to indicate the starting point of the test case. If neither is set, the test case will start with start page on the default start flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/dialogflow_cx_test_case#page DialogflowCxTestCase#page}
	Page *string `field:"optional" json:"page" yaml:"page"`
	// Session parameters to be compared when calculating differences.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/dialogflow_cx_test_case#tracking_parameters DialogflowCxTestCase#tracking_parameters}
	TrackingParameters *[]*string `field:"optional" json:"trackingParameters" yaml:"trackingParameters"`
}

