package bigquerydatatransferconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v7/bigquerydatatransferconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/4.68.0/docs/resources/bigquery_data_transfer_config google_bigquery_data_transfer_config}.
type BigqueryDataTransferConfig interface {
	cdktf.TerraformResource
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
	DataRefreshWindowDays() *float64
	SetDataRefreshWindowDays(val *float64)
	DataRefreshWindowDaysInput() *float64
	DataSourceId() *string
	SetDataSourceId(val *string)
	DataSourceIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DestinationDatasetId() *string
	SetDestinationDatasetId(val *string)
	DestinationDatasetIdInput() *string
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EmailPreferences() BigqueryDataTransferConfigEmailPreferencesOutputReference
	EmailPreferencesInput() *BigqueryDataTransferConfigEmailPreferences
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	NotificationPubsubTopic() *string
	SetNotificationPubsubTopic(val *string)
	NotificationPubsubTopicInput() *string
	Params() *map[string]*string
	SetParams(val *map[string]*string)
	ParamsInput() *map[string]*string
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
	Schedule() *string
	SetSchedule(val *string)
	ScheduleInput() *string
	ScheduleOptions() BigqueryDataTransferConfigScheduleOptionsOutputReference
	ScheduleOptionsInput() *BigqueryDataTransferConfigScheduleOptions
	SensitiveParams() BigqueryDataTransferConfigSensitiveParamsOutputReference
	SensitiveParamsInput() *BigqueryDataTransferConfigSensitiveParams
	ServiceAccountName() *string
	SetServiceAccountName(val *string)
	ServiceAccountNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BigqueryDataTransferConfigTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutEmailPreferences(value *BigqueryDataTransferConfigEmailPreferences)
	PutScheduleOptions(value *BigqueryDataTransferConfigScheduleOptions)
	PutSensitiveParams(value *BigqueryDataTransferConfigSensitiveParams)
	PutTimeouts(value *BigqueryDataTransferConfigTimeouts)
	ResetDataRefreshWindowDays()
	ResetDestinationDatasetId()
	ResetDisabled()
	ResetEmailPreferences()
	ResetId()
	ResetLocation()
	ResetNotificationPubsubTopic()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetSchedule()
	ResetScheduleOptions()
	ResetSensitiveParams()
	ResetServiceAccountName()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for BigqueryDataTransferConfig
type jsiiProxy_BigqueryDataTransferConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BigqueryDataTransferConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) DataRefreshWindowDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataRefreshWindowDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) DataRefreshWindowDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataRefreshWindowDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) DataSourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) DataSourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) DestinationDatasetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationDatasetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) DestinationDatasetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationDatasetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) EmailPreferences() BigqueryDataTransferConfigEmailPreferencesOutputReference {
	var returns BigqueryDataTransferConfigEmailPreferencesOutputReference
	_jsii_.Get(
		j,
		"emailPreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) EmailPreferencesInput() *BigqueryDataTransferConfigEmailPreferences {
	var returns *BigqueryDataTransferConfigEmailPreferences
	_jsii_.Get(
		j,
		"emailPreferencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) NotificationPubsubTopic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPubsubTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) NotificationPubsubTopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPubsubTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) Params() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"params",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) ParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"paramsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) ScheduleOptions() BigqueryDataTransferConfigScheduleOptionsOutputReference {
	var returns BigqueryDataTransferConfigScheduleOptionsOutputReference
	_jsii_.Get(
		j,
		"scheduleOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) ScheduleOptionsInput() *BigqueryDataTransferConfigScheduleOptions {
	var returns *BigqueryDataTransferConfigScheduleOptions
	_jsii_.Get(
		j,
		"scheduleOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) SensitiveParams() BigqueryDataTransferConfigSensitiveParamsOutputReference {
	var returns BigqueryDataTransferConfigSensitiveParamsOutputReference
	_jsii_.Get(
		j,
		"sensitiveParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) SensitiveParamsInput() *BigqueryDataTransferConfigSensitiveParams {
	var returns *BigqueryDataTransferConfigSensitiveParams
	_jsii_.Get(
		j,
		"sensitiveParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) ServiceAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) Timeouts() BigqueryDataTransferConfigTimeoutsOutputReference {
	var returns BigqueryDataTransferConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryDataTransferConfig) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.68.0/docs/resources/bigquery_data_transfer_config google_bigquery_data_transfer_config} Resource.
func NewBigqueryDataTransferConfig(scope constructs.Construct, id *string, config *BigqueryDataTransferConfigConfig) BigqueryDataTransferConfig {
	_init_.Initialize()

	if err := validateNewBigqueryDataTransferConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryDataTransferConfig{}

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryDataTransferConfig.BigqueryDataTransferConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.68.0/docs/resources/bigquery_data_transfer_config google_bigquery_data_transfer_config} Resource.
func NewBigqueryDataTransferConfig_Override(b BigqueryDataTransferConfig, scope constructs.Construct, id *string, config *BigqueryDataTransferConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryDataTransferConfig.BigqueryDataTransferConfig",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfig)SetDataRefreshWindowDays(val *float64) {
	if err := j.validateSetDataRefreshWindowDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataRefreshWindowDays",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfig)SetDataSourceId(val *string) {
	if err := j.validateSetDataSourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceId",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfig)SetDestinationDatasetId(val *string) {
	if err := j.validateSetDestinationDatasetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationDatasetId",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfig)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfig)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfig)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfig)SetNotificationPubsubTopic(val *string) {
	if err := j.validateSetNotificationPubsubTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationPubsubTopic",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfig)SetParams(val *map[string]*string) {
	if err := j.validateSetParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"params",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfig)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfig)SetSchedule(val *string) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_BigqueryDataTransferConfig)SetServiceAccountName(val *string) {
	if err := j.validateSetServiceAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountName",
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
func BigqueryDataTransferConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBigqueryDataTransferConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.bigqueryDataTransferConfig.BigqueryDataTransferConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BigqueryDataTransferConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBigqueryDataTransferConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.bigqueryDataTransferConfig.BigqueryDataTransferConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BigqueryDataTransferConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBigqueryDataTransferConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.bigqueryDataTransferConfig.BigqueryDataTransferConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BigqueryDataTransferConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.bigqueryDataTransferConfig.BigqueryDataTransferConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfig) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfig) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfig) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) PutEmailPreferences(value *BigqueryDataTransferConfigEmailPreferences) {
	if err := b.validatePutEmailPreferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEmailPreferences",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) PutScheduleOptions(value *BigqueryDataTransferConfigScheduleOptions) {
	if err := b.validatePutScheduleOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putScheduleOptions",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) PutSensitiveParams(value *BigqueryDataTransferConfigSensitiveParams) {
	if err := b.validatePutSensitiveParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSensitiveParams",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) PutTimeouts(value *BigqueryDataTransferConfigTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) ResetDataRefreshWindowDays() {
	_jsii_.InvokeVoid(
		b,
		"resetDataRefreshWindowDays",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) ResetDestinationDatasetId() {
	_jsii_.InvokeVoid(
		b,
		"resetDestinationDatasetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) ResetDisabled() {
	_jsii_.InvokeVoid(
		b,
		"resetDisabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) ResetEmailPreferences() {
	_jsii_.InvokeVoid(
		b,
		"resetEmailPreferences",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) ResetLocation() {
	_jsii_.InvokeVoid(
		b,
		"resetLocation",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) ResetNotificationPubsubTopic() {
	_jsii_.InvokeVoid(
		b,
		"resetNotificationPubsubTopic",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) ResetProject() {
	_jsii_.InvokeVoid(
		b,
		"resetProject",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) ResetSchedule() {
	_jsii_.InvokeVoid(
		b,
		"resetSchedule",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) ResetScheduleOptions() {
	_jsii_.InvokeVoid(
		b,
		"resetScheduleOptions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) ResetSensitiveParams() {
	_jsii_.InvokeVoid(
		b,
		"resetSensitiveParams",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) ResetServiceAccountName() {
	_jsii_.InvokeVoid(
		b,
		"resetServiceAccountName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryDataTransferConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryDataTransferConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

