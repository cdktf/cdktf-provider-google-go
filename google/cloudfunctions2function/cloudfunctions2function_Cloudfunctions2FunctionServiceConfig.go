package cloudfunctions2function


type Cloudfunctions2FunctionServiceConfig struct {
	// Whether 100% of traffic is routed to the latest revision. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#all_traffic_on_latest_revision Cloudfunctions2Function#all_traffic_on_latest_revision}
	AllTrafficOnLatestRevision interface{} `field:"optional" json:"allTrafficOnLatestRevision" yaml:"allTrafficOnLatestRevision"`
	// The amount of memory available for a function.
	//
	// Defaults to 256M. Supported units are k, M, G, Mi, Gi. If no unit is
	// supplied the value is interpreted as bytes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#available_memory Cloudfunctions2Function#available_memory}
	AvailableMemory *string `field:"optional" json:"availableMemory" yaml:"availableMemory"`
	// Environment variables that shall be available during function execution.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#environment_variables Cloudfunctions2Function#environment_variables}
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Available ingress settings. Defaults to "ALLOW_ALL" if unspecified. Default value: "ALLOW_ALL" Possible values: ["ALLOW_ALL", "ALLOW_INTERNAL_ONLY", "ALLOW_INTERNAL_AND_GCLB"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#ingress_settings Cloudfunctions2Function#ingress_settings}
	IngressSettings *string `field:"optional" json:"ingressSettings" yaml:"ingressSettings"`
	// The limit on the maximum number of function instances that may coexist at a given time.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#max_instance_count Cloudfunctions2Function#max_instance_count}
	MaxInstanceCount *float64 `field:"optional" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// The limit on the minimum number of function instances that may coexist at a given time.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#min_instance_count Cloudfunctions2Function#min_instance_count}
	MinInstanceCount *float64 `field:"optional" json:"minInstanceCount" yaml:"minInstanceCount"`
	// secret_environment_variables block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#secret_environment_variables Cloudfunctions2Function#secret_environment_variables}
	SecretEnvironmentVariables interface{} `field:"optional" json:"secretEnvironmentVariables" yaml:"secretEnvironmentVariables"`
	// secret_volumes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#secret_volumes Cloudfunctions2Function#secret_volumes}
	SecretVolumes interface{} `field:"optional" json:"secretVolumes" yaml:"secretVolumes"`
	// Name of the service associated with a Function.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#service Cloudfunctions2Function#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// The email of the service account for this function.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#service_account_email Cloudfunctions2Function#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// The function execution timeout.
	//
	// Execution is considered failed and
	// can be terminated if the function is not completed at the end of the
	// timeout period. Defaults to 60 seconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#timeout_seconds Cloudfunctions2Function#timeout_seconds}
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// The Serverless VPC Access connector that this cloud function can connect to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#vpc_connector Cloudfunctions2Function#vpc_connector}
	VpcConnector *string `field:"optional" json:"vpcConnector" yaml:"vpcConnector"`
	// Available egress settings. Possible values: ["VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED", "PRIVATE_RANGES_ONLY", "ALL_TRAFFIC"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#vpc_connector_egress_settings Cloudfunctions2Function#vpc_connector_egress_settings}
	VpcConnectorEgressSettings *string `field:"optional" json:"vpcConnectorEgressSettings" yaml:"vpcConnectorEgressSettings"`
}

