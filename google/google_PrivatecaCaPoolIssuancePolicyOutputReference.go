// Prebuilt google Provider for Terraform CDK (cdktf)
package google

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-google-go/google/v2/jsii"

	"github.com/hashicorp/cdktf-provider-google-go/google/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivatecaCaPoolIssuancePolicyOutputReference interface {
	cdktf.ComplexObject
	AllowedIssuanceModes() PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference
	AllowedIssuanceModesInput() *PrivatecaCaPoolIssuancePolicyAllowedIssuanceModes
	AllowedKeyTypes() PrivatecaCaPoolIssuancePolicyAllowedKeyTypesList
	AllowedKeyTypesInput() interface{}
	BaselineValues() PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference
	BaselineValuesInput() *PrivatecaCaPoolIssuancePolicyBaselineValues
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
	IdentityConstraints() PrivatecaCaPoolIssuancePolicyIdentityConstraintsOutputReference
	IdentityConstraintsInput() *PrivatecaCaPoolIssuancePolicyIdentityConstraints
	InternalValue() *PrivatecaCaPoolIssuancePolicy
	SetInternalValue(val *PrivatecaCaPoolIssuancePolicy)
	MaximumLifetime() *string
	SetMaximumLifetime(val *string)
	MaximumLifetimeInput() *string
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
	PutAllowedIssuanceModes(value *PrivatecaCaPoolIssuancePolicyAllowedIssuanceModes)
	PutAllowedKeyTypes(value interface{})
	PutBaselineValues(value *PrivatecaCaPoolIssuancePolicyBaselineValues)
	PutIdentityConstraints(value *PrivatecaCaPoolIssuancePolicyIdentityConstraints)
	ResetAllowedIssuanceModes()
	ResetAllowedKeyTypes()
	ResetBaselineValues()
	ResetIdentityConstraints()
	ResetMaximumLifetime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrivatecaCaPoolIssuancePolicyOutputReference
type jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) AllowedIssuanceModes() PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference {
	var returns PrivatecaCaPoolIssuancePolicyAllowedIssuanceModesOutputReference
	_jsii_.Get(
		j,
		"allowedIssuanceModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) AllowedIssuanceModesInput() *PrivatecaCaPoolIssuancePolicyAllowedIssuanceModes {
	var returns *PrivatecaCaPoolIssuancePolicyAllowedIssuanceModes
	_jsii_.Get(
		j,
		"allowedIssuanceModesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) AllowedKeyTypes() PrivatecaCaPoolIssuancePolicyAllowedKeyTypesList {
	var returns PrivatecaCaPoolIssuancePolicyAllowedKeyTypesList
	_jsii_.Get(
		j,
		"allowedKeyTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) AllowedKeyTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedKeyTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) BaselineValues() PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference {
	var returns PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference
	_jsii_.Get(
		j,
		"baselineValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) BaselineValuesInput() *PrivatecaCaPoolIssuancePolicyBaselineValues {
	var returns *PrivatecaCaPoolIssuancePolicyBaselineValues
	_jsii_.Get(
		j,
		"baselineValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) IdentityConstraints() PrivatecaCaPoolIssuancePolicyIdentityConstraintsOutputReference {
	var returns PrivatecaCaPoolIssuancePolicyIdentityConstraintsOutputReference
	_jsii_.Get(
		j,
		"identityConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) IdentityConstraintsInput() *PrivatecaCaPoolIssuancePolicyIdentityConstraints {
	var returns *PrivatecaCaPoolIssuancePolicyIdentityConstraints
	_jsii_.Get(
		j,
		"identityConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) InternalValue() *PrivatecaCaPoolIssuancePolicy {
	var returns *PrivatecaCaPoolIssuancePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) MaximumLifetime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) MaximumLifetimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPrivatecaCaPoolIssuancePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PrivatecaCaPoolIssuancePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewPrivatecaCaPoolIssuancePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.PrivatecaCaPoolIssuancePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPrivatecaCaPoolIssuancePolicyOutputReference_Override(p PrivatecaCaPoolIssuancePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.PrivatecaCaPoolIssuancePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference)SetInternalValue(val *PrivatecaCaPoolIssuancePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference)SetMaximumLifetime(val *string) {
	if err := j.validateSetMaximumLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumLifetime",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) PutAllowedIssuanceModes(value *PrivatecaCaPoolIssuancePolicyAllowedIssuanceModes) {
	if err := p.validatePutAllowedIssuanceModesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAllowedIssuanceModes",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) PutAllowedKeyTypes(value interface{}) {
	if err := p.validatePutAllowedKeyTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAllowedKeyTypes",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) PutBaselineValues(value *PrivatecaCaPoolIssuancePolicyBaselineValues) {
	if err := p.validatePutBaselineValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putBaselineValues",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) PutIdentityConstraints(value *PrivatecaCaPoolIssuancePolicyIdentityConstraints) {
	if err := p.validatePutIdentityConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putIdentityConstraints",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) ResetAllowedIssuanceModes() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedIssuanceModes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) ResetAllowedKeyTypes() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedKeyTypes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) ResetBaselineValues() {
	_jsii_.InvokeVoid(
		p,
		"resetBaselineValues",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) ResetIdentityConstraints() {
	_jsii_.InvokeVoid(
		p,
		"resetIdentityConstraints",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) ResetMaximumLifetime() {
	_jsii_.InvokeVoid(
		p,
		"resetMaximumLifetime",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

