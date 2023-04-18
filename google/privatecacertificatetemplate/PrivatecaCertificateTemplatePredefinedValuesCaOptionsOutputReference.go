package privatecacertificatetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v7/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v7/privatecacertificatetemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference interface {
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
	InternalValue() *PrivatecaCertificateTemplatePredefinedValuesCaOptions
	SetInternalValue(val *PrivatecaCertificateTemplatePredefinedValuesCaOptions)
	IsCa() interface{}
	SetIsCa(val interface{})
	IsCaInput() interface{}
	MaxIssuerPathLength() *float64
	SetMaxIssuerPathLength(val *float64)
	MaxIssuerPathLengthInput() *float64
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
	ResetIsCa()
	ResetMaxIssuerPathLength()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference
type jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) InternalValue() *PrivatecaCertificateTemplatePredefinedValuesCaOptions {
	var returns *PrivatecaCertificateTemplatePredefinedValuesCaOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) IsCa() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) IsCaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) MaxIssuerPathLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIssuerPathLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) MaxIssuerPathLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIssuerPathLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewPrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCertificateTemplate.PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference_Override(p PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCertificateTemplate.PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference)SetInternalValue(val *PrivatecaCertificateTemplatePredefinedValuesCaOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference)SetIsCa(val interface{}) {
	if err := j.validateSetIsCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCa",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference)SetMaxIssuerPathLength(val *float64) {
	if err := j.validateSetMaxIssuerPathLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIssuerPathLength",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) ResetIsCa() {
	_jsii_.InvokeVoid(
		p,
		"resetIsCa",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) ResetMaxIssuerPathLength() {
	_jsii_.InvokeVoid(
		p,
		"resetMaxIssuerPathLength",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

