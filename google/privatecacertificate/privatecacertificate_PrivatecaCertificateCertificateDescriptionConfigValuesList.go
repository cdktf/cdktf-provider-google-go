package privatecacertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v4/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v4/privatecacertificate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivatecaCertificateCertificateDescriptionConfigValuesList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) PrivatecaCertificateCertificateDescriptionConfigValuesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrivatecaCertificateCertificateDescriptionConfigValuesList
type jsiiProxy_PrivatecaCertificateCertificateDescriptionConfigValuesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_PrivatecaCertificateCertificateDescriptionConfigValuesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateCertificateDescriptionConfigValuesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateCertificateDescriptionConfigValuesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateCertificateDescriptionConfigValuesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateCertificateDescriptionConfigValuesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewPrivatecaCertificateCertificateDescriptionConfigValuesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) PrivatecaCertificateCertificateDescriptionConfigValuesList {
	_init_.Initialize()

	if err := validateNewPrivatecaCertificateCertificateDescriptionConfigValuesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivatecaCertificateCertificateDescriptionConfigValuesList{}

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCertificate.PrivatecaCertificateCertificateDescriptionConfigValuesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewPrivatecaCertificateCertificateDescriptionConfigValuesList_Override(p PrivatecaCertificateCertificateDescriptionConfigValuesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCertificate.PrivatecaCertificateCertificateDescriptionConfigValuesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		p,
	)
}

func (j *jsiiProxy_PrivatecaCertificateCertificateDescriptionConfigValuesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateCertificateDescriptionConfigValuesList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateCertificateDescriptionConfigValuesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (p *jsiiProxy_PrivatecaCertificateCertificateDescriptionConfigValuesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateCertificateDescriptionConfigValuesList) Get(index *float64) PrivatecaCertificateCertificateDescriptionConfigValuesOutputReference {
	if err := p.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns PrivatecaCertificateCertificateDescriptionConfigValuesOutputReference

	_jsii_.Invoke(
		p,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateCertificateDescriptionConfigValuesList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PrivatecaCertificateCertificateDescriptionConfigValuesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

