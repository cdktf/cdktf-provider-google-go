package recaptchaenterprisekey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-google-go/google/v3/jsii"

	"github.com/hashicorp/cdktf-provider-google-go/google/v3/recaptchaenterprisekey/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RecaptchaEnterpriseKeyAndroidSettingsOutputReference interface {
	cdktf.ComplexObject
	AllowAllPackageNames() interface{}
	SetAllowAllPackageNames(val interface{})
	AllowAllPackageNamesInput() interface{}
	AllowedPackageNames() *[]*string
	SetAllowedPackageNames(val *[]*string)
	AllowedPackageNamesInput() *[]*string
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
	InternalValue() *RecaptchaEnterpriseKeyAndroidSettings
	SetInternalValue(val *RecaptchaEnterpriseKeyAndroidSettings)
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
	ResetAllowAllPackageNames()
	ResetAllowedPackageNames()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RecaptchaEnterpriseKeyAndroidSettingsOutputReference
type jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) AllowAllPackageNames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAllPackageNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) AllowAllPackageNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAllPackageNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) AllowedPackageNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedPackageNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) AllowedPackageNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedPackageNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) InternalValue() *RecaptchaEnterpriseKeyAndroidSettings {
	var returns *RecaptchaEnterpriseKeyAndroidSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRecaptchaEnterpriseKeyAndroidSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RecaptchaEnterpriseKeyAndroidSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewRecaptchaEnterpriseKeyAndroidSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.recaptchaEnterpriseKey.RecaptchaEnterpriseKeyAndroidSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRecaptchaEnterpriseKeyAndroidSettingsOutputReference_Override(r RecaptchaEnterpriseKeyAndroidSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.recaptchaEnterpriseKey.RecaptchaEnterpriseKeyAndroidSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference)SetAllowAllPackageNames(val interface{}) {
	if err := j.validateSetAllowAllPackageNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAllPackageNames",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference)SetAllowedPackageNames(val *[]*string) {
	if err := j.validateSetAllowedPackageNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedPackageNames",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference)SetInternalValue(val *RecaptchaEnterpriseKeyAndroidSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) ResetAllowAllPackageNames() {
	_jsii_.InvokeVoid(
		r,
		"resetAllowAllPackageNames",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) ResetAllowedPackageNames() {
	_jsii_.InvokeVoid(
		r,
		"resetAllowedPackageNames",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecaptchaEnterpriseKeyAndroidSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

