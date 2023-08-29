// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v9/dialogflowcxagent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/4.80.0/docs/resources/dialogflow_cx_agent google_dialogflow_cx_agent}.
type DialogflowCxAgent interface {
	cdktf.TerraformResource
	AvatarUri() *string
	SetAvatarUri(val *string)
	AvatarUriInput() *string
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
	DefaultLanguageCode() *string
	SetDefaultLanguageCode(val *string)
	DefaultLanguageCodeInput() *string
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
	EnableSpellCorrection() interface{}
	SetEnableSpellCorrection(val interface{})
	EnableSpellCorrectionInput() interface{}
	EnableStackdriverLogging() interface{}
	SetEnableStackdriverLogging(val interface{})
	EnableStackdriverLoggingInput() interface{}
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
	SecuritySettings() *string
	SetSecuritySettings(val *string)
	SecuritySettingsInput() *string
	SpeechToTextSettings() DialogflowCxAgentSpeechToTextSettingsOutputReference
	SpeechToTextSettingsInput() *DialogflowCxAgentSpeechToTextSettings
	StartFlow() *string
	SupportedLanguageCodes() *[]*string
	SetSupportedLanguageCodes(val *[]*string)
	SupportedLanguageCodesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DialogflowCxAgentTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
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
	PutSpeechToTextSettings(value *DialogflowCxAgentSpeechToTextSettings)
	PutTimeouts(value *DialogflowCxAgentTimeouts)
	ResetAvatarUri()
	ResetDescription()
	ResetEnableSpellCorrection()
	ResetEnableStackdriverLogging()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetSecuritySettings()
	ResetSpeechToTextSettings()
	ResetSupportedLanguageCodes()
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

// The jsii proxy struct for DialogflowCxAgent
type jsiiProxy_DialogflowCxAgent struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DialogflowCxAgent) AvatarUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"avatarUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) AvatarUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"avatarUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) DefaultLanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLanguageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) DefaultLanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLanguageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) EnableSpellCorrection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSpellCorrection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) EnableSpellCorrectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSpellCorrectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) EnableStackdriverLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStackdriverLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) EnableStackdriverLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStackdriverLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) SecuritySettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securitySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) SecuritySettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securitySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) SpeechToTextSettings() DialogflowCxAgentSpeechToTextSettingsOutputReference {
	var returns DialogflowCxAgentSpeechToTextSettingsOutputReference
	_jsii_.Get(
		j,
		"speechToTextSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) SpeechToTextSettingsInput() *DialogflowCxAgentSpeechToTextSettings {
	var returns *DialogflowCxAgentSpeechToTextSettings
	_jsii_.Get(
		j,
		"speechToTextSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) StartFlow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startFlow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) SupportedLanguageCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedLanguageCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) SupportedLanguageCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedLanguageCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Timeouts() DialogflowCxAgentTimeoutsOutputReference {
	var returns DialogflowCxAgentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.80.0/docs/resources/dialogflow_cx_agent google_dialogflow_cx_agent} Resource.
func NewDialogflowCxAgent(scope constructs.Construct, id *string, config *DialogflowCxAgentConfig) DialogflowCxAgent {
	_init_.Initialize()

	if err := validateNewDialogflowCxAgentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxAgent{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxAgent.DialogflowCxAgent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.80.0/docs/resources/dialogflow_cx_agent google_dialogflow_cx_agent} Resource.
func NewDialogflowCxAgent_Override(d DialogflowCxAgent, scope constructs.Construct, id *string, config *DialogflowCxAgentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxAgent.DialogflowCxAgent",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetAvatarUri(val *string) {
	if err := j.validateSetAvatarUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"avatarUri",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetDefaultLanguageCode(val *string) {
	if err := j.validateSetDefaultLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLanguageCode",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetEnableSpellCorrection(val interface{}) {
	if err := j.validateSetEnableSpellCorrectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSpellCorrection",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetEnableStackdriverLogging(val interface{}) {
	if err := j.validateSetEnableStackdriverLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableStackdriverLogging",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetSecuritySettings(val *string) {
	if err := j.validateSetSecuritySettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securitySettings",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetSupportedLanguageCodes(val *[]*string) {
	if err := j.validateSetSupportedLanguageCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedLanguageCodes",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
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
func DialogflowCxAgent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDialogflowCxAgent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dialogflowCxAgent.DialogflowCxAgent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DialogflowCxAgent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDialogflowCxAgent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dialogflowCxAgent.DialogflowCxAgent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DialogflowCxAgent_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDialogflowCxAgent_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dialogflowCxAgent.DialogflowCxAgent",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DialogflowCxAgent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.dialogflowCxAgent.DialogflowCxAgent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowCxAgent) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxAgent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxAgent) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxAgent) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxAgent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxAgent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxAgent) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxAgent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxAgent) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxAgent) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) PutSpeechToTextSettings(value *DialogflowCxAgentSpeechToTextSettings) {
	if err := d.validatePutSpeechToTextSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSpeechToTextSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) PutTimeouts(value *DialogflowCxAgentTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetAvatarUri() {
	_jsii_.InvokeVoid(
		d,
		"resetAvatarUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetEnableSpellCorrection() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableSpellCorrection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetEnableStackdriverLogging() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableStackdriverLogging",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetSecuritySettings() {
	_jsii_.InvokeVoid(
		d,
		"resetSecuritySettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetSpeechToTextSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetSpeechToTextSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetSupportedLanguageCodes() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportedLanguageCodes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

