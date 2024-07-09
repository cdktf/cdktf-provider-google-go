// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxentitytype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/dialogflowcxentitytype/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/dialogflow_cx_entity_type google_dialogflow_cx_entity_type}.
type DialogflowCxEntityType interface {
	cdktf.TerraformResource
	AutoExpansionMode() *string
	SetAutoExpansionMode(val *string)
	AutoExpansionModeInput() *string
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
	EnableFuzzyExtraction() interface{}
	SetEnableFuzzyExtraction(val interface{})
	EnableFuzzyExtractionInput() interface{}
	Entities() DialogflowCxEntityTypeEntitiesList
	EntitiesInput() interface{}
	ExcludedPhrases() DialogflowCxEntityTypeExcludedPhrasesList
	ExcludedPhrasesInput() interface{}
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
	Kind() *string
	SetKind(val *string)
	KindInput() *string
	LanguageCode() *string
	SetLanguageCode(val *string)
	LanguageCodeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
	Parent() *string
	SetParent(val *string)
	ParentInput() *string
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
	Redact() interface{}
	SetRedact(val interface{})
	RedactInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DialogflowCxEntityTypeTimeoutsOutputReference
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
	PutEntities(value interface{})
	PutExcludedPhrases(value interface{})
	PutTimeouts(value *DialogflowCxEntityTypeTimeouts)
	ResetAutoExpansionMode()
	ResetEnableFuzzyExtraction()
	ResetExcludedPhrases()
	ResetId()
	ResetLanguageCode()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParent()
	ResetRedact()
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

// The jsii proxy struct for DialogflowCxEntityType
type jsiiProxy_DialogflowCxEntityType struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DialogflowCxEntityType) AutoExpansionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoExpansionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) AutoExpansionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoExpansionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) EnableFuzzyExtraction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFuzzyExtraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) EnableFuzzyExtractionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFuzzyExtractionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) Entities() DialogflowCxEntityTypeEntitiesList {
	var returns DialogflowCxEntityTypeEntitiesList
	_jsii_.Get(
		j,
		"entities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) EntitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) ExcludedPhrases() DialogflowCxEntityTypeExcludedPhrasesList {
	var returns DialogflowCxEntityTypeExcludedPhrasesList
	_jsii_.Get(
		j,
		"excludedPhrases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) ExcludedPhrasesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludedPhrasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) KindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) LanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) LanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) Parent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) ParentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) Redact() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) RedactInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) Timeouts() DialogflowCxEntityTypeTimeoutsOutputReference {
	var returns DialogflowCxEntityTypeTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxEntityType) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/dialogflow_cx_entity_type google_dialogflow_cx_entity_type} Resource.
func NewDialogflowCxEntityType(scope constructs.Construct, id *string, config *DialogflowCxEntityTypeConfig) DialogflowCxEntityType {
	_init_.Initialize()

	if err := validateNewDialogflowCxEntityTypeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxEntityType{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxEntityType.DialogflowCxEntityType",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/dialogflow_cx_entity_type google_dialogflow_cx_entity_type} Resource.
func NewDialogflowCxEntityType_Override(d DialogflowCxEntityType, scope constructs.Construct, id *string, config *DialogflowCxEntityTypeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxEntityType.DialogflowCxEntityType",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxEntityType)SetAutoExpansionMode(val *string) {
	if err := j.validateSetAutoExpansionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoExpansionMode",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxEntityType)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxEntityType)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxEntityType)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxEntityType)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxEntityType)SetEnableFuzzyExtraction(val interface{}) {
	if err := j.validateSetEnableFuzzyExtractionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableFuzzyExtraction",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxEntityType)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxEntityType)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxEntityType)SetKind(val *string) {
	if err := j.validateSetKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxEntityType)SetLanguageCode(val *string) {
	if err := j.validateSetLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageCode",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxEntityType)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxEntityType)SetParent(val *string) {
	if err := j.validateSetParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parent",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxEntityType)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxEntityType)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxEntityType)SetRedact(val interface{}) {
	if err := j.validateSetRedactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redact",
		val,
	)
}

// Generates CDKTF code for importing a DialogflowCxEntityType resource upon running "cdktf plan <stack-name>".
func DialogflowCxEntityType_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDialogflowCxEntityType_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dialogflowCxEntityType.DialogflowCxEntityType",
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
func DialogflowCxEntityType_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDialogflowCxEntityType_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dialogflowCxEntityType.DialogflowCxEntityType",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DialogflowCxEntityType_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDialogflowCxEntityType_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dialogflowCxEntityType.DialogflowCxEntityType",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DialogflowCxEntityType_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDialogflowCxEntityType_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dialogflowCxEntityType.DialogflowCxEntityType",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DialogflowCxEntityType_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.dialogflowCxEntityType.DialogflowCxEntityType",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DialogflowCxEntityType) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DialogflowCxEntityType) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DialogflowCxEntityType) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowCxEntityType) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxEntityType) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxEntityType) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxEntityType) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxEntityType) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxEntityType) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxEntityType) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxEntityType) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxEntityType) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxEntityType) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DialogflowCxEntityType) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxEntityType) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DialogflowCxEntityType) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DialogflowCxEntityType) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DialogflowCxEntityType) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DialogflowCxEntityType) PutEntities(value interface{}) {
	if err := d.validatePutEntitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEntities",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxEntityType) PutExcludedPhrases(value interface{}) {
	if err := d.validatePutExcludedPhrasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExcludedPhrases",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxEntityType) PutTimeouts(value *DialogflowCxEntityTypeTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxEntityType) ResetAutoExpansionMode() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoExpansionMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxEntityType) ResetEnableFuzzyExtraction() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableFuzzyExtraction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxEntityType) ResetExcludedPhrases() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludedPhrases",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxEntityType) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxEntityType) ResetLanguageCode() {
	_jsii_.InvokeVoid(
		d,
		"resetLanguageCode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxEntityType) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxEntityType) ResetParent() {
	_jsii_.InvokeVoid(
		d,
		"resetParent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxEntityType) ResetRedact() {
	_jsii_.InvokeVoid(
		d,
		"resetRedact",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxEntityType) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxEntityType) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxEntityType) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxEntityType) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxEntityType) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxEntityType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxEntityType) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

