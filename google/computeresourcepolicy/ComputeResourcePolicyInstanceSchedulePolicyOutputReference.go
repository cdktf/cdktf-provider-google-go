package computeresourcepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v6/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v6/computeresourcepolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeResourcePolicyInstanceSchedulePolicyOutputReference interface {
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
	ExpirationTime() *string
	SetExpirationTime(val *string)
	ExpirationTimeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeResourcePolicyInstanceSchedulePolicy
	SetInternalValue(val *ComputeResourcePolicyInstanceSchedulePolicy)
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
	VmStartSchedule() ComputeResourcePolicyInstanceSchedulePolicyVmStartScheduleOutputReference
	VmStartScheduleInput() *ComputeResourcePolicyInstanceSchedulePolicyVmStartSchedule
	VmStopSchedule() ComputeResourcePolicyInstanceSchedulePolicyVmStopScheduleOutputReference
	VmStopScheduleInput() *ComputeResourcePolicyInstanceSchedulePolicyVmStopSchedule
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
	PutVmStartSchedule(value *ComputeResourcePolicyInstanceSchedulePolicyVmStartSchedule)
	PutVmStopSchedule(value *ComputeResourcePolicyInstanceSchedulePolicyVmStopSchedule)
	ResetExpirationTime()
	ResetStartTime()
	ResetVmStartSchedule()
	ResetVmStopSchedule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeResourcePolicyInstanceSchedulePolicyOutputReference
type jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) ExpirationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) ExpirationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) InternalValue() *ComputeResourcePolicyInstanceSchedulePolicy {
	var returns *ComputeResourcePolicyInstanceSchedulePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) VmStartSchedule() ComputeResourcePolicyInstanceSchedulePolicyVmStartScheduleOutputReference {
	var returns ComputeResourcePolicyInstanceSchedulePolicyVmStartScheduleOutputReference
	_jsii_.Get(
		j,
		"vmStartSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) VmStartScheduleInput() *ComputeResourcePolicyInstanceSchedulePolicyVmStartSchedule {
	var returns *ComputeResourcePolicyInstanceSchedulePolicyVmStartSchedule
	_jsii_.Get(
		j,
		"vmStartScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) VmStopSchedule() ComputeResourcePolicyInstanceSchedulePolicyVmStopScheduleOutputReference {
	var returns ComputeResourcePolicyInstanceSchedulePolicyVmStopScheduleOutputReference
	_jsii_.Get(
		j,
		"vmStopSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) VmStopScheduleInput() *ComputeResourcePolicyInstanceSchedulePolicyVmStopSchedule {
	var returns *ComputeResourcePolicyInstanceSchedulePolicyVmStopSchedule
	_jsii_.Get(
		j,
		"vmStopScheduleInput",
		&returns,
	)
	return returns
}


func NewComputeResourcePolicyInstanceSchedulePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeResourcePolicyInstanceSchedulePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewComputeResourcePolicyInstanceSchedulePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeResourcePolicy.ComputeResourcePolicyInstanceSchedulePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeResourcePolicyInstanceSchedulePolicyOutputReference_Override(c ComputeResourcePolicyInstanceSchedulePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeResourcePolicy.ComputeResourcePolicyInstanceSchedulePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference)SetExpirationTime(val *string) {
	if err := j.validateSetExpirationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationTime",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference)SetInternalValue(val *ComputeResourcePolicyInstanceSchedulePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) PutVmStartSchedule(value *ComputeResourcePolicyInstanceSchedulePolicyVmStartSchedule) {
	if err := c.validatePutVmStartScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVmStartSchedule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) PutVmStopSchedule(value *ComputeResourcePolicyInstanceSchedulePolicyVmStopSchedule) {
	if err := c.validatePutVmStopScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVmStopSchedule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) ResetExpirationTime() {
	_jsii_.InvokeVoid(
		c,
		"resetExpirationTime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		c,
		"resetStartTime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) ResetVmStartSchedule() {
	_jsii_.InvokeVoid(
		c,
		"resetVmStartSchedule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) ResetVmStopSchedule() {
	_jsii_.InvokeVoid(
		c,
		"resetVmStopSchedule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeResourcePolicyInstanceSchedulePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

