package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v8/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/4.75.1/docs google}.
type GoogleProvider interface {
	cdktf.TerraformProvider
	AccessApprovalCustomEndpoint() *string
	SetAccessApprovalCustomEndpoint(val *string)
	AccessApprovalCustomEndpointInput() *string
	AccessContextManagerCustomEndpoint() *string
	SetAccessContextManagerCustomEndpoint(val *string)
	AccessContextManagerCustomEndpointInput() *string
	AccessToken() *string
	SetAccessToken(val *string)
	AccessTokenInput() *string
	ActiveDirectoryCustomEndpoint() *string
	SetActiveDirectoryCustomEndpoint(val *string)
	ActiveDirectoryCustomEndpointInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AlloydbCustomEndpoint() *string
	SetAlloydbCustomEndpoint(val *string)
	AlloydbCustomEndpointInput() *string
	ApigeeCustomEndpoint() *string
	SetApigeeCustomEndpoint(val *string)
	ApigeeCustomEndpointInput() *string
	ApikeysCustomEndpoint() *string
	SetApikeysCustomEndpoint(val *string)
	ApikeysCustomEndpointInput() *string
	AppEngineCustomEndpoint() *string
	SetAppEngineCustomEndpoint(val *string)
	AppEngineCustomEndpointInput() *string
	ArtifactRegistryCustomEndpoint() *string
	SetArtifactRegistryCustomEndpoint(val *string)
	ArtifactRegistryCustomEndpointInput() *string
	AssuredWorkloadsCustomEndpoint() *string
	SetAssuredWorkloadsCustomEndpoint(val *string)
	AssuredWorkloadsCustomEndpointInput() *string
	Batching() *GoogleProviderBatching
	SetBatching(val *GoogleProviderBatching)
	BatchingInput() *GoogleProviderBatching
	BeyondcorpCustomEndpoint() *string
	SetBeyondcorpCustomEndpoint(val *string)
	BeyondcorpCustomEndpointInput() *string
	BigqueryAnalyticsHubCustomEndpoint() *string
	SetBigqueryAnalyticsHubCustomEndpoint(val *string)
	BigqueryAnalyticsHubCustomEndpointInput() *string
	BigqueryConnectionCustomEndpoint() *string
	SetBigqueryConnectionCustomEndpoint(val *string)
	BigqueryConnectionCustomEndpointInput() *string
	BigQueryCustomEndpoint() *string
	SetBigQueryCustomEndpoint(val *string)
	BigQueryCustomEndpointInput() *string
	BigqueryDatapolicyCustomEndpoint() *string
	SetBigqueryDatapolicyCustomEndpoint(val *string)
	BigqueryDatapolicyCustomEndpointInput() *string
	BigqueryDataTransferCustomEndpoint() *string
	SetBigqueryDataTransferCustomEndpoint(val *string)
	BigqueryDataTransferCustomEndpointInput() *string
	BigqueryReservationCustomEndpoint() *string
	SetBigqueryReservationCustomEndpoint(val *string)
	BigqueryReservationCustomEndpointInput() *string
	BigtableCustomEndpoint() *string
	SetBigtableCustomEndpoint(val *string)
	BigtableCustomEndpointInput() *string
	BillingCustomEndpoint() *string
	SetBillingCustomEndpoint(val *string)
	BillingCustomEndpointInput() *string
	BillingProject() *string
	SetBillingProject(val *string)
	BillingProjectInput() *string
	BinaryAuthorizationCustomEndpoint() *string
	SetBinaryAuthorizationCustomEndpoint(val *string)
	BinaryAuthorizationCustomEndpointInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateManagerCustomEndpoint() *string
	SetCertificateManagerCustomEndpoint(val *string)
	CertificateManagerCustomEndpointInput() *string
	CloudAssetCustomEndpoint() *string
	SetCloudAssetCustomEndpoint(val *string)
	CloudAssetCustomEndpointInput() *string
	CloudBillingCustomEndpoint() *string
	SetCloudBillingCustomEndpoint(val *string)
	CloudBillingCustomEndpointInput() *string
	CloudBuildCustomEndpoint() *string
	SetCloudBuildCustomEndpoint(val *string)
	CloudBuildCustomEndpointInput() *string
	Cloudbuildv2CustomEndpoint() *string
	SetCloudbuildv2CustomEndpoint(val *string)
	Cloudbuildv2CustomEndpointInput() *string
	CloudBuildWorkerPoolCustomEndpoint() *string
	SetCloudBuildWorkerPoolCustomEndpoint(val *string)
	CloudBuildWorkerPoolCustomEndpointInput() *string
	ClouddeployCustomEndpoint() *string
	SetClouddeployCustomEndpoint(val *string)
	ClouddeployCustomEndpointInput() *string
	Cloudfunctions2CustomEndpoint() *string
	SetCloudfunctions2CustomEndpoint(val *string)
	Cloudfunctions2CustomEndpointInput() *string
	CloudFunctionsCustomEndpoint() *string
	SetCloudFunctionsCustomEndpoint(val *string)
	CloudFunctionsCustomEndpointInput() *string
	CloudIdentityCustomEndpoint() *string
	SetCloudIdentityCustomEndpoint(val *string)
	CloudIdentityCustomEndpointInput() *string
	CloudIdsCustomEndpoint() *string
	SetCloudIdsCustomEndpoint(val *string)
	CloudIdsCustomEndpointInput() *string
	CloudIotCustomEndpoint() *string
	SetCloudIotCustomEndpoint(val *string)
	CloudIotCustomEndpointInput() *string
	CloudResourceManagerCustomEndpoint() *string
	SetCloudResourceManagerCustomEndpoint(val *string)
	CloudResourceManagerCustomEndpointInput() *string
	CloudRunCustomEndpoint() *string
	SetCloudRunCustomEndpoint(val *string)
	CloudRunCustomEndpointInput() *string
	CloudRunV2CustomEndpoint() *string
	SetCloudRunV2CustomEndpoint(val *string)
	CloudRunV2CustomEndpointInput() *string
	CloudSchedulerCustomEndpoint() *string
	SetCloudSchedulerCustomEndpoint(val *string)
	CloudSchedulerCustomEndpointInput() *string
	CloudTasksCustomEndpoint() *string
	SetCloudTasksCustomEndpoint(val *string)
	CloudTasksCustomEndpointInput() *string
	ComposerCustomEndpoint() *string
	SetComposerCustomEndpoint(val *string)
	ComposerCustomEndpointInput() *string
	ComputeCustomEndpoint() *string
	SetComputeCustomEndpoint(val *string)
	ComputeCustomEndpointInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerAnalysisCustomEndpoint() *string
	SetContainerAnalysisCustomEndpoint(val *string)
	ContainerAnalysisCustomEndpointInput() *string
	ContainerAttachedCustomEndpoint() *string
	SetContainerAttachedCustomEndpoint(val *string)
	ContainerAttachedCustomEndpointInput() *string
	ContainerAwsCustomEndpoint() *string
	SetContainerAwsCustomEndpoint(val *string)
	ContainerAwsCustomEndpointInput() *string
	ContainerAzureCustomEndpoint() *string
	SetContainerAzureCustomEndpoint(val *string)
	ContainerAzureCustomEndpointInput() *string
	ContainerCustomEndpoint() *string
	SetContainerCustomEndpoint(val *string)
	ContainerCustomEndpointInput() *string
	Credentials() *string
	SetCredentials(val *string)
	CredentialsInput() *string
	DatabaseMigrationServiceCustomEndpoint() *string
	SetDatabaseMigrationServiceCustomEndpoint(val *string)
	DatabaseMigrationServiceCustomEndpointInput() *string
	DataCatalogCustomEndpoint() *string
	SetDataCatalogCustomEndpoint(val *string)
	DataCatalogCustomEndpointInput() *string
	DataflowCustomEndpoint() *string
	SetDataflowCustomEndpoint(val *string)
	DataflowCustomEndpointInput() *string
	DataFusionCustomEndpoint() *string
	SetDataFusionCustomEndpoint(val *string)
	DataFusionCustomEndpointInput() *string
	DataLossPreventionCustomEndpoint() *string
	SetDataLossPreventionCustomEndpoint(val *string)
	DataLossPreventionCustomEndpointInput() *string
	DataplexCustomEndpoint() *string
	SetDataplexCustomEndpoint(val *string)
	DataplexCustomEndpointInput() *string
	DataprocCustomEndpoint() *string
	SetDataprocCustomEndpoint(val *string)
	DataprocCustomEndpointInput() *string
	DataprocMetastoreCustomEndpoint() *string
	SetDataprocMetastoreCustomEndpoint(val *string)
	DataprocMetastoreCustomEndpointInput() *string
	DatastoreCustomEndpoint() *string
	SetDatastoreCustomEndpoint(val *string)
	DatastoreCustomEndpointInput() *string
	DatastreamCustomEndpoint() *string
	SetDatastreamCustomEndpoint(val *string)
	DatastreamCustomEndpointInput() *string
	DeploymentManagerCustomEndpoint() *string
	SetDeploymentManagerCustomEndpoint(val *string)
	DeploymentManagerCustomEndpointInput() *string
	DialogflowCustomEndpoint() *string
	SetDialogflowCustomEndpoint(val *string)
	DialogflowCustomEndpointInput() *string
	DialogflowCxCustomEndpoint() *string
	SetDialogflowCxCustomEndpoint(val *string)
	DialogflowCxCustomEndpointInput() *string
	DnsCustomEndpoint() *string
	SetDnsCustomEndpoint(val *string)
	DnsCustomEndpointInput() *string
	DocumentAiCustomEndpoint() *string
	SetDocumentAiCustomEndpoint(val *string)
	DocumentAiCustomEndpointInput() *string
	EssentialContactsCustomEndpoint() *string
	SetEssentialContactsCustomEndpoint(val *string)
	EssentialContactsCustomEndpointInput() *string
	EventarcCustomEndpoint() *string
	SetEventarcCustomEndpoint(val *string)
	EventarcCustomEndpointInput() *string
	FilestoreCustomEndpoint() *string
	SetFilestoreCustomEndpoint(val *string)
	FilestoreCustomEndpointInput() *string
	FirebaserulesCustomEndpoint() *string
	SetFirebaserulesCustomEndpoint(val *string)
	FirebaserulesCustomEndpointInput() *string
	FirestoreCustomEndpoint() *string
	SetFirestoreCustomEndpoint(val *string)
	FirestoreCustomEndpointInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GameServicesCustomEndpoint() *string
	SetGameServicesCustomEndpoint(val *string)
	GameServicesCustomEndpointInput() *string
	GkeBackupCustomEndpoint() *string
	SetGkeBackupCustomEndpoint(val *string)
	GkeBackupCustomEndpointInput() *string
	GkeHub2CustomEndpoint() *string
	SetGkeHub2CustomEndpoint(val *string)
	GkeHub2CustomEndpointInput() *string
	GkeHubCustomEndpoint() *string
	SetGkeHubCustomEndpoint(val *string)
	GkeHubCustomEndpointInput() *string
	HealthcareCustomEndpoint() *string
	SetHealthcareCustomEndpoint(val *string)
	HealthcareCustomEndpointInput() *string
	Iam2CustomEndpoint() *string
	SetIam2CustomEndpoint(val *string)
	Iam2CustomEndpointInput() *string
	IamBetaCustomEndpoint() *string
	SetIamBetaCustomEndpoint(val *string)
	IamBetaCustomEndpointInput() *string
	IamCredentialsCustomEndpoint() *string
	SetIamCredentialsCustomEndpoint(val *string)
	IamCredentialsCustomEndpointInput() *string
	IamCustomEndpoint() *string
	SetIamCustomEndpoint(val *string)
	IamCustomEndpointInput() *string
	IamWorkforcePoolCustomEndpoint() *string
	SetIamWorkforcePoolCustomEndpoint(val *string)
	IamWorkforcePoolCustomEndpointInput() *string
	IapCustomEndpoint() *string
	SetIapCustomEndpoint(val *string)
	IapCustomEndpointInput() *string
	IdentityPlatformCustomEndpoint() *string
	SetIdentityPlatformCustomEndpoint(val *string)
	IdentityPlatformCustomEndpointInput() *string
	ImpersonateServiceAccount() *string
	SetImpersonateServiceAccount(val *string)
	ImpersonateServiceAccountDelegates() *[]*string
	SetImpersonateServiceAccountDelegates(val *[]*string)
	ImpersonateServiceAccountDelegatesInput() *[]*string
	ImpersonateServiceAccountInput() *string
	KmsCustomEndpoint() *string
	SetKmsCustomEndpoint(val *string)
	KmsCustomEndpointInput() *string
	LoggingCustomEndpoint() *string
	SetLoggingCustomEndpoint(val *string)
	LoggingCustomEndpointInput() *string
	LookerCustomEndpoint() *string
	SetLookerCustomEndpoint(val *string)
	LookerCustomEndpointInput() *string
	MemcacheCustomEndpoint() *string
	SetMemcacheCustomEndpoint(val *string)
	MemcacheCustomEndpointInput() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	MlEngineCustomEndpoint() *string
	SetMlEngineCustomEndpoint(val *string)
	MlEngineCustomEndpointInput() *string
	MonitoringCustomEndpoint() *string
	SetMonitoringCustomEndpoint(val *string)
	MonitoringCustomEndpointInput() *string
	NetworkConnectivityCustomEndpoint() *string
	SetNetworkConnectivityCustomEndpoint(val *string)
	NetworkConnectivityCustomEndpointInput() *string
	NetworkManagementCustomEndpoint() *string
	SetNetworkManagementCustomEndpoint(val *string)
	NetworkManagementCustomEndpointInput() *string
	NetworkSecurityCustomEndpoint() *string
	SetNetworkSecurityCustomEndpoint(val *string)
	NetworkSecurityCustomEndpointInput() *string
	NetworkServicesCustomEndpoint() *string
	SetNetworkServicesCustomEndpoint(val *string)
	NetworkServicesCustomEndpointInput() *string
	// The tree node.
	Node() constructs.Node
	NotebooksCustomEndpoint() *string
	SetNotebooksCustomEndpoint(val *string)
	NotebooksCustomEndpointInput() *string
	OrgPolicyCustomEndpoint() *string
	SetOrgPolicyCustomEndpoint(val *string)
	OrgPolicyCustomEndpointInput() *string
	OsConfigCustomEndpoint() *string
	SetOsConfigCustomEndpoint(val *string)
	OsConfigCustomEndpointInput() *string
	OsLoginCustomEndpoint() *string
	SetOsLoginCustomEndpoint(val *string)
	OsLoginCustomEndpointInput() *string
	PrivatecaCustomEndpoint() *string
	SetPrivatecaCustomEndpoint(val *string)
	PrivatecaCustomEndpointInput() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	PublicCaCustomEndpoint() *string
	SetPublicCaCustomEndpoint(val *string)
	PublicCaCustomEndpointInput() *string
	PubsubCustomEndpoint() *string
	SetPubsubCustomEndpoint(val *string)
	PubsubCustomEndpointInput() *string
	PubsubLiteCustomEndpoint() *string
	SetPubsubLiteCustomEndpoint(val *string)
	PubsubLiteCustomEndpointInput() *string
	// Experimental.
	RawOverrides() interface{}
	RecaptchaEnterpriseCustomEndpoint() *string
	SetRecaptchaEnterpriseCustomEndpoint(val *string)
	RecaptchaEnterpriseCustomEndpointInput() *string
	RedisCustomEndpoint() *string
	SetRedisCustomEndpoint(val *string)
	RedisCustomEndpointInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RequestReason() *string
	SetRequestReason(val *string)
	RequestReasonInput() *string
	RequestTimeout() *string
	SetRequestTimeout(val *string)
	RequestTimeoutInput() *string
	ResourceManagerCustomEndpoint() *string
	SetResourceManagerCustomEndpoint(val *string)
	ResourceManagerCustomEndpointInput() *string
	ResourceManagerV3CustomEndpoint() *string
	SetResourceManagerV3CustomEndpoint(val *string)
	ResourceManagerV3CustomEndpointInput() *string
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	SecretManagerCustomEndpoint() *string
	SetSecretManagerCustomEndpoint(val *string)
	SecretManagerCustomEndpointInput() *string
	SecurityCenterCustomEndpoint() *string
	SetSecurityCenterCustomEndpoint(val *string)
	SecurityCenterCustomEndpointInput() *string
	ServiceManagementCustomEndpoint() *string
	SetServiceManagementCustomEndpoint(val *string)
	ServiceManagementCustomEndpointInput() *string
	ServiceNetworkingCustomEndpoint() *string
	SetServiceNetworkingCustomEndpoint(val *string)
	ServiceNetworkingCustomEndpointInput() *string
	ServiceUsageCustomEndpoint() *string
	SetServiceUsageCustomEndpoint(val *string)
	ServiceUsageCustomEndpointInput() *string
	SourceRepoCustomEndpoint() *string
	SetSourceRepoCustomEndpoint(val *string)
	SourceRepoCustomEndpointInput() *string
	SpannerCustomEndpoint() *string
	SetSpannerCustomEndpoint(val *string)
	SpannerCustomEndpointInput() *string
	SqlCustomEndpoint() *string
	SetSqlCustomEndpoint(val *string)
	SqlCustomEndpointInput() *string
	StorageCustomEndpoint() *string
	SetStorageCustomEndpoint(val *string)
	StorageCustomEndpointInput() *string
	StorageTransferCustomEndpoint() *string
	SetStorageTransferCustomEndpoint(val *string)
	StorageTransferCustomEndpointInput() *string
	TagsCustomEndpoint() *string
	SetTagsCustomEndpoint(val *string)
	TagsCustomEndpointInput() *string
	TagsLocationCustomEndpoint() *string
	SetTagsLocationCustomEndpoint(val *string)
	TagsLocationCustomEndpointInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	TpuCustomEndpoint() *string
	SetTpuCustomEndpoint(val *string)
	TpuCustomEndpointInput() *string
	UserProjectOverride() interface{}
	SetUserProjectOverride(val interface{})
	UserProjectOverrideInput() interface{}
	VertexAiCustomEndpoint() *string
	SetVertexAiCustomEndpoint(val *string)
	VertexAiCustomEndpointInput() *string
	VpcAccessCustomEndpoint() *string
	SetVpcAccessCustomEndpoint(val *string)
	VpcAccessCustomEndpointInput() *string
	WorkflowsCustomEndpoint() *string
	SetWorkflowsCustomEndpoint(val *string)
	WorkflowsCustomEndpointInput() *string
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAccessApprovalCustomEndpoint()
	ResetAccessContextManagerCustomEndpoint()
	ResetAccessToken()
	ResetActiveDirectoryCustomEndpoint()
	ResetAlias()
	ResetAlloydbCustomEndpoint()
	ResetApigeeCustomEndpoint()
	ResetApikeysCustomEndpoint()
	ResetAppEngineCustomEndpoint()
	ResetArtifactRegistryCustomEndpoint()
	ResetAssuredWorkloadsCustomEndpoint()
	ResetBatching()
	ResetBeyondcorpCustomEndpoint()
	ResetBigqueryAnalyticsHubCustomEndpoint()
	ResetBigqueryConnectionCustomEndpoint()
	ResetBigQueryCustomEndpoint()
	ResetBigqueryDatapolicyCustomEndpoint()
	ResetBigqueryDataTransferCustomEndpoint()
	ResetBigqueryReservationCustomEndpoint()
	ResetBigtableCustomEndpoint()
	ResetBillingCustomEndpoint()
	ResetBillingProject()
	ResetBinaryAuthorizationCustomEndpoint()
	ResetCertificateManagerCustomEndpoint()
	ResetCloudAssetCustomEndpoint()
	ResetCloudBillingCustomEndpoint()
	ResetCloudBuildCustomEndpoint()
	ResetCloudbuildv2CustomEndpoint()
	ResetCloudBuildWorkerPoolCustomEndpoint()
	ResetClouddeployCustomEndpoint()
	ResetCloudfunctions2CustomEndpoint()
	ResetCloudFunctionsCustomEndpoint()
	ResetCloudIdentityCustomEndpoint()
	ResetCloudIdsCustomEndpoint()
	ResetCloudIotCustomEndpoint()
	ResetCloudResourceManagerCustomEndpoint()
	ResetCloudRunCustomEndpoint()
	ResetCloudRunV2CustomEndpoint()
	ResetCloudSchedulerCustomEndpoint()
	ResetCloudTasksCustomEndpoint()
	ResetComposerCustomEndpoint()
	ResetComputeCustomEndpoint()
	ResetContainerAnalysisCustomEndpoint()
	ResetContainerAttachedCustomEndpoint()
	ResetContainerAwsCustomEndpoint()
	ResetContainerAzureCustomEndpoint()
	ResetContainerCustomEndpoint()
	ResetCredentials()
	ResetDatabaseMigrationServiceCustomEndpoint()
	ResetDataCatalogCustomEndpoint()
	ResetDataflowCustomEndpoint()
	ResetDataFusionCustomEndpoint()
	ResetDataLossPreventionCustomEndpoint()
	ResetDataplexCustomEndpoint()
	ResetDataprocCustomEndpoint()
	ResetDataprocMetastoreCustomEndpoint()
	ResetDatastoreCustomEndpoint()
	ResetDatastreamCustomEndpoint()
	ResetDeploymentManagerCustomEndpoint()
	ResetDialogflowCustomEndpoint()
	ResetDialogflowCxCustomEndpoint()
	ResetDnsCustomEndpoint()
	ResetDocumentAiCustomEndpoint()
	ResetEssentialContactsCustomEndpoint()
	ResetEventarcCustomEndpoint()
	ResetFilestoreCustomEndpoint()
	ResetFirebaserulesCustomEndpoint()
	ResetFirestoreCustomEndpoint()
	ResetGameServicesCustomEndpoint()
	ResetGkeBackupCustomEndpoint()
	ResetGkeHub2CustomEndpoint()
	ResetGkeHubCustomEndpoint()
	ResetHealthcareCustomEndpoint()
	ResetIam2CustomEndpoint()
	ResetIamBetaCustomEndpoint()
	ResetIamCredentialsCustomEndpoint()
	ResetIamCustomEndpoint()
	ResetIamWorkforcePoolCustomEndpoint()
	ResetIapCustomEndpoint()
	ResetIdentityPlatformCustomEndpoint()
	ResetImpersonateServiceAccount()
	ResetImpersonateServiceAccountDelegates()
	ResetKmsCustomEndpoint()
	ResetLoggingCustomEndpoint()
	ResetLookerCustomEndpoint()
	ResetMemcacheCustomEndpoint()
	ResetMlEngineCustomEndpoint()
	ResetMonitoringCustomEndpoint()
	ResetNetworkConnectivityCustomEndpoint()
	ResetNetworkManagementCustomEndpoint()
	ResetNetworkSecurityCustomEndpoint()
	ResetNetworkServicesCustomEndpoint()
	ResetNotebooksCustomEndpoint()
	ResetOrgPolicyCustomEndpoint()
	ResetOsConfigCustomEndpoint()
	ResetOsLoginCustomEndpoint()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivatecaCustomEndpoint()
	ResetProject()
	ResetPublicCaCustomEndpoint()
	ResetPubsubCustomEndpoint()
	ResetPubsubLiteCustomEndpoint()
	ResetRecaptchaEnterpriseCustomEndpoint()
	ResetRedisCustomEndpoint()
	ResetRegion()
	ResetRequestReason()
	ResetRequestTimeout()
	ResetResourceManagerCustomEndpoint()
	ResetResourceManagerV3CustomEndpoint()
	ResetScopes()
	ResetSecretManagerCustomEndpoint()
	ResetSecurityCenterCustomEndpoint()
	ResetServiceManagementCustomEndpoint()
	ResetServiceNetworkingCustomEndpoint()
	ResetServiceUsageCustomEndpoint()
	ResetSourceRepoCustomEndpoint()
	ResetSpannerCustomEndpoint()
	ResetSqlCustomEndpoint()
	ResetStorageCustomEndpoint()
	ResetStorageTransferCustomEndpoint()
	ResetTagsCustomEndpoint()
	ResetTagsLocationCustomEndpoint()
	ResetTpuCustomEndpoint()
	ResetUserProjectOverride()
	ResetVertexAiCustomEndpoint()
	ResetVpcAccessCustomEndpoint()
	ResetWorkflowsCustomEndpoint()
	ResetZone()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleProvider
type jsiiProxy_GoogleProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_GoogleProvider) AccessApprovalCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessApprovalCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) AccessApprovalCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessApprovalCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) AccessContextManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessContextManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) AccessContextManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessContextManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) AccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ActiveDirectoryCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ActiveDirectoryCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) AlloydbCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alloydbCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) AlloydbCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alloydbCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ApigeeCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apigeeCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ApigeeCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apigeeCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ApikeysCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apikeysCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ApikeysCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apikeysCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) AppEngineCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appEngineCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) AppEngineCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appEngineCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ArtifactRegistryCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactRegistryCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ArtifactRegistryCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactRegistryCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) AssuredWorkloadsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assuredWorkloadsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) AssuredWorkloadsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assuredWorkloadsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) Batching() *GoogleProviderBatching {
	var returns *GoogleProviderBatching
	_jsii_.Get(
		j,
		"batching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BatchingInput() *GoogleProviderBatching {
	var returns *GoogleProviderBatching
	_jsii_.Get(
		j,
		"batchingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BeyondcorpCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beyondcorpCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BeyondcorpCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beyondcorpCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BigqueryAnalyticsHubCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryAnalyticsHubCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BigqueryAnalyticsHubCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryAnalyticsHubCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BigqueryConnectionCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryConnectionCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BigqueryConnectionCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryConnectionCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BigQueryCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigQueryCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BigQueryCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigQueryCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BigqueryDatapolicyCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryDatapolicyCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BigqueryDatapolicyCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryDatapolicyCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BigqueryDataTransferCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryDataTransferCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BigqueryDataTransferCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryDataTransferCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BigqueryReservationCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryReservationCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BigqueryReservationCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryReservationCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BigtableCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigtableCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BigtableCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigtableCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BillingCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BillingCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BillingProject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BillingProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingProjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BinaryAuthorizationCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryAuthorizationCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) BinaryAuthorizationCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryAuthorizationCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CertificateManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CertificateManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudAssetCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudAssetCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudAssetCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudAssetCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudBillingCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudBillingCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudBillingCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudBillingCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudBuildCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudBuildCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudBuildCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudBuildCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) Cloudbuildv2CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudbuildv2CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) Cloudbuildv2CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudbuildv2CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudBuildWorkerPoolCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudBuildWorkerPoolCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudBuildWorkerPoolCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudBuildWorkerPoolCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ClouddeployCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clouddeployCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ClouddeployCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clouddeployCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) Cloudfunctions2CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudfunctions2CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) Cloudfunctions2CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudfunctions2CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudFunctionsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFunctionsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudFunctionsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFunctionsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudIdentityCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudIdentityCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudIdentityCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudIdentityCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudIdsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudIdsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudIdsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudIdsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudIotCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudIotCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudIotCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudIotCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudResourceManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudResourceManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudResourceManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudResourceManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudRunCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudRunCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudRunCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudRunCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudRunV2CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudRunV2CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudRunV2CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudRunV2CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudSchedulerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudSchedulerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudSchedulerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudSchedulerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudTasksCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudTasksCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CloudTasksCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudTasksCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ComposerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"composerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ComposerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"composerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ComputeCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ComputeCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ContainerAnalysisCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAnalysisCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ContainerAnalysisCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAnalysisCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ContainerAttachedCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAttachedCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ContainerAttachedCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAttachedCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ContainerAwsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAwsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ContainerAwsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAwsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ContainerAzureCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAzureCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ContainerAzureCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAzureCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ContainerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ContainerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) Credentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) CredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DatabaseMigrationServiceCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseMigrationServiceCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DatabaseMigrationServiceCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseMigrationServiceCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DataCatalogCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCatalogCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DataCatalogCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCatalogCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DataflowCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataflowCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DataflowCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataflowCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DataFusionCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFusionCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DataFusionCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFusionCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DataLossPreventionCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLossPreventionCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DataLossPreventionCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLossPreventionCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DataplexCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataplexCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DataplexCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataplexCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DataprocCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DataprocCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DataprocMetastoreCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocMetastoreCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DataprocMetastoreCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocMetastoreCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DatastoreCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DatastoreCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DatastreamCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastreamCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DatastreamCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastreamCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DeploymentManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DeploymentManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DialogflowCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialogflowCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DialogflowCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialogflowCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DialogflowCxCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialogflowCxCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DialogflowCxCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialogflowCxCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DnsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DnsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DocumentAiCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentAiCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) DocumentAiCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentAiCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) EssentialContactsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"essentialContactsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) EssentialContactsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"essentialContactsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) EventarcCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventarcCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) EventarcCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventarcCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) FilestoreCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filestoreCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) FilestoreCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filestoreCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) FirebaserulesCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaserulesCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) FirebaserulesCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaserulesCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) FirestoreCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firestoreCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) FirestoreCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firestoreCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) GameServicesCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServicesCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) GameServicesCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServicesCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) GkeBackupCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeBackupCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) GkeBackupCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeBackupCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) GkeHub2CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeHub2CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) GkeHub2CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeHub2CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) GkeHubCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeHubCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) GkeHubCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeHubCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) HealthcareCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthcareCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) HealthcareCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthcareCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) Iam2CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iam2CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) Iam2CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iam2CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) IamBetaCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamBetaCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) IamBetaCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamBetaCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) IamCredentialsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamCredentialsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) IamCredentialsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamCredentialsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) IamCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) IamCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) IamWorkforcePoolCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamWorkforcePoolCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) IamWorkforcePoolCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamWorkforcePoolCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) IapCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iapCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) IapCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iapCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) IdentityPlatformCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPlatformCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) IdentityPlatformCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPlatformCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ImpersonateServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"impersonateServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ImpersonateServiceAccountDelegates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"impersonateServiceAccountDelegates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ImpersonateServiceAccountDelegatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"impersonateServiceAccountDelegatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ImpersonateServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"impersonateServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) KmsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) KmsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) LoggingCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) LoggingCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) LookerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) LookerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) MemcacheCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memcacheCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) MemcacheCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memcacheCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) MlEngineCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlEngineCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) MlEngineCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlEngineCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) MonitoringCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) MonitoringCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) NetworkConnectivityCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectivityCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) NetworkConnectivityCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectivityCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) NetworkManagementCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkManagementCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) NetworkManagementCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkManagementCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) NetworkSecurityCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSecurityCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) NetworkSecurityCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSecurityCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) NetworkServicesCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkServicesCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) NetworkServicesCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkServicesCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) NotebooksCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebooksCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) NotebooksCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebooksCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) OrgPolicyCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgPolicyCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) OrgPolicyCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgPolicyCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) OsConfigCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osConfigCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) OsConfigCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osConfigCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) OsLoginCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osLoginCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) OsLoginCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osLoginCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) PrivatecaCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privatecaCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) PrivatecaCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privatecaCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) PublicCaCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicCaCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) PublicCaCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicCaCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) PubsubCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) PubsubCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) PubsubLiteCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubLiteCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) PubsubLiteCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubLiteCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) RecaptchaEnterpriseCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recaptchaEnterpriseCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) RecaptchaEnterpriseCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recaptchaEnterpriseCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) RedisCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) RedisCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) RequestReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) RequestReasonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestReasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) RequestTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) RequestTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ResourceManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ResourceManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ResourceManagerV3CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceManagerV3CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ResourceManagerV3CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceManagerV3CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) SecretManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) SecretManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) SecurityCenterCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityCenterCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) SecurityCenterCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityCenterCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ServiceManagementCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceManagementCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ServiceManagementCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceManagementCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ServiceNetworkingCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNetworkingCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ServiceNetworkingCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNetworkingCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ServiceUsageCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUsageCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ServiceUsageCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUsageCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) SourceRepoCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRepoCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) SourceRepoCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRepoCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) SpannerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spannerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) SpannerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spannerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) SqlCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) SqlCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) StorageCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) StorageCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) StorageTransferCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTransferCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) StorageTransferCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTransferCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) TagsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) TagsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) TagsLocationCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagsLocationCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) TagsLocationCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagsLocationCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) TpuCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpuCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) TpuCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpuCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) UserProjectOverride() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userProjectOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) UserProjectOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userProjectOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) VertexAiCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vertexAiCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) VertexAiCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vertexAiCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) VpcAccessCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcAccessCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) VpcAccessCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcAccessCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) WorkflowsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) WorkflowsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleProvider) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.75.1/docs google} Resource.
func NewGoogleProvider(scope constructs.Construct, id *string, config *GoogleProviderConfig) GoogleProvider {
	_init_.Initialize()

	if err := validateNewGoogleProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleProvider{}

	_jsii_.Create(
		"@cdktf/provider-google.provider.GoogleProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.75.1/docs google} Resource.
func NewGoogleProvider_Override(g GoogleProvider, scope constructs.Construct, id *string, config *GoogleProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.provider.GoogleProvider",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleProvider)SetAccessApprovalCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"accessApprovalCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetAccessContextManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"accessContextManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetAccessToken(val *string) {
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetActiveDirectoryCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"activeDirectoryCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetAlloydbCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"alloydbCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetApigeeCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"apigeeCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetApikeysCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"apikeysCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetAppEngineCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"appEngineCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetArtifactRegistryCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"artifactRegistryCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetAssuredWorkloadsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"assuredWorkloadsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetBatching(val *GoogleProviderBatching) {
	if err := j.validateSetBatchingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batching",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetBeyondcorpCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"beyondcorpCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetBigqueryAnalyticsHubCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigqueryAnalyticsHubCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetBigqueryConnectionCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigqueryConnectionCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetBigQueryCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigQueryCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetBigqueryDatapolicyCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigqueryDatapolicyCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetBigqueryDataTransferCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigqueryDataTransferCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetBigqueryReservationCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigqueryReservationCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetBigtableCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigtableCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetBillingCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"billingCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetBillingProject(val *string) {
	_jsii_.Set(
		j,
		"billingProject",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetBinaryAuthorizationCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"binaryAuthorizationCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetCertificateManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"certificateManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetCloudAssetCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudAssetCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetCloudBillingCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudBillingCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetCloudBuildCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudBuildCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetCloudbuildv2CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudbuildv2CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetCloudBuildWorkerPoolCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudBuildWorkerPoolCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetClouddeployCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"clouddeployCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetCloudfunctions2CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudfunctions2CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetCloudFunctionsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudFunctionsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetCloudIdentityCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudIdentityCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetCloudIdsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudIdsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetCloudIotCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudIotCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetCloudResourceManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudResourceManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetCloudRunCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudRunCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetCloudRunV2CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudRunV2CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetCloudSchedulerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudSchedulerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetCloudTasksCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudTasksCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetComposerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"composerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetComputeCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"computeCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetContainerAnalysisCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"containerAnalysisCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetContainerAttachedCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"containerAttachedCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetContainerAwsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"containerAwsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetContainerAzureCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"containerAzureCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetContainerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"containerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetCredentials(val *string) {
	_jsii_.Set(
		j,
		"credentials",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetDatabaseMigrationServiceCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"databaseMigrationServiceCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetDataCatalogCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataCatalogCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetDataflowCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataflowCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetDataFusionCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataFusionCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetDataLossPreventionCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataLossPreventionCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetDataplexCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataplexCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetDataprocCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataprocCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetDataprocMetastoreCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataprocMetastoreCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetDatastoreCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"datastoreCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetDatastreamCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"datastreamCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetDeploymentManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"deploymentManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetDialogflowCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dialogflowCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetDialogflowCxCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dialogflowCxCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetDnsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dnsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetDocumentAiCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"documentAiCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetEssentialContactsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"essentialContactsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetEventarcCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"eventarcCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetFilestoreCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"filestoreCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetFirebaserulesCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"firebaserulesCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetFirestoreCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"firestoreCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetGameServicesCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"gameServicesCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetGkeBackupCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"gkeBackupCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetGkeHub2CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"gkeHub2CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetGkeHubCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"gkeHubCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetHealthcareCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"healthcareCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetIam2CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"iam2CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetIamBetaCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"iamBetaCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetIamCredentialsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"iamCredentialsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetIamCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"iamCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetIamWorkforcePoolCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"iamWorkforcePoolCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetIapCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"iapCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetIdentityPlatformCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"identityPlatformCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetImpersonateServiceAccount(val *string) {
	_jsii_.Set(
		j,
		"impersonateServiceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetImpersonateServiceAccountDelegates(val *[]*string) {
	_jsii_.Set(
		j,
		"impersonateServiceAccountDelegates",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetKmsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"kmsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetLoggingCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"loggingCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetLookerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"lookerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetMemcacheCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"memcacheCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetMlEngineCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"mlEngineCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetMonitoringCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"monitoringCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetNetworkConnectivityCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"networkConnectivityCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetNetworkManagementCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"networkManagementCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetNetworkSecurityCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"networkSecurityCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetNetworkServicesCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"networkServicesCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetNotebooksCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"notebooksCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetOrgPolicyCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"orgPolicyCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetOsConfigCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"osConfigCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetOsLoginCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"osLoginCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetPrivatecaCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"privatecaCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetProject(val *string) {
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetPublicCaCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"publicCaCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetPubsubCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"pubsubCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetPubsubLiteCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"pubsubLiteCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetRecaptchaEnterpriseCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"recaptchaEnterpriseCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetRedisCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"redisCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetRequestReason(val *string) {
	_jsii_.Set(
		j,
		"requestReason",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetRequestTimeout(val *string) {
	_jsii_.Set(
		j,
		"requestTimeout",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetResourceManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"resourceManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetResourceManagerV3CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"resourceManagerV3CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetSecretManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"secretManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetSecurityCenterCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"securityCenterCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetServiceManagementCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"serviceManagementCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetServiceNetworkingCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"serviceNetworkingCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetServiceUsageCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"serviceUsageCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetSourceRepoCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"sourceRepoCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetSpannerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"spannerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetSqlCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"sqlCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetStorageCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"storageCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetStorageTransferCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"storageTransferCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetTagsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"tagsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetTagsLocationCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"tagsLocationCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetTpuCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"tpuCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetUserProjectOverride(val interface{}) {
	if err := j.validateSetUserProjectOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userProjectOverride",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetVertexAiCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"vertexAiCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetVpcAccessCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"vpcAccessCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetWorkflowsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"workflowsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleProvider)SetZone(val *string) {
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func GoogleProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.provider.GoogleProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.provider.GoogleProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.provider.GoogleProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.provider.GoogleProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleProvider) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleProvider) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleProvider) ResetAccessApprovalCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessApprovalCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetAccessContextManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessContextManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetAccessToken() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetActiveDirectoryCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetActiveDirectoryCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		g,
		"resetAlias",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetAlloydbCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetAlloydbCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetApigeeCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetApigeeCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetApikeysCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetApikeysCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetAppEngineCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetAppEngineCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetArtifactRegistryCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetArtifactRegistryCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetAssuredWorkloadsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetAssuredWorkloadsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetBatching() {
	_jsii_.InvokeVoid(
		g,
		"resetBatching",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetBeyondcorpCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBeyondcorpCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetBigqueryAnalyticsHubCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryAnalyticsHubCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetBigqueryConnectionCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryConnectionCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetBigQueryCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigQueryCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetBigqueryDatapolicyCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryDatapolicyCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetBigqueryDataTransferCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryDataTransferCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetBigqueryReservationCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryReservationCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetBigtableCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigtableCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetBillingCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBillingCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetBillingProject() {
	_jsii_.InvokeVoid(
		g,
		"resetBillingProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetBinaryAuthorizationCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBinaryAuthorizationCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetCertificateManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCertificateManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetCloudAssetCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudAssetCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetCloudBillingCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudBillingCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetCloudBuildCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudBuildCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetCloudbuildv2CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudbuildv2CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetCloudBuildWorkerPoolCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudBuildWorkerPoolCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetClouddeployCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetClouddeployCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetCloudfunctions2CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudfunctions2CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetCloudFunctionsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudFunctionsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetCloudIdentityCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudIdentityCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetCloudIdsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudIdsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetCloudIotCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudIotCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetCloudResourceManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudResourceManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetCloudRunCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudRunCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetCloudRunV2CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudRunV2CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetCloudSchedulerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudSchedulerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetCloudTasksCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudTasksCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetComposerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetComposerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetComputeCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetComputeCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetContainerAnalysisCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerAnalysisCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetContainerAttachedCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerAttachedCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetContainerAwsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerAwsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetContainerAzureCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerAzureCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetContainerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetCredentials() {
	_jsii_.InvokeVoid(
		g,
		"resetCredentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetDatabaseMigrationServiceCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDatabaseMigrationServiceCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetDataCatalogCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataCatalogCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetDataflowCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataflowCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetDataFusionCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataFusionCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetDataLossPreventionCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataLossPreventionCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetDataplexCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataplexCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetDataprocCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataprocCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetDataprocMetastoreCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataprocMetastoreCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetDatastoreCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDatastoreCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetDatastreamCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDatastreamCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetDeploymentManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDeploymentManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetDialogflowCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDialogflowCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetDialogflowCxCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDialogflowCxCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetDnsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDnsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetDocumentAiCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDocumentAiCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetEssentialContactsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetEssentialContactsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetEventarcCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetEventarcCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetFilestoreCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetFilestoreCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetFirebaserulesCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetFirebaserulesCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetFirestoreCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetFirestoreCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetGameServicesCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetGameServicesCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetGkeBackupCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetGkeBackupCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetGkeHub2CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetGkeHub2CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetGkeHubCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetGkeHubCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetHealthcareCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthcareCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetIam2CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIam2CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetIamBetaCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIamBetaCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetIamCredentialsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIamCredentialsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetIamCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIamCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetIamWorkforcePoolCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIamWorkforcePoolCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetIapCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIapCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetIdentityPlatformCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIdentityPlatformCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetImpersonateServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetImpersonateServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetImpersonateServiceAccountDelegates() {
	_jsii_.InvokeVoid(
		g,
		"resetImpersonateServiceAccountDelegates",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetKmsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetLoggingCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetLoggingCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetLookerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetLookerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetMemcacheCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetMemcacheCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetMlEngineCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetMlEngineCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetMonitoringCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetMonitoringCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetNetworkConnectivityCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkConnectivityCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetNetworkManagementCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkManagementCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetNetworkSecurityCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkSecurityCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetNetworkServicesCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkServicesCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetNotebooksCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetNotebooksCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetOrgPolicyCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetOrgPolicyCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetOsConfigCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetOsConfigCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetOsLoginCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetOsLoginCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetPrivatecaCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivatecaCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetPublicCaCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetPublicCaCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetPubsubCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetPubsubCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetPubsubLiteCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetPubsubLiteCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetRecaptchaEnterpriseCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetRecaptchaEnterpriseCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetRedisCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetRedisCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetRequestReason() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestReason",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetRequestTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetResourceManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetResourceManagerV3CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceManagerV3CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetScopes() {
	_jsii_.InvokeVoid(
		g,
		"resetScopes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetSecretManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetSecurityCenterCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityCenterCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetServiceManagementCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceManagementCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetServiceNetworkingCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceNetworkingCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetServiceUsageCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceUsageCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetSourceRepoCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceRepoCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetSpannerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSpannerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetSqlCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSqlCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetStorageCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetStorageTransferCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageTransferCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetTagsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetTagsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetTagsLocationCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetTagsLocationCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetTpuCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetTpuCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetUserProjectOverride() {
	_jsii_.InvokeVoid(
		g,
		"resetUserProjectOverride",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetVertexAiCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetVertexAiCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetVpcAccessCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcAccessCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetWorkflowsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkflowsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) ResetZone() {
	_jsii_.InvokeVoid(
		g,
		"resetZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

