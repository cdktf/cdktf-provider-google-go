package cloudtasksqueue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v3/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v3/cloudtasksqueue/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudTasksQueueRetryConfigOutputReference interface {
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
	InternalValue() *CloudTasksQueueRetryConfig
	SetInternalValue(val *CloudTasksQueueRetryConfig)
	MaxAttempts() *float64
	SetMaxAttempts(val *float64)
	MaxAttemptsInput() *float64
	MaxBackoff() *string
	SetMaxBackoff(val *string)
	MaxBackoffInput() *string
	MaxDoublings() *float64
	SetMaxDoublings(val *float64)
	MaxDoublingsInput() *float64
	MaxRetryDuration() *string
	SetMaxRetryDuration(val *string)
	MaxRetryDurationInput() *string
	MinBackoff() *string
	SetMinBackoff(val *string)
	MinBackoffInput() *string
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
	ResetMaxAttempts()
	ResetMaxBackoff()
	ResetMaxDoublings()
	ResetMaxRetryDuration()
	ResetMinBackoff()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudTasksQueueRetryConfigOutputReference
type jsiiProxy_CloudTasksQueueRetryConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) InternalValue() *CloudTasksQueueRetryConfig {
	var returns *CloudTasksQueueRetryConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) MaxAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) MaxAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) MaxBackoff() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxBackoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) MaxBackoffInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxBackoffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) MaxDoublings() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDoublings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) MaxDoublingsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDoublingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) MaxRetryDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRetryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) MaxRetryDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRetryDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) MinBackoff() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minBackoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) MinBackoffInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minBackoffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudTasksQueueRetryConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudTasksQueueRetryConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudTasksQueueRetryConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudTasksQueueRetryConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudTasksQueue.CloudTasksQueueRetryConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudTasksQueueRetryConfigOutputReference_Override(c CloudTasksQueueRetryConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudTasksQueue.CloudTasksQueueRetryConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference)SetInternalValue(val *CloudTasksQueueRetryConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference)SetMaxAttempts(val *float64) {
	if err := j.validateSetMaxAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAttempts",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference)SetMaxBackoff(val *string) {
	if err := j.validateSetMaxBackoffParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBackoff",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference)SetMaxDoublings(val *float64) {
	if err := j.validateSetMaxDoublingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDoublings",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference)SetMaxRetryDuration(val *string) {
	if err := j.validateSetMaxRetryDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetryDuration",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference)SetMinBackoff(val *string) {
	if err := j.validateSetMinBackoffParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minBackoff",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueRetryConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) ResetMaxAttempts() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxAttempts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) ResetMaxBackoff() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxBackoff",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) ResetMaxDoublings() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxDoublings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) ResetMaxRetryDuration() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxRetryDuration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) ResetMinBackoff() {
	_jsii_.InvokeVoid(
		c,
		"resetMinBackoff",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudTasksQueueRetryConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

