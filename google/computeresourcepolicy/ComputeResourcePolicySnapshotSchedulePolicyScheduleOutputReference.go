package computeresourcepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v8/computeresourcepolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference interface {
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
	DailySchedule() ComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleOutputReference
	DailyScheduleInput() *ComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule
	// Experimental.
	Fqn() *string
	HourlySchedule() ComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleOutputReference
	HourlyScheduleInput() *ComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule
	InternalValue() *ComputeResourcePolicySnapshotSchedulePolicySchedule
	SetInternalValue(val *ComputeResourcePolicySnapshotSchedulePolicySchedule)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeeklySchedule() ComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleOutputReference
	WeeklyScheduleInput() *ComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule
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
	PutDailySchedule(value *ComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule)
	PutHourlySchedule(value *ComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule)
	PutWeeklySchedule(value *ComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule)
	ResetDailySchedule()
	ResetHourlySchedule()
	ResetWeeklySchedule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference
type jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) DailySchedule() ComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleOutputReference {
	var returns ComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleOutputReference
	_jsii_.Get(
		j,
		"dailySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) DailyScheduleInput() *ComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule {
	var returns *ComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule
	_jsii_.Get(
		j,
		"dailyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) HourlySchedule() ComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleOutputReference {
	var returns ComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleOutputReference
	_jsii_.Get(
		j,
		"hourlySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) HourlyScheduleInput() *ComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule {
	var returns *ComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule
	_jsii_.Get(
		j,
		"hourlyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) InternalValue() *ComputeResourcePolicySnapshotSchedulePolicySchedule {
	var returns *ComputeResourcePolicySnapshotSchedulePolicySchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) WeeklySchedule() ComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleOutputReference {
	var returns ComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleOutputReference
	_jsii_.Get(
		j,
		"weeklySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) WeeklyScheduleInput() *ComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule {
	var returns *ComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule
	_jsii_.Get(
		j,
		"weeklyScheduleInput",
		&returns,
	)
	return returns
}


func NewComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeResourcePolicy.ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference_Override(c ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeResourcePolicy.ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference)SetInternalValue(val *ComputeResourcePolicySnapshotSchedulePolicySchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) PutDailySchedule(value *ComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule) {
	if err := c.validatePutDailyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDailySchedule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) PutHourlySchedule(value *ComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule) {
	if err := c.validatePutHourlyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHourlySchedule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) PutWeeklySchedule(value *ComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule) {
	if err := c.validatePutWeeklyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWeeklySchedule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) ResetDailySchedule() {
	_jsii_.InvokeVoid(
		c,
		"resetDailySchedule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) ResetHourlySchedule() {
	_jsii_.InvokeVoid(
		c,
		"resetHourlySchedule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) ResetWeeklySchedule() {
	_jsii_.InvokeVoid(
		c,
		"resetWeeklySchedule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

