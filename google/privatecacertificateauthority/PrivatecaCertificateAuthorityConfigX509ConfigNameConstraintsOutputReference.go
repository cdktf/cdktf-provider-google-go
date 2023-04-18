package privatecacertificateauthority

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v7/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v7/privatecacertificateauthority/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference interface {
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
	Critical() interface{}
	SetCritical(val interface{})
	CriticalInput() interface{}
	ExcludedDnsNames() *[]*string
	SetExcludedDnsNames(val *[]*string)
	ExcludedDnsNamesInput() *[]*string
	ExcludedEmailAddresses() *[]*string
	SetExcludedEmailAddresses(val *[]*string)
	ExcludedEmailAddressesInput() *[]*string
	ExcludedIpRanges() *[]*string
	SetExcludedIpRanges(val *[]*string)
	ExcludedIpRangesInput() *[]*string
	ExcludedUris() *[]*string
	SetExcludedUris(val *[]*string)
	ExcludedUrisInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *PrivatecaCertificateAuthorityConfigX509ConfigNameConstraints
	SetInternalValue(val *PrivatecaCertificateAuthorityConfigX509ConfigNameConstraints)
	PermittedDnsNames() *[]*string
	SetPermittedDnsNames(val *[]*string)
	PermittedDnsNamesInput() *[]*string
	PermittedEmailAddresses() *[]*string
	SetPermittedEmailAddresses(val *[]*string)
	PermittedEmailAddressesInput() *[]*string
	PermittedIpRanges() *[]*string
	SetPermittedIpRanges(val *[]*string)
	PermittedIpRangesInput() *[]*string
	PermittedUris() *[]*string
	SetPermittedUris(val *[]*string)
	PermittedUrisInput() *[]*string
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
	ResetExcludedDnsNames()
	ResetExcludedEmailAddresses()
	ResetExcludedIpRanges()
	ResetExcludedUris()
	ResetPermittedDnsNames()
	ResetPermittedEmailAddresses()
	ResetPermittedIpRanges()
	ResetPermittedUris()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference
type jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) Critical() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"critical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) CriticalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"criticalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ExcludedDnsNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedDnsNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ExcludedDnsNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedDnsNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ExcludedEmailAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedEmailAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ExcludedEmailAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedEmailAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ExcludedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ExcludedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ExcludedUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ExcludedUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) InternalValue() *PrivatecaCertificateAuthorityConfigX509ConfigNameConstraints {
	var returns *PrivatecaCertificateAuthorityConfigX509ConfigNameConstraints
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) PermittedDnsNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedDnsNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) PermittedDnsNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedDnsNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) PermittedEmailAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedEmailAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) PermittedEmailAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedEmailAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) PermittedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) PermittedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) PermittedUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) PermittedUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference {
	_init_.Initialize()

	if err := validateNewPrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCertificateAuthority.PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference_Override(p PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCertificateAuthority.PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference)SetCritical(val interface{}) {
	if err := j.validateSetCriticalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"critical",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference)SetExcludedDnsNames(val *[]*string) {
	if err := j.validateSetExcludedDnsNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedDnsNames",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference)SetExcludedEmailAddresses(val *[]*string) {
	if err := j.validateSetExcludedEmailAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedEmailAddresses",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference)SetExcludedIpRanges(val *[]*string) {
	if err := j.validateSetExcludedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedIpRanges",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference)SetExcludedUris(val *[]*string) {
	if err := j.validateSetExcludedUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedUris",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference)SetInternalValue(val *PrivatecaCertificateAuthorityConfigX509ConfigNameConstraints) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference)SetPermittedDnsNames(val *[]*string) {
	if err := j.validateSetPermittedDnsNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedDnsNames",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference)SetPermittedEmailAddresses(val *[]*string) {
	if err := j.validateSetPermittedEmailAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedEmailAddresses",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference)SetPermittedIpRanges(val *[]*string) {
	if err := j.validateSetPermittedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedIpRanges",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference)SetPermittedUris(val *[]*string) {
	if err := j.validateSetPermittedUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedUris",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ResetExcludedDnsNames() {
	_jsii_.InvokeVoid(
		p,
		"resetExcludedDnsNames",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ResetExcludedEmailAddresses() {
	_jsii_.InvokeVoid(
		p,
		"resetExcludedEmailAddresses",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ResetExcludedIpRanges() {
	_jsii_.InvokeVoid(
		p,
		"resetExcludedIpRanges",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ResetExcludedUris() {
	_jsii_.InvokeVoid(
		p,
		"resetExcludedUris",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ResetPermittedDnsNames() {
	_jsii_.InvokeVoid(
		p,
		"resetPermittedDnsNames",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ResetPermittedEmailAddresses() {
	_jsii_.InvokeVoid(
		p,
		"resetPermittedEmailAddresses",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ResetPermittedIpRanges() {
	_jsii_.InvokeVoid(
		p,
		"resetPermittedIpRanges",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ResetPermittedUris() {
	_jsii_.InvokeVoid(
		p,
		"resetPermittedUris",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

