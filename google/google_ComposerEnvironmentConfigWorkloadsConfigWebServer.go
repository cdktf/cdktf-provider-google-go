// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComposerEnvironmentConfigWorkloadsConfigWebServer struct {
	// CPU request and limit for Airflow web server.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/composer_environment#cpu ComposerEnvironment#cpu}
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Memory (GB) request and limit for Airflow web server.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/composer_environment#memory_gb ComposerEnvironment#memory_gb}
	MemoryGb *float64 `field:"optional" json:"memoryGb" yaml:"memoryGb"`
	// Storage (GB) request and limit for Airflow web server.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/composer_environment#storage_gb ComposerEnvironment#storage_gb}
	StorageGb *float64 `field:"optional" json:"storageGb" yaml:"storageGb"`
}

