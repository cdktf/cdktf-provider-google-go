package accesscontextmanagerserviceperimeters

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v5/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v5/accesscontextmanagerserviceperimeters/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference interface {
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
	Identities() *[]*string
	SetIdentities(val *[]*string)
	IdentitiesInput() *[]*string
	IdentityType() *string
	SetIdentityType(val *string)
	IdentityTypeInput() *string
	InternalValue() *AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFrom
	SetInternalValue(val *AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFrom)
	Sources() AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromSourcesList
	SourcesInput() interface{}
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
	PutSources(value interface{})
	ResetIdentities()
	ResetIdentityType()
	ResetSources()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference
type jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) Identities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) IdentitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) IdentityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) IdentityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) InternalValue() *AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFrom {
	var returns *AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFrom
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) Sources() AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromSourcesList {
	var returns AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromSourcesList
	_jsii_.Get(
		j,
		"sources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) SourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference {
	_init_.Initialize()

	if err := validateNewAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.accessContextManagerServicePerimeters.AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference_Override(a AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.accessContextManagerServicePerimeters.AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference)SetIdentities(val *[]*string) {
	if err := j.validateSetIdentitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identities",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference)SetIdentityType(val *string) {
	if err := j.validateSetIdentityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityType",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference)SetInternalValue(val *AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFrom) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) PutSources(value interface{}) {
	if err := a.validatePutSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSources",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) ResetIdentities() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentities",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) ResetIdentityType() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentityType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) ResetSources() {
	_jsii_.InvokeVoid(
		a,
		"resetSources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFromOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
