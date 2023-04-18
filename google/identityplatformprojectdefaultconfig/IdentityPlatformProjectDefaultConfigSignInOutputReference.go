package identityplatformprojectdefaultconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v6/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v6/identityplatformprojectdefaultconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdentityPlatformProjectDefaultConfigSignInOutputReference interface {
	cdktf.ComplexObject
	AllowDuplicateEmails() interface{}
	SetAllowDuplicateEmails(val interface{})
	AllowDuplicateEmailsInput() interface{}
	Anonymous() IdentityPlatformProjectDefaultConfigSignInAnonymousOutputReference
	AnonymousInput() *IdentityPlatformProjectDefaultConfigSignInAnonymous
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
	Email() IdentityPlatformProjectDefaultConfigSignInEmailOutputReference
	EmailInput() *IdentityPlatformProjectDefaultConfigSignInEmail
	// Experimental.
	Fqn() *string
	HashConfig() IdentityPlatformProjectDefaultConfigSignInHashConfigList
	InternalValue() *IdentityPlatformProjectDefaultConfigSignIn
	SetInternalValue(val *IdentityPlatformProjectDefaultConfigSignIn)
	PhoneNumber() IdentityPlatformProjectDefaultConfigSignInPhoneNumberOutputReference
	PhoneNumberInput() *IdentityPlatformProjectDefaultConfigSignInPhoneNumber
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
	PutAnonymous(value *IdentityPlatformProjectDefaultConfigSignInAnonymous)
	PutEmail(value *IdentityPlatformProjectDefaultConfigSignInEmail)
	PutPhoneNumber(value *IdentityPlatformProjectDefaultConfigSignInPhoneNumber)
	ResetAllowDuplicateEmails()
	ResetAnonymous()
	ResetEmail()
	ResetPhoneNumber()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IdentityPlatformProjectDefaultConfigSignInOutputReference
type jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) AllowDuplicateEmails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDuplicateEmails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) AllowDuplicateEmailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowDuplicateEmailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) Anonymous() IdentityPlatformProjectDefaultConfigSignInAnonymousOutputReference {
	var returns IdentityPlatformProjectDefaultConfigSignInAnonymousOutputReference
	_jsii_.Get(
		j,
		"anonymous",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) AnonymousInput() *IdentityPlatformProjectDefaultConfigSignInAnonymous {
	var returns *IdentityPlatformProjectDefaultConfigSignInAnonymous
	_jsii_.Get(
		j,
		"anonymousInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) Email() IdentityPlatformProjectDefaultConfigSignInEmailOutputReference {
	var returns IdentityPlatformProjectDefaultConfigSignInEmailOutputReference
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) EmailInput() *IdentityPlatformProjectDefaultConfigSignInEmail {
	var returns *IdentityPlatformProjectDefaultConfigSignInEmail
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) HashConfig() IdentityPlatformProjectDefaultConfigSignInHashConfigList {
	var returns IdentityPlatformProjectDefaultConfigSignInHashConfigList
	_jsii_.Get(
		j,
		"hashConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) InternalValue() *IdentityPlatformProjectDefaultConfigSignIn {
	var returns *IdentityPlatformProjectDefaultConfigSignIn
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) PhoneNumber() IdentityPlatformProjectDefaultConfigSignInPhoneNumberOutputReference {
	var returns IdentityPlatformProjectDefaultConfigSignInPhoneNumberOutputReference
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) PhoneNumberInput() *IdentityPlatformProjectDefaultConfigSignInPhoneNumber {
	var returns *IdentityPlatformProjectDefaultConfigSignInPhoneNumber
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIdentityPlatformProjectDefaultConfigSignInOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IdentityPlatformProjectDefaultConfigSignInOutputReference {
	_init_.Initialize()

	if err := validateNewIdentityPlatformProjectDefaultConfigSignInOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.identityPlatformProjectDefaultConfig.IdentityPlatformProjectDefaultConfigSignInOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIdentityPlatformProjectDefaultConfigSignInOutputReference_Override(i IdentityPlatformProjectDefaultConfigSignInOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.identityPlatformProjectDefaultConfig.IdentityPlatformProjectDefaultConfigSignInOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference)SetAllowDuplicateEmails(val interface{}) {
	if err := j.validateSetAllowDuplicateEmailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowDuplicateEmails",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference)SetInternalValue(val *IdentityPlatformProjectDefaultConfigSignIn) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) PutAnonymous(value *IdentityPlatformProjectDefaultConfigSignInAnonymous) {
	if err := i.validatePutAnonymousParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAnonymous",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) PutEmail(value *IdentityPlatformProjectDefaultConfigSignInEmail) {
	if err := i.validatePutEmailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEmail",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) PutPhoneNumber(value *IdentityPlatformProjectDefaultConfigSignInPhoneNumber) {
	if err := i.validatePutPhoneNumberParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPhoneNumber",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) ResetAllowDuplicateEmails() {
	_jsii_.InvokeVoid(
		i,
		"resetAllowDuplicateEmails",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) ResetAnonymous() {
	_jsii_.InvokeVoid(
		i,
		"resetAnonymous",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		i,
		"resetEmail",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) ResetPhoneNumber() {
	_jsii_.InvokeVoid(
		i,
		"resetPhoneNumber",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformProjectDefaultConfigSignInOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

