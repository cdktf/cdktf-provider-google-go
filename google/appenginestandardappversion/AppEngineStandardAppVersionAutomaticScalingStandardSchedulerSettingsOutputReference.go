package appenginestandardappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v5/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v5/appenginestandardappversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference interface {
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
	InternalValue() *AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings
	SetInternalValue(val *AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings)
	MaxInstances() *float64
	SetMaxInstances(val *float64)
	MaxInstancesInput() *float64
	MinInstances() *float64
	SetMinInstances(val *float64)
	MinInstancesInput() *float64
	TargetCpuUtilization() *float64
	SetTargetCpuUtilization(val *float64)
	TargetCpuUtilizationInput() *float64
	TargetThroughputUtilization() *float64
	SetTargetThroughputUtilization(val *float64)
	TargetThroughputUtilizationInput() *float64
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
	ResetMaxInstances()
	ResetMinInstances()
	ResetTargetCpuUtilization()
	ResetTargetThroughputUtilization()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference
type jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) InternalValue() *AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings {
	var returns *AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) MaxInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) MaxInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) MinInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) MinInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) TargetCpuUtilization() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetCpuUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) TargetCpuUtilizationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetCpuUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) TargetThroughputUtilization() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetThroughputUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) TargetThroughputUtilizationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetThroughputUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.appEngineStandardAppVersion.AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference_Override(a AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.appEngineStandardAppVersion.AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference)SetInternalValue(val *AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference)SetMaxInstances(val *float64) {
	if err := j.validateSetMaxInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInstances",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference)SetMinInstances(val *float64) {
	if err := j.validateSetMinInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minInstances",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference)SetTargetCpuUtilization(val *float64) {
	if err := j.validateSetTargetCpuUtilizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetCpuUtilization",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference)SetTargetThroughputUtilization(val *float64) {
	if err := j.validateSetTargetThroughputUtilizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetThroughputUtilization",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) ResetMaxInstances() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxInstances",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) ResetMinInstances() {
	_jsii_.InvokeVoid(
		a,
		"resetMinInstances",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) ResetTargetCpuUtilization() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetCpuUtilization",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) ResetTargetThroughputUtilization() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetThroughputUtilization",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
