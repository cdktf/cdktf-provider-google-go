package cloudfunctionsfunction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v6/cloudfunctionsfunction/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions_function google_cloudfunctions_function}.
type CloudfunctionsFunction interface {
	cdktf.TerraformResource
	AvailableMemoryMb() *float64
	SetAvailableMemoryMb(val *float64)
	AvailableMemoryMbInput() *float64
	BuildEnvironmentVariables() *map[string]*string
	SetBuildEnvironmentVariables(val *map[string]*string)
	BuildEnvironmentVariablesInput() *map[string]*string
	BuildWorkerPool() *string
	SetBuildWorkerPool(val *string)
	BuildWorkerPoolInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DockerRegistry() *string
	SetDockerRegistry(val *string)
	DockerRegistryInput() *string
	DockerRepository() *string
	SetDockerRepository(val *string)
	DockerRepositoryInput() *string
	EntryPoint() *string
	SetEntryPoint(val *string)
	EntryPointInput() *string
	EnvironmentVariables() *map[string]*string
	SetEnvironmentVariables(val *map[string]*string)
	EnvironmentVariablesInput() *map[string]*string
	EventTrigger() CloudfunctionsFunctionEventTriggerOutputReference
	EventTriggerInput() *CloudfunctionsFunctionEventTrigger
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpsTriggerSecurityLevel() *string
	SetHttpsTriggerSecurityLevel(val *string)
	HttpsTriggerSecurityLevelInput() *string
	HttpsTriggerUrl() *string
	SetHttpsTriggerUrl(val *string)
	HttpsTriggerUrlInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IngressSettings() *string
	SetIngressSettings(val *string)
	IngressSettingsInput() *string
	KmsKeyName() *string
	SetKmsKeyName(val *string)
	KmsKeyNameInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxInstances() *float64
	SetMaxInstances(val *float64)
	MaxInstancesInput() *float64
	MinInstances() *float64
	SetMinInstances(val *float64)
	MinInstancesInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Runtime() *string
	SetRuntime(val *string)
	RuntimeInput() *string
	SecretEnvironmentVariables() CloudfunctionsFunctionSecretEnvironmentVariablesList
	SecretEnvironmentVariablesInput() interface{}
	SecretVolumes() CloudfunctionsFunctionSecretVolumesList
	SecretVolumesInput() interface{}
	ServiceAccountEmail() *string
	SetServiceAccountEmail(val *string)
	ServiceAccountEmailInput() *string
	SourceArchiveBucket() *string
	SetSourceArchiveBucket(val *string)
	SourceArchiveBucketInput() *string
	SourceArchiveObject() *string
	SetSourceArchiveObject(val *string)
	SourceArchiveObjectInput() *string
	SourceRepository() CloudfunctionsFunctionSourceRepositoryOutputReference
	SourceRepositoryInput() *CloudfunctionsFunctionSourceRepository
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	Timeouts() CloudfunctionsFunctionTimeoutsOutputReference
	TimeoutsInput() interface{}
	TriggerHttp() interface{}
	SetTriggerHttp(val interface{})
	TriggerHttpInput() interface{}
	VpcConnector() *string
	SetVpcConnector(val *string)
	VpcConnectorEgressSettings() *string
	SetVpcConnectorEgressSettings(val *string)
	VpcConnectorEgressSettingsInput() *string
	VpcConnectorInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutEventTrigger(value *CloudfunctionsFunctionEventTrigger)
	PutSecretEnvironmentVariables(value interface{})
	PutSecretVolumes(value interface{})
	PutSourceRepository(value *CloudfunctionsFunctionSourceRepository)
	PutTimeouts(value *CloudfunctionsFunctionTimeouts)
	ResetAvailableMemoryMb()
	ResetBuildEnvironmentVariables()
	ResetBuildWorkerPool()
	ResetDescription()
	ResetDockerRegistry()
	ResetDockerRepository()
	ResetEntryPoint()
	ResetEnvironmentVariables()
	ResetEventTrigger()
	ResetHttpsTriggerSecurityLevel()
	ResetHttpsTriggerUrl()
	ResetId()
	ResetIngressSettings()
	ResetKmsKeyName()
	ResetLabels()
	ResetMaxInstances()
	ResetMinInstances()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRegion()
	ResetSecretEnvironmentVariables()
	ResetSecretVolumes()
	ResetServiceAccountEmail()
	ResetSourceArchiveBucket()
	ResetSourceArchiveObject()
	ResetSourceRepository()
	ResetTimeout()
	ResetTimeouts()
	ResetTriggerHttp()
	ResetVpcConnector()
	ResetVpcConnectorEgressSettings()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudfunctionsFunction
type jsiiProxy_CloudfunctionsFunction struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudfunctionsFunction) AvailableMemoryMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableMemoryMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) AvailableMemoryMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableMemoryMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) BuildEnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"buildEnvironmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) BuildEnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"buildEnvironmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) BuildWorkerPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildWorkerPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) BuildWorkerPoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildWorkerPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) DockerRegistry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) DockerRegistryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRegistryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) DockerRepository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) DockerRepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) EntryPoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) EntryPointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entryPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) EnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) EventTrigger() CloudfunctionsFunctionEventTriggerOutputReference {
	var returns CloudfunctionsFunctionEventTriggerOutputReference
	_jsii_.Get(
		j,
		"eventTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) EventTriggerInput() *CloudfunctionsFunctionEventTrigger {
	var returns *CloudfunctionsFunctionEventTrigger
	_jsii_.Get(
		j,
		"eventTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) HttpsTriggerSecurityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsTriggerSecurityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) HttpsTriggerSecurityLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsTriggerSecurityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) HttpsTriggerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsTriggerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) HttpsTriggerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsTriggerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) IngressSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingressSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) IngressSettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingressSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) KmsKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) KmsKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) MaxInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) MaxInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) MinInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) MinInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SecretEnvironmentVariables() CloudfunctionsFunctionSecretEnvironmentVariablesList {
	var returns CloudfunctionsFunctionSecretEnvironmentVariablesList
	_jsii_.Get(
		j,
		"secretEnvironmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SecretEnvironmentVariablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretEnvironmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SecretVolumes() CloudfunctionsFunctionSecretVolumesList {
	var returns CloudfunctionsFunctionSecretVolumesList
	_jsii_.Get(
		j,
		"secretVolumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SecretVolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretVolumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) ServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) ServiceAccountEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SourceArchiveBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArchiveBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SourceArchiveBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArchiveBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SourceArchiveObject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArchiveObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SourceArchiveObjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArchiveObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SourceRepository() CloudfunctionsFunctionSourceRepositoryOutputReference {
	var returns CloudfunctionsFunctionSourceRepositoryOutputReference
	_jsii_.Get(
		j,
		"sourceRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) SourceRepositoryInput() *CloudfunctionsFunctionSourceRepository {
	var returns *CloudfunctionsFunctionSourceRepository
	_jsii_.Get(
		j,
		"sourceRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) Timeouts() CloudfunctionsFunctionTimeoutsOutputReference {
	var returns CloudfunctionsFunctionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) TriggerHttp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggerHttp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) TriggerHttpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggerHttpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) VpcConnector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) VpcConnectorEgressSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorEgressSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) VpcConnectorEgressSettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorEgressSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfunctionsFunction) VpcConnectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions_function google_cloudfunctions_function} Resource.
