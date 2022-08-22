// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210701

import "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

type ComputeResource_STATUSARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Identity: The identity of the resource.
	Identity *Identity_STATUSARM `json:"identity,omitempty"`

	// Location: Specifies the location of the resource.
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Compute properties
	Properties *Compute_STATUSARM `json:"properties,omitempty"`

	// Sku: The sku of the workspace.
	Sku *Sku_STATUSARM `json:"sku,omitempty"`

	// SystemData: System data
	SystemData *SystemData_STATUSARM `json:"systemData,omitempty"`

	// Tags: Contains resource tags defined as key/value pairs.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type Compute_STATUSARM struct {
	// ComputeLocation: Location for the underlying compute
	ComputeLocation *string `json:"computeLocation,omitempty"`

	// ComputeType: The type of compute
	ComputeType *ComputeType_STATUS `json:"computeType,omitempty"`

	// CreatedOn: The time at which the compute was created.
	CreatedOn *string `json:"createdOn,omitempty"`

	// Description: The description of the Machine Learning compute.
	Description *string `json:"description,omitempty"`

	// DisableLocalAuth: Opt-out of local authentication and ensure customers can use only MSI and AAD exclusively for
	// authentication.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// IsAttachedCompute: Indicating whether the compute was provisioned by user and brought from outside if true, or machine
	// learning service provisioned it if false.
	IsAttachedCompute *bool `json:"isAttachedCompute,omitempty"`

	// ModifiedOn: The time at which the compute was last modified.
	ModifiedOn *string `json:"modifiedOn,omitempty"`

	// ProvisioningErrors: Errors during provisioning
	ProvisioningErrors []ErrorResponse_STATUSARM `json:"provisioningErrors,omitempty"`

	// ProvisioningState: The provision state of the cluster. Valid values are Unknown, Updating, Provisioning, Succeeded, and
	// Failed.
	ProvisioningState *ComputeSTATUSProvisioningState `json:"provisioningState,omitempty"`

	// ResourceId: ARM resource id of the underlying compute
	ResourceId *string `json:"resourceId,omitempty"`
}

type Identity_STATUSARM struct {
	// PrincipalId: The principal ID of resource identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The tenant ID of resource.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The identity type.
	Type *IdentitySTATUSType `json:"type,omitempty"`

	// UserAssignedIdentities: The user assigned identities associated with the resource.
	UserAssignedIdentities map[string]UserAssignedIdentity_STATUSARM `json:"userAssignedIdentities,omitempty"`
}

type Sku_STATUSARM struct {
	// Name: Name of the sku
	Name *string `json:"name,omitempty"`

	// Tier: Tier of the sku like Basic or Enterprise
	Tier *string `json:"tier,omitempty"`
}

