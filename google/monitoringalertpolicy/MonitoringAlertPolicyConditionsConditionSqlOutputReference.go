// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringalertpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/monitoringalertpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitoringAlertPolicyConditionsConditionSqlOutputReference interface {
	cdktf.ComplexObject
	BooleanTest() MonitoringAlertPolicyConditionsConditionSqlBooleanTestOutputReference
	BooleanTestInput() *MonitoringAlertPolicyConditionsConditionSqlBooleanTest
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Daily() MonitoringAlertPolicyConditionsConditionSqlDailyOutputReference
	DailyInput() *MonitoringAlertPolicyConditionsConditionSqlDaily
	// Experimental.
	Fqn() *string
	Hourly() MonitoringAlertPolicyConditionsConditionSqlHourlyOutputReference
	HourlyInput() *MonitoringAlertPolicyConditionsConditionSqlHourly
	InternalValue() *MonitoringAlertPolicyConditionsConditionSql
	SetInternalValue(val *MonitoringAlertPolicyConditionsConditionSql)
	Minutes() MonitoringAlertPolicyConditionsConditionSqlMinutesOutputReference
	MinutesInput() *MonitoringAlertPolicyConditionsConditionSqlMinutes
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
	RowCountTest() MonitoringAlertPolicyConditionsConditionSqlRowCountTestOutputReference
	RowCountTestInput() *MonitoringAlertPolicyConditionsConditionSqlRowCountTest
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutBooleanTest(value *MonitoringAlertPolicyConditionsConditionSqlBooleanTest)
	PutDaily(value *MonitoringAlertPolicyConditionsConditionSqlDaily)
	PutHourly(value *MonitoringAlertPolicyConditionsConditionSqlHourly)
	PutMinutes(value *MonitoringAlertPolicyConditionsConditionSqlMinutes)
	PutRowCountTest(value *MonitoringAlertPolicyConditionsConditionSqlRowCountTest)
	ResetBooleanTest()
	ResetDaily()
	ResetHourly()
	ResetMinutes()
	ResetRowCountTest()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitoringAlertPolicyConditionsConditionSqlOutputReference
type jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) BooleanTest() MonitoringAlertPolicyConditionsConditionSqlBooleanTestOutputReference {
	var returns MonitoringAlertPolicyConditionsConditionSqlBooleanTestOutputReference
	_jsii_.Get(
		j,
		"booleanTest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) BooleanTestInput() *MonitoringAlertPolicyConditionsConditionSqlBooleanTest {
	var returns *MonitoringAlertPolicyConditionsConditionSqlBooleanTest
	_jsii_.Get(
		j,
		"booleanTestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) Daily() MonitoringAlertPolicyConditionsConditionSqlDailyOutputReference {
	var returns MonitoringAlertPolicyConditionsConditionSqlDailyOutputReference
	_jsii_.Get(
		j,
		"daily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) DailyInput() *MonitoringAlertPolicyConditionsConditionSqlDaily {
	var returns *MonitoringAlertPolicyConditionsConditionSqlDaily
	_jsii_.Get(
		j,
		"dailyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) Hourly() MonitoringAlertPolicyConditionsConditionSqlHourlyOutputReference {
	var returns MonitoringAlertPolicyConditionsConditionSqlHourlyOutputReference
	_jsii_.Get(
		j,
		"hourly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) HourlyInput() *MonitoringAlertPolicyConditionsConditionSqlHourly {
	var returns *MonitoringAlertPolicyConditionsConditionSqlHourly
	_jsii_.Get(
		j,
		"hourlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) InternalValue() *MonitoringAlertPolicyConditionsConditionSql {
	var returns *MonitoringAlertPolicyConditionsConditionSql
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) Minutes() MonitoringAlertPolicyConditionsConditionSqlMinutesOutputReference {
	var returns MonitoringAlertPolicyConditionsConditionSqlMinutesOutputReference
	_jsii_.Get(
		j,
		"minutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) MinutesInput() *MonitoringAlertPolicyConditionsConditionSqlMinutes {
	var returns *MonitoringAlertPolicyConditionsConditionSqlMinutes
	_jsii_.Get(
		j,
		"minutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) RowCountTest() MonitoringAlertPolicyConditionsConditionSqlRowCountTestOutputReference {
	var returns MonitoringAlertPolicyConditionsConditionSqlRowCountTestOutputReference
	_jsii_.Get(
		j,
		"rowCountTest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) RowCountTestInput() *MonitoringAlertPolicyConditionsConditionSqlRowCountTest {
	var returns *MonitoringAlertPolicyConditionsConditionSqlRowCountTest
	_jsii_.Get(
		j,
		"rowCountTestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitoringAlertPolicyConditionsConditionSqlOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitoringAlertPolicyConditionsConditionSqlOutputReference {
	_init_.Initialize()

	if err := validateNewMonitoringAlertPolicyConditionsConditionSqlOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.monitoringAlertPolicy.MonitoringAlertPolicyConditionsConditionSqlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitoringAlertPolicyConditionsConditionSqlOutputReference_Override(m MonitoringAlertPolicyConditionsConditionSqlOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.monitoringAlertPolicy.MonitoringAlertPolicyConditionsConditionSqlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference)SetInternalValue(val *MonitoringAlertPolicyConditionsConditionSql) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) PutBooleanTest(value *MonitoringAlertPolicyConditionsConditionSqlBooleanTest) {
	if err := m.validatePutBooleanTestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBooleanTest",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) PutDaily(value *MonitoringAlertPolicyConditionsConditionSqlDaily) {
	if err := m.validatePutDailyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDaily",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) PutHourly(value *MonitoringAlertPolicyConditionsConditionSqlHourly) {
	if err := m.validatePutHourlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putHourly",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) PutMinutes(value *MonitoringAlertPolicyConditionsConditionSqlMinutes) {
	if err := m.validatePutMinutesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMinutes",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) PutRowCountTest(value *MonitoringAlertPolicyConditionsConditionSqlRowCountTest) {
	if err := m.validatePutRowCountTestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRowCountTest",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) ResetBooleanTest() {
	_jsii_.InvokeVoid(
		m,
		"resetBooleanTest",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) ResetDaily() {
	_jsii_.InvokeVoid(
		m,
		"resetDaily",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) ResetHourly() {
	_jsii_.InvokeVoid(
		m,
		"resetHourly",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) ResetMinutes() {
	_jsii_.InvokeVoid(
		m,
		"resetMinutes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) ResetRowCountTest() {
	_jsii_.InvokeVoid(
		m,
		"resetRowCountTest",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionSqlOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

