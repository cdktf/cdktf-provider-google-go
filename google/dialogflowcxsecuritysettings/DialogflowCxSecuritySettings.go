// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxsecuritysettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v16/dialogflowcxsecuritysettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_cx_security_settings google_dialogflow_cx_security_settings}.
type DialogflowCxSecuritySettings interface {
	cdktf.TerraformResource
	AudioExportSettings() DialogflowCxSecuritySettingsAudioExportSettingsOutputReference
	AudioExportSettingsInput() *DialogflowCxSecuritySettingsAudioExportSettings
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
	DeidentifyTemplate() *string
	SetDeidentifyTemplate(val *string)
	DeidentifyTemplateInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
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
	InsightsExportSettings() DialogflowCxSecuritySettingsInsightsExportSettingsOutputReference
	InsightsExportSettingsInput() *DialogflowCxSecuritySettingsInsightsExportSettings
	InspectTemplate() *string
	SetInspectTemplate(val *string)
	InspectTemplateInput() *string
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
	PurgeDataTypes() *[]*string
	SetPurgeDataTypes(val *[]*string)
	PurgeDataTypesInput() *[]*string
	// Experimental.
	RawOverrides() interface{}
	RedactionScope() *string
	SetRedactionScope(val *string)
	RedactionScopeInput() *string
	RedactionStrategy() *string
	SetRedactionStrategy(val *string)
	RedactionStrategyInput() *string
	RetentionStrategy() *string
	SetRetentionStrategy(val *string)
	RetentionStrategyInput() *string
	RetentionWindowDays() *float64
	SetRetentionWindowDays(val *float64)
	RetentionWindowDaysInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DialogflowCxSecuritySettingsTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutAudioExportSettings(value *DialogflowCxSecuritySettingsAudioExportSettings)
	PutInsightsExportSettings(value *DialogflowCxSecuritySettingsInsightsExportSettings)
	PutTimeouts(value *DialogflowCxSecuritySettingsTimeouts)
	ResetAudioExportSettings()
	ResetDeidentifyTemplate()
	ResetId()
	ResetInsightsExportSettings()
	ResetInspectTemplate()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetPurgeDataTypes()
	ResetRedactionScope()
	ResetRedactionStrategy()
	ResetRetentionStrategy()
	ResetRetentionWindowDays()
	ResetTimeouts()
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

// The jsii proxy struct for DialogflowCxSecuritySettings
type jsiiProxy_DialogflowCxSecuritySettings struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) AudioExportSettings() DialogflowCxSecuritySettingsAudioExportSettingsOutputReference {
	var returns DialogflowCxSecuritySettingsAudioExportSettingsOutputReference
	_jsii_.Get(
		j,
		"audioExportSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) AudioExportSettingsInput() *DialogflowCxSecuritySettingsAudioExportSettings {
	var returns *DialogflowCxSecuritySettingsAudioExportSettings
	_jsii_.Get(
		j,
		"audioExportSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) DeidentifyTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deidentifyTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) DeidentifyTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deidentifyTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) InsightsExportSettings() DialogflowCxSecuritySettingsInsightsExportSettingsOutputReference {
	var returns DialogflowCxSecuritySettingsInsightsExportSettingsOutputReference
	_jsii_.Get(
		j,
		"insightsExportSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) InsightsExportSettingsInput() *DialogflowCxSecuritySettingsInsightsExportSettings {
	var returns *DialogflowCxSecuritySettingsInsightsExportSettings
	_jsii_.Get(
		j,
		"insightsExportSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) InspectTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inspectTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) InspectTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inspectTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) PurgeDataTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"purgeDataTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) PurgeDataTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"purgeDataTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) RedactionScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redactionScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) RedactionScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redactionScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) RedactionStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redactionStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) RedactionStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redactionStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) RetentionStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) RetentionStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) RetentionWindowDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionWindowDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) RetentionWindowDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionWindowDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) Timeouts() DialogflowCxSecuritySettingsTimeoutsOutputReference {
	var returns DialogflowCxSecuritySettingsTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettings) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_cx_security_settings google_dialogflow_cx_security_settings} Resource.
func NewDialogflowCxSecuritySettings(scope constructs.Construct, id *string, config *DialogflowCxSecuritySettingsConfig) DialogflowCxSecuritySettings {
	_init_.Initialize()

	if err := validateNewDialogflowCxSecuritySettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxSecuritySettings{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxSecuritySettings.DialogflowCxSecuritySettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_cx_security_settings google_dialogflow_cx_security_settings} Resource.
func NewDialogflowCxSecuritySettings_Override(d DialogflowCxSecuritySettings, scope constructs.Construct, id *string, config *DialogflowCxSecuritySettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxSecuritySettings.DialogflowCxSecuritySettings",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettings)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettings)SetDeidentifyTemplate(val *string) {
	if err := j.validateSetDeidentifyTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deidentifyTemplate",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettings)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettings)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettings)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettings)SetInspectTemplate(val *string) {
	if err := j.validateSetInspectTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inspectTemplate",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettings)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettings)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettings)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettings)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettings)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettings)SetPurgeDataTypes(val *[]*string) {
	if err := j.validateSetPurgeDataTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purgeDataTypes",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettings)SetRedactionScope(val *string) {
	if err := j.validateSetRedactionScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redactionScope",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettings)SetRedactionStrategy(val *string) {
	if err := j.validateSetRedactionStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redactionStrategy",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettings)SetRetentionStrategy(val *string) {
	if err := j.validateSetRetentionStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionStrategy",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettings)SetRetentionWindowDays(val *float64) {
	if err := j.validateSetRetentionWindowDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionWindowDays",
		val,
	)
}

// Generates CDKTF code for importing a DialogflowCxSecuritySettings resource upon running "cdktf plan <stack-name>".
func DialogflowCxSecuritySettings_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDialogflowCxSecuritySettings_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dialogflowCxSecuritySettings.DialogflowCxSecuritySettings",
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
func DialogflowCxSecuritySettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDialogflowCxSecuritySettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dialogflowCxSecuritySettings.DialogflowCxSecuritySettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DialogflowCxSecuritySettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDialogflowCxSecuritySettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dialogflowCxSecuritySettings.DialogflowCxSecuritySettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DialogflowCxSecuritySettings_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDialogflowCxSecuritySettings_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dialogflowCxSecuritySettings.DialogflowCxSecuritySettings",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DialogflowCxSecuritySettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.dialogflowCxSecuritySettings.DialogflowCxSecuritySettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowCxSecuritySettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxSecuritySettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxSecuritySettings) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxSecuritySettings) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxSecuritySettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxSecuritySettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxSecuritySettings) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxSecuritySettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxSecuritySettings) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxSecuritySettings) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) PutAudioExportSettings(value *DialogflowCxSecuritySettingsAudioExportSettings) {
	if err := d.validatePutAudioExportSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAudioExportSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) PutInsightsExportSettings(value *DialogflowCxSecuritySettingsInsightsExportSettings) {
	if err := d.validatePutInsightsExportSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInsightsExportSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) PutTimeouts(value *DialogflowCxSecuritySettingsTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) ResetAudioExportSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetAudioExportSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) ResetDeidentifyTemplate() {
	_jsii_.InvokeVoid(
		d,
		"resetDeidentifyTemplate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) ResetInsightsExportSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetInsightsExportSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) ResetInspectTemplate() {
	_jsii_.InvokeVoid(
		d,
		"resetInspectTemplate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) ResetPurgeDataTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetPurgeDataTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) ResetRedactionScope() {
	_jsii_.InvokeVoid(
		d,
		"resetRedactionScope",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) ResetRedactionStrategy() {
	_jsii_.InvokeVoid(
		d,
		"resetRedactionStrategy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) ResetRetentionStrategy() {
	_jsii_.InvokeVoid(
		d,
		"resetRetentionStrategy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) ResetRetentionWindowDays() {
	_jsii_.InvokeVoid(
		d,
		"resetRetentionWindowDays",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxSecuritySettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

