package datagoogleprivatecacertificateauthority

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v5/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v5/datagoogleprivatecacertificateauthority/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference interface {
	cdktf.ComplexObject
	AdditionalExtensions() DataGooglePrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsList
	AiaOcspServers() *[]*string
	CaOptions() DataGooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsList
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
	InternalValue() *DataGooglePrivatecaCertificateAuthorityConfigX509Config
	SetInternalValue(val *DataGooglePrivatecaCertificateAuthorityConfigX509Config)
	KeyUsage() DataGooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageList
	PolicyIds() DataGooglePrivatecaCertificateAuthorityConfigX509ConfigPolicyIdsList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference
type jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) AdditionalExtensions() DataGooglePrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsList {
	var returns DataGooglePrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsList
	_jsii_.Get(
		j,
		"additionalExtensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) AiaOcspServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aiaOcspServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) CaOptions() DataGooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsList {
	var returns DataGooglePrivatecaCertificateAuthorityConfigX509ConfigCaOptionsList
	_jsii_.Get(
		j,
		"caOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) InternalValue() *DataGooglePrivatecaCertificateAuthorityConfigX509Config {
	var returns *DataGooglePrivatecaCertificateAuthorityConfigX509Config
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) KeyUsage() DataGooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageList {
	var returns DataGooglePrivatecaCertificateAuthorityConfigX509ConfigKeyUsageList
	_jsii_.Get(
		j,
		"keyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) PolicyIds() DataGooglePrivatecaCertificateAuthorityConfigX509ConfigPolicyIdsList {
	var returns DataGooglePrivatecaCertificateAuthorityConfigX509ConfigPolicyIdsList
	_jsii_.Get(
		j,
		"policyIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGooglePrivatecaCertificateAuthority.DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference_Override(d DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGooglePrivatecaCertificateAuthority.DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference)SetInternalValue(val *DataGooglePrivatecaCertificateAuthorityConfigX509Config) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGooglePrivatecaCertificateAuthorityConfigX509ConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

