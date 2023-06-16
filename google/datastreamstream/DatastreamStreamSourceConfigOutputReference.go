package datastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v8/datastreamstream/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatastreamStreamSourceConfigOutputReference interface {
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
	InternalValue() *DatastreamStreamSourceConfig
	SetInternalValue(val *DatastreamStreamSourceConfig)
	MysqlSourceConfig() DatastreamStreamSourceConfigMysqlSourceConfigOutputReference
	MysqlSourceConfigInput() *DatastreamStreamSourceConfigMysqlSourceConfig
	OracleSourceConfig() DatastreamStreamSourceConfigOracleSourceConfigOutputReference
	OracleSourceConfigInput() *DatastreamStreamSourceConfigOracleSourceConfig
	PostgresqlSourceConfig() DatastreamStreamSourceConfigPostgresqlSourceConfigOutputReference
	PostgresqlSourceConfigInput() *DatastreamStreamSourceConfigPostgresqlSourceConfig
	SourceConnectionProfile() *string
	SetSourceConnectionProfile(val *string)
	SourceConnectionProfileInput() *string
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
	PutMysqlSourceConfig(value *DatastreamStreamSourceConfigMysqlSourceConfig)
	PutOracleSourceConfig(value *DatastreamStreamSourceConfigOracleSourceConfig)
	PutPostgresqlSourceConfig(value *DatastreamStreamSourceConfigPostgresqlSourceConfig)
	ResetMysqlSourceConfig()
	ResetOracleSourceConfig()
	ResetPostgresqlSourceConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatastreamStreamSourceConfigOutputReference
type jsiiProxy_DatastreamStreamSourceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference) InternalValue() *DatastreamStreamSourceConfig {
	var returns *DatastreamStreamSourceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference) MysqlSourceConfig() DatastreamStreamSourceConfigMysqlSourceConfigOutputReference {
	var returns DatastreamStreamSourceConfigMysqlSourceConfigOutputReference
	_jsii_.Get(
		j,
		"mysqlSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference) MysqlSourceConfigInput() *DatastreamStreamSourceConfigMysqlSourceConfig {
	var returns *DatastreamStreamSourceConfigMysqlSourceConfig
	_jsii_.Get(
		j,
		"mysqlSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference) OracleSourceConfig() DatastreamStreamSourceConfigOracleSourceConfigOutputReference {
	var returns DatastreamStreamSourceConfigOracleSourceConfigOutputReference
	_jsii_.Get(
		j,
		"oracleSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference) OracleSourceConfigInput() *DatastreamStreamSourceConfigOracleSourceConfig {
	var returns *DatastreamStreamSourceConfigOracleSourceConfig
	_jsii_.Get(
		j,
		"oracleSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference) PostgresqlSourceConfig() DatastreamStreamSourceConfigPostgresqlSourceConfigOutputReference {
	var returns DatastreamStreamSourceConfigPostgresqlSourceConfigOutputReference
	_jsii_.Get(
		j,
		"postgresqlSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference) PostgresqlSourceConfigInput() *DatastreamStreamSourceConfigPostgresqlSourceConfig {
	var returns *DatastreamStreamSourceConfigPostgresqlSourceConfig
	_jsii_.Get(
		j,
		"postgresqlSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference) SourceConnectionProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceConnectionProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference) SourceConnectionProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceConnectionProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatastreamStreamSourceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatastreamStreamSourceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDatastreamStreamSourceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatastreamStreamSourceConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatastreamStreamSourceConfigOutputReference_Override(d DatastreamStreamSourceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference)SetInternalValue(val *DatastreamStreamSourceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference)SetSourceConnectionProfile(val *string) {
	if err := j.validateSetSourceConnectionProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceConnectionProfile",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) PutMysqlSourceConfig(value *DatastreamStreamSourceConfigMysqlSourceConfig) {
	if err := d.validatePutMysqlSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMysqlSourceConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) PutOracleSourceConfig(value *DatastreamStreamSourceConfigOracleSourceConfig) {
	if err := d.validatePutOracleSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOracleSourceConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) PutPostgresqlSourceConfig(value *DatastreamStreamSourceConfigPostgresqlSourceConfig) {
	if err := d.validatePutPostgresqlSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPostgresqlSourceConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) ResetMysqlSourceConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetMysqlSourceConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) ResetOracleSourceConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetOracleSourceConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) ResetPostgresqlSourceConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetPostgresqlSourceConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

