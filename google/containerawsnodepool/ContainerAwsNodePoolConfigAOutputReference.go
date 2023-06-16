package containerawsnodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v8/containerawsnodepool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerAwsNodePoolConfigAOutputReference interface {
	cdktf.ComplexObject
	AutoscalingMetricsCollection() ContainerAwsNodePoolConfigAutoscalingMetricsCollectionOutputReference
	AutoscalingMetricsCollectionInput() *ContainerAwsNodePoolConfigAutoscalingMetricsCollection
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
	ConfigEncryption() ContainerAwsNodePoolConfigConfigEncryptionOutputReference
	ConfigEncryptionInput() *ContainerAwsNodePoolConfigConfigEncryption
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IamInstanceProfile() *string
	SetIamInstanceProfile(val *string)
	IamInstanceProfileInput() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() *ContainerAwsNodePoolConfigA
	SetInternalValue(val *ContainerAwsNodePoolConfigA)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	ProxyConfig() ContainerAwsNodePoolConfigProxyConfigOutputReference
	ProxyConfigInput() *ContainerAwsNodePoolConfigProxyConfig
	RootVolume() ContainerAwsNodePoolConfigRootVolumeOutputReference
	RootVolumeInput() *ContainerAwsNodePoolConfigRootVolume
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SshConfig() ContainerAwsNodePoolConfigSshConfigOutputReference
	SshConfigInput() *ContainerAwsNodePoolConfigSshConfig
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	Taints() ContainerAwsNodePoolConfigTaintsList
	TaintsInput() interface{}
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
	PutAutoscalingMetricsCollection(value *ContainerAwsNodePoolConfigAutoscalingMetricsCollection)
	PutConfigEncryption(value *ContainerAwsNodePoolConfigConfigEncryption)
	PutProxyConfig(value *ContainerAwsNodePoolConfigProxyConfig)
	PutRootVolume(value *ContainerAwsNodePoolConfigRootVolume)
	PutSshConfig(value *ContainerAwsNodePoolConfigSshConfig)
	PutTaints(value interface{})
	ResetAutoscalingMetricsCollection()
	ResetInstanceType()
	ResetLabels()
	ResetProxyConfig()
	ResetRootVolume()
	ResetSecurityGroupIds()
	ResetSshConfig()
	ResetTags()
	ResetTaints()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerAwsNodePoolConfigAOutputReference
type jsiiProxy_ContainerAwsNodePoolConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) AutoscalingMetricsCollection() ContainerAwsNodePoolConfigAutoscalingMetricsCollectionOutputReference {
	var returns ContainerAwsNodePoolConfigAutoscalingMetricsCollectionOutputReference
	_jsii_.Get(
		j,
		"autoscalingMetricsCollection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) AutoscalingMetricsCollectionInput() *ContainerAwsNodePoolConfigAutoscalingMetricsCollection {
	var returns *ContainerAwsNodePoolConfigAutoscalingMetricsCollection
	_jsii_.Get(
		j,
		"autoscalingMetricsCollectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) ConfigEncryption() ContainerAwsNodePoolConfigConfigEncryptionOutputReference {
	var returns ContainerAwsNodePoolConfigConfigEncryptionOutputReference
	_jsii_.Get(
		j,
		"configEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) ConfigEncryptionInput() *ContainerAwsNodePoolConfigConfigEncryption {
	var returns *ContainerAwsNodePoolConfigConfigEncryption
	_jsii_.Get(
		j,
		"configEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) IamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) IamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) InternalValue() *ContainerAwsNodePoolConfigA {
	var returns *ContainerAwsNodePoolConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) ProxyConfig() ContainerAwsNodePoolConfigProxyConfigOutputReference {
	var returns ContainerAwsNodePoolConfigProxyConfigOutputReference
	_jsii_.Get(
		j,
		"proxyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) ProxyConfigInput() *ContainerAwsNodePoolConfigProxyConfig {
	var returns *ContainerAwsNodePoolConfigProxyConfig
	_jsii_.Get(
		j,
		"proxyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) RootVolume() ContainerAwsNodePoolConfigRootVolumeOutputReference {
	var returns ContainerAwsNodePoolConfigRootVolumeOutputReference
	_jsii_.Get(
		j,
		"rootVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) RootVolumeInput() *ContainerAwsNodePoolConfigRootVolume {
	var returns *ContainerAwsNodePoolConfigRootVolume
	_jsii_.Get(
		j,
		"rootVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) SshConfig() ContainerAwsNodePoolConfigSshConfigOutputReference {
	var returns ContainerAwsNodePoolConfigSshConfigOutputReference
	_jsii_.Get(
		j,
		"sshConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) SshConfigInput() *ContainerAwsNodePoolConfigSshConfig {
	var returns *ContainerAwsNodePoolConfigSshConfig
	_jsii_.Get(
		j,
		"sshConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) Taints() ContainerAwsNodePoolConfigTaintsList {
	var returns ContainerAwsNodePoolConfigTaintsList
	_jsii_.Get(
		j,
		"taints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) TaintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerAwsNodePoolConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerAwsNodePoolConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewContainerAwsNodePoolConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerAwsNodePoolConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerAwsNodePool.ContainerAwsNodePoolConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerAwsNodePoolConfigAOutputReference_Override(c ContainerAwsNodePoolConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerAwsNodePool.ContainerAwsNodePoolConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference)SetIamInstanceProfile(val *string) {
	if err := j.validateSetIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference)SetInternalValue(val *ContainerAwsNodePoolConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) PutAutoscalingMetricsCollection(value *ContainerAwsNodePoolConfigAutoscalingMetricsCollection) {
	if err := c.validatePutAutoscalingMetricsCollectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAutoscalingMetricsCollection",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) PutConfigEncryption(value *ContainerAwsNodePoolConfigConfigEncryption) {
	if err := c.validatePutConfigEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putConfigEncryption",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) PutProxyConfig(value *ContainerAwsNodePoolConfigProxyConfig) {
	if err := c.validatePutProxyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putProxyConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) PutRootVolume(value *ContainerAwsNodePoolConfigRootVolume) {
	if err := c.validatePutRootVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRootVolume",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) PutSshConfig(value *ContainerAwsNodePoolConfigSshConfig) {
	if err := c.validatePutSshConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSshConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) PutTaints(value interface{}) {
	if err := c.validatePutTaintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTaints",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) ResetAutoscalingMetricsCollection() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoscalingMetricsCollection",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		c,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) ResetProxyConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetProxyConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) ResetRootVolume() {
	_jsii_.InvokeVoid(
		c,
		"resetRootVolume",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		c,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) ResetSshConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetSshConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) ResetTaints() {
	_jsii_.InvokeVoid(
		c,
		"resetTaints",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerAwsNodePoolConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

