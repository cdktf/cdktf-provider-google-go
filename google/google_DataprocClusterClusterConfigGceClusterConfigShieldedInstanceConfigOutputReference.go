// Prebuilt google Provider for Terraform CDK (cdktf)
package google

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-google-go/google/v2/jsii"

	"github.com/hashicorp/cdktf-provider-google-go/google/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference interface {
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
	EnableIntegrityMonitoring() interface{}
	SetEnableIntegrityMonitoring(val interface{})
	EnableIntegrityMonitoringInput() interface{}
	EnableSecureBoot() interface{}
	SetEnableSecureBoot(val interface{})
	EnableSecureBootInput() interface{}
	EnableVtpm() interface{}
	SetEnableVtpm(val interface{})
	EnableVtpmInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfig
	SetInternalValue(val *DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfig)
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
	ResetEnableIntegrityMonitoring()
	ResetEnableSecureBoot()
	ResetEnableVtpm()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference
type jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) EnableIntegrityMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIntegrityMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) EnableIntegrityMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIntegrityMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) EnableSecureBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSecureBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) EnableSecureBootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSecureBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) EnableVtpm() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableVtpm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) EnableVtpmInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableVtpmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) InternalValue() *DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfig {
	var returns *DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference_Override(d DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference)SetEnableIntegrityMonitoring(val interface{}) {
	if err := j.validateSetEnableIntegrityMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIntegrityMonitoring",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference)SetEnableSecureBoot(val interface{}) {
	if err := j.validateSetEnableSecureBootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSecureBoot",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference)SetEnableVtpm(val interface{}) {
	if err := j.validateSetEnableVtpmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableVtpm",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference)SetInternalValue(val *DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) ResetEnableIntegrityMonitoring() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableIntegrityMonitoring",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) ResetEnableSecureBoot() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableSecureBoot",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) ResetEnableVtpm() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableVtpm",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

