// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec_ARM struct {
	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties to create and update Azure Cosmos DB storedProcedure.
	Properties *SqlStoredProcedureCreateUpdateProperties_ARM `json:"properties,omitempty"`
	Tags       map[string]string                             `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (procedure DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec_ARM) GetName() string {
	return procedure.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/storedProcedures"
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec_ARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/storedProcedures"
}

// Properties to create and update Azure Cosmos DB storedProcedure.
type SqlStoredProcedureCreateUpdateProperties_ARM struct {
	// Options: A key-value pair of options to be applied for the request. This corresponds to the headers sent with the
	// request.
	Options *CreateUpdateOptions_ARM `json:"options,omitempty"`

	// Resource: The standard JSON format of a storedProcedure
	Resource *SqlStoredProcedureResource_ARM `json:"resource,omitempty"`
}

// Cosmos DB SQL storedProcedure resource object
type SqlStoredProcedureResource_ARM struct {
	// Body: Body of the Stored Procedure
	Body *string `json:"body,omitempty"`

	// Id: Name of the Cosmos DB SQL storedProcedure
	Id *string `json:"id,omitempty"`
}
