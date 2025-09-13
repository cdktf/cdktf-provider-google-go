// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v16/dialogflowconversationprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile google_dialogflow_conversation_profile}.
type DialogflowConversationProfile interface {
	cdktf.TerraformResource
	AutomatedAgentConfig() DialogflowConversationProfileAutomatedAgentConfigOutputReference
	AutomatedAgentConfigInput() *DialogflowConversationProfileAutomatedAgentConfig
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
	HumanAgentAssistantConfig() DialogflowConversationProfileHumanAgentAssistantConfigOutputReference
	HumanAgentAssistantConfigInput() *DialogflowConversationProfileHumanAgentAssistantConfig
	HumanAgentHandoffConfig() DialogflowConversationProfileHumanAgentHandoffConfigOutputReference
	HumanAgentHandoffConfigInput() *DialogflowConversationProfileHumanAgentHandoffConfig
	Id() *string
	SetId(val *string)
	IdInput() *string
	LanguageCode() *string
	SetLanguageCode(val *string)
	LanguageCodeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LoggingConfig() DialogflowConversationProfileLoggingConfigOutputReference
	LoggingConfigInput() *DialogflowConversationProfileLoggingConfig
	Name() *string
	NewMessageEventNotificationConfig() DialogflowConversationProfileNewMessageEventNotificationConfigOutputReference
	NewMessageEventNotificationConfigInput() *DialogflowConversationProfileNewMessageEventNotificationConfig
	// The tree node.
	Node() constructs.Node
	NotificationConfig() DialogflowConversationProfileNotificationConfigOutputReference
	NotificationConfigInput() *DialogflowConversationProfileNotificationConfig
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
	SttConfig() DialogflowConversationProfileSttConfigOutputReference
	SttConfigInput() *DialogflowConversationProfileSttConfig
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DialogflowConversationProfileTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
	TtsConfig() DialogflowConversationProfileTtsConfigOutputReference
	TtsConfigInput() *DialogflowConversationProfileTtsConfig
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
	PutAutomatedAgentConfig(value *DialogflowConversationProfileAutomatedAgentConfig)
	PutHumanAgentAssistantConfig(value *DialogflowConversationProfileHumanAgentAssistantConfig)
	PutHumanAgentHandoffConfig(value *DialogflowConversationProfileHumanAgentHandoffConfig)
	PutLoggingConfig(value *DialogflowConversationProfileLoggingConfig)
	PutNewMessageEventNotificationConfig(value *DialogflowConversationProfileNewMessageEventNotificationConfig)
	PutNotificationConfig(value *DialogflowConversationProfileNotificationConfig)
	PutSttConfig(value *DialogflowConversationProfileSttConfig)
	PutTimeouts(value *DialogflowConversationProfileTimeouts)
	PutTtsConfig(value *DialogflowConversationProfileTtsConfig)
	ResetAutomatedAgentConfig()
	ResetHumanAgentAssistantConfig()
	ResetHumanAgentHandoffConfig()
	ResetId()
	ResetLanguageCode()
	ResetLoggingConfig()
	ResetNewMessageEventNotificationConfig()
	ResetNotificationConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetSecuritySettings()
	ResetSttConfig()
	ResetTimeouts()
	ResetTimeZone()
	ResetTtsConfig()
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

// The jsii proxy struct for DialogflowConversationProfile
type jsiiProxy_DialogflowConversationProfile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DialogflowConversationProfile) AutomatedAgentConfig() DialogflowConversationProfileAutomatedAgentConfigOutputReference {
	var returns DialogflowConversationProfileAutomatedAgentConfigOutputReference
	_jsii_.Get(
		j,
		"automatedAgentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) AutomatedAgentConfigInput() *DialogflowConversationProfileAutomatedAgentConfig {
	var returns *DialogflowConversationProfileAutomatedAgentConfig
	_jsii_.Get(
		j,
		"automatedAgentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) HumanAgentAssistantConfig() DialogflowConversationProfileHumanAgentAssistantConfigOutputReference {
	var returns DialogflowConversationProfileHumanAgentAssistantConfigOutputReference
	_jsii_.Get(
		j,
		"humanAgentAssistantConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) HumanAgentAssistantConfigInput() *DialogflowConversationProfileHumanAgentAssistantConfig {
	var returns *DialogflowConversationProfileHumanAgentAssistantConfig
	_jsii_.Get(
		j,
		"humanAgentAssistantConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) HumanAgentHandoffConfig() DialogflowConversationProfileHumanAgentHandoffConfigOutputReference {
	var returns DialogflowConversationProfileHumanAgentHandoffConfigOutputReference
	_jsii_.Get(
		j,
		"humanAgentHandoffConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) HumanAgentHandoffConfigInput() *DialogflowConversationProfileHumanAgentHandoffConfig {
	var returns *DialogflowConversationProfileHumanAgentHandoffConfig
	_jsii_.Get(
		j,
		"humanAgentHandoffConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) LanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) LanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) LoggingConfig() DialogflowConversationProfileLoggingConfigOutputReference {
	var returns DialogflowConversationProfileLoggingConfigOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) LoggingConfigInput() *DialogflowConversationProfileLoggingConfig {
	var returns *DialogflowConversationProfileLoggingConfig
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) NewMessageEventNotificationConfig() DialogflowConversationProfileNewMessageEventNotificationConfigOutputReference {
	var returns DialogflowConversationProfileNewMessageEventNotificationConfigOutputReference
	_jsii_.Get(
		j,
		"newMessageEventNotificationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) NewMessageEventNotificationConfigInput() *DialogflowConversationProfileNewMessageEventNotificationConfig {
	var returns *DialogflowConversationProfileNewMessageEventNotificationConfig
	_jsii_.Get(
		j,
		"newMessageEventNotificationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) NotificationConfig() DialogflowConversationProfileNotificationConfigOutputReference {
	var returns DialogflowConversationProfileNotificationConfigOutputReference
	_jsii_.Get(
		j,
		"notificationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) NotificationConfigInput() *DialogflowConversationProfileNotificationConfig {
	var returns *DialogflowConversationProfileNotificationConfig
	_jsii_.Get(
		j,
		"notificationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) SecuritySettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securitySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) SecuritySettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securitySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) SttConfig() DialogflowConversationProfileSttConfigOutputReference {
	var returns DialogflowConversationProfileSttConfigOutputReference
	_jsii_.Get(
		j,
		"sttConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) SttConfigInput() *DialogflowConversationProfileSttConfig {
	var returns *DialogflowConversationProfileSttConfig
	_jsii_.Get(
		j,
		"sttConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) Timeouts() DialogflowConversationProfileTimeoutsOutputReference {
	var returns DialogflowConversationProfileTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) TtsConfig() DialogflowConversationProfileTtsConfigOutputReference {
	var returns DialogflowConversationProfileTtsConfigOutputReference
	_jsii_.Get(
		j,
		"ttsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfile) TtsConfigInput() *DialogflowConversationProfileTtsConfig {
	var returns *DialogflowConversationProfileTtsConfig
	_jsii_.Get(
		j,
		"ttsConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile google_dialogflow_conversation_profile} Resource.
func NewDialogflowConversationProfile(scope constructs.Construct, id *string, config *DialogflowConversationProfileConfig) DialogflowConversationProfile {
	_init_.Initialize()

	if err := validateNewDialogflowConversationProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowConversationProfile{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile google_dialogflow_conversation_profile} Resource.
func NewDialogflowConversationProfile_Override(d DialogflowConversationProfile, scope constructs.Construct, id *string, config *DialogflowConversationProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfile",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DialogflowConversationProfile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfile)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfile)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfile)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfile)SetLanguageCode(val *string) {
	if err := j.validateSetLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageCode",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfile)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfile)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfile)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfile)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfile)SetSecuritySettings(val *string) {
	if err := j.validateSetSecuritySettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securitySettings",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfile)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

// Generates CDKTF code for importing a DialogflowConversationProfile resource upon running "cdktf plan <stack-name>".
func DialogflowConversationProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDialogflowConversationProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfile",
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
func DialogflowConversationProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDialogflowConversationProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DialogflowConversationProfile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDialogflowConversationProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DialogflowConversationProfile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDialogflowConversationProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DialogflowConversationProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DialogflowConversationProfile) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowConversationProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowConversationProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowConversationProfile) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowConversationProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowConversationProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowConversationProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowConversationProfile) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowConversationProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowConversationProfile) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfile) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowConversationProfile) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) PutAutomatedAgentConfig(value *DialogflowConversationProfileAutomatedAgentConfig) {
	if err := d.validatePutAutomatedAgentConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAutomatedAgentConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) PutHumanAgentAssistantConfig(value *DialogflowConversationProfileHumanAgentAssistantConfig) {
	if err := d.validatePutHumanAgentAssistantConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHumanAgentAssistantConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) PutHumanAgentHandoffConfig(value *DialogflowConversationProfileHumanAgentHandoffConfig) {
	if err := d.validatePutHumanAgentHandoffConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHumanAgentHandoffConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) PutLoggingConfig(value *DialogflowConversationProfileLoggingConfig) {
	if err := d.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) PutNewMessageEventNotificationConfig(value *DialogflowConversationProfileNewMessageEventNotificationConfig) {
	if err := d.validatePutNewMessageEventNotificationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNewMessageEventNotificationConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) PutNotificationConfig(value *DialogflowConversationProfileNotificationConfig) {
	if err := d.validatePutNotificationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNotificationConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) PutSttConfig(value *DialogflowConversationProfileSttConfig) {
	if err := d.validatePutSttConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSttConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) PutTimeouts(value *DialogflowConversationProfileTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) PutTtsConfig(value *DialogflowConversationProfileTtsConfig) {
	if err := d.validatePutTtsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTtsConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) ResetAutomatedAgentConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetAutomatedAgentConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) ResetHumanAgentAssistantConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetHumanAgentAssistantConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) ResetHumanAgentHandoffConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetHumanAgentHandoffConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) ResetLanguageCode() {
	_jsii_.InvokeVoid(
		d,
		"resetLanguageCode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) ResetNewMessageEventNotificationConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetNewMessageEventNotificationConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) ResetNotificationConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetNotificationConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) ResetSecuritySettings() {
	_jsii_.InvokeVoid(
		d,
		"resetSecuritySettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) ResetSttConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetSttConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) ResetTimeZone() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) ResetTtsConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetTtsConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfile) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfile) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

