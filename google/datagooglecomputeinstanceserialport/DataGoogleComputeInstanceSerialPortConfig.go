// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagooglecomputeinstanceserialport

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleComputeInstanceSerialPortConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.80.0/docs/data-sources/compute_instance_serial_port#instance DataGoogleComputeInstanceSerialPort#instance}.
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.80.0/docs/data-sources/compute_instance_serial_port#port DataGoogleComputeInstanceSerialPort#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.80.0/docs/data-sources/compute_instance_serial_port#id DataGoogleComputeInstanceSerialPort#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.80.0/docs/data-sources/compute_instance_serial_port#project DataGoogleComputeInstanceSerialPort#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.80.0/docs/data-sources/compute_instance_serial_port#zone DataGoogleComputeInstanceSerialPort#zone}.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

