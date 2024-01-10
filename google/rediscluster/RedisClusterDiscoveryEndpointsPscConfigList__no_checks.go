// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rediscluster

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RedisClusterDiscoveryEndpointsPscConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RedisClusterDiscoveryEndpointsPscConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RedisClusterDiscoveryEndpointsPscConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RedisClusterDiscoveryEndpointsPscConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RedisClusterDiscoveryEndpointsPscConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RedisClusterDiscoveryEndpointsPscConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRedisClusterDiscoveryEndpointsPscConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

