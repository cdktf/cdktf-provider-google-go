package appenginestandardappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v4/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v4/appenginestandardappversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppEngineStandardAppVersionAutomaticScalingOutputReference interface {
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
	InternalValue() *AppEngineStandardAppVersionAutomaticScaling
	SetInternalValue(val *AppEngineStandardAppVersionAutomaticScaling)
	MaxConcurrentRequests() *float64
	SetMaxConcurrentRequests(val *float64)
	MaxConcurrentRequestsInput() *float64
	MaxIdleInstances() *float64
	SetMaxIdleInstances(val *float64)
	MaxIdleInstancesInput() *float64
	MaxPendingLatency() *string
	SetMaxPendingLatency(val *string)
	MaxPendingLatencyInput() *string
	MinIdleInstances() *float64
	SetMinIdleInstances(val *float64)
	MinIdleInstancesInput() *float64
	MinPendingLatency() *string
	SetMinPendingLatency(val *string)
	MinPendingLatencyInput() *string
	StandardSchedulerSettings() AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference
	StandardSchedulerSettingsInput() *AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings
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
	PutStandardSchedulerSettings(value *AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings)
	ResetMaxConcurrentRequests()
	ResetMaxIdleInstances()
	ResetMaxPendingLatency()
	ResetMinIdleInstances()
	ResetMinPendingLatency()
	ResetStandardSchedulerSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppEngineStandardAppVersionAutomaticScalingOutputReference
type jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) InternalValue() *AppEngineStandardAppVersionAutomaticScaling {
	var returns *AppEngineStandardAppVersionAutomaticScaling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) MaxConcurrentRequests() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) MaxConcurrentRequestsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentRequestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) MaxIdleInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) MaxIdleInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) MaxPendingLatency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPendingLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) MaxPendingLatencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPendingLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) MinIdleInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIdleInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) MinIdleInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIdleInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) MinPendingLatency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minPendingLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) MinPendingLatencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minPendingLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) StandardSchedulerSettings() AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference {
	var returns AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettingsOutputReference
	_jsii_.Get(
		j,
		"standardSchedulerSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) StandardSchedulerSettingsInput() *AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings {
	var returns *AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings
	_jsii_.Get(
		j,
		"standardSchedulerSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppEngineStandardAppVersionAutomaticScalingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppEngineStandardAppVersionAutomaticScalingOutputReference {
	_init_.Initialize()

	if err := validateNewAppEngineStandardAppVersionAutomaticScalingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.appEngineStandardAppVersion.AppEngineStandardAppVersionAutomaticScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppEngineStandardAppVersionAutomaticScalingOutputReference_Override(a AppEngineStandardAppVersionAutomaticScalingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.appEngineStandardAppVersion.AppEngineStandardAppVersionAutomaticScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference)SetInternalValue(val *AppEngineStandardAppVersionAutomaticScaling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference)SetMaxConcurrentRequests(val *float64) {
	if err := j.validateSetMaxConcurrentRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentRequests",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference)SetMaxIdleInstances(val *float64) {
	if err := j.validateSetMaxIdleInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIdleInstances",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference)SetMaxPendingLatency(val *string) {
	if err := j.validateSetMaxPendingLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPendingLatency",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference)SetMinIdleInstances(val *float64) {
	if err := j.validateSetMinIdleInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minIdleInstances",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference)SetMinPendingLatency(val *string) {
	if err := j.validateSetMinPendingLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minPendingLatency",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) PutStandardSchedulerSettings(value *AppEngineStandardAppVersionAutomaticScalingStandardSchedulerSettings) {
	if err := a.validatePutStandardSchedulerSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStandardSchedulerSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) ResetMaxConcurrentRequests() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxConcurrentRequests",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) ResetMaxIdleInstances() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxIdleInstances",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) ResetMaxPendingLatency() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxPendingLatency",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) ResetMinIdleInstances() {
	_jsii_.InvokeVoid(
		a,
		"resetMinIdleInstances",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) ResetMinPendingLatency() {
	_jsii_.InvokeVoid(
		a,
		"resetMinPendingLatency",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) ResetStandardSchedulerSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetStandardSchedulerSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppEngineStandardAppVersionAutomaticScalingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

