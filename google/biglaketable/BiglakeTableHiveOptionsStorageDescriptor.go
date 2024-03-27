// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package biglaketable


type BiglakeTableHiveOptionsStorageDescriptor struct {
	// The fully qualified Java class name of the input format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/biglake_table#input_format BiglakeTable#input_format}
	InputFormat *string `field:"optional" json:"inputFormat" yaml:"inputFormat"`
	// Cloud Storage folder URI where the table data is stored, starting with "gs://".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/biglake_table#location_uri BiglakeTable#location_uri}
	LocationUri *string `field:"optional" json:"locationUri" yaml:"locationUri"`
	// The fully qualified Java class name of the output format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/biglake_table#output_format BiglakeTable#output_format}
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

