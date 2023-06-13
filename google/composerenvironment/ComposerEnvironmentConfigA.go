package composerenvironment


type ComposerEnvironmentConfigA struct {
	// database_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/composer_environment#database_config ComposerEnvironment#database_config}
	DatabaseConfig *ComposerEnvironmentConfigDatabaseConfig `field:"optional" json:"databaseConfig" yaml:"databaseConfig"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/composer_environment#encryption_config ComposerEnvironment#encryption_config}
	EncryptionConfig *ComposerEnvironmentConfigEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// The size of the Cloud Composer environment.
	//
	// This field is supported for Cloud Composer environments in versions composer-2.*.*-airflow-*.*.* and newer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/composer_environment#environment_size ComposerEnvironment#environment_size}
	EnvironmentSize *string `field:"optional" json:"environmentSize" yaml:"environmentSize"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/composer_environment#maintenance_window ComposerEnvironment#maintenance_window}
	MaintenanceWindow *ComposerEnvironmentConfigMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// master_authorized_networks_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/composer_environment#master_authorized_networks_config ComposerEnvironment#master_authorized_networks_config}
	MasterAuthorizedNetworksConfig *ComposerEnvironmentConfigMasterAuthorizedNetworksConfig `field:"optional" json:"masterAuthorizedNetworksConfig" yaml:"masterAuthorizedNetworksConfig"`
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/composer_environment#node_config ComposerEnvironment#node_config}
	NodeConfig *ComposerEnvironmentConfigNodeConfig `field:"optional" json:"nodeConfig" yaml:"nodeConfig"`
	// The number of nodes in the Kubernetes Engine cluster that will be used to run this environment.
	//
	// This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/composer_environment#node_count ComposerEnvironment#node_count}
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// private_environment_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/composer_environment#private_environment_config ComposerEnvironment#private_environment_config}
	PrivateEnvironmentConfig *ComposerEnvironmentConfigPrivateEnvironmentConfig `field:"optional" json:"privateEnvironmentConfig" yaml:"privateEnvironmentConfig"`
	// recovery_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/composer_environment#recovery_config ComposerEnvironment#recovery_config}
	RecoveryConfig *ComposerEnvironmentConfigRecoveryConfig `field:"optional" json:"recoveryConfig" yaml:"recoveryConfig"`
	// software_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/composer_environment#software_config ComposerEnvironment#software_config}
	SoftwareConfig *ComposerEnvironmentConfigSoftwareConfig `field:"optional" json:"softwareConfig" yaml:"softwareConfig"`
	// web_server_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/composer_environment#web_server_config ComposerEnvironment#web_server_config}
	WebServerConfig *ComposerEnvironmentConfigWebServerConfig `field:"optional" json:"webServerConfig" yaml:"webServerConfig"`
	// web_server_network_access_control block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/composer_environment#web_server_network_access_control ComposerEnvironment#web_server_network_access_control}
	WebServerNetworkAccessControl *ComposerEnvironmentConfigWebServerNetworkAccessControl `field:"optional" json:"webServerNetworkAccessControl" yaml:"webServerNetworkAccessControl"`
	// workloads_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/composer_environment#workloads_config ComposerEnvironment#workloads_config}
	WorkloadsConfig *ComposerEnvironmentConfigWorkloadsConfig `field:"optional" json:"workloadsConfig" yaml:"workloadsConfig"`
}

