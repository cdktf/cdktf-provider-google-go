package monitoringalertpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v3/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v3/monitoringalertpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitoringAlertPolicyConditionsConditionThresholdAggregationsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) MonitoringAlertPolicyConditionsConditionThresholdAggregationsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitoringAlertPolicyConditionsConditionThresholdAggregationsList
type jsiiProxy_MonitoringAlertPolicyConditionsConditionThresholdAggregationsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionThresholdAggregationsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionThresholdAggregationsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionThresholdAggregationsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionThresholdAggregationsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionThresholdAggregationsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionThresholdAggregationsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewMonitoringAlertPolicyConditionsConditionThresholdAggregationsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) MonitoringAlertPolicyConditionsConditionThresholdAggregationsList {
	_init_.Initialize()

	if err := validateNewMonitoringAlertPolicyConditionsConditionThresholdAggregationsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitoringAlertPolicyConditionsConditionThresholdAggregationsList{}

	_jsii_.Create(
		"@cdktf/provider-google.monitoringAlertPolicy.MonitoringAlertPolicyConditionsConditionThresholdAggregationsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewMonitoringAlertPolicyConditionsConditionThresholdAggregationsList_Override(m MonitoringAlertPolicyConditionsConditionThresholdAggregationsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.monitoringAlertPolicy.MonitoringAlertPolicyConditionsConditionThresholdAggregationsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionThresholdAggregationsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionThresholdAggregationsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionThresholdAggregationsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitoringAlertPolicyConditionsConditionThresholdAggregationsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionThresholdAggregationsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionThresholdAggregationsList) Get(index *float64) MonitoringAlertPolicyConditionsConditionThresholdAggregationsOutputReference {
	if err := m.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns MonitoringAlertPolicyConditionsConditionThresholdAggregationsOutputReference

	_jsii_.Invoke(
		m,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionThresholdAggregationsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MonitoringAlertPolicyConditionsConditionThresholdAggregationsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