type SystemData_STATUSARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemDataSTATUSCreatedByType `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemDataSTATUSLastModifiedByType `json:"lastModifiedByType,omitempty"`
}

type ComputeSTATUSProvisioningState string

const (
	ComputeSTATUSProvisioningState_Canceled  = ComputeSTATUSProvisioningState("Canceled")
	ComputeSTATUSProvisioningState_Creating  = ComputeSTATUSProvisioningState("Creating")
	ComputeSTATUSProvisioningState_Deleting  = ComputeSTATUSProvisioningState("Deleting")
	ComputeSTATUSProvisioningState_Failed    = ComputeSTATUSProvisioningState("Failed")
	ComputeSTATUSProvisioningState_Succeeded = ComputeSTATUSProvisioningState("Succeeded")
	ComputeSTATUSProvisioningState_Unknown   = ComputeSTATUSProvisioningState("Unknown")
	ComputeSTATUSProvisioningState_Updating  = ComputeSTATUSProvisioningState("Updating")
)

type ComputeType_STATUS string

const (
	ComputeType_STATUS_AKS               = ComputeType_STATUS("AKS")
	ComputeType_STATUS_AmlCompute        = ComputeType_STATUS("AmlCompute")
	ComputeType_STATUS_ComputeInstance   = ComputeType_STATUS("ComputeInstance")
	ComputeType_STATUS_DataFactory       = ComputeType_STATUS("DataFactory")
	ComputeType_STATUS_DataLakeAnalytics = ComputeType_STATUS("DataLakeAnalytics")
	ComputeType_STATUS_Databricks        = ComputeType_STATUS("Databricks")
	ComputeType_STATUS_HDInsight         = ComputeType_STATUS("HDInsight")
	ComputeType_STATUS_Kubernetes        = ComputeType_STATUS("Kubernetes")
	ComputeType_STATUS_SynapseSpark      = ComputeType_STATUS("SynapseSpark")
	ComputeType_STATUS_VirtualMachine    = ComputeType_STATUS("VirtualMachine")
)

type ErrorResponse_STATUSARM struct {
	// Error: The error object.
	Error *ErrorDetail_STATUSARM `json:"error,omitempty"`
}

type IdentitySTATUSType string

const (
	IdentitySTATUSType_None                       = IdentitySTATUSType("None")
	IdentitySTATUSType_SystemAssigned             = IdentitySTATUSType("SystemAssigned")
	IdentitySTATUSType_SystemAssignedUserAssigned = IdentitySTATUSType("SystemAssigned,UserAssigned")
	IdentitySTATUSType_UserAssigned               = IdentitySTATUSType("UserAssigned")
)

type SystemDataSTATUSCreatedByType string

const (
	SystemDataSTATUSCreatedByType_Application     = SystemDataSTATUSCreatedByType("Application")
	SystemDataSTATUSCreatedByType_Key             = SystemDataSTATUSCreatedByType("Key")
	SystemDataSTATUSCreatedByType_ManagedIdentity = SystemDataSTATUSCreatedByType("ManagedIdentity")
	SystemDataSTATUSCreatedByType_User            = SystemDataSTATUSCreatedByType("User")
)

type SystemDataSTATUSLastModifiedByType string

const (
	SystemDataSTATUSLastModifiedByType_Application     = SystemDataSTATUSLastModifiedByType("Application")
	SystemDataSTATUSLastModifiedByType_Key             = SystemDataSTATUSLastModifiedByType("Key")
	SystemDataSTATUSLastModifiedByType_ManagedIdentity = SystemDataSTATUSLastModifiedByType("ManagedIdentity")
	SystemDataSTATUSLastModifiedByType_User            = SystemDataSTATUSLastModifiedByType("User")
)

type UserAssignedIdentity_STATUSARM struct {
	// ClientId: The clientId(aka appId) of the user assigned identity.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: The principal ID of the user assigned identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The tenant ID of the user assigned identity.
	TenantId *string `json:"tenantId,omitempty"`
}

type ErrorDetail_STATUSARM struct {
	// AdditionalInfo: The error additional info.
	AdditionalInfo []ErrorAdditionalInfo_STATUSARM `json:"additionalInfo,omitempty"`

	// Code: The error code.
	Code *string `json:"code,omitempty"`

	// Details: The error details.
	Details []ErrorDetail_STATUS_UnrolledARM `json:"details,omitempty"`

	// Message: The error message.
	Message *string `json:"message,omitempty"`

	// Target: The error target.
	Target *string `json:"target,omitempty"`
}

type ErrorAdditionalInfo_STATUSARM struct {
	// Info: The additional info.
	Info map[string]v1.JSON `json:"info,omitempty"`

	// Type: The additional info type.
	Type *string `json:"type,omitempty"`
}

type ErrorDetail_STATUS_UnrolledARM struct {
	// AdditionalInfo: The error additional info.
	AdditionalInfo []ErrorAdditionalInfo_STATUSARM `json:"additionalInfo,omitempty"`

	// Code: The error code.
	Code *string `json:"code,omitempty"`

	// Message: The error message.
	Message *string `json:"message,omitempty"`

	// Target: The error target.
	Target *string `json:"target,omitempty"`
}