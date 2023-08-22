package composerenvironment


type ComposerEnvironmentConfigDatabaseConfig struct {
	// Optional.
	//
	// Cloud SQL machine type used by Airflow database. It has to be one of: db-n1-standard-2, db-n1-standard-4, db-n1-standard-8 or db-n1-standard-16. If not specified, db-n1-standard-2 will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/composer_environment#machine_type ComposerEnvironment#machine_type}
	MachineType *string `field:"required" json:"machineType" yaml:"machineType"`
}

