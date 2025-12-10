// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package pubsubsubscription

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PubsubSubscriptionMessageTransformsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PubsubSubscriptionMessageTransformsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PubsubSubscriptionMessageTransformsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PubsubSubscriptionMessageTransformsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PubsubSubscriptionMessageTransformsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PubsubSubscriptionMessageTransformsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PubsubSubscriptionMessageTransformsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPubsubSubscriptionMessageTransformsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

