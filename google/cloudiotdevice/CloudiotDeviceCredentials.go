// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudiotdevice


type CloudiotDeviceCredentials struct {
	// public_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.83.0/docs/resources/cloudiot_device#public_key CloudiotDevice#public_key}
	PublicKey *CloudiotDeviceCredentialsPublicKey `field:"required" json:"publicKey" yaml:"publicKey"`
	// The time at which this credential becomes invalid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.83.0/docs/resources/cloudiot_device#expiration_time CloudiotDevice#expiration_time}
	ExpirationTime *string `field:"optional" json:"expirationTime" yaml:"expirationTime"`
}

