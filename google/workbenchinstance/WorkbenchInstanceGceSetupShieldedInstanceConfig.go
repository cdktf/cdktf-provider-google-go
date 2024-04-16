// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workbenchinstance


type WorkbenchInstanceGceSetupShieldedInstanceConfig struct {
	// Optional.
	//
	// Defines whether the VM instance has integrity monitoring
	// enabled. Enables monitoring and attestation of the boot integrity of the VM
	// instance. The attestation is performed against the integrity policy baseline.
	// This baseline is initially derived from the implicitly trusted boot image
	// when the VM instance is created. Enabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.25.0/docs/resources/workbench_instance#enable_integrity_monitoring WorkbenchInstance#enable_integrity_monitoring}
	EnableIntegrityMonitoring interface{} `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Optional.
	//
	// Defines whether the VM instance has Secure Boot enabled.
	// Secure Boot helps ensure that the system only runs authentic software by verifying
	// the digital signature of all boot components, and halting the boot process
	// if signature verification fails. Disabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.25.0/docs/resources/workbench_instance#enable_secure_boot WorkbenchInstance#enable_secure_boot}
	EnableSecureBoot interface{} `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
	// Optional. Defines whether the VM instance has the vTPM enabled. Enabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.25.0/docs/resources/workbench_instance#enable_vtpm WorkbenchInstance#enable_vtpm}
	EnableVtpm interface{} `field:"optional" json:"enableVtpm" yaml:"enableVtpm"`
}

