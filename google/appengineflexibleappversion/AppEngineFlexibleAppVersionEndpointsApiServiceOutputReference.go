package appengineflexibleappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v8/appengineflexibleappversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference interface {
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
	ConfigId() *string
	SetConfigId(val *string)
	ConfigIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisableTraceSampling() interface{}
	SetDisableTraceSampling(val interface{})
	DisableTraceSamplingInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *AppEngineFlexibleAppVersionEndpointsApiService
	SetInternalValue(val *AppEngineFlexibleAppVersionEndpointsApiService)
	Name() *string
	SetName(val *string)
	NameInput() *string
	RolloutStrategy() *string
	SetRolloutStrategy(val *string)
	RolloutStrategyInput() *string
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
	ResetConfigId()
	ResetDisableTraceSampling()
	ResetRolloutStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference
type jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) ConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) ConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) DisableTraceSampling() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableTraceSampling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) DisableTraceSamplingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableTraceSamplingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) InternalValue() *AppEngineFlexibleAppVersionEndpointsApiService {
	var returns *AppEngineFlexibleAppVersionEndpointsApiService
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) RolloutStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolloutStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) RolloutStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolloutStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppEngineFlexibleAppVersionEndpointsApiServiceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference {
	_init_.Initialize()

	if err := validateNewAppEngineFlexibleAppVersionEndpointsApiServiceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppEngineFlexibleAppVersionEndpointsApiServiceOutputReference_Override(a AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference)SetConfigId(val *string) {
	if err := j.validateSetConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configId",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference)SetDisableTraceSampling(val interface{}) {
	if err := j.validateSetDisableTraceSamplingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableTraceSampling",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference)SetInternalValue(val *AppEngineFlexibleAppVersionEndpointsApiService) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference)SetRolloutStrategy(val *string) {
	if err := j.validateSetRolloutStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rolloutStrategy",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) ResetConfigId() {
	_jsii_.InvokeVoid(
		a,
		"resetConfigId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) ResetDisableTraceSampling() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableTraceSampling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) ResetRolloutStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetRolloutStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

