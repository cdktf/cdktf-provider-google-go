package dataprocmetastoreservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v5/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v5/dataprocmetastoreservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocMetastoreServiceHiveMetastoreConfigOutputReference interface {
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
	ConfigOverrides() *map[string]*string
	SetConfigOverrides(val *map[string]*string)
	ConfigOverridesInput() *map[string]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataprocMetastoreServiceHiveMetastoreConfig
	SetInternalValue(val *DataprocMetastoreServiceHiveMetastoreConfig)
	KerberosConfig() DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference
	KerberosConfigInput() *DataprocMetastoreServiceHiveMetastoreConfigKerberosConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutKerberosConfig(value *DataprocMetastoreServiceHiveMetastoreConfigKerberosConfig)
	ResetConfigOverrides()
	ResetKerberosConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocMetastoreServiceHiveMetastoreConfigOutputReference
type jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) ConfigOverrides() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) ConfigOverridesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) InternalValue() *DataprocMetastoreServiceHiveMetastoreConfig {
	var returns *DataprocMetastoreServiceHiveMetastoreConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) KerberosConfig() DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference {
	var returns DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigOutputReference
	_jsii_.Get(
		j,
		"kerberosConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) KerberosConfigInput() *DataprocMetastoreServiceHiveMetastoreConfigKerberosConfig {
	var returns *DataprocMetastoreServiceHiveMetastoreConfigKerberosConfig
	_jsii_.Get(
		j,
		"kerberosConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewDataprocMetastoreServiceHiveMetastoreConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocMetastoreServiceHiveMetastoreConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocMetastoreServiceHiveMetastoreConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocMetastoreService.DataprocMetastoreServiceHiveMetastoreConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocMetastoreServiceHiveMetastoreConfigOutputReference_Override(d DataprocMetastoreServiceHiveMetastoreConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocMetastoreService.DataprocMetastoreServiceHiveMetastoreConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference)SetConfigOverrides(val *map[string]*string) {
	if err := j.validateSetConfigOverridesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configOverrides",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference)SetInternalValue(val *DataprocMetastoreServiceHiveMetastoreConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) PutKerberosConfig(value *DataprocMetastoreServiceHiveMetastoreConfigKerberosConfig) {
	if err := d.validatePutKerberosConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKerberosConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) ResetConfigOverrides() {
	_jsii_.InvokeVoid(
		d,
		"resetConfigOverrides",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) ResetKerberosConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetKerberosConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataprocMetastoreServiceHiveMetastoreConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

