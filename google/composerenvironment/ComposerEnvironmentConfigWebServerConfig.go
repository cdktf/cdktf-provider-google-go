package composerenvironment


type ComposerEnvironmentConfigWebServerConfig struct {
	// Optional.
	//
	// Machine type on which Airflow web server is running. It has to be one of: composer-n1-webserver-2, composer-n1-webserver-4 or composer-n1-webserver-8. If not specified, composer-n1-webserver-2 will be used. Value custom is returned only in response, if Airflow web server parameters were manually changed to a non-standard values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/composer_environment#machine_type ComposerEnvironment#machine_type}
	MachineType *string `field:"required" json:"machineType" yaml:"machineType"`
}

