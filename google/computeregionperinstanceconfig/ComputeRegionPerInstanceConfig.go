// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionperinstanceconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/computeregionperinstanceconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/compute_region_per_instance_config google_compute_region_per_instance_config}.
type ComputeRegionPerInstanceConfig interface {
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	MinimalAction() *string
	SetMinimalAction(val *string)
	MinimalActionInput() *string
	MostDisruptiveAllowedAction() *string
	SetMostDisruptiveAllowedAction(val *string)
	MostDisruptiveAllowedActionInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PreservedState() ComputeRegionPerInstanceConfigPreservedStateOutputReference
	PreservedStateInput() *ComputeRegionPerInstanceConfigPreservedState
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
	RegionInstanceGroupManager() *string
	SetRegionInstanceGroupManager(val *string)
	RegionInstanceGroupManagerInput() *string
	RemoveInstanceOnDestroy() interface{}
	SetRemoveInstanceOnDestroy(val interface{})
	RemoveInstanceOnDestroyInput() interface{}
	RemoveInstanceStateOnDestroy() interface{}
	SetRemoveInstanceStateOnDestroy(val interface{})
	RemoveInstanceStateOnDestroyInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeRegionPerInstanceConfigTimeoutsOutputReference
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
	PutPreservedState(value *ComputeRegionPerInstanceConfigPreservedState)
	PutTimeouts(value *ComputeRegionPerInstanceConfigTimeouts)
	ResetId()
	ResetMinimalAction()
	ResetMostDisruptiveAllowedAction()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreservedState()
	ResetProject()
	ResetRegion()
	ResetRemoveInstanceOnDestroy()
	ResetRemoveInstanceStateOnDestroy()
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

// The jsii proxy struct for ComputeRegionPerInstanceConfig
type jsiiProxy_ComputeRegionPerInstanceConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) MinimalAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimalAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) MinimalActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimalActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) MostDisruptiveAllowedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mostDisruptiveAllowedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) MostDisruptiveAllowedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mostDisruptiveAllowedActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) PreservedState() ComputeRegionPerInstanceConfigPreservedStateOutputReference {
	var returns ComputeRegionPerInstanceConfigPreservedStateOutputReference
	_jsii_.Get(
		j,
		"preservedState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) PreservedStateInput() *ComputeRegionPerInstanceConfigPreservedState {
	var returns *ComputeRegionPerInstanceConfigPreservedState
	_jsii_.Get(
		j,
		"preservedStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) RegionInstanceGroupManager() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInstanceGroupManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) RegionInstanceGroupManagerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInstanceGroupManagerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) RemoveInstanceOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeInstanceOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) RemoveInstanceOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeInstanceOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) RemoveInstanceStateOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeInstanceStateOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) RemoveInstanceStateOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeInstanceStateOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) Timeouts() ComputeRegionPerInstanceConfigTimeoutsOutputReference {
	var returns ComputeRegionPerInstanceConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/compute_region_per_instance_config google_compute_region_per_instance_config} Resource.
func NewComputeRegionPerInstanceConfig(scope constructs.Construct, id *string, config *ComputeRegionPerInstanceConfigConfig) ComputeRegionPerInstanceConfig {
	_init_.Initialize()

	if err := validateNewComputeRegionPerInstanceConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionPerInstanceConfig{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionPerInstanceConfig.ComputeRegionPerInstanceConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/compute_region_per_instance_config google_compute_region_per_instance_config} Resource.
func NewComputeRegionPerInstanceConfig_Override(c ComputeRegionPerInstanceConfig, scope constructs.Construct, id *string, config *ComputeRegionPerInstanceConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionPerInstanceConfig.ComputeRegionPerInstanceConfig",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig)SetMinimalAction(val *string) {
	if err := j.validateSetMinimalActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimalAction",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig)SetMostDisruptiveAllowedAction(val *string) {
	if err := j.validateSetMostDisruptiveAllowedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mostDisruptiveAllowedAction",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig)SetRegionInstanceGroupManager(val *string) {
	if err := j.validateSetRegionInstanceGroupManagerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionInstanceGroupManager",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig)SetRemoveInstanceOnDestroy(val interface{}) {
	if err := j.validateSetRemoveInstanceOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removeInstanceOnDestroy",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionPerInstanceConfig)SetRemoveInstanceStateOnDestroy(val interface{}) {
	if err := j.validateSetRemoveInstanceStateOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removeInstanceStateOnDestroy",
		val,
	)
}

// Generates CDKTF code for importing a ComputeRegionPerInstanceConfig resource upon running "cdktf plan <stack-name>".
func ComputeRegionPerInstanceConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateComputeRegionPerInstanceConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionPerInstanceConfig.ComputeRegionPerInstanceConfig",
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
func ComputeRegionPerInstanceConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeRegionPerInstanceConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionPerInstanceConfig.ComputeRegionPerInstanceConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeRegionPerInstanceConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeRegionPerInstanceConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionPerInstanceConfig.ComputeRegionPerInstanceConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeRegionPerInstanceConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeRegionPerInstanceConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionPerInstanceConfig.ComputeRegionPerInstanceConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeRegionPerInstanceConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.computeRegionPerInstanceConfig.ComputeRegionPerInstanceConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) PutPreservedState(value *ComputeRegionPerInstanceConfigPreservedState) {
	if err := c.validatePutPreservedStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPreservedState",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) PutTimeouts(value *ComputeRegionPerInstanceConfigTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) ResetMinimalAction() {
	_jsii_.InvokeVoid(
		c,
		"resetMinimalAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) ResetMostDisruptiveAllowedAction() {
	_jsii_.InvokeVoid(
		c,
		"resetMostDisruptiveAllowedAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) ResetPreservedState() {
	_jsii_.InvokeVoid(
		c,
		"resetPreservedState",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) ResetRemoveInstanceOnDestroy() {
	_jsii_.InvokeVoid(
		c,
		"resetRemoveInstanceOnDestroy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) ResetRemoveInstanceStateOnDestroy() {
	_jsii_.InvokeVoid(
		c,
		"resetRemoveInstanceStateOnDestroy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionPerInstanceConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

