package dataprocmetastoreservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v8/dataprocmetastoreservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference interface {
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
	InternalValue() *DataprocMetastoreServiceHiveMetastoreConfigKerberosConfig
	SetInternalValue(val *DataprocMetastoreServiceHiveMetastoreConfigKerberosConfig)
	Keytab() DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytabOutputReference
	KeytabInput() *DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytab
	Krb5ConfigGcsUri() *string
	SetKrb5ConfigGcsUri(val *string)
	Krb5ConfigGcsUriInput() *string
	Principal() *string
	SetPrincipal(val *string)
	PrincipalInput() *string
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
	PutKeytab(value *DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytab)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference
type jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) InternalValue() *DataprocMetastoreServiceHiveMetastoreConfigKerberosConfig {
	var returns *DataprocMetastoreServiceHiveMetastoreConfigKerberosConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) Keytab() DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytabOutputReference {
	var returns DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytabOutputReference
	_jsii_.Get(
		j,
		"keytab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) KeytabInput() *DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytab {
	var returns *DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytab
	_jsii_.Get(
		j,
		"keytabInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) Krb5ConfigGcsUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krb5ConfigGcsUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) Krb5ConfigGcsUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krb5ConfigGcsUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) PrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocMetastoreService.DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference_Override(d DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocMetastoreService.DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference)SetInternalValue(val *DataprocMetastoreServiceHiveMetastoreConfigKerberosConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference)SetKrb5ConfigGcsUri(val *string) {
	if err := j.validateSetKrb5ConfigGcsUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"krb5ConfigGcsUri",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference)SetPrincipal(val *string) {
	if err := j.validateSetPrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) PutKeytab(value *DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytab) {
	if err := d.validatePutKeytabParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKeytab",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

