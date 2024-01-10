// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasemigrationserviceconnectionprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/databasemigrationserviceconnectionprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference interface {
	cdktf.ComplexObject
	CaCertificate() *string
	SetCaCertificate(val *string)
	CaCertificateInput() *string
	ClientCertificate() *string
	SetClientCertificate(val *string)
	ClientCertificateInput() *string
	ClientKey() *string
	SetClientKey(val *string)
	ClientKeyInput() *string
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
	InternalValue() *DatabaseMigrationServiceConnectionProfilePostgresqlSsl
	SetInternalValue(val *DatabaseMigrationServiceConnectionProfilePostgresqlSsl)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
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
	ResetClientCertificate()
	ResetClientKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference
type jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) CaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) CaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) ClientCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) ClientCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) ClientKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) ClientKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) InternalValue() *DatabaseMigrationServiceConnectionProfilePostgresqlSsl {
	var returns *DatabaseMigrationServiceConnectionProfilePostgresqlSsl
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewDatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.databaseMigrationServiceConnectionProfile.DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference_Override(d DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.databaseMigrationServiceConnectionProfile.DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference)SetCaCertificate(val *string) {
	if err := j.validateSetCaCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificate",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference)SetClientCertificate(val *string) {
	if err := j.validateSetClientCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificate",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference)SetClientKey(val *string) {
	if err := j.validateSetClientKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientKey",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference)SetInternalValue(val *DatabaseMigrationServiceConnectionProfilePostgresqlSsl) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) ResetClientCertificate() {
	_jsii_.InvokeVoid(
		d,
		"resetClientCertificate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) ResetClientKey() {
	_jsii_.InvokeVoid(
		d,
		"resetClientKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfilePostgresqlSslOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

