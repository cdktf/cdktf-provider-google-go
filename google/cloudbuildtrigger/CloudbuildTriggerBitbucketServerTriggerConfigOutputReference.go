package cloudbuildtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v6/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v6/cloudbuildtrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudbuildTriggerBitbucketServerTriggerConfigOutputReference interface {
	cdktf.ComplexObject
	BitbucketServerConfigResource() *string
	SetBitbucketServerConfigResource(val *string)
	BitbucketServerConfigResourceInput() *string
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
	InternalValue() *CloudbuildTriggerBitbucketServerTriggerConfig
	SetInternalValue(val *CloudbuildTriggerBitbucketServerTriggerConfig)
	ProjectKey() *string
	SetProjectKey(val *string)
	ProjectKeyInput() *string
	PullRequest() CloudbuildTriggerBitbucketServerTriggerConfigPullRequestOutputReference
	PullRequestInput() *CloudbuildTriggerBitbucketServerTriggerConfigPullRequest
	Push() CloudbuildTriggerBitbucketServerTriggerConfigPushOutputReference
	PushInput() *CloudbuildTriggerBitbucketServerTriggerConfigPush
	RepoSlug() *string
	SetRepoSlug(val *string)
	RepoSlugInput() *string
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
	PutPullRequest(value *CloudbuildTriggerBitbucketServerTriggerConfigPullRequest)
	PutPush(value *CloudbuildTriggerBitbucketServerTriggerConfigPush)
	ResetPullRequest()
	ResetPush()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudbuildTriggerBitbucketServerTriggerConfigOutputReference
type jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) BitbucketServerConfigResource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitbucketServerConfigResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) BitbucketServerConfigResourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitbucketServerConfigResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) InternalValue() *CloudbuildTriggerBitbucketServerTriggerConfig {
	var returns *CloudbuildTriggerBitbucketServerTriggerConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) ProjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) ProjectKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) PullRequest() CloudbuildTriggerBitbucketServerTriggerConfigPullRequestOutputReference {
	var returns CloudbuildTriggerBitbucketServerTriggerConfigPullRequestOutputReference
	_jsii_.Get(
		j,
		"pullRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) PullRequestInput() *CloudbuildTriggerBitbucketServerTriggerConfigPullRequest {
	var returns *CloudbuildTriggerBitbucketServerTriggerConfigPullRequest
	_jsii_.Get(
		j,
		"pullRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) Push() CloudbuildTriggerBitbucketServerTriggerConfigPushOutputReference {
	var returns CloudbuildTriggerBitbucketServerTriggerConfigPushOutputReference
	_jsii_.Get(
		j,
		"push",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) PushInput() *CloudbuildTriggerBitbucketServerTriggerConfigPush {
	var returns *CloudbuildTriggerBitbucketServerTriggerConfigPush
	_jsii_.Get(
		j,
		"pushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) RepoSlug() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoSlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) RepoSlugInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoSlugInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudbuildTriggerBitbucketServerTriggerConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudbuildTriggerBitbucketServerTriggerConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudbuildTriggerBitbucketServerTriggerConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTriggerBitbucketServerTriggerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudbuildTriggerBitbucketServerTriggerConfigOutputReference_Override(c CloudbuildTriggerBitbucketServerTriggerConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTriggerBitbucketServerTriggerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference)SetBitbucketServerConfigResource(val *string) {
	if err := j.validateSetBitbucketServerConfigResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitbucketServerConfigResource",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference)SetInternalValue(val *CloudbuildTriggerBitbucketServerTriggerConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference)SetProjectKey(val *string) {
	if err := j.validateSetProjectKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectKey",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference)SetRepoSlug(val *string) {
	if err := j.validateSetRepoSlugParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repoSlug",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) PutPullRequest(value *CloudbuildTriggerBitbucketServerTriggerConfigPullRequest) {
	if err := c.validatePutPullRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPullRequest",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) PutPush(value *CloudbuildTriggerBitbucketServerTriggerConfigPush) {
	if err := c.validatePutPushParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPush",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) ResetPullRequest() {
	_jsii_.InvokeVoid(
		c,
		"resetPullRequest",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) ResetPush() {
	_jsii_.InvokeVoid(
		c,
		"resetPush",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudbuildTriggerBitbucketServerTriggerConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

