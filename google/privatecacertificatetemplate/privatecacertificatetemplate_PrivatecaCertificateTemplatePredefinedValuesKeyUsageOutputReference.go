package privatecacertificatetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v5/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v5/privatecacertificatetemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference interface {
	cdktf.ComplexObject
	BaseKeyUsage() PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference
	BaseKeyUsageInput() *PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage
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
	ExtendedKeyUsage() PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageOutputReference
	ExtendedKeyUsageInput() *PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage
	// Experimental.
	Fqn() *string
	InternalValue() *PrivatecaCertificateTemplatePredefinedValuesKeyUsage
	SetInternalValue(val *PrivatecaCertificateTemplatePredefinedValuesKeyUsage)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UnknownExtendedKeyUsages() PrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesList
	UnknownExtendedKeyUsagesInput() interface{}
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
	PutBaseKeyUsage(value *PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage)
	PutExtendedKeyUsage(value *PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage)
	PutUnknownExtendedKeyUsages(value interface{})
	ResetBaseKeyUsage()
	ResetExtendedKeyUsage()
	ResetUnknownExtendedKeyUsages()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference
type jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) BaseKeyUsage() PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference {
	var returns PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference
	_jsii_.Get(
		j,
		"baseKeyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) BaseKeyUsageInput() *PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage {
	var returns *PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage
	_jsii_.Get(
		j,
		"baseKeyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) ExtendedKeyUsage() PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageOutputReference {
	var returns PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageOutputReference
	_jsii_.Get(
		j,
		"extendedKeyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) ExtendedKeyUsageInput() *PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage {
	var returns *PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage
	_jsii_.Get(
		j,
		"extendedKeyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) InternalValue() *PrivatecaCertificateTemplatePredefinedValuesKeyUsage {
	var returns *PrivatecaCertificateTemplatePredefinedValuesKeyUsage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) UnknownExtendedKeyUsages() PrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesList {
	var returns PrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesList
	_jsii_.Get(
		j,
		"unknownExtendedKeyUsages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) UnknownExtendedKeyUsagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unknownExtendedKeyUsagesInput",
		&returns,
	)
	return returns
}


func NewPrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference {
	_init_.Initialize()

	if err := validateNewPrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCertificateTemplate.PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference_Override(p PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCertificateTemplate.PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference)SetInternalValue(val *PrivatecaCertificateTemplatePredefinedValuesKeyUsage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) PutBaseKeyUsage(value *PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage) {
	if err := p.validatePutBaseKeyUsageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putBaseKeyUsage",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) PutExtendedKeyUsage(value *PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage) {
	if err := p.validatePutExtendedKeyUsageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putExtendedKeyUsage",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) PutUnknownExtendedKeyUsages(value interface{}) {
	if err := p.validatePutUnknownExtendedKeyUsagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putUnknownExtendedKeyUsages",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) ResetBaseKeyUsage() {
	_jsii_.InvokeVoid(
		p,
		"resetBaseKeyUsage",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) ResetExtendedKeyUsage() {
	_jsii_.InvokeVoid(
		p,
		"resetExtendedKeyUsage",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) ResetUnknownExtendedKeyUsages() {
	_jsii_.InvokeVoid(
		p,
		"resetUnknownExtendedKeyUsages",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

