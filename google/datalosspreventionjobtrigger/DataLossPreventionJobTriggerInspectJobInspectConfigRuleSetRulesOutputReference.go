package datalosspreventionjobtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v7/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v7/datalosspreventionjobtrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference interface {
	cdktf.ComplexObject
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
	ExclusionRule() DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference
	ExclusionRuleInput() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule
	// Experimental.
	Fqn() *string
	HotwordRule() DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference
	HotwordRuleInput() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	PutExclusionRule(value *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule)
	PutHotwordRule(value *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule)
	ResetExclusionRule()
	ResetHotwordRule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference
type jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) ExclusionRule() DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleOutputReference
	_jsii_.Get(
		j,
		"exclusionRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) ExclusionRuleInput() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule {
	var returns *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule
	_jsii_.Get(
		j,
		"exclusionRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) HotwordRule() DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleOutputReference
	_jsii_.Get(
		j,
		"hotwordRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) HotwordRuleInput() *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule {
	var returns *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule
	_jsii_.Get(
		j,
		"hotwordRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionJobTrigger.DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference_Override(d DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionJobTrigger.DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) PutExclusionRule(value *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule) {
	if err := d.validatePutExclusionRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExclusionRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) PutHotwordRule(value *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule) {
	if err := d.validatePutHotwordRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHotwordRule",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) ResetExclusionRule() {
	_jsii_.InvokeVoid(
		d,
		"resetExclusionRule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) ResetHotwordRule() {
	_jsii_.InvokeVoid(
		d,
		"resetHotwordRule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

