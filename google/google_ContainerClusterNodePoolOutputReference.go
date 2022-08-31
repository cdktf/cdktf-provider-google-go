// Prebuilt google Provider for Terraform CDK (cdktf)
package google

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-google-go/google/v2/jsii"

	"github.com/hashicorp/cdktf-provider-google-go/google/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerClusterNodePoolOutputReference interface {
	cdktf.ComplexObject
	Autoscaling() ContainerClusterNodePoolAutoscalingOutputReference
	AutoscalingInput() *ContainerClusterNodePoolAutoscaling
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
	InitialNodeCount() *float64
	SetInitialNodeCount(val *float64)
	InitialNodeCountInput() *float64
	InstanceGroupUrls() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ManagedInstanceGroupUrls() *[]*string
	Management() ContainerClusterNodePoolManagementOutputReference
	ManagementInput() *ContainerClusterNodePoolManagement
	MaxPodsPerNode() *float64
	SetMaxPodsPerNode(val *float64)
	MaxPodsPerNodeInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	NodeConfig() ContainerClusterNodePoolNodeConfigOutputReference
	NodeConfigInput() *ContainerClusterNodePoolNodeConfig
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
	NodeLocations() *[]*string
	SetNodeLocations(val *[]*string)
	NodeLocationsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpgradeSettings() ContainerClusterNodePoolUpgradeSettingsOutputReference
	UpgradeSettingsInput() *ContainerClusterNodePoolUpgradeSettings
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
	PutAutoscaling(value *ContainerClusterNodePoolAutoscaling)
	PutManagement(value *ContainerClusterNodePoolManagement)
	PutNodeConfig(value *ContainerClusterNodePoolNodeConfig)
	PutUpgradeSettings(value *ContainerClusterNodePoolUpgradeSettings)
	ResetAutoscaling()
	ResetInitialNodeCount()
	ResetManagement()
	ResetMaxPodsPerNode()
	ResetName()
	ResetNamePrefix()
	ResetNodeConfig()
	ResetNodeCount()
	ResetNodeLocations()
	ResetUpgradeSettings()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterNodePoolOutputReference
type jsiiProxy_ContainerClusterNodePoolOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) Autoscaling() ContainerClusterNodePoolAutoscalingOutputReference {
	var returns ContainerClusterNodePoolAutoscalingOutputReference
	_jsii_.Get(
		j,
		"autoscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) AutoscalingInput() *ContainerClusterNodePoolAutoscaling {
	var returns *ContainerClusterNodePoolAutoscaling
	_jsii_.Get(
		j,
		"autoscalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) InitialNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) InitialNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) InstanceGroupUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGroupUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) ManagedInstanceGroupUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"managedInstanceGroupUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) Management() ContainerClusterNodePoolManagementOutputReference {
	var returns ContainerClusterNodePoolManagementOutputReference
	_jsii_.Get(
		j,
		"management",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) ManagementInput() *ContainerClusterNodePoolManagement {
	var returns *ContainerClusterNodePoolManagement
	_jsii_.Get(
		j,
		"managementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) MaxPodsPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPodsPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) MaxPodsPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPodsPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) NodeConfig() ContainerClusterNodePoolNodeConfigOutputReference {
	var returns ContainerClusterNodePoolNodeConfigOutputReference
	_jsii_.Get(
		j,
		"nodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) NodeConfigInput() *ContainerClusterNodePoolNodeConfig {
	var returns *ContainerClusterNodePoolNodeConfig
	_jsii_.Get(
		j,
		"nodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) NodeLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodeLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) NodeLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodeLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) UpgradeSettings() ContainerClusterNodePoolUpgradeSettingsOutputReference {
	var returns ContainerClusterNodePoolUpgradeSettingsOutputReference
	_jsii_.Get(
		j,
		"upgradeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) UpgradeSettingsInput() *ContainerClusterNodePoolUpgradeSettings {
	var returns *ContainerClusterNodePoolUpgradeSettings
	_jsii_.Get(
		j,
		"upgradeSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewContainerClusterNodePoolOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ContainerClusterNodePoolOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterNodePoolOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterNodePoolOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.ContainerClusterNodePoolOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewContainerClusterNodePoolOutputReference_Override(c ContainerClusterNodePoolOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.ContainerClusterNodePoolOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference)SetInitialNodeCount(val *float64) {
	if err := j.validateSetInitialNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialNodeCount",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference)SetMaxPodsPerNode(val *float64) {
	if err := j.validateSetMaxPodsPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPodsPerNode",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference)SetNodeLocations(val *[]*string) {
	if err := j.validateSetNodeLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeLocations",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) PutAutoscaling(value *ContainerClusterNodePoolAutoscaling) {
	if err := c.validatePutAutoscalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAutoscaling",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) PutManagement(value *ContainerClusterNodePoolManagement) {
	if err := c.validatePutManagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putManagement",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) PutNodeConfig(value *ContainerClusterNodePoolNodeConfig) {
	if err := c.validatePutNodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNodeConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) PutUpgradeSettings(value *ContainerClusterNodePoolUpgradeSettings) {
	if err := c.validatePutUpgradeSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUpgradeSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) ResetAutoscaling() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoscaling",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) ResetInitialNodeCount() {
	_jsii_.InvokeVoid(
		c,
		"resetInitialNodeCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) ResetManagement() {
	_jsii_.InvokeVoid(
		c,
		"resetManagement",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) ResetMaxPodsPerNode() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxPodsPerNode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		c,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) ResetNodeConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) ResetNodeCount() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) ResetNodeLocations() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeLocations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) ResetUpgradeSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetUpgradeSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerClusterNodePoolOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

