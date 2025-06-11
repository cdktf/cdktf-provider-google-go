// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package pubsubtopic

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PubsubTopicMessageTransformsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PubsubTopicMessageTransformsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PubsubTopicMessageTransformsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PubsubTopicMessageTransformsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PubsubTopicMessageTransformsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PubsubTopicMessageTransformsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PubsubTopicMessageTransformsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPubsubTopicMessageTransformsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

