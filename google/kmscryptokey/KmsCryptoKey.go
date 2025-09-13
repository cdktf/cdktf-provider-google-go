// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmscryptokey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v16/kmscryptokey/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/kms_crypto_key google_kms_crypto_key}.
type KmsCryptoKey interface {
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
	CryptoKeyBackend() *string
	SetCryptoKeyBackend(val *string)
	CryptoKeyBackendInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DestroyScheduledDuration() *string
	SetDestroyScheduledDuration(val *string)
	DestroyScheduledDurationInput() *string
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
	ImportOnly() interface{}
	SetImportOnly(val interface{})
	ImportOnlyInput() interface{}
	KeyRing() *string
	SetKeyRing(val *string)
	KeyRingInput() *string
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
	// The tree node.
	Node() constructs.Node
	Primary() KmsCryptoKeyPrimaryList
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Purpose() *string
	SetPurpose(val *string)
	PurposeInput() *string
	// Experimental.
	RawOverrides() interface{}
	RotationPeriod() *string
	SetRotationPeriod(val *string)
	RotationPeriodInput() *string
	SkipInitialVersionCreation() interface{}
	SetSkipInitialVersionCreation(val interface{})
	SkipInitialVersionCreationInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() KmsCryptoKeyTimeoutsOutputReference
	TimeoutsInput() interface{}
	VersionTemplate() KmsCryptoKeyVersionTemplateOutputReference
	VersionTemplateInput() *KmsCryptoKeyVersionTemplate
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
	PutTimeouts(value *KmsCryptoKeyTimeouts)
	PutVersionTemplate(value *KmsCryptoKeyVersionTemplate)
	ResetCryptoKeyBackend()
	ResetDestroyScheduledDuration()
	ResetId()
	ResetImportOnly()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPurpose()
	ResetRotationPeriod()
	ResetSkipInitialVersionCreation()
	ResetTimeouts()
	ResetVersionTemplate()
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

// The jsii proxy struct for KmsCryptoKey
type jsiiProxy_KmsCryptoKey struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KmsCryptoKey) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) CryptoKeyBackend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cryptoKeyBackend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) CryptoKeyBackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cryptoKeyBackendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) DestroyScheduledDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destroyScheduledDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) DestroyScheduledDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destroyScheduledDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) ImportOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) ImportOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) KeyRing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyRing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) KeyRingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyRingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) Primary() KmsCryptoKeyPrimaryList {
	var returns KmsCryptoKeyPrimaryList
	_jsii_.Get(
		j,
		"primary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) Purpose() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purpose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) PurposeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purposeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) RotationPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) RotationPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rotationPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) SkipInitialVersionCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipInitialVersionCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) SkipInitialVersionCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipInitialVersionCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) Timeouts() KmsCryptoKeyTimeoutsOutputReference {
	var returns KmsCryptoKeyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) VersionTemplate() KmsCryptoKeyVersionTemplateOutputReference {
	var returns KmsCryptoKeyVersionTemplateOutputReference
	_jsii_.Get(
		j,
		"versionTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKey) VersionTemplateInput() *KmsCryptoKeyVersionTemplate {
	var returns *KmsCryptoKeyVersionTemplate
	_jsii_.Get(
		j,
		"versionTemplateInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/kms_crypto_key google_kms_crypto_key} Resource.
func NewKmsCryptoKey(scope constructs.Construct, id *string, config *KmsCryptoKeyConfig) KmsCryptoKey {
	_init_.Initialize()

	if err := validateNewKmsCryptoKeyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KmsCryptoKey{}

	_jsii_.Create(
		"@cdktf/provider-google.kmsCryptoKey.KmsCryptoKey",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/kms_crypto_key google_kms_crypto_key} Resource.
func NewKmsCryptoKey_Override(k KmsCryptoKey, scope constructs.Construct, id *string, config *KmsCryptoKeyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.kmsCryptoKey.KmsCryptoKey",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KmsCryptoKey)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKey)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKey)SetCryptoKeyBackend(val *string) {
	if err := j.validateSetCryptoKeyBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cryptoKeyBackend",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKey)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKey)SetDestroyScheduledDuration(val *string) {
	if err := j.validateSetDestroyScheduledDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destroyScheduledDuration",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKey)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKey)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKey)SetImportOnly(val interface{}) {
	if err := j.validateSetImportOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importOnly",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKey)SetKeyRing(val *string) {
	if err := j.validateSetKeyRingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyRing",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKey)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKey)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKey)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKey)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKey)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKey)SetPurpose(val *string) {
	if err := j.validateSetPurposeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purpose",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKey)SetRotationPeriod(val *string) {
	if err := j.validateSetRotationPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotationPeriod",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKey)SetSkipInitialVersionCreation(val interface{}) {
	if err := j.validateSetSkipInitialVersionCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipInitialVersionCreation",
		val,
	)
}

// Generates CDKTF code for importing a KmsCryptoKey resource upon running "cdktf plan <stack-name>".
func KmsCryptoKey_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateKmsCryptoKey_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.kmsCryptoKey.KmsCryptoKey",
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
func KmsCryptoKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKmsCryptoKey_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.kmsCryptoKey.KmsCryptoKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KmsCryptoKey_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKmsCryptoKey_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.kmsCryptoKey.KmsCryptoKey",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KmsCryptoKey_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKmsCryptoKey_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.kmsCryptoKey.KmsCryptoKey",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KmsCryptoKey_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.kmsCryptoKey.KmsCryptoKey",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KmsCryptoKey) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KmsCryptoKey) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KmsCryptoKey) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKey) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKey) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKey) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKey) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKey) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKey) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKey) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKey) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKey) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKey) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KmsCryptoKey) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKey) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KmsCryptoKey) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KmsCryptoKey) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KmsCryptoKey) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KmsCryptoKey) PutTimeouts(value *KmsCryptoKeyTimeouts) {
	if err := k.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KmsCryptoKey) PutVersionTemplate(value *KmsCryptoKeyVersionTemplate) {
	if err := k.validatePutVersionTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putVersionTemplate",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KmsCryptoKey) ResetCryptoKeyBackend() {
	_jsii_.InvokeVoid(
		k,
		"resetCryptoKeyBackend",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsCryptoKey) ResetDestroyScheduledDuration() {
	_jsii_.InvokeVoid(
		k,
		"resetDestroyScheduledDuration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsCryptoKey) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsCryptoKey) ResetImportOnly() {
	_jsii_.InvokeVoid(
		k,
		"resetImportOnly",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsCryptoKey) ResetLabels() {
	_jsii_.InvokeVoid(
		k,
		"resetLabels",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsCryptoKey) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsCryptoKey) ResetPurpose() {
	_jsii_.InvokeVoid(
		k,
		"resetPurpose",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsCryptoKey) ResetRotationPeriod() {
	_jsii_.InvokeVoid(
		k,
		"resetRotationPeriod",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsCryptoKey) ResetSkipInitialVersionCreation() {
	_jsii_.InvokeVoid(
		k,
		"resetSkipInitialVersionCreation",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsCryptoKey) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsCryptoKey) ResetVersionTemplate() {
	_jsii_.InvokeVoid(
		k,
		"resetVersionTemplate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsCryptoKey) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKey) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKey) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKey) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKey) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

