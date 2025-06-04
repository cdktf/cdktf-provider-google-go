// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigpatchdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v16/osconfigpatchdeployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_patch_deployment google_os_config_patch_deployment}.
type OsConfigPatchDeployment interface {
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
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Duration() *string
	SetDuration(val *string)
	DurationInput() *string
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
	InstanceFilter() OsConfigPatchDeploymentInstanceFilterOutputReference
	InstanceFilterInput() *OsConfigPatchDeploymentInstanceFilter
	LastExecuteTime() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
	OneTimeSchedule() OsConfigPatchDeploymentOneTimeScheduleOutputReference
	OneTimeScheduleInput() *OsConfigPatchDeploymentOneTimeSchedule
	PatchConfig() OsConfigPatchDeploymentPatchConfigOutputReference
	PatchConfigInput() *OsConfigPatchDeploymentPatchConfig
	PatchDeploymentId() *string
	SetPatchDeploymentId(val *string)
	PatchDeploymentIdInput() *string
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
	RecurringSchedule() OsConfigPatchDeploymentRecurringScheduleOutputReference
	RecurringScheduleInput() *OsConfigPatchDeploymentRecurringSchedule
	Rollout() OsConfigPatchDeploymentRolloutOutputReference
	RolloutInput() *OsConfigPatchDeploymentRollout
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() OsConfigPatchDeploymentTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
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
	PutInstanceFilter(value *OsConfigPatchDeploymentInstanceFilter)
	PutOneTimeSchedule(value *OsConfigPatchDeploymentOneTimeSchedule)
	PutPatchConfig(value *OsConfigPatchDeploymentPatchConfig)
	PutRecurringSchedule(value *OsConfigPatchDeploymentRecurringSchedule)
	PutRollout(value *OsConfigPatchDeploymentRollout)
	PutTimeouts(value *OsConfigPatchDeploymentTimeouts)
	ResetDescription()
	ResetDuration()
	ResetId()
	ResetOneTimeSchedule()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPatchConfig()
	ResetProject()
	ResetRecurringSchedule()
	ResetRollout()
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

// The jsii proxy struct for OsConfigPatchDeployment
type jsiiProxy_OsConfigPatchDeployment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OsConfigPatchDeployment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) InstanceFilter() OsConfigPatchDeploymentInstanceFilterOutputReference {
	var returns OsConfigPatchDeploymentInstanceFilterOutputReference
	_jsii_.Get(
		j,
		"instanceFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) InstanceFilterInput() *OsConfigPatchDeploymentInstanceFilter {
	var returns *OsConfigPatchDeploymentInstanceFilter
	_jsii_.Get(
		j,
		"instanceFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) LastExecuteTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastExecuteTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) OneTimeSchedule() OsConfigPatchDeploymentOneTimeScheduleOutputReference {
	var returns OsConfigPatchDeploymentOneTimeScheduleOutputReference
	_jsii_.Get(
		j,
		"oneTimeSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) OneTimeScheduleInput() *OsConfigPatchDeploymentOneTimeSchedule {
	var returns *OsConfigPatchDeploymentOneTimeSchedule
	_jsii_.Get(
		j,
		"oneTimeScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) PatchConfig() OsConfigPatchDeploymentPatchConfigOutputReference {
	var returns OsConfigPatchDeploymentPatchConfigOutputReference
	_jsii_.Get(
		j,
		"patchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) PatchConfigInput() *OsConfigPatchDeploymentPatchConfig {
	var returns *OsConfigPatchDeploymentPatchConfig
	_jsii_.Get(
		j,
		"patchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) PatchDeploymentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchDeploymentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) PatchDeploymentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchDeploymentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) RecurringSchedule() OsConfigPatchDeploymentRecurringScheduleOutputReference {
	var returns OsConfigPatchDeploymentRecurringScheduleOutputReference
	_jsii_.Get(
		j,
		"recurringSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) RecurringScheduleInput() *OsConfigPatchDeploymentRecurringSchedule {
	var returns *OsConfigPatchDeploymentRecurringSchedule
	_jsii_.Get(
		j,
		"recurringScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) Rollout() OsConfigPatchDeploymentRolloutOutputReference {
	var returns OsConfigPatchDeploymentRolloutOutputReference
	_jsii_.Get(
		j,
		"rollout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) RolloutInput() *OsConfigPatchDeploymentRollout {
	var returns *OsConfigPatchDeploymentRollout
	_jsii_.Get(
		j,
		"rolloutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) Timeouts() OsConfigPatchDeploymentTimeoutsOutputReference {
	var returns OsConfigPatchDeploymentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeployment) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_patch_deployment google_os_config_patch_deployment} Resource.
func NewOsConfigPatchDeployment(scope constructs.Construct, id *string, config *OsConfigPatchDeploymentConfig) OsConfigPatchDeployment {
	_init_.Initialize()

	if err := validateNewOsConfigPatchDeploymentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OsConfigPatchDeployment{}

	_jsii_.Create(
		"@cdktf/provider-google.osConfigPatchDeployment.OsConfigPatchDeployment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/os_config_patch_deployment google_os_config_patch_deployment} Resource.
func NewOsConfigPatchDeployment_Override(o OsConfigPatchDeployment, scope constructs.Construct, id *string, config *OsConfigPatchDeploymentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.osConfigPatchDeployment.OsConfigPatchDeployment",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OsConfigPatchDeployment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeployment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeployment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeployment)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeployment)SetDuration(val *string) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeployment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeployment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeployment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeployment)SetPatchDeploymentId(val *string) {
	if err := j.validateSetPatchDeploymentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patchDeploymentId",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeployment)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeployment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeployment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a OsConfigPatchDeployment resource upon running "cdktf plan <stack-name>".
func OsConfigPatchDeployment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOsConfigPatchDeployment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.osConfigPatchDeployment.OsConfigPatchDeployment",
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
func OsConfigPatchDeployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOsConfigPatchDeployment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.osConfigPatchDeployment.OsConfigPatchDeployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OsConfigPatchDeployment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOsConfigPatchDeployment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.osConfigPatchDeployment.OsConfigPatchDeployment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OsConfigPatchDeployment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOsConfigPatchDeployment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.osConfigPatchDeployment.OsConfigPatchDeployment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OsConfigPatchDeployment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.osConfigPatchDeployment.OsConfigPatchDeployment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OsConfigPatchDeployment) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeployment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeployment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeployment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeployment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeployment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeployment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeployment) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeployment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeployment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeployment) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeployment) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) PutInstanceFilter(value *OsConfigPatchDeploymentInstanceFilter) {
	if err := o.validatePutInstanceFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putInstanceFilter",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) PutOneTimeSchedule(value *OsConfigPatchDeploymentOneTimeSchedule) {
	if err := o.validatePutOneTimeScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putOneTimeSchedule",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) PutPatchConfig(value *OsConfigPatchDeploymentPatchConfig) {
	if err := o.validatePutPatchConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putPatchConfig",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) PutRecurringSchedule(value *OsConfigPatchDeploymentRecurringSchedule) {
	if err := o.validatePutRecurringScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRecurringSchedule",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) PutRollout(value *OsConfigPatchDeploymentRollout) {
	if err := o.validatePutRolloutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRollout",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) PutTimeouts(value *OsConfigPatchDeploymentTimeouts) {
	if err := o.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) ResetDescription() {
	_jsii_.InvokeVoid(
		o,
		"resetDescription",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) ResetDuration() {
	_jsii_.InvokeVoid(
		o,
		"resetDuration",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) ResetOneTimeSchedule() {
	_jsii_.InvokeVoid(
		o,
		"resetOneTimeSchedule",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) ResetPatchConfig() {
	_jsii_.InvokeVoid(
		o,
		"resetPatchConfig",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) ResetProject() {
	_jsii_.InvokeVoid(
		o,
		"resetProject",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) ResetRecurringSchedule() {
	_jsii_.InvokeVoid(
		o,
		"resetRecurringSchedule",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) ResetRollout() {
	_jsii_.InvokeVoid(
		o,
		"resetRollout",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) ResetTimeouts() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeployment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeployment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeployment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeployment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeployment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeployment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

