// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringslo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/monitoringslo/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/monitoring_slo google_monitoring_slo}.
type MonitoringSlo interface {
	cdktf.TerraformResource
	BasicSli() MonitoringSloBasicSliOutputReference
	BasicSliInput() *MonitoringSloBasicSli
	CalendarPeriod() *string
	SetCalendarPeriod(val *string)
	CalendarPeriodInput() *string
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
	Goal() *float64
	SetGoal(val *float64)
	GoalInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
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
	RequestBasedSli() MonitoringSloRequestBasedSliOutputReference
	RequestBasedSliInput() *MonitoringSloRequestBasedSli
	RollingPeriodDays() *float64
	SetRollingPeriodDays(val *float64)
	RollingPeriodDaysInput() *float64
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	SloId() *string
	SetSloId(val *string)
	SloIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MonitoringSloTimeoutsOutputReference
	TimeoutsInput() interface{}
	UserLabels() *map[string]*string
	SetUserLabels(val *map[string]*string)
	UserLabelsInput() *map[string]*string
	WindowsBasedSli() MonitoringSloWindowsBasedSliOutputReference
	WindowsBasedSliInput() *MonitoringSloWindowsBasedSli
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
	PutBasicSli(value *MonitoringSloBasicSli)
	PutRequestBasedSli(value *MonitoringSloRequestBasedSli)
	PutTimeouts(value *MonitoringSloTimeouts)
	PutWindowsBasedSli(value *MonitoringSloWindowsBasedSli)
	ResetBasicSli()
	ResetCalendarPeriod()
	ResetDisplayName()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRequestBasedSli()
	ResetRollingPeriodDays()
	ResetSloId()
	ResetTimeouts()
	ResetUserLabels()
	ResetWindowsBasedSli()
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

// The jsii proxy struct for MonitoringSlo
type jsiiProxy_MonitoringSlo struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MonitoringSlo) BasicSli() MonitoringSloBasicSliOutputReference {
	var returns MonitoringSloBasicSliOutputReference
	_jsii_.Get(
		j,
		"basicSli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) BasicSliInput() *MonitoringSloBasicSli {
	var returns *MonitoringSloBasicSli
	_jsii_.Get(
		j,
		"basicSliInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) CalendarPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"calendarPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) CalendarPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"calendarPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) Goal() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"goal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) GoalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"goalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) RequestBasedSli() MonitoringSloRequestBasedSliOutputReference {
	var returns MonitoringSloRequestBasedSliOutputReference
	_jsii_.Get(
		j,
		"requestBasedSli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) RequestBasedSliInput() *MonitoringSloRequestBasedSli {
	var returns *MonitoringSloRequestBasedSli
	_jsii_.Get(
		j,
		"requestBasedSliInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) RollingPeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rollingPeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) RollingPeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rollingPeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) SloId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sloId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) SloIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sloIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) Timeouts() MonitoringSloTimeoutsOutputReference {
	var returns MonitoringSloTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) UserLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) UserLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) WindowsBasedSli() MonitoringSloWindowsBasedSliOutputReference {
	var returns MonitoringSloWindowsBasedSliOutputReference
	_jsii_.Get(
		j,
		"windowsBasedSli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSlo) WindowsBasedSliInput() *MonitoringSloWindowsBasedSli {
	var returns *MonitoringSloWindowsBasedSli
	_jsii_.Get(
		j,
		"windowsBasedSliInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/monitoring_slo google_monitoring_slo} Resource.
func NewMonitoringSlo(scope constructs.Construct, id *string, config *MonitoringSloConfig) MonitoringSlo {
	_init_.Initialize()

	if err := validateNewMonitoringSloParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitoringSlo{}

	_jsii_.Create(
		"@cdktf/provider-google.monitoringSlo.MonitoringSlo",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/monitoring_slo google_monitoring_slo} Resource.
func NewMonitoringSlo_Override(m MonitoringSlo, scope constructs.Construct, id *string, config *MonitoringSloConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.monitoringSlo.MonitoringSlo",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MonitoringSlo)SetCalendarPeriod(val *string) {
	if err := j.validateSetCalendarPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"calendarPeriod",
		val,
	)
}

func (j *jsiiProxy_MonitoringSlo)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MonitoringSlo)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MonitoringSlo)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MonitoringSlo)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_MonitoringSlo)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MonitoringSlo)SetGoal(val *float64) {
	if err := j.validateSetGoalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"goal",
		val,
	)
}

