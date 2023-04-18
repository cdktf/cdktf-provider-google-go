package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v6/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v6/containercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference interface {
	cdktf.ComplexObject
	AutoRepair() interface{}
	SetAutoRepair(val interface{})
	AutoRepairInput() interface{}
	AutoUpgrade() interface{}
	SetAutoUpgrade(val interface{})
	AutoUpgradeInput() interface{}
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
	InternalValue() *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagement
	SetInternalValue(val *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagement)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpgradeOptions() ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementUpgradeOptionsList
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
	ResetAutoRepair()
	ResetAutoUpgrade()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference
type jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) AutoRepair() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRepair",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) AutoRepairInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRepairInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) AutoUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) AutoUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) InternalValue() *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagement {
	var returns *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagement
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) UpgradeOptions() ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementUpgradeOptionsList {
	var returns ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementUpgradeOptionsList
	_jsii_.Get(
		j,
		"upgradeOptions",
		&returns,
	)
	return returns
}


func NewContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference_Override(c ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference)SetAutoRepair(val interface{}) {
	if err := j.validateSetAutoRepairParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRepair",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference)SetAutoUpgrade(val interface{}) {
	if err := j.validateSetAutoUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoUpgrade",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference)SetInternalValue(val *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagement) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) ResetAutoRepair() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoRepair",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) ResetAutoUpgrade() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoUpgrade",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

