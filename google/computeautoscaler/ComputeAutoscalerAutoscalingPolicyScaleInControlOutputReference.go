package computeautoscaler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v6/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v6/computeautoscaler/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeAutoscalerAutoscalingPolicyScaleInControl
	SetInternalValue(val *ComputeAutoscalerAutoscalingPolicyScaleInControl)
	MaxScaledInReplicas() ComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference
	MaxScaledInReplicasInput() *ComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeWindowSec() *float64
	SetTimeWindowSec(val *float64)
	TimeWindowSecInput() *float64
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
	PutMaxScaledInReplicas(value *ComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas)
	ResetMaxScaledInReplicas()
	ResetTimeWindowSec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference
type jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) InternalValue() *ComputeAutoscalerAutoscalingPolicyScaleInControl {
	var returns *ComputeAutoscalerAutoscalingPolicyScaleInControl
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) MaxScaledInReplicas() ComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference {
	var returns ComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference
	_jsii_.Get(
		j,
		"maxScaledInReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) MaxScaledInReplicasInput() *ComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas {
	var returns *ComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas
	_jsii_.Get(
		j,
		"maxScaledInReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) TimeWindowSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeWindowSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) TimeWindowSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeWindowSecInput",
		&returns,
	)
	return returns
}


func NewComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference {
	_init_.Initialize()

	if err := validateNewComputeAutoscalerAutoscalingPolicyScaleInControlOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeAutoscaler.ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference_Override(c ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeAutoscaler.ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference)SetInternalValue(val *ComputeAutoscalerAutoscalingPolicyScaleInControl) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference)SetTimeWindowSec(val *float64) {
	if err := j.validateSetTimeWindowSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeWindowSec",
		val,
	)
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) PutMaxScaledInReplicas(value *ComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) {
	if err := c.validatePutMaxScaledInReplicasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMaxScaledInReplicas",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) ResetMaxScaledInReplicas() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxScaledInReplicas",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) ResetTimeWindowSec() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeWindowSec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

