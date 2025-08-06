// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowintent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v16/dialogflowintent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_intent google_dialogflow_intent}.
type DialogflowIntent interface {
	cdktf.TerraformResource
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	DefaultResponsePlatforms() *[]*string
	SetDefaultResponsePlatforms(val *[]*string)
	DefaultResponsePlatformsInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Events() *[]*string
	SetEvents(val *[]*string)
	EventsInput() *[]*string
	FollowupIntentInfo() DialogflowIntentFollowupIntentInfoList
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
	InputContextNames() *[]*string
	SetInputContextNames(val *[]*string)
	InputContextNamesInput() *[]*string
	IsFallback() interface{}
	SetIsFallback(val interface{})
	IsFallbackInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MlDisabled() interface{}
	SetMlDisabled(val interface{})
	MlDisabledInput() interface{}
	Name() *string
	// The tree node.
	Node() constructs.Node
	ParentFollowupIntentName() *string
	SetParentFollowupIntentName(val *string)
	ParentFollowupIntentNameInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
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
	ResetContexts() interface{}
	SetResetContexts(val interface{})
	ResetContextsInput() interface{}
	RootFollowupIntentName() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DialogflowIntentTimeoutsOutputReference
	TimeoutsInput() interface{}
	WebhookState() *string
	SetWebhookState(val *string)
	WebhookStateInput() *string
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
	PutTimeouts(value *DialogflowIntentTimeouts)
	ResetAction()
	ResetDefaultResponsePlatforms()
	ResetEvents()
	ResetId()
	ResetInputContextNames()
	ResetIsFallback()
	ResetMlDisabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParentFollowupIntentName()
	ResetPriority()
	ResetProject()
	ResetResetContexts()
	ResetTimeouts()
	ResetWebhookState()
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

// The jsii proxy struct for DialogflowIntent
type jsiiProxy_DialogflowIntent struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DialogflowIntent) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) DefaultResponsePlatforms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultResponsePlatforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) DefaultResponsePlatformsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultResponsePlatformsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) Events() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"events",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) EventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) FollowupIntentInfo() DialogflowIntentFollowupIntentInfoList {
	var returns DialogflowIntentFollowupIntentInfoList
	_jsii_.Get(
		j,
		"followupIntentInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) InputContextNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputContextNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) InputContextNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputContextNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) IsFallback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isFallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) IsFallbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isFallbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) MlDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mlDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) MlDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mlDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) ParentFollowupIntentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentFollowupIntentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) ParentFollowupIntentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentFollowupIntentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) ResetContexts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resetContexts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) ResetContextsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resetContextsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) RootFollowupIntentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootFollowupIntentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) Timeouts() DialogflowIntentTimeoutsOutputReference {
	var returns DialogflowIntentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) WebhookState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowIntent) WebhookStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookStateInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_intent google_dialogflow_intent} Resource.
func NewDialogflowIntent(scope constructs.Construct, id *string, config *DialogflowIntentConfig) DialogflowIntent {
	_init_.Initialize()

	if err := validateNewDialogflowIntentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowIntent{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowIntent.DialogflowIntent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/dialogflow_intent google_dialogflow_intent} Resource.
func NewDialogflowIntent_Override(d DialogflowIntent, scope constructs.Construct, id *string, config *DialogflowIntentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowIntent.DialogflowIntent",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetDefaultResponsePlatforms(val *[]*string) {
	if err := j.validateSetDefaultResponsePlatformsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultResponsePlatforms",
		val,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetEvents(val *[]*string) {
	if err := j.validateSetEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"events",
		val,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetInputContextNames(val *[]*string) {
	if err := j.validateSetInputContextNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputContextNames",
		val,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetIsFallback(val interface{}) {
	if err := j.validateSetIsFallbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isFallback",
		val,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetMlDisabled(val interface{}) {
	if err := j.validateSetMlDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mlDisabled",
		val,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetParentFollowupIntentName(val *string) {
	if err := j.validateSetParentFollowupIntentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentFollowupIntentName",
		val,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetResetContexts(val interface{}) {
	if err := j.validateSetResetContextsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resetContexts",
		val,
	)
}

func (j *jsiiProxy_DialogflowIntent)SetWebhookState(val *string) {
	if err := j.validateSetWebhookStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhookState",
		val,
	)
}

// Generates CDKTF code for importing a DialogflowIntent resource upon running "cdktf plan <stack-name>".
func DialogflowIntent_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDialogflowIntent_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dialogflowIntent.DialogflowIntent",
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
func DialogflowIntent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDialogflowIntent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dialogflowIntent.DialogflowIntent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DialogflowIntent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDialogflowIntent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dialogflowIntent.DialogflowIntent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DialogflowIntent_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDialogflowIntent_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dialogflowIntent.DialogflowIntent",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DialogflowIntent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.dialogflowIntent.DialogflowIntent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DialogflowIntent) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DialogflowIntent) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DialogflowIntent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowIntent) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowIntent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowIntent) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowIntent) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowIntent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowIntent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowIntent) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowIntent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowIntent) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowIntent) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DialogflowIntent) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowIntent) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DialogflowIntent) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DialogflowIntent) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DialogflowIntent) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DialogflowIntent) PutTimeouts(value *DialogflowIntentTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowIntent) ResetAction() {
	_jsii_.InvokeVoid(
		d,
		"resetAction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowIntent) ResetDefaultResponsePlatforms() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultResponsePlatforms",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowIntent) ResetEvents() {
	_jsii_.InvokeVoid(
		d,
		"resetEvents",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowIntent) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowIntent) ResetInputContextNames() {
	_jsii_.InvokeVoid(
		d,
		"resetInputContextNames",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowIntent) ResetIsFallback() {
	_jsii_.InvokeVoid(
		d,
		"resetIsFallback",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowIntent) ResetMlDisabled() {
	_jsii_.InvokeVoid(
		d,
		"resetMlDisabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowIntent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowIntent) ResetParentFollowupIntentName() {
	_jsii_.InvokeVoid(
		d,
		"resetParentFollowupIntentName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowIntent) ResetPriority() {
	_jsii_.InvokeVoid(
		d,
		"resetPriority",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowIntent) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowIntent) ResetResetContexts() {
	_jsii_.InvokeVoid(
		d,
		"resetResetContexts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowIntent) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowIntent) ResetWebhookState() {
	_jsii_.InvokeVoid(
		d,
		"resetWebhookState",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowIntent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowIntent) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowIntent) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowIntent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowIntent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowIntent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

