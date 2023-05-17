package computeinstancetemplate


type ComputeInstanceTemplateShieldedInstanceConfig struct {
	// Compare the most recent boot measurements to the integrity policy baseline and return a pair of pass/fail results depending on whether they match or not.
	//
	// Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_template#enable_integrity_monitoring ComputeInstanceTemplate#enable_integrity_monitoring}
	EnableIntegrityMonitoring interface{} `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Verify the digital signature of all boot components, and halt the boot process if signature verification fails.
	//
	// Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_template#enable_secure_boot ComputeInstanceTemplate#enable_secure_boot}
	EnableSecureBoot interface{} `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
	// Use a virtualized trusted platform module, which is a specialized computer chip you can use to encrypt objects like keys and certificates.
	//
	// Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_template#enable_vtpm ComputeInstanceTemplate#enable_vtpm}
	EnableVtpm interface{} `field:"optional" json:"enableVtpm" yaml:"enableVtpm"`
}

