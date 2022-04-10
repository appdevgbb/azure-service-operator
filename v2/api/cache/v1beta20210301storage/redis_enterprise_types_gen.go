// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210301storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=cache.azure.com,resources=redisenterprises,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.azure.com,resources={redisenterprises/status,redisenterprises/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1beta20210301.RedisEnterprise
//Generated from: https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/resourceDefinitions/redisEnterprise
type RedisEnterprise struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisEnterprise_Spec `json:"spec,omitempty"`
	Status            Cluster_Status       `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisEnterprise{}

// GetConditions returns the conditions of the resource
func (enterprise *RedisEnterprise) GetConditions() conditions.Conditions {
	return enterprise.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (enterprise *RedisEnterprise) SetConditions(conditions conditions.Conditions) {
	enterprise.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &RedisEnterprise{}

// AzureName returns the Azure name of the resource
func (enterprise *RedisEnterprise) AzureName() string {
	return enterprise.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-03-01"
func (enterprise RedisEnterprise) GetAPIVersion() string {
	return "2021-03-01"
}

// GetResourceKind returns the kind of the resource
func (enterprise *RedisEnterprise) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (enterprise *RedisEnterprise) GetSpec() genruntime.ConvertibleSpec {
	return &enterprise.Spec
}

// GetStatus returns the status of this resource
func (enterprise *RedisEnterprise) GetStatus() genruntime.ConvertibleStatus {
	return &enterprise.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redisEnterprise"
func (enterprise *RedisEnterprise) GetType() string {
	return "Microsoft.Cache/redisEnterprise"
}

// NewEmptyStatus returns a new empty (blank) status
func (enterprise *RedisEnterprise) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Cluster_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (enterprise *RedisEnterprise) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(enterprise.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  enterprise.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (enterprise *RedisEnterprise) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Cluster_Status); ok {
		enterprise.Status = *st
		return nil
	}

	// Convert status to required version
	var st Cluster_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	enterprise.Status = st
	return nil
}

// Hub marks that this RedisEnterprise is the hub type for conversion
func (enterprise *RedisEnterprise) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (enterprise *RedisEnterprise) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: enterprise.Spec.OriginalVersion,
		Kind:    "RedisEnterprise",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1beta20210301.RedisEnterprise
//Generated from: https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/resourceDefinitions/redisEnterprise
type RedisEnterpriseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisEnterprise `json:"items"`
}

//Storage version of v1beta20210301.Cluster_Status
type Cluster_Status struct {
	Conditions                 []conditions.Condition                                 `json:"conditions,omitempty"`
	HostName                   *string                                                `json:"hostName,omitempty"`
	Id                         *string                                                `json:"id,omitempty"`
	Location                   *string                                                `json:"location,omitempty"`
	MinimumTlsVersion          *string                                                `json:"minimumTlsVersion,omitempty"`
	Name                       *string                                                `json:"name,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_Status_SubResourceEmbedded `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag                                 `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                                                `json:"provisioningState,omitempty"`
	RedisVersion               *string                                                `json:"redisVersion,omitempty"`
	ResourceState              *string                                                `json:"resourceState,omitempty"`
	Sku                        *Sku_Status                                            `json:"sku,omitempty"`
	Tags                       map[string]string                                      `json:"tags,omitempty"`
	Type                       *string                                                `json:"type,omitempty"`
	Zones                      []string                                               `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Cluster_Status{}

// ConvertStatusFrom populates our Cluster_Status from the provided source
func (cluster *Cluster_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == cluster {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(cluster)
}

// ConvertStatusTo populates the provided destination from our Cluster_Status
func (cluster *Cluster_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == cluster {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(cluster)
}

//Storage version of v1beta20210301.RedisEnterprise_Spec
type RedisEnterprise_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName         string  `json:"azureName,omitempty"`
	Location          *string `json:"location,omitempty"`
	MinimumTlsVersion *string `json:"minimumTlsVersion,omitempty"`
	OriginalVersion   string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	//Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	//controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	//reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Sku         *Sku                               `json:"sku,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
	Zones       []string                           `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RedisEnterprise_Spec{}

// ConvertSpecFrom populates our RedisEnterprise_Spec from the provided source
func (enterprise *RedisEnterprise_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == enterprise {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(enterprise)
}

// ConvertSpecTo populates the provided destination from our RedisEnterprise_Spec
func (enterprise *RedisEnterprise_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == enterprise {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(enterprise)
}

//Storage version of v1beta20210301.PrivateEndpointConnection_Status_SubResourceEmbedded
type PrivateEndpointConnection_Status_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1beta20210301.Sku
//Generated from: https://schema.management.azure.com/schemas/2021-03-01/Microsoft.Cache.Enterprise.json#/definitions/Sku
type Sku struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1beta20210301.Sku_Status
type Sku_Status struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&RedisEnterprise{}, &RedisEnterpriseList{})
}