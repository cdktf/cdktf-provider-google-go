// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolNodeConfigShieldedInstanceConfig struct {
	// Defines whether the instance has integrity monitoring enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/container_node_pool#enable_integrity_monitoring ContainerNodePool#enable_integrity_monitoring}
	EnableIntegrityMonitoring interface{} `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Defines whether the instance has Secure Boot enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/container_node_pool#enable_secure_boot ContainerNodePool#enable_secure_boot}
	EnableSecureBoot interface{} `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
}

