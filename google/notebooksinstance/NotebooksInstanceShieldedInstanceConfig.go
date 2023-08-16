package notebooksinstance


type NotebooksInstanceShieldedInstanceConfig struct {
	// Defines whether the instance has integrity monitoring enabled.
	//
	// Enables monitoring and attestation of the
	// boot integrity of the instance. The attestation is performed against the integrity policy baseline.
	// This baseline is initially derived from the implicitly trusted boot image when the instance is created.
	// Enabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/notebooks_instance#enable_integrity_monitoring NotebooksInstance#enable_integrity_monitoring}
	EnableIntegrityMonitoring interface{} `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Defines whether the instance has Secure Boot enabled.
	//
	// Secure Boot helps ensure that the system only runs
	// authentic software by verifying the digital signature of all boot components, and halting the boot process
	// if signature verification fails.
	// Disabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/notebooks_instance#enable_secure_boot NotebooksInstance#enable_secure_boot}
	EnableSecureBoot interface{} `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
	// Defines whether the instance has the vTPM enabled. Enabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/notebooks_instance#enable_vtpm NotebooksInstance#enable_vtpm}
	EnableVtpm interface{} `field:"optional" json:"enableVtpm" yaml:"enableVtpm"`
}

