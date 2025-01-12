// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DatabaseAccounts_SqlDatabases_Container_Spec_ARM struct {
	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties to create and update Azure Cosmos DB container.
	Properties *SqlContainerCreateUpdateProperties_ARM `json:"properties,omitempty"`
	Tags       map[string]string                       `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccounts_SqlDatabases_Container_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (container DatabaseAccounts_SqlDatabases_Container_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (container *DatabaseAccounts_SqlDatabases_Container_Spec_ARM) GetName() string {
	return container.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers"
func (container *DatabaseAccounts_SqlDatabases_Container_Spec_ARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers"
}

// Properties to create and update Azure Cosmos DB container.
type SqlContainerCreateUpdateProperties_ARM struct {
	// Options: A key-value pair of options to be applied for the request. This corresponds to the headers sent with the
	// request.
	Options *CreateUpdateOptions_ARM `json:"options,omitempty"`

	// Resource: The standard JSON format of a container
	Resource *SqlContainerResource_ARM `json:"resource,omitempty"`
}

// Cosmos DB SQL container resource object
type SqlContainerResource_ARM struct {
	// AnalyticalStorageTtl: Analytical TTL.
	AnalyticalStorageTtl *int `json:"analyticalStorageTtl,omitempty"`

	// ConflictResolutionPolicy: The conflict resolution policy for the container.
	ConflictResolutionPolicy *ConflictResolutionPolicy_ARM `json:"conflictResolutionPolicy,omitempty"`

	// DefaultTtl: Default time to live
	DefaultTtl *int `json:"defaultTtl,omitempty"`

	// Id: Name of the Cosmos DB SQL container
	Id *string `json:"id,omitempty"`

	// IndexingPolicy: The configuration of the indexing policy. By default, the indexing is automatic for all document paths
	// within the container
	IndexingPolicy *IndexingPolicy_ARM `json:"indexingPolicy,omitempty"`

	// PartitionKey: The configuration of the partition key to be used for partitioning data into multiple partitions
	PartitionKey *ContainerPartitionKey_ARM `json:"partitionKey,omitempty"`

	// UniqueKeyPolicy: The unique key policy configuration for specifying uniqueness constraints on documents in the
	// collection in the Azure Cosmos DB service.
	UniqueKeyPolicy *UniqueKeyPolicy_ARM `json:"uniqueKeyPolicy,omitempty"`
}

// The conflict resolution policy for the container.
type ConflictResolutionPolicy_ARM struct {
	// ConflictResolutionPath: The conflict resolution path in the case of LastWriterWins mode.
	ConflictResolutionPath *string `json:"conflictResolutionPath,omitempty"`

	// ConflictResolutionProcedure: The procedure to resolve conflicts in the case of custom mode.
	ConflictResolutionProcedure *string `json:"conflictResolutionProcedure,omitempty"`

	// Mode: Indicates the conflict resolution mode.
	Mode *ConflictResolutionPolicy_Mode `json:"mode,omitempty"`
}

// The configuration of the partition key to be used for partitioning data into multiple partitions
type ContainerPartitionKey_ARM struct {
	// Kind: Indicates the kind of algorithm used for partitioning. For MultiHash, multiple partition keys (upto three maximum)
	// are supported for container create
	Kind *ContainerPartitionKey_Kind `json:"kind,omitempty"`

	// Paths: List of paths using which data within the container can be partitioned
	Paths []string `json:"paths,omitempty"`

	// Version: Indicates the version of the partition key definition
	Version *int `json:"version,omitempty"`
}

// Cosmos DB indexing policy
type IndexingPolicy_ARM struct {
	// Automatic: Indicates if the indexing policy is automatic
	Automatic *bool `json:"automatic,omitempty"`

	// CompositeIndexes: List of composite path list
	CompositeIndexes [][]CompositePath_ARM `json:"compositeIndexes,omitempty"`

	// ExcludedPaths: List of paths to exclude from indexing
	ExcludedPaths []ExcludedPath_ARM `json:"excludedPaths,omitempty"`

	// IncludedPaths: List of paths to include in the indexing
	IncludedPaths []IncludedPath_ARM `json:"includedPaths,omitempty"`

	// IndexingMode: Indicates the indexing mode.
	IndexingMode *IndexingPolicy_IndexingMode `json:"indexingMode,omitempty"`

	// SpatialIndexes: List of spatial specifics
	SpatialIndexes []SpatialSpec_ARM `json:"spatialIndexes,omitempty"`
}

// The unique key policy configuration for specifying uniqueness constraints on documents in the collection in the Azure
// Cosmos DB service.
type UniqueKeyPolicy_ARM struct {
	// UniqueKeys: List of unique keys on that enforces uniqueness constraint on documents in the collection in the Azure
	// Cosmos DB service.
	UniqueKeys []UniqueKey_ARM `json:"uniqueKeys,omitempty"`
}

type CompositePath_ARM struct {
	// Order: Sort order for composite paths.
	Order *CompositePath_Order `json:"order,omitempty"`

	// Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
	// (/path/*)
	Path *string `json:"path,omitempty"`
}

// +kubebuilder:validation:Enum={"Custom","LastWriterWins"}
type ConflictResolutionPolicy_Mode string

const (
	ConflictResolutionPolicy_Mode_Custom         = ConflictResolutionPolicy_Mode("Custom")
	ConflictResolutionPolicy_Mode_LastWriterWins = ConflictResolutionPolicy_Mode("LastWriterWins")
)

// +kubebuilder:validation:Enum={"Hash","MultiHash","Range"}
type ContainerPartitionKey_Kind string

const (
	ContainerPartitionKey_Kind_Hash      = ContainerPartitionKey_Kind("Hash")
	ContainerPartitionKey_Kind_MultiHash = ContainerPartitionKey_Kind("MultiHash")
	ContainerPartitionKey_Kind_Range     = ContainerPartitionKey_Kind("Range")
)

type ExcludedPath_ARM struct {
	// Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
	// (/path/*)
	Path *string `json:"path,omitempty"`
}

// The paths that are included in indexing
type IncludedPath_ARM struct {
	// Indexes: List of indexes for this path
	Indexes []Indexes_ARM `json:"indexes,omitempty"`

	// Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
	// (/path/*)
	Path *string `json:"path,omitempty"`
}

// +kubebuilder:validation:Enum={"consistent","lazy","none"}
type IndexingPolicy_IndexingMode string

const (
	IndexingPolicy_IndexingMode_Consistent = IndexingPolicy_IndexingMode("consistent")
	IndexingPolicy_IndexingMode_Lazy       = IndexingPolicy_IndexingMode("lazy")
	IndexingPolicy_IndexingMode_None       = IndexingPolicy_IndexingMode("none")
)

type SpatialSpec_ARM struct {
	// Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
	// (/path/*)
	Path *string `json:"path,omitempty"`

	// Types: List of path's spatial type
	Types []SpatialType `json:"types,omitempty"`
}

// The unique key on that enforces uniqueness constraint on documents in the collection in the Azure Cosmos DB service.
type UniqueKey_ARM struct {
	// Paths: List of paths must be unique for each document in the Azure Cosmos DB service
	Paths []string `json:"paths,omitempty"`
}

// +kubebuilder:validation:Enum={"ascending","descending"}
type CompositePath_Order string

const (
	CompositePath_Order_Ascending  = CompositePath_Order("ascending")
	CompositePath_Order_Descending = CompositePath_Order("descending")
)

// The indexes for the path.
type Indexes_ARM struct {
	// DataType: The datatype for which the indexing behavior is applied to.
	DataType *Indexes_DataType `json:"dataType,omitempty"`

	// Kind: Indicates the type of index.
	Kind *Indexes_Kind `json:"kind,omitempty"`

	// Precision: The precision of the index. -1 is maximum precision.
	Precision *int `json:"precision,omitempty"`
}

// Indicates the spatial type of index.
// +kubebuilder:validation:Enum={"LineString","MultiPolygon","Point","Polygon"}
type SpatialType string

const (
	SpatialType_LineString   = SpatialType("LineString")
	SpatialType_MultiPolygon = SpatialType("MultiPolygon")
	SpatialType_Point        = SpatialType("Point")
	SpatialType_Polygon      = SpatialType("Polygon")
)

// +kubebuilder:validation:Enum={"LineString","MultiPolygon","Number","Point","Polygon","String"}
type Indexes_DataType string

const (
	Indexes_DataType_LineString   = Indexes_DataType("LineString")
	Indexes_DataType_MultiPolygon = Indexes_DataType("MultiPolygon")
	Indexes_DataType_Number       = Indexes_DataType("Number")
	Indexes_DataType_Point        = Indexes_DataType("Point")
	Indexes_DataType_Polygon      = Indexes_DataType("Polygon")
	Indexes_DataType_String       = Indexes_DataType("String")
)

// +kubebuilder:validation:Enum={"Hash","Range","Spatial"}
type Indexes_Kind string

const (
	Indexes_Kind_Hash    = Indexes_Kind("Hash")
	Indexes_Kind_Range   = Indexes_Kind("Range")
	Indexes_Kind_Spatial = Indexes_Kind("Spatial")
)
