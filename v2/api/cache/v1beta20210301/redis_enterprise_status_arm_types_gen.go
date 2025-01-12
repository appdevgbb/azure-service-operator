// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210301

type RedisEnterprise_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Other properties of the cluster.
	Properties *ClusterProperties_STATUS_ARM `json:"properties,omitempty"`

	// Sku: The SKU to create, which affects price, performance, and features.
	Sku *Sku_STATUS_ARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`

	// Zones: The Availability Zones where this cluster will be deployed.
	Zones []string `json:"zones,omitempty"`
}

// Properties of RedisEnterprise clusters, as opposed to general resource properties like location, tags
type ClusterProperties_STATUS_ARM struct {
	// HostName: DNS name of the cluster endpoint
	HostName *string `json:"hostName,omitempty"`

	// MinimumTlsVersion: The minimum TLS version for the cluster to support, e.g. '1.2'
	MinimumTlsVersion *ClusterProperties_MinimumTlsVersion_STATUS `json:"minimumTlsVersion,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connections associated with the specified RedisEnterprise cluster
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_ARM `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: Current provisioning status of the cluster
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// RedisVersion: Version of redis the cluster supports, e.g. '6'
	RedisVersion *string `json:"redisVersion,omitempty"`

	// ResourceState: Current resource status of the cluster
	ResourceState *ResourceState_STATUS `json:"resourceState,omitempty"`
}

// SKU parameters supplied to the create RedisEnterprise operation.
type Sku_STATUS_ARM struct {
	// Capacity: The size of the RedisEnterprise cluster. Defaults to 2 or 3 depending on SKU. Valid values are (2, 4, 6, ...)
	// for Enterprise SKUs and (3, 9, 15, ...) for Flash SKUs.
	Capacity *int `json:"capacity,omitempty"`

	// Name: The type of RedisEnterprise cluster to deploy. Possible values: (Enterprise_E10, EnterpriseFlash_F300 etc.)
	Name *Sku_Name_STATUS `json:"name,omitempty"`
}

// The Private Endpoint Connection resource.
type PrivateEndpointConnection_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`
}

type Sku_Name_STATUS string

const (
	Sku_Name_STATUS_EnterpriseFlash_F1500 = Sku_Name_STATUS("EnterpriseFlash_F1500")
	Sku_Name_STATUS_EnterpriseFlash_F300  = Sku_Name_STATUS("EnterpriseFlash_F300")
	Sku_Name_STATUS_EnterpriseFlash_F700  = Sku_Name_STATUS("EnterpriseFlash_F700")
	Sku_Name_STATUS_Enterprise_E10        = Sku_Name_STATUS("Enterprise_E10")
	Sku_Name_STATUS_Enterprise_E100       = Sku_Name_STATUS("Enterprise_E100")
	Sku_Name_STATUS_Enterprise_E20        = Sku_Name_STATUS("Enterprise_E20")
	Sku_Name_STATUS_Enterprise_E50        = Sku_Name_STATUS("Enterprise_E50")
)
