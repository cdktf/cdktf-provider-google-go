// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2job


type CloudRunV2JobTemplateTemplateContainersStartupProbeHttpGet struct {
	// http_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.82.0/docs/resources/cloud_run_v2_job#http_headers CloudRunV2Job#http_headers}
	HttpHeaders interface{} `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	// Path to access on the HTTP server. Defaults to '/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.82.0/docs/resources/cloud_run_v2_job#path CloudRunV2Job#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

