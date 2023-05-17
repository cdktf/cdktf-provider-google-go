package composerenvironment


type ComposerEnvironmentConfigWorkloadsConfigWebServer struct {
	// CPU request and limit for Airflow web server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/composer_environment#cpu ComposerEnvironment#cpu}
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Memory (GB) request and limit for Airflow web server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/composer_environment#memory_gb ComposerEnvironment#memory_gb}
	MemoryGb *float64 `field:"optional" json:"memoryGb" yaml:"memoryGb"`
	// Storage (GB) request and limit for Airflow web server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/composer_environment#storage_gb ComposerEnvironment#storage_gb}
	StorageGb *float64 `field:"optional" json:"storageGb" yaml:"storageGb"`
}