func NewCloudfunctionsFunction(scope constructs.Construct, id *string, config *CloudfunctionsFunctionConfig) CloudfunctionsFunction {
	_init_.Initialize()

	if err := validateNewCloudfunctionsFunctionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfunctionsFunction{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudfunctionsFunction.CloudfunctionsFunction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions_function google_cloudfunctions_function} Resource.
func NewCloudfunctionsFunction_Override(c CloudfunctionsFunction, scope constructs.Construct, id *string, config *CloudfunctionsFunctionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudfunctionsFunction.CloudfunctionsFunction",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetAvailableMemoryMb(val *float64) {
	if err := j.validateSetAvailableMemoryMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availableMemoryMb",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetBuildEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetBuildEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildEnvironmentVariables",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetBuildWorkerPool(val *string) {
	if err := j.validateSetBuildWorkerPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildWorkerPool",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetDockerRegistry(val *string) {
	if err := j.validateSetDockerRegistryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerRegistry",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetDockerRepository(val *string) {
	if err := j.validateSetDockerRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerRepository",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetEntryPoint(val *string) {
	if err := j.validateSetEntryPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entryPoint",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetHttpsTriggerSecurityLevel(val *string) {
	if err := j.validateSetHttpsTriggerSecurityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsTriggerSecurityLevel",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetHttpsTriggerUrl(val *string) {
	if err := j.validateSetHttpsTriggerUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsTriggerUrl",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetIngressSettings(val *string) {
	if err := j.validateSetIngressSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressSettings",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetKmsKeyName(val *string) {
	if err := j.validateSetKmsKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyName",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetMaxInstances(val *float64) {
	if err := j.validateSetMaxInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInstances",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetMinInstances(val *float64) {
	if err := j.validateSetMinInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minInstances",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetServiceAccountEmail(val *string) {
	if err := j.validateSetServiceAccountEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountEmail",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetSourceArchiveBucket(val *string) {
	if err := j.validateSetSourceArchiveBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceArchiveBucket",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetSourceArchiveObject(val *string) {
	if err := j.validateSetSourceArchiveObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceArchiveObject",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetTriggerHttp(val interface{}) {
	if err := j.validateSetTriggerHttpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerHttp",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetVpcConnector(val *string) {
	if err := j.validateSetVpcConnectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcConnector",
		val,
	)
}

func (j *jsiiProxy_CloudfunctionsFunction)SetVpcConnectorEgressSettings(val *string) {
	if err := j.validateSetVpcConnectorEgressSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcConnectorEgressSettings",
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
func CloudfunctionsFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudfunctionsFunction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cloudfunctionsFunction.CloudfunctionsFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudfunctionsFunction_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudfunctionsFunction_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cloudfunctionsFunction.CloudfunctionsFunction",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudfunctionsFunction_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudfunctionsFunction_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cloudfunctionsFunction.CloudfunctionsFunction",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudfunctionsFunction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.cloudfunctionsFunction.CloudfunctionsFunction",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) PutEventTrigger(value *CloudfunctionsFunctionEventTrigger) {
	if err := c.validatePutEventTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEventTrigger",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) PutSecretEnvironmentVariables(value interface{}) {
	if err := c.validatePutSecretEnvironmentVariablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecretEnvironmentVariables",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) PutSecretVolumes(value interface{}) {
	if err := c.validatePutSecretVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecretVolumes",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) PutSourceRepository(value *CloudfunctionsFunctionSourceRepository) {
	if err := c.validatePutSourceRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSourceRepository",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) PutTimeouts(value *CloudfunctionsFunctionTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetAvailableMemoryMb() {
	_jsii_.InvokeVoid(
		c,
		"resetAvailableMemoryMb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetBuildEnvironmentVariables() {
	_jsii_.InvokeVoid(
		c,
		"resetBuildEnvironmentVariables",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetBuildWorkerPool() {
	_jsii_.InvokeVoid(
		c,
		"resetBuildWorkerPool",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetDockerRegistry() {
	_jsii_.InvokeVoid(
		c,
		"resetDockerRegistry",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetDockerRepository() {
	_jsii_.InvokeVoid(
		c,
		"resetDockerRepository",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetEntryPoint() {
	_jsii_.InvokeVoid(
		c,
		"resetEntryPoint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		c,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetEventTrigger() {
	_jsii_.InvokeVoid(
		c,
		"resetEventTrigger",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetHttpsTriggerSecurityLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpsTriggerSecurityLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetHttpsTriggerUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpsTriggerUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetIngressSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetIngressSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetKmsKeyName() {
	_jsii_.InvokeVoid(
		c,
		"resetKmsKeyName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetMaxInstances() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxInstances",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetMinInstances() {
	_jsii_.InvokeVoid(
		c,
		"resetMinInstances",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetSecretEnvironmentVariables() {
	_jsii_.InvokeVoid(
		c,
		"resetSecretEnvironmentVariables",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetSecretVolumes() {
	_jsii_.InvokeVoid(
		c,
		"resetSecretVolumes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetServiceAccountEmail() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccountEmail",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetSourceArchiveBucket() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceArchiveBucket",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetSourceArchiveObject() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceArchiveObject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetSourceRepository() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceRepository",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetTriggerHttp() {
	_jsii_.InvokeVoid(
		c,
		"resetTriggerHttp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetVpcConnector() {
	_jsii_.InvokeVoid(
		c,
		"resetVpcConnector",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) ResetVpcConnectorEgressSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetVpcConnectorEgressSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfunctionsFunction) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfunctionsFunction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

