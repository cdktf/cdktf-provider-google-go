package apikeyskey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v8/apikeyskey/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApikeysKeyRestrictionsOutputReference interface {
	cdktf.ComplexObject
	AndroidKeyRestrictions() ApikeysKeyRestrictionsAndroidKeyRestrictionsOutputReference
	AndroidKeyRestrictionsInput() *ApikeysKeyRestrictionsAndroidKeyRestrictions
	ApiTargets() ApikeysKeyRestrictionsApiTargetsList
	ApiTargetsInput() interface{}
	BrowserKeyRestrictions() ApikeysKeyRestrictionsBrowserKeyRestrictionsOutputReference
	BrowserKeyRestrictionsInput() *ApikeysKeyRestrictionsBrowserKeyRestrictions
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
	InternalValue() *ApikeysKeyRestrictions
	SetInternalValue(val *ApikeysKeyRestrictions)
	IosKeyRestrictions() ApikeysKeyRestrictionsIosKeyRestrictionsOutputReference
	IosKeyRestrictionsInput() *ApikeysKeyRestrictionsIosKeyRestrictions
	ServerKeyRestrictions() ApikeysKeyRestrictionsServerKeyRestrictionsOutputReference
	ServerKeyRestrictionsInput() *ApikeysKeyRestrictionsServerKeyRestrictions
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
	PutAndroidKeyRestrictions(value *ApikeysKeyRestrictionsAndroidKeyRestrictions)
	PutApiTargets(value interface{})
	PutBrowserKeyRestrictions(value *ApikeysKeyRestrictionsBrowserKeyRestrictions)
	PutIosKeyRestrictions(value *ApikeysKeyRestrictionsIosKeyRestrictions)
	PutServerKeyRestrictions(value *ApikeysKeyRestrictionsServerKeyRestrictions)
	ResetAndroidKeyRestrictions()
	ResetApiTargets()
	ResetBrowserKeyRestrictions()
	ResetIosKeyRestrictions()
	ResetServerKeyRestrictions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApikeysKeyRestrictionsOutputReference
type jsiiProxy_ApikeysKeyRestrictionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference) AndroidKeyRestrictions() ApikeysKeyRestrictionsAndroidKeyRestrictionsOutputReference {
	var returns ApikeysKeyRestrictionsAndroidKeyRestrictionsOutputReference
	_jsii_.Get(
		j,
		"androidKeyRestrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference) AndroidKeyRestrictionsInput() *ApikeysKeyRestrictionsAndroidKeyRestrictions {
	var returns *ApikeysKeyRestrictionsAndroidKeyRestrictions
	_jsii_.Get(
		j,
		"androidKeyRestrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference) ApiTargets() ApikeysKeyRestrictionsApiTargetsList {
	var returns ApikeysKeyRestrictionsApiTargetsList
	_jsii_.Get(
		j,
		"apiTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference) ApiTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference) BrowserKeyRestrictions() ApikeysKeyRestrictionsBrowserKeyRestrictionsOutputReference {
	var returns ApikeysKeyRestrictionsBrowserKeyRestrictionsOutputReference
	_jsii_.Get(
		j,
		"browserKeyRestrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference) BrowserKeyRestrictionsInput() *ApikeysKeyRestrictionsBrowserKeyRestrictions {
	var returns *ApikeysKeyRestrictionsBrowserKeyRestrictions
	_jsii_.Get(
		j,
		"browserKeyRestrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference) InternalValue() *ApikeysKeyRestrictions {
	var returns *ApikeysKeyRestrictions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference) IosKeyRestrictions() ApikeysKeyRestrictionsIosKeyRestrictionsOutputReference {
	var returns ApikeysKeyRestrictionsIosKeyRestrictionsOutputReference
	_jsii_.Get(
		j,
		"iosKeyRestrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference) IosKeyRestrictionsInput() *ApikeysKeyRestrictionsIosKeyRestrictions {
	var returns *ApikeysKeyRestrictionsIosKeyRestrictions
	_jsii_.Get(
		j,
		"iosKeyRestrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference) ServerKeyRestrictions() ApikeysKeyRestrictionsServerKeyRestrictionsOutputReference {
	var returns ApikeysKeyRestrictionsServerKeyRestrictionsOutputReference
	_jsii_.Get(
		j,
		"serverKeyRestrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference) ServerKeyRestrictionsInput() *ApikeysKeyRestrictionsServerKeyRestrictions {
	var returns *ApikeysKeyRestrictionsServerKeyRestrictions
	_jsii_.Get(
		j,
		"serverKeyRestrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApikeysKeyRestrictionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApikeysKeyRestrictionsOutputReference {
	_init_.Initialize()

	if err := validateNewApikeysKeyRestrictionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApikeysKeyRestrictionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.apikeysKey.ApikeysKeyRestrictionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApikeysKeyRestrictionsOutputReference_Override(a ApikeysKeyRestrictionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.apikeysKey.ApikeysKeyRestrictionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference)SetInternalValue(val *ApikeysKeyRestrictions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApikeysKeyRestrictionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) PutAndroidKeyRestrictions(value *ApikeysKeyRestrictionsAndroidKeyRestrictions) {
	if err := a.validatePutAndroidKeyRestrictionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAndroidKeyRestrictions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) PutApiTargets(value interface{}) {
	if err := a.validatePutApiTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApiTargets",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) PutBrowserKeyRestrictions(value *ApikeysKeyRestrictionsBrowserKeyRestrictions) {
	if err := a.validatePutBrowserKeyRestrictionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBrowserKeyRestrictions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) PutIosKeyRestrictions(value *ApikeysKeyRestrictionsIosKeyRestrictions) {
	if err := a.validatePutIosKeyRestrictionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIosKeyRestrictions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) PutServerKeyRestrictions(value *ApikeysKeyRestrictionsServerKeyRestrictions) {
	if err := a.validatePutServerKeyRestrictionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServerKeyRestrictions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) ResetAndroidKeyRestrictions() {
	_jsii_.InvokeVoid(
		a,
		"resetAndroidKeyRestrictions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) ResetApiTargets() {
	_jsii_.InvokeVoid(
		a,
		"resetApiTargets",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) ResetBrowserKeyRestrictions() {
	_jsii_.InvokeVoid(
		a,
		"resetBrowserKeyRestrictions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) ResetIosKeyRestrictions() {
	_jsii_.InvokeVoid(
		a,
		"resetIosKeyRestrictions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) ResetServerKeyRestrictions() {
	_jsii_.InvokeVoid(
		a,
		"resetServerKeyRestrictions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApikeysKeyRestrictionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

