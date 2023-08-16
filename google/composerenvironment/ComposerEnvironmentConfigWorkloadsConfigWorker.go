package composerenvironment


type ComposerEnvironmentConfigWorkloadsConfigWorker struct {
	// CPU request and limit for a single Airflow worker replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/composer_environment#cpu ComposerEnvironment#cpu}
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Maximum number of workers for autoscaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/composer_environment#max_count ComposerEnvironment#max_count}
	MaxCount *float64 `field:"optional" json:"maxCount" yaml:"maxCount"`
	// Memory (GB) request and limit for a single Airflow worker replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/composer_environment#memory_gb ComposerEnvironment#memory_gb}
	MemoryGb *float64 `field:"optional" json:"memoryGb" yaml:"memoryGb"`
	// Minimum number of workers for autoscaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/composer_environment#min_count ComposerEnvironment#min_count}
	MinCount *float64 `field:"optional" json:"minCount" yaml:"minCount"`
	// Storage (GB) request and limit for a single Airflow worker replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/composer_environment#storage_gb ComposerEnvironment#storage_gb}
	StorageGb *float64 `field:"optional" json:"storageGb" yaml:"storageGb"`
}

