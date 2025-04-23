// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package recaptchaenterprisekey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/recaptchaenterprisekey/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/recaptcha_enterprise_key google_recaptcha_enterprise_key}.
type RecaptchaEnterpriseKey interface {
	cdktf.TerraformResource
	AndroidSettings() RecaptchaEnterpriseKeyAndroidSettingsOutputReference
	AndroidSettingsInput() *RecaptchaEnterpriseKeyAndroidSettings
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EffectiveLabels() cdktf.StringMap
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
	IosSettings() RecaptchaEnterpriseKeyIosSettingsOutputReference
	IosSettingsInput() *RecaptchaEnterpriseKeyIosSettings
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TestingOptions() RecaptchaEnterpriseKeyTestingOptionsOutputReference
	TestingOptionsInput() *RecaptchaEnterpriseKeyTestingOptions
	Timeouts() RecaptchaEnterpriseKeyTimeoutsOutputReference
	TimeoutsInput() interface{}
	WafSettings() RecaptchaEnterpriseKeyWafSettingsOutputReference
	WafSettingsInput() *RecaptchaEnterpriseKeyWafSettings
	WebSettings() RecaptchaEnterpriseKeyWebSettingsOutputReference
	WebSettingsInput() *RecaptchaEnterpriseKeyWebSettings
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
	PutAndroidSettings(value *RecaptchaEnterpriseKeyAndroidSettings)
	PutIosSettings(value *RecaptchaEnterpriseKeyIosSettings)
	PutTestingOptions(value *RecaptchaEnterpriseKeyTestingOptions)
	PutTimeouts(value *RecaptchaEnterpriseKeyTimeouts)
	PutWafSettings(value *RecaptchaEnterpriseKeyWafSettings)
	PutWebSettings(value *RecaptchaEnterpriseKeyWebSettings)
	ResetAndroidSettings()
	ResetId()
	ResetIosSettings()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetTestingOptions()
	ResetTimeouts()
	ResetWafSettings()
	ResetWebSettings()
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

// The jsii proxy struct for RecaptchaEnterpriseKey
type jsiiProxy_RecaptchaEnterpriseKey struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) AndroidSettings() RecaptchaEnterpriseKeyAndroidSettingsOutputReference {
	var returns RecaptchaEnterpriseKeyAndroidSettingsOutputReference
	_jsii_.Get(
		j,
		"androidSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) AndroidSettingsInput() *RecaptchaEnterpriseKeyAndroidSettings {
	var returns *RecaptchaEnterpriseKeyAndroidSettings
	_jsii_.Get(
		j,
		"androidSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) IosSettings() RecaptchaEnterpriseKeyIosSettingsOutputReference {
	var returns RecaptchaEnterpriseKeyIosSettingsOutputReference
	_jsii_.Get(
		j,
		"iosSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) IosSettingsInput() *RecaptchaEnterpriseKeyIosSettings {
	var returns *RecaptchaEnterpriseKeyIosSettings
	_jsii_.Get(
		j,
		"iosSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) TestingOptions() RecaptchaEnterpriseKeyTestingOptionsOutputReference {
	var returns RecaptchaEnterpriseKeyTestingOptionsOutputReference
	_jsii_.Get(
		j,
		"testingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) TestingOptionsInput() *RecaptchaEnterpriseKeyTestingOptions {
	var returns *RecaptchaEnterpriseKeyTestingOptions
	_jsii_.Get(
		j,
		"testingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) Timeouts() RecaptchaEnterpriseKeyTimeoutsOutputReference {
	var returns RecaptchaEnterpriseKeyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) WafSettings() RecaptchaEnterpriseKeyWafSettingsOutputReference {
	var returns RecaptchaEnterpriseKeyWafSettingsOutputReference
	_jsii_.Get(
		j,
		"wafSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) WafSettingsInput() *RecaptchaEnterpriseKeyWafSettings {
	var returns *RecaptchaEnterpriseKeyWafSettings
	_jsii_.Get(
		j,
		"wafSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) WebSettings() RecaptchaEnterpriseKeyWebSettingsOutputReference {
	var returns RecaptchaEnterpriseKeyWebSettingsOutputReference
	_jsii_.Get(
		j,
		"webSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKey) WebSettingsInput() *RecaptchaEnterpriseKeyWebSettings {
	var returns *RecaptchaEnterpriseKeyWebSettings
	_jsii_.Get(
		j,
		"webSettingsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/recaptcha_enterprise_key google_recaptcha_enterprise_key} Resource.
func NewRecaptchaEnterpriseKey(scope constructs.Construct, id *string, config *RecaptchaEnterpriseKeyConfig) RecaptchaEnterpriseKey {
	_init_.Initialize()

	if err := validateNewRecaptchaEnterpriseKeyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RecaptchaEnterpriseKey{}

	_jsii_.Create(
		"@cdktf/provider-google.recaptchaEnterpriseKey.RecaptchaEnterpriseKey",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/recaptcha_enterprise_key google_recaptcha_enterprise_key} Resource.
func NewRecaptchaEnterpriseKey_Override(r RecaptchaEnterpriseKey, scope constructs.Construct, id *string, config *RecaptchaEnterpriseKeyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.recaptchaEnterpriseKey.RecaptchaEnterpriseKey",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKey)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKey)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKey)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKey)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKey)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKey)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKey)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKey)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKey)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKey)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKey)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a RecaptchaEnterpriseKey resource upon running "cdktf plan <stack-name>".
func RecaptchaEnterpriseKey_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRecaptchaEnterpriseKey_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.recaptchaEnterpriseKey.RecaptchaEnterpriseKey",
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
func RecaptchaEnterpriseKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRecaptchaEnterpriseKey_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.recaptchaEnterpriseKey.RecaptchaEnterpriseKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RecaptchaEnterpriseKey_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRecaptchaEnterpriseKey_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.recaptchaEnterpriseKey.RecaptchaEnterpriseKey",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RecaptchaEnterpriseKey_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRecaptchaEnterpriseKey_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.recaptchaEnterpriseKey.RecaptchaEnterpriseKey",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RecaptchaEnterpriseKey_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.recaptchaEnterpriseKey.RecaptchaEnterpriseKey",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) PutAndroidSettings(value *RecaptchaEnterpriseKeyAndroidSettings) {
	if err := r.validatePutAndroidSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAndroidSettings",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) PutIosSettings(value *RecaptchaEnterpriseKeyIosSettings) {
	if err := r.validatePutIosSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putIosSettings",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) PutTestingOptions(value *RecaptchaEnterpriseKeyTestingOptions) {
	if err := r.validatePutTestingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTestingOptions",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) PutTimeouts(value *RecaptchaEnterpriseKeyTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) PutWafSettings(value *RecaptchaEnterpriseKeyWafSettings) {
	if err := r.validatePutWafSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putWafSettings",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) PutWebSettings(value *RecaptchaEnterpriseKeyWebSettings) {
	if err := r.validatePutWebSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putWebSettings",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) ResetAndroidSettings() {
	_jsii_.InvokeVoid(
		r,
		"resetAndroidSettings",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) ResetIosSettings() {
	_jsii_.InvokeVoid(
		r,
		"resetIosSettings",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) ResetLabels() {
	_jsii_.InvokeVoid(
		r,
		"resetLabels",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) ResetProject() {
	_jsii_.InvokeVoid(
		r,
		"resetProject",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) ResetTestingOptions() {
	_jsii_.InvokeVoid(
		r,
		"resetTestingOptions",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) ResetWafSettings() {
	_jsii_.InvokeVoid(
		r,
		"resetWafSettings",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) ResetWebSettings() {
	_jsii_.InvokeVoid(
		r,
		"resetWebSettings",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKey) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

