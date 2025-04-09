// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafusioninstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/datafusioninstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/data_fusion_instance google_data_fusion_instance}.
type DataFusionInstance interface {
	cdktf.TerraformResource
	Accelerators() DataFusionInstanceAcceleratorsList
	AcceleratorsInput() interface{}
	ApiEndpoint() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	CryptoKeyConfig() DataFusionInstanceCryptoKeyConfigOutputReference
	CryptoKeyConfigInput() *DataFusionInstanceCryptoKeyConfig
	DataprocServiceAccount() *string
	SetDataprocServiceAccount(val *string)
	DataprocServiceAccountInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EffectiveLabels() cdktf.StringMap
	EnableRbac() interface{}
	SetEnableRbac(val interface{})
	EnableRbacInput() interface{}
	EnableStackdriverLogging() interface{}
	SetEnableStackdriverLogging(val interface{})
	EnableStackdriverLoggingInput() interface{}
	EnableStackdriverMonitoring() interface{}
	SetEnableStackdriverMonitoring(val interface{})
	EnableStackdriverMonitoringInput() interface{}
	EventPublishConfig() DataFusionInstanceEventPublishConfigOutputReference
	EventPublishConfigInput() *DataFusionInstanceEventPublishConfig
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcsBucket() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkConfig() DataFusionInstanceNetworkConfigOutputReference
	NetworkConfigInput() *DataFusionInstanceNetworkConfig
	// The tree node.
	Node() constructs.Node
	Options() *map[string]*string
	SetOptions(val *map[string]*string)
	OptionsInput() *map[string]*string
	P4ServiceAccount() *string
	PrivateInstance() interface{}
	SetPrivateInstance(val interface{})
	PrivateInstanceInput() interface{}
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
	ServiceEndpoint() *string
	State() *string
	StateMessage() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TenantProjectId() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataFusionInstanceTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UpdateTime() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAccelerators(value interface{})
	PutCryptoKeyConfig(value *DataFusionInstanceCryptoKeyConfig)
	PutEventPublishConfig(value *DataFusionInstanceEventPublishConfig)
	PutNetworkConfig(value *DataFusionInstanceNetworkConfig)
	PutTimeouts(value *DataFusionInstanceTimeouts)
	ResetAccelerators()
	ResetCryptoKeyConfig()
	ResetDataprocServiceAccount()
	ResetDescription()
	ResetDisplayName()
	ResetEnableRbac()
	ResetEnableStackdriverLogging()
	ResetEnableStackdriverMonitoring()
	ResetEventPublishConfig()
	ResetId()
	ResetLabels()
	ResetNetworkConfig()
	ResetOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateInstance()
	ResetProject()
	ResetRegion()
	ResetTags()
	ResetTimeouts()
	ResetVersion()
	ResetZone()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataFusionInstance
type jsiiProxy_DataFusionInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DataFusionInstance) Accelerators() DataFusionInstanceAcceleratorsList {
	var returns DataFusionInstanceAcceleratorsList
	_jsii_.Get(
		j,
		"accelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) AcceleratorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) ApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) CryptoKeyConfig() DataFusionInstanceCryptoKeyConfigOutputReference {
	var returns DataFusionInstanceCryptoKeyConfigOutputReference
	_jsii_.Get(
		j,
		"cryptoKeyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) CryptoKeyConfigInput() *DataFusionInstanceCryptoKeyConfig {
	var returns *DataFusionInstanceCryptoKeyConfig
	_jsii_.Get(
		j,
		"cryptoKeyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) DataprocServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) DataprocServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) EnableRbac() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRbac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) EnableRbacInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRbacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) EnableStackdriverLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStackdriverLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) EnableStackdriverLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStackdriverLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) EnableStackdriverMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStackdriverMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) EnableStackdriverMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStackdriverMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) EventPublishConfig() DataFusionInstanceEventPublishConfigOutputReference {
	var returns DataFusionInstanceEventPublishConfigOutputReference
	_jsii_.Get(
		j,
		"eventPublishConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) EventPublishConfigInput() *DataFusionInstanceEventPublishConfig {
	var returns *DataFusionInstanceEventPublishConfig
	_jsii_.Get(
		j,
		"eventPublishConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) GcsBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) NetworkConfig() DataFusionInstanceNetworkConfigOutputReference {
	var returns DataFusionInstanceNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) NetworkConfigInput() *DataFusionInstanceNetworkConfig {
	var returns *DataFusionInstanceNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) Options() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) OptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) P4ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"p4ServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) PrivateInstance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) PrivateInstanceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) ServiceEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) StateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) TenantProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) Timeouts() DataFusionInstanceTimeoutsOutputReference {
	var returns DataFusionInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFusionInstance) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/data_fusion_instance google_data_fusion_instance} Resource.
func NewDataFusionInstance(scope constructs.Construct, id *string, config *DataFusionInstanceConfig) DataFusionInstance {
	_init_.Initialize()

	if err := validateNewDataFusionInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataFusionInstance{}

	_jsii_.Create(
		"@cdktf/provider-google.dataFusionInstance.DataFusionInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/data_fusion_instance google_data_fusion_instance} Resource.
func NewDataFusionInstance_Override(d DataFusionInstance, scope constructs.Construct, id *string, config *DataFusionInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataFusionInstance.DataFusionInstance",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetDataprocServiceAccount(val *string) {
	if err := j.validateSetDataprocServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataprocServiceAccount",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetEnableRbac(val interface{}) {
	if err := j.validateSetEnableRbacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableRbac",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetEnableStackdriverLogging(val interface{}) {
	if err := j.validateSetEnableStackdriverLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableStackdriverLogging",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetEnableStackdriverMonitoring(val interface{}) {
	if err := j.validateSetEnableStackdriverMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableStackdriverMonitoring",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetOptions(val *map[string]*string) {
	if err := j.validateSetOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetPrivateInstance(val interface{}) {
	if err := j.validateSetPrivateInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateInstance",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_DataFusionInstance)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Generates CDKTF code for importing a DataFusionInstance resource upon running "cdktf plan <stack-name>".
func DataFusionInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataFusionInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataFusionInstance.DataFusionInstance",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func DataFusionInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataFusionInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataFusionInstance.DataFusionInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataFusionInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataFusionInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataFusionInstance.DataFusionInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataFusionInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataFusionInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataFusionInstance.DataFusionInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataFusionInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.dataFusionInstance.DataFusionInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataFusionInstance) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DataFusionInstance) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataFusionInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFusionInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFusionInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFusionInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFusionInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFusionInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFusionInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFusionInstance) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFusionInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFusionInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFusionInstance) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DataFusionInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFusionInstance) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataFusionInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DataFusionInstance) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataFusionInstance) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataFusionInstance) PutAccelerators(value interface{}) {
	if err := d.validatePutAcceleratorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAccelerators",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFusionInstance) PutCryptoKeyConfig(value *DataFusionInstanceCryptoKeyConfig) {
	if err := d.validatePutCryptoKeyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCryptoKeyConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFusionInstance) PutEventPublishConfig(value *DataFusionInstanceEventPublishConfig) {
	if err := d.validatePutEventPublishConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventPublishConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFusionInstance) PutNetworkConfig(value *DataFusionInstanceNetworkConfig) {
	if err := d.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFusionInstance) PutTimeouts(value *DataFusionInstanceTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetAccelerators() {
	_jsii_.InvokeVoid(
		d,
		"resetAccelerators",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetCryptoKeyConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetCryptoKeyConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetDataprocServiceAccount() {
	_jsii_.InvokeVoid(
		d,
		"resetDataprocServiceAccount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetDisplayName() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetEnableRbac() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableRbac",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetEnableStackdriverLogging() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableStackdriverLogging",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetEnableStackdriverMonitoring() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableStackdriverMonitoring",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetEventPublishConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetEventPublishConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetLabels() {
	_jsii_.InvokeVoid(
		d,
		"resetLabels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetPrivateInstance() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivateInstance",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) ResetZone() {
	_jsii_.InvokeVoid(
		d,
		"resetZone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFusionInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFusionInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFusionInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFusionInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFusionInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFusionInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