func (j *jsiiProxy_MonitoringSlo)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MonitoringSlo)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MonitoringSlo)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_MonitoringSlo)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MonitoringSlo)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MonitoringSlo)SetRollingPeriodDays(val *float64) {
	if err := j.validateSetRollingPeriodDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rollingPeriodDays",
		val,
	)
}

func (j *jsiiProxy_MonitoringSlo)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_MonitoringSlo)SetSloId(val *string) {
	if err := j.validateSetSloIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sloId",
		val,
	)
}

func (j *jsiiProxy_MonitoringSlo)SetUserLabels(val *map[string]*string) {
	if err := j.validateSetUserLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userLabels",
		val,
	)
}

// Generates CDKTF code for importing a MonitoringSlo resource upon running "cdktf plan <stack-name>".
func MonitoringSlo_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMonitoringSlo_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.monitoringSlo.MonitoringSlo",
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
func MonitoringSlo_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitoringSlo_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.monitoringSlo.MonitoringSlo",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MonitoringSlo_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitoringSlo_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.monitoringSlo.MonitoringSlo",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MonitoringSlo_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitoringSlo_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.monitoringSlo.MonitoringSlo",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MonitoringSlo_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.monitoringSlo.MonitoringSlo",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MonitoringSlo) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MonitoringSlo) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MonitoringSlo) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSlo) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSlo) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSlo) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSlo) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSlo) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSlo) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSlo) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSlo) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSlo) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSlo) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MonitoringSlo) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSlo) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MonitoringSlo) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MonitoringSlo) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MonitoringSlo) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MonitoringSlo) PutBasicSli(value *MonitoringSloBasicSli) {
	if err := m.validatePutBasicSliParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBasicSli",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitoringSlo) PutRequestBasedSli(value *MonitoringSloRequestBasedSli) {
	if err := m.validatePutRequestBasedSliParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRequestBasedSli",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitoringSlo) PutTimeouts(value *MonitoringSloTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitoringSlo) PutWindowsBasedSli(value *MonitoringSloWindowsBasedSli) {
	if err := m.validatePutWindowsBasedSliParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putWindowsBasedSli",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitoringSlo) ResetBasicSli() {
	_jsii_.InvokeVoid(
		m,
		"resetBasicSli",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringSlo) ResetCalendarPeriod() {
	_jsii_.InvokeVoid(
		m,
		"resetCalendarPeriod",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringSlo) ResetDisplayName() {
	_jsii_.InvokeVoid(
		m,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringSlo) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringSlo) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringSlo) ResetProject() {
	_jsii_.InvokeVoid(
		m,
		"resetProject",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringSlo) ResetRequestBasedSli() {
	_jsii_.InvokeVoid(
		m,
		"resetRequestBasedSli",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringSlo) ResetRollingPeriodDays() {
	_jsii_.InvokeVoid(
		m,
		"resetRollingPeriodDays",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringSlo) ResetSloId() {
	_jsii_.InvokeVoid(
		m,
		"resetSloId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringSlo) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringSlo) ResetUserLabels() {
	_jsii_.InvokeVoid(
		m,
		"resetUserLabels",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringSlo) ResetWindowsBasedSli() {
	_jsii_.InvokeVoid(
		m,
		"resetWindowsBasedSli",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringSlo) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSlo) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSlo) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSlo) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSlo) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSlo) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

