// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2job


type CloudRunV2JobTemplateTemplateContainersStartupProbeHttpGet struct {
	// http_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/cloud_run_v2_job#http_headers CloudRunV2Job#http_headers}
	HttpHeaders interface{} `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	// Path to access on the HTTP server. If set, it should not be empty string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/cloud_run_v2_job#path CloudRunV2Job#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Port number to access on the container.
	//
	// Number must be in the range 1 to 65535.
	// If not specified, defaults to the same value as container.ports[0].containerPort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/cloud_run_v2_job#port CloudRunV2Job#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

