// Prebuilt google Provider for Terraform CDK (cdktf)
package google

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-google-go/google/v2/jsii"

	"github.com/hashicorp/cdktf-provider-google-go/google/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeHealthCheckGrpcHealthCheckOutputReference interface {
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
	GrpcServiceName() *string
	SetGrpcServiceName(val *string)
	GrpcServiceNameInput() *string
	InternalValue() *ComputeHealthCheckGrpcHealthCheck
	SetInternalValue(val *ComputeHealthCheckGrpcHealthCheck)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PortName() *string
	SetPortName(val *string)
	PortNameInput() *string
	PortSpecification() *string
	SetPortSpecification(val *string)
	PortSpecificationInput() *string
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
	ResetGrpcServiceName()
	ResetPort()
	ResetPortName()
	ResetPortSpecification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeHealthCheckGrpcHealthCheckOutputReference
type jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) GrpcServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grpcServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) GrpcServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grpcServiceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) InternalValue() *ComputeHealthCheckGrpcHealthCheck {
	var returns *ComputeHealthCheckGrpcHealthCheck
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) PortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) PortNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) PortSpecification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) PortSpecificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeHealthCheckGrpcHealthCheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeHealthCheckGrpcHealthCheckOutputReference {
	_init_.Initialize()

	if err := validateNewComputeHealthCheckGrpcHealthCheckOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.ComputeHealthCheckGrpcHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeHealthCheckGrpcHealthCheckOutputReference_Override(c ComputeHealthCheckGrpcHealthCheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.ComputeHealthCheckGrpcHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference)SetGrpcServiceName(val *string) {
	if err := j.validateSetGrpcServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grpcServiceName",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference)SetInternalValue(val *ComputeHealthCheckGrpcHealthCheck) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference)SetPortName(val *string) {
	if err := j.validateSetPortNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portName",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference)SetPortSpecification(val *string) {
	if err := j.validateSetPortSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portSpecification",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) ResetGrpcServiceName() {
	_jsii_.InvokeVoid(
		c,
		"resetGrpcServiceName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		c,
		"resetPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) ResetPortName() {
	_jsii_.InvokeVoid(
		c,
		"resetPortName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) ResetPortSpecification() {
	_jsii_.InvokeVoid(
		c,
		"resetPortSpecification",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheckGrpcHealthCheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

