package cloudrunv2service

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v4/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v4/cloudrunv2service/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference interface {
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
	HttpHeaders() CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetHttpHeadersList
	HttpHeadersInput() interface{}
	InternalValue() *CloudRunV2ServiceTemplateContainersLivenessProbeHttpGet
	SetInternalValue(val *CloudRunV2ServiceTemplateContainersLivenessProbeHttpGet)
	Path() *string
	SetPath(val *string)
	PathInput() *string
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
	PutHttpHeaders(value interface{})
	ResetHttpHeaders()
	ResetPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference
type jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) HttpHeaders() CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetHttpHeadersList {
	var returns CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetHttpHeadersList
	_jsii_.Get(
		j,
		"httpHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) HttpHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) InternalValue() *CloudRunV2ServiceTemplateContainersLivenessProbeHttpGet {
	var returns *CloudRunV2ServiceTemplateContainersLivenessProbeHttpGet
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference {
	_init_.Initialize()

	if err := validateNewCloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudRunV2Service.CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference_Override(c CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudRunV2Service.CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference)SetInternalValue(val *CloudRunV2ServiceTemplateContainersLivenessProbeHttpGet) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) PutHttpHeaders(value interface{}) {
	if err := c.validatePutHttpHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpHeaders",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) ResetHttpHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		c,
		"resetPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudRunV2ServiceTemplateContainersLivenessProbeHttpGetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

