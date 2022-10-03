package provider


type GoogleProviderConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#access_approval_custom_endpoint GoogleProvider#access_approval_custom_endpoint}.
	AccessApprovalCustomEndpoint *string `field:"optional" json:"accessApprovalCustomEndpoint" yaml:"accessApprovalCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#access_context_manager_custom_endpoint GoogleProvider#access_context_manager_custom_endpoint}.
	AccessContextManagerCustomEndpoint *string `field:"optional" json:"accessContextManagerCustomEndpoint" yaml:"accessContextManagerCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#access_token GoogleProvider#access_token}.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#active_directory_custom_endpoint GoogleProvider#active_directory_custom_endpoint}.
	ActiveDirectoryCustomEndpoint *string `field:"optional" json:"activeDirectoryCustomEndpoint" yaml:"activeDirectoryCustomEndpoint"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#alias GoogleProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#apigee_custom_endpoint GoogleProvider#apigee_custom_endpoint}.
	ApigeeCustomEndpoint *string `field:"optional" json:"apigeeCustomEndpoint" yaml:"apigeeCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#apikeys_custom_endpoint GoogleProvider#apikeys_custom_endpoint}.
	ApikeysCustomEndpoint *string `field:"optional" json:"apikeysCustomEndpoint" yaml:"apikeysCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#app_engine_custom_endpoint GoogleProvider#app_engine_custom_endpoint}.
	AppEngineCustomEndpoint *string `field:"optional" json:"appEngineCustomEndpoint" yaml:"appEngineCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#artifact_registry_custom_endpoint GoogleProvider#artifact_registry_custom_endpoint}.
	ArtifactRegistryCustomEndpoint *string `field:"optional" json:"artifactRegistryCustomEndpoint" yaml:"artifactRegistryCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#assured_workloads_custom_endpoint GoogleProvider#assured_workloads_custom_endpoint}.
	AssuredWorkloadsCustomEndpoint *string `field:"optional" json:"assuredWorkloadsCustomEndpoint" yaml:"assuredWorkloadsCustomEndpoint"`
	// batching block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#batching GoogleProvider#batching}
	Batching *GoogleProviderBatching `field:"optional" json:"batching" yaml:"batching"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#bigquery_connection_custom_endpoint GoogleProvider#bigquery_connection_custom_endpoint}.
	BigqueryConnectionCustomEndpoint *string `field:"optional" json:"bigqueryConnectionCustomEndpoint" yaml:"bigqueryConnectionCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#big_query_custom_endpoint GoogleProvider#big_query_custom_endpoint}.
	BigQueryCustomEndpoint *string `field:"optional" json:"bigQueryCustomEndpoint" yaml:"bigQueryCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#bigquery_data_transfer_custom_endpoint GoogleProvider#bigquery_data_transfer_custom_endpoint}.
	BigqueryDataTransferCustomEndpoint *string `field:"optional" json:"bigqueryDataTransferCustomEndpoint" yaml:"bigqueryDataTransferCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#bigquery_reservation_custom_endpoint GoogleProvider#bigquery_reservation_custom_endpoint}.
	BigqueryReservationCustomEndpoint *string `field:"optional" json:"bigqueryReservationCustomEndpoint" yaml:"bigqueryReservationCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#bigtable_custom_endpoint GoogleProvider#bigtable_custom_endpoint}.
	BigtableCustomEndpoint *string `field:"optional" json:"bigtableCustomEndpoint" yaml:"bigtableCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#billing_custom_endpoint GoogleProvider#billing_custom_endpoint}.
	BillingCustomEndpoint *string `field:"optional" json:"billingCustomEndpoint" yaml:"billingCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#billing_project GoogleProvider#billing_project}.
	BillingProject *string `field:"optional" json:"billingProject" yaml:"billingProject"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#binary_authorization_custom_endpoint GoogleProvider#binary_authorization_custom_endpoint}.
	BinaryAuthorizationCustomEndpoint *string `field:"optional" json:"binaryAuthorizationCustomEndpoint" yaml:"binaryAuthorizationCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#certificate_manager_custom_endpoint GoogleProvider#certificate_manager_custom_endpoint}.
	CertificateManagerCustomEndpoint *string `field:"optional" json:"certificateManagerCustomEndpoint" yaml:"certificateManagerCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#cloud_asset_custom_endpoint GoogleProvider#cloud_asset_custom_endpoint}.
	CloudAssetCustomEndpoint *string `field:"optional" json:"cloudAssetCustomEndpoint" yaml:"cloudAssetCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#cloud_billing_custom_endpoint GoogleProvider#cloud_billing_custom_endpoint}.
	CloudBillingCustomEndpoint *string `field:"optional" json:"cloudBillingCustomEndpoint" yaml:"cloudBillingCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#cloud_build_custom_endpoint GoogleProvider#cloud_build_custom_endpoint}.
	CloudBuildCustomEndpoint *string `field:"optional" json:"cloudBuildCustomEndpoint" yaml:"cloudBuildCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#cloud_build_worker_pool_custom_endpoint GoogleProvider#cloud_build_worker_pool_custom_endpoint}.
	CloudBuildWorkerPoolCustomEndpoint *string `field:"optional" json:"cloudBuildWorkerPoolCustomEndpoint" yaml:"cloudBuildWorkerPoolCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#clouddeploy_custom_endpoint GoogleProvider#clouddeploy_custom_endpoint}.
	ClouddeployCustomEndpoint *string `field:"optional" json:"clouddeployCustomEndpoint" yaml:"clouddeployCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#cloudfunctions2_custom_endpoint GoogleProvider#cloudfunctions2_custom_endpoint}.
	Cloudfunctions2CustomEndpoint *string `field:"optional" json:"cloudfunctions2CustomEndpoint" yaml:"cloudfunctions2CustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#cloud_functions_custom_endpoint GoogleProvider#cloud_functions_custom_endpoint}.
	CloudFunctionsCustomEndpoint *string `field:"optional" json:"cloudFunctionsCustomEndpoint" yaml:"cloudFunctionsCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#cloud_identity_custom_endpoint GoogleProvider#cloud_identity_custom_endpoint}.
	CloudIdentityCustomEndpoint *string `field:"optional" json:"cloudIdentityCustomEndpoint" yaml:"cloudIdentityCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#cloud_iot_custom_endpoint GoogleProvider#cloud_iot_custom_endpoint}.
	CloudIotCustomEndpoint *string `field:"optional" json:"cloudIotCustomEndpoint" yaml:"cloudIotCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#cloud_resource_manager_custom_endpoint GoogleProvider#cloud_resource_manager_custom_endpoint}.
	CloudResourceManagerCustomEndpoint *string `field:"optional" json:"cloudResourceManagerCustomEndpoint" yaml:"cloudResourceManagerCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#cloud_run_custom_endpoint GoogleProvider#cloud_run_custom_endpoint}.
	CloudRunCustomEndpoint *string `field:"optional" json:"cloudRunCustomEndpoint" yaml:"cloudRunCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#cloud_scheduler_custom_endpoint GoogleProvider#cloud_scheduler_custom_endpoint}.
	CloudSchedulerCustomEndpoint *string `field:"optional" json:"cloudSchedulerCustomEndpoint" yaml:"cloudSchedulerCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#cloud_tasks_custom_endpoint GoogleProvider#cloud_tasks_custom_endpoint}.
	CloudTasksCustomEndpoint *string `field:"optional" json:"cloudTasksCustomEndpoint" yaml:"cloudTasksCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#composer_custom_endpoint GoogleProvider#composer_custom_endpoint}.
	ComposerCustomEndpoint *string `field:"optional" json:"composerCustomEndpoint" yaml:"composerCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#compute_custom_endpoint GoogleProvider#compute_custom_endpoint}.
	ComputeCustomEndpoint *string `field:"optional" json:"computeCustomEndpoint" yaml:"computeCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#container_analysis_custom_endpoint GoogleProvider#container_analysis_custom_endpoint}.
	ContainerAnalysisCustomEndpoint *string `field:"optional" json:"containerAnalysisCustomEndpoint" yaml:"containerAnalysisCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#container_aws_custom_endpoint GoogleProvider#container_aws_custom_endpoint}.
	ContainerAwsCustomEndpoint *string `field:"optional" json:"containerAwsCustomEndpoint" yaml:"containerAwsCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#container_azure_custom_endpoint GoogleProvider#container_azure_custom_endpoint}.
	ContainerAzureCustomEndpoint *string `field:"optional" json:"containerAzureCustomEndpoint" yaml:"containerAzureCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#container_custom_endpoint GoogleProvider#container_custom_endpoint}.
	ContainerCustomEndpoint *string `field:"optional" json:"containerCustomEndpoint" yaml:"containerCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#credentials GoogleProvider#credentials}.
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#data_catalog_custom_endpoint GoogleProvider#data_catalog_custom_endpoint}.
	DataCatalogCustomEndpoint *string `field:"optional" json:"dataCatalogCustomEndpoint" yaml:"dataCatalogCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#dataflow_custom_endpoint GoogleProvider#dataflow_custom_endpoint}.
	DataflowCustomEndpoint *string `field:"optional" json:"dataflowCustomEndpoint" yaml:"dataflowCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#data_fusion_custom_endpoint GoogleProvider#data_fusion_custom_endpoint}.
	DataFusionCustomEndpoint *string `field:"optional" json:"dataFusionCustomEndpoint" yaml:"dataFusionCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#data_loss_prevention_custom_endpoint GoogleProvider#data_loss_prevention_custom_endpoint}.
	DataLossPreventionCustomEndpoint *string `field:"optional" json:"dataLossPreventionCustomEndpoint" yaml:"dataLossPreventionCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#dataplex_custom_endpoint GoogleProvider#dataplex_custom_endpoint}.
	DataplexCustomEndpoint *string `field:"optional" json:"dataplexCustomEndpoint" yaml:"dataplexCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#dataproc_custom_endpoint GoogleProvider#dataproc_custom_endpoint}.
	DataprocCustomEndpoint *string `field:"optional" json:"dataprocCustomEndpoint" yaml:"dataprocCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#dataproc_metastore_custom_endpoint GoogleProvider#dataproc_metastore_custom_endpoint}.
	DataprocMetastoreCustomEndpoint *string `field:"optional" json:"dataprocMetastoreCustomEndpoint" yaml:"dataprocMetastoreCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#datastore_custom_endpoint GoogleProvider#datastore_custom_endpoint}.
	DatastoreCustomEndpoint *string `field:"optional" json:"datastoreCustomEndpoint" yaml:"datastoreCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#datastream_custom_endpoint GoogleProvider#datastream_custom_endpoint}.
	DatastreamCustomEndpoint *string `field:"optional" json:"datastreamCustomEndpoint" yaml:"datastreamCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#deployment_manager_custom_endpoint GoogleProvider#deployment_manager_custom_endpoint}.
	DeploymentManagerCustomEndpoint *string `field:"optional" json:"deploymentManagerCustomEndpoint" yaml:"deploymentManagerCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#dialogflow_custom_endpoint GoogleProvider#dialogflow_custom_endpoint}.
	DialogflowCustomEndpoint *string `field:"optional" json:"dialogflowCustomEndpoint" yaml:"dialogflowCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#dialogflow_cx_custom_endpoint GoogleProvider#dialogflow_cx_custom_endpoint}.
	DialogflowCxCustomEndpoint *string `field:"optional" json:"dialogflowCxCustomEndpoint" yaml:"dialogflowCxCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#dns_custom_endpoint GoogleProvider#dns_custom_endpoint}.
	DnsCustomEndpoint *string `field:"optional" json:"dnsCustomEndpoint" yaml:"dnsCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#document_ai_custom_endpoint GoogleProvider#document_ai_custom_endpoint}.
	DocumentAiCustomEndpoint *string `field:"optional" json:"documentAiCustomEndpoint" yaml:"documentAiCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#essential_contacts_custom_endpoint GoogleProvider#essential_contacts_custom_endpoint}.
	EssentialContactsCustomEndpoint *string `field:"optional" json:"essentialContactsCustomEndpoint" yaml:"essentialContactsCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#eventarc_custom_endpoint GoogleProvider#eventarc_custom_endpoint}.
	EventarcCustomEndpoint *string `field:"optional" json:"eventarcCustomEndpoint" yaml:"eventarcCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#filestore_custom_endpoint GoogleProvider#filestore_custom_endpoint}.
	FilestoreCustomEndpoint *string `field:"optional" json:"filestoreCustomEndpoint" yaml:"filestoreCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#firebaserules_custom_endpoint GoogleProvider#firebaserules_custom_endpoint}.
	FirebaserulesCustomEndpoint *string `field:"optional" json:"firebaserulesCustomEndpoint" yaml:"firebaserulesCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#firestore_custom_endpoint GoogleProvider#firestore_custom_endpoint}.
	FirestoreCustomEndpoint *string `field:"optional" json:"firestoreCustomEndpoint" yaml:"firestoreCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#game_services_custom_endpoint GoogleProvider#game_services_custom_endpoint}.
	GameServicesCustomEndpoint *string `field:"optional" json:"gameServicesCustomEndpoint" yaml:"gameServicesCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#gke_hub_custom_endpoint GoogleProvider#gke_hub_custom_endpoint}.
	GkeHubCustomEndpoint *string `field:"optional" json:"gkeHubCustomEndpoint" yaml:"gkeHubCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#healthcare_custom_endpoint GoogleProvider#healthcare_custom_endpoint}.
	HealthcareCustomEndpoint *string `field:"optional" json:"healthcareCustomEndpoint" yaml:"healthcareCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#iam_beta_custom_endpoint GoogleProvider#iam_beta_custom_endpoint}.
	IamBetaCustomEndpoint *string `field:"optional" json:"iamBetaCustomEndpoint" yaml:"iamBetaCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#iam_credentials_custom_endpoint GoogleProvider#iam_credentials_custom_endpoint}.
	IamCredentialsCustomEndpoint *string `field:"optional" json:"iamCredentialsCustomEndpoint" yaml:"iamCredentialsCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#iam_custom_endpoint GoogleProvider#iam_custom_endpoint}.
	IamCustomEndpoint *string `field:"optional" json:"iamCustomEndpoint" yaml:"iamCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#iap_custom_endpoint GoogleProvider#iap_custom_endpoint}.
	IapCustomEndpoint *string `field:"optional" json:"iapCustomEndpoint" yaml:"iapCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#identity_platform_custom_endpoint GoogleProvider#identity_platform_custom_endpoint}.
	IdentityPlatformCustomEndpoint *string `field:"optional" json:"identityPlatformCustomEndpoint" yaml:"identityPlatformCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#impersonate_service_account GoogleProvider#impersonate_service_account}.
	ImpersonateServiceAccount *string `field:"optional" json:"impersonateServiceAccount" yaml:"impersonateServiceAccount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#impersonate_service_account_delegates GoogleProvider#impersonate_service_account_delegates}.
	ImpersonateServiceAccountDelegates *[]*string `field:"optional" json:"impersonateServiceAccountDelegates" yaml:"impersonateServiceAccountDelegates"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#kms_custom_endpoint GoogleProvider#kms_custom_endpoint}.
	KmsCustomEndpoint *string `field:"optional" json:"kmsCustomEndpoint" yaml:"kmsCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#logging_custom_endpoint GoogleProvider#logging_custom_endpoint}.
	LoggingCustomEndpoint *string `field:"optional" json:"loggingCustomEndpoint" yaml:"loggingCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#memcache_custom_endpoint GoogleProvider#memcache_custom_endpoint}.
	MemcacheCustomEndpoint *string `field:"optional" json:"memcacheCustomEndpoint" yaml:"memcacheCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#ml_engine_custom_endpoint GoogleProvider#ml_engine_custom_endpoint}.
	MlEngineCustomEndpoint *string `field:"optional" json:"mlEngineCustomEndpoint" yaml:"mlEngineCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#monitoring_custom_endpoint GoogleProvider#monitoring_custom_endpoint}.
	MonitoringCustomEndpoint *string `field:"optional" json:"monitoringCustomEndpoint" yaml:"monitoringCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#network_connectivity_custom_endpoint GoogleProvider#network_connectivity_custom_endpoint}.
	NetworkConnectivityCustomEndpoint *string `field:"optional" json:"networkConnectivityCustomEndpoint" yaml:"networkConnectivityCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#network_management_custom_endpoint GoogleProvider#network_management_custom_endpoint}.
	NetworkManagementCustomEndpoint *string `field:"optional" json:"networkManagementCustomEndpoint" yaml:"networkManagementCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#network_services_custom_endpoint GoogleProvider#network_services_custom_endpoint}.
	NetworkServicesCustomEndpoint *string `field:"optional" json:"networkServicesCustomEndpoint" yaml:"networkServicesCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#notebooks_custom_endpoint GoogleProvider#notebooks_custom_endpoint}.
	NotebooksCustomEndpoint *string `field:"optional" json:"notebooksCustomEndpoint" yaml:"notebooksCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#org_policy_custom_endpoint GoogleProvider#org_policy_custom_endpoint}.
	OrgPolicyCustomEndpoint *string `field:"optional" json:"orgPolicyCustomEndpoint" yaml:"orgPolicyCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#os_config_custom_endpoint GoogleProvider#os_config_custom_endpoint}.
	OsConfigCustomEndpoint *string `field:"optional" json:"osConfigCustomEndpoint" yaml:"osConfigCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#os_login_custom_endpoint GoogleProvider#os_login_custom_endpoint}.
	OsLoginCustomEndpoint *string `field:"optional" json:"osLoginCustomEndpoint" yaml:"osLoginCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#privateca_custom_endpoint GoogleProvider#privateca_custom_endpoint}.
	PrivatecaCustomEndpoint *string `field:"optional" json:"privatecaCustomEndpoint" yaml:"privatecaCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#project GoogleProvider#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#pubsub_custom_endpoint GoogleProvider#pubsub_custom_endpoint}.
	PubsubCustomEndpoint *string `field:"optional" json:"pubsubCustomEndpoint" yaml:"pubsubCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#pubsub_lite_custom_endpoint GoogleProvider#pubsub_lite_custom_endpoint}.
	PubsubLiteCustomEndpoint *string `field:"optional" json:"pubsubLiteCustomEndpoint" yaml:"pubsubLiteCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#recaptcha_enterprise_custom_endpoint GoogleProvider#recaptcha_enterprise_custom_endpoint}.
	RecaptchaEnterpriseCustomEndpoint *string `field:"optional" json:"recaptchaEnterpriseCustomEndpoint" yaml:"recaptchaEnterpriseCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#redis_custom_endpoint GoogleProvider#redis_custom_endpoint}.
	RedisCustomEndpoint *string `field:"optional" json:"redisCustomEndpoint" yaml:"redisCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#region GoogleProvider#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#request_reason GoogleProvider#request_reason}.
	RequestReason *string `field:"optional" json:"requestReason" yaml:"requestReason"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#request_timeout GoogleProvider#request_timeout}.
	RequestTimeout *string `field:"optional" json:"requestTimeout" yaml:"requestTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#resource_manager_custom_endpoint GoogleProvider#resource_manager_custom_endpoint}.
	ResourceManagerCustomEndpoint *string `field:"optional" json:"resourceManagerCustomEndpoint" yaml:"resourceManagerCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#resource_manager_v3_custom_endpoint GoogleProvider#resource_manager_v3_custom_endpoint}.
	ResourceManagerV3CustomEndpoint *string `field:"optional" json:"resourceManagerV3CustomEndpoint" yaml:"resourceManagerV3CustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#scopes GoogleProvider#scopes}.
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#secret_manager_custom_endpoint GoogleProvider#secret_manager_custom_endpoint}.
	SecretManagerCustomEndpoint *string `field:"optional" json:"secretManagerCustomEndpoint" yaml:"secretManagerCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#security_center_custom_endpoint GoogleProvider#security_center_custom_endpoint}.
	SecurityCenterCustomEndpoint *string `field:"optional" json:"securityCenterCustomEndpoint" yaml:"securityCenterCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#service_management_custom_endpoint GoogleProvider#service_management_custom_endpoint}.
	ServiceManagementCustomEndpoint *string `field:"optional" json:"serviceManagementCustomEndpoint" yaml:"serviceManagementCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#service_networking_custom_endpoint GoogleProvider#service_networking_custom_endpoint}.
	ServiceNetworkingCustomEndpoint *string `field:"optional" json:"serviceNetworkingCustomEndpoint" yaml:"serviceNetworkingCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#service_usage_custom_endpoint GoogleProvider#service_usage_custom_endpoint}.
	ServiceUsageCustomEndpoint *string `field:"optional" json:"serviceUsageCustomEndpoint" yaml:"serviceUsageCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#source_repo_custom_endpoint GoogleProvider#source_repo_custom_endpoint}.
	SourceRepoCustomEndpoint *string `field:"optional" json:"sourceRepoCustomEndpoint" yaml:"sourceRepoCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#spanner_custom_endpoint GoogleProvider#spanner_custom_endpoint}.
	SpannerCustomEndpoint *string `field:"optional" json:"spannerCustomEndpoint" yaml:"spannerCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#sql_custom_endpoint GoogleProvider#sql_custom_endpoint}.
	SqlCustomEndpoint *string `field:"optional" json:"sqlCustomEndpoint" yaml:"sqlCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#storage_custom_endpoint GoogleProvider#storage_custom_endpoint}.
	StorageCustomEndpoint *string `field:"optional" json:"storageCustomEndpoint" yaml:"storageCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#storage_transfer_custom_endpoint GoogleProvider#storage_transfer_custom_endpoint}.
	StorageTransferCustomEndpoint *string `field:"optional" json:"storageTransferCustomEndpoint" yaml:"storageTransferCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#tags_custom_endpoint GoogleProvider#tags_custom_endpoint}.
	TagsCustomEndpoint *string `field:"optional" json:"tagsCustomEndpoint" yaml:"tagsCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#tpu_custom_endpoint GoogleProvider#tpu_custom_endpoint}.
	TpuCustomEndpoint *string `field:"optional" json:"tpuCustomEndpoint" yaml:"tpuCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#user_project_override GoogleProvider#user_project_override}.
	UserProjectOverride interface{} `field:"optional" json:"userProjectOverride" yaml:"userProjectOverride"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#vertex_ai_custom_endpoint GoogleProvider#vertex_ai_custom_endpoint}.
	VertexAiCustomEndpoint *string `field:"optional" json:"vertexAiCustomEndpoint" yaml:"vertexAiCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#vpc_access_custom_endpoint GoogleProvider#vpc_access_custom_endpoint}.
	VpcAccessCustomEndpoint *string `field:"optional" json:"vpcAccessCustomEndpoint" yaml:"vpcAccessCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#workflows_custom_endpoint GoogleProvider#workflows_custom_endpoint}.
	WorkflowsCustomEndpoint *string `field:"optional" json:"workflowsCustomEndpoint" yaml:"workflowsCustomEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google#zone GoogleProvider#zone}.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

