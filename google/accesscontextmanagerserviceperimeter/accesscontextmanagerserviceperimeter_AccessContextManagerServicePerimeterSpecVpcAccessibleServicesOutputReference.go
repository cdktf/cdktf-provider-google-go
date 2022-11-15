package accesscontextmanagerserviceperimeter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v4/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v4/accesscontextmanagerserviceperimeter/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference interface {
	cdktf.ComplexObject
	AllowedServices() *[]*string
	SetAllowedServices(val *[]*string)
	AllowedServicesInput() *[]*string
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
	EnableRestriction() interface{}
	SetEnableRestriction(val interface{})
	EnableRestrictionInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *AccessContextManagerServicePerimeterSpecVpcAccessibleServices
	SetInternalValue(val *AccessContextManagerServicePerimeterSpecVpcAccessibleServices)
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
	ResetAllowedServices()
	ResetEnableRestriction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference
type jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) AllowedServices() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) AllowedServicesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) EnableRestriction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) EnableRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) InternalValue() *AccessContextManagerServicePerimeterSpecVpcAccessibleServices {
	var returns *AccessContextManagerServicePerimeterSpecVpcAccessibleServices
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference {
	_init_.Initialize()

	if err := validateNewAccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.accessContextManagerServicePerimeter.AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference_Override(a AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.accessContextManagerServicePerimeter.AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference)SetAllowedServices(val *[]*string) {
	if err := j.validateSetAllowedServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedServices",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference)SetEnableRestriction(val interface{}) {
	if err := j.validateSetEnableRestrictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableRestriction",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference)SetInternalValue(val *AccessContextManagerServicePerimeterSpecVpcAccessibleServices) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) ResetAllowedServices() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedServices",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) ResetEnableRestriction() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableRestriction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccessContextManagerServicePerimeterSpecVpcAccessibleServicesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

