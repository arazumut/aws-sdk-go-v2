// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AgentStatus string

// Enum values for AgentStatus
const (
	AgentStatusHealthy     AgentStatus = "HEALTHY"
	AgentStatusUnhealthy   AgentStatus = "UNHEALTHY"
	AgentStatusRunning     AgentStatus = "RUNNING"
	AgentStatusUnknown     AgentStatus = "UNKNOWN"
	AgentStatusBlacklisted AgentStatus = "BLACKLISTED"
	AgentStatusShutdown    AgentStatus = "SHUTDOWN"
)

// Values returns all known values for AgentStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AgentStatus) Values() []AgentStatus {
	return []AgentStatus{
		"HEALTHY",
		"UNHEALTHY",
		"RUNNING",
		"UNKNOWN",
		"BLACKLISTED",
		"SHUTDOWN",
	}
}

type BatchDeleteConfigurationTaskStatus string

// Enum values for BatchDeleteConfigurationTaskStatus
const (
	BatchDeleteConfigurationTaskStatusInitializing BatchDeleteConfigurationTaskStatus = "INITIALIZING"
	BatchDeleteConfigurationTaskStatusValidating   BatchDeleteConfigurationTaskStatus = "VALIDATING"
	BatchDeleteConfigurationTaskStatusDeleting     BatchDeleteConfigurationTaskStatus = "DELETING"
	BatchDeleteConfigurationTaskStatusCompleted    BatchDeleteConfigurationTaskStatus = "COMPLETED"
	BatchDeleteConfigurationTaskStatusFailed       BatchDeleteConfigurationTaskStatus = "FAILED"
)

// Values returns all known values for BatchDeleteConfigurationTaskStatus. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (BatchDeleteConfigurationTaskStatus) Values() []BatchDeleteConfigurationTaskStatus {
	return []BatchDeleteConfigurationTaskStatus{
		"INITIALIZING",
		"VALIDATING",
		"DELETING",
		"COMPLETED",
		"FAILED",
	}
}

type BatchDeleteImportDataErrorCode string

// Enum values for BatchDeleteImportDataErrorCode
const (
	BatchDeleteImportDataErrorCodeNotFound            BatchDeleteImportDataErrorCode = "NOT_FOUND"
	BatchDeleteImportDataErrorCodeInternalServerError BatchDeleteImportDataErrorCode = "INTERNAL_SERVER_ERROR"
	BatchDeleteImportDataErrorCodeOverLimit           BatchDeleteImportDataErrorCode = "OVER_LIMIT"
)

// Values returns all known values for BatchDeleteImportDataErrorCode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (BatchDeleteImportDataErrorCode) Values() []BatchDeleteImportDataErrorCode {
	return []BatchDeleteImportDataErrorCode{
		"NOT_FOUND",
		"INTERNAL_SERVER_ERROR",
		"OVER_LIMIT",
	}
}

type ConfigurationItemType string

// Enum values for ConfigurationItemType
const (
	ConfigurationItemTypeServer      ConfigurationItemType = "SERVER"
	ConfigurationItemTypeProcess     ConfigurationItemType = "PROCESS"
	ConfigurationItemTypeConnections ConfigurationItemType = "CONNECTION"
	ConfigurationItemTypeApplication ConfigurationItemType = "APPLICATION"
)

// Values returns all known values for ConfigurationItemType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConfigurationItemType) Values() []ConfigurationItemType {
	return []ConfigurationItemType{
		"SERVER",
		"PROCESS",
		"CONNECTION",
		"APPLICATION",
	}
}

type ContinuousExportStatus string

// Enum values for ContinuousExportStatus
const (
	ContinuousExportStatusStartInProgress ContinuousExportStatus = "START_IN_PROGRESS"
	ContinuousExportStatusStartFailed     ContinuousExportStatus = "START_FAILED"
	ContinuousExportStatusActive          ContinuousExportStatus = "ACTIVE"
	ContinuousExportStatusError           ContinuousExportStatus = "ERROR"
	ContinuousExportStatusStopInProgress  ContinuousExportStatus = "STOP_IN_PROGRESS"
	ContinuousExportStatusStopFailed      ContinuousExportStatus = "STOP_FAILED"
	ContinuousExportStatusInactive        ContinuousExportStatus = "INACTIVE"
)

// Values returns all known values for ContinuousExportStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ContinuousExportStatus) Values() []ContinuousExportStatus {
	return []ContinuousExportStatus{
		"START_IN_PROGRESS",
		"START_FAILED",
		"ACTIVE",
		"ERROR",
		"STOP_IN_PROGRESS",
		"STOP_FAILED",
		"INACTIVE",
	}
}

type DataSource string

// Enum values for DataSource
const (
	DataSourceAgent DataSource = "AGENT"
)

// Values returns all known values for DataSource. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataSource) Values() []DataSource {
	return []DataSource{
		"AGENT",
	}
}

type DeleteAgentErrorCode string

// Enum values for DeleteAgentErrorCode
const (
	DeleteAgentErrorCodeNotFound            DeleteAgentErrorCode = "NOT_FOUND"
	DeleteAgentErrorCodeInternalServerError DeleteAgentErrorCode = "INTERNAL_SERVER_ERROR"
	DeleteAgentErrorCodeAgentInUse          DeleteAgentErrorCode = "AGENT_IN_USE"
)

// Values returns all known values for DeleteAgentErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeleteAgentErrorCode) Values() []DeleteAgentErrorCode {
	return []DeleteAgentErrorCode{
		"NOT_FOUND",
		"INTERNAL_SERVER_ERROR",
		"AGENT_IN_USE",
	}
}

type DeletionConfigurationItemType string

// Enum values for DeletionConfigurationItemType
const (
	DeletionConfigurationItemTypeServer DeletionConfigurationItemType = "SERVER"
)

// Values returns all known values for DeletionConfigurationItemType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeletionConfigurationItemType) Values() []DeletionConfigurationItemType {
	return []DeletionConfigurationItemType{
		"SERVER",
	}
}

type ExportDataFormat string

// Enum values for ExportDataFormat
const (
	ExportDataFormatCsv ExportDataFormat = "CSV"
)

// Values returns all known values for ExportDataFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExportDataFormat) Values() []ExportDataFormat {
	return []ExportDataFormat{
		"CSV",
	}
}

type ExportStatus string

// Enum values for ExportStatus
const (
	ExportStatusFailed     ExportStatus = "FAILED"
	ExportStatusSucceeded  ExportStatus = "SUCCEEDED"
	ExportStatusInProgress ExportStatus = "IN_PROGRESS"
)

// Values returns all known values for ExportStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExportStatus) Values() []ExportStatus {
	return []ExportStatus{
		"FAILED",
		"SUCCEEDED",
		"IN_PROGRESS",
	}
}

type FileClassification string

// Enum values for FileClassification
const (
	FileClassificationModelizeitExport FileClassification = "MODELIZEIT_EXPORT"
	FileClassificationRvtoolsExport    FileClassification = "RVTOOLS_EXPORT"
	FileClassificationVmwareNsxExport  FileClassification = "VMWARE_NSX_EXPORT"
	FileClassificationImportTemplate   FileClassification = "IMPORT_TEMPLATE"
)

// Values returns all known values for FileClassification. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FileClassification) Values() []FileClassification {
	return []FileClassification{
		"MODELIZEIT_EXPORT",
		"RVTOOLS_EXPORT",
		"VMWARE_NSX_EXPORT",
		"IMPORT_TEMPLATE",
	}
}

type ImportStatus string

// Enum values for ImportStatus
const (
	ImportStatusImportInProgress                ImportStatus = "IMPORT_IN_PROGRESS"
	ImportStatusImportComplete                  ImportStatus = "IMPORT_COMPLETE"
	ImportStatusImportCompleteWithErrors        ImportStatus = "IMPORT_COMPLETE_WITH_ERRORS"
	ImportStatusImportFailed                    ImportStatus = "IMPORT_FAILED"
	ImportStatusImportFailedServerLimitExceeded ImportStatus = "IMPORT_FAILED_SERVER_LIMIT_EXCEEDED"
	ImportStatusImportFailedRecordLimitExceeded ImportStatus = "IMPORT_FAILED_RECORD_LIMIT_EXCEEDED"
	ImportStatusImportFailedUnsupportedFileType ImportStatus = "IMPORT_FAILED_UNSUPPORTED_FILE_TYPE"
	ImportStatusDeleteInProgress                ImportStatus = "DELETE_IN_PROGRESS"
	ImportStatusDeleteComplete                  ImportStatus = "DELETE_COMPLETE"
	ImportStatusDeleteFailed                    ImportStatus = "DELETE_FAILED"
	ImportStatusDeleteFailedLimitExceeded       ImportStatus = "DELETE_FAILED_LIMIT_EXCEEDED"
	ImportStatusInternalError                   ImportStatus = "INTERNAL_ERROR"
)

// Values returns all known values for ImportStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ImportStatus) Values() []ImportStatus {
	return []ImportStatus{
		"IMPORT_IN_PROGRESS",
		"IMPORT_COMPLETE",
		"IMPORT_COMPLETE_WITH_ERRORS",
		"IMPORT_FAILED",
		"IMPORT_FAILED_SERVER_LIMIT_EXCEEDED",
		"IMPORT_FAILED_RECORD_LIMIT_EXCEEDED",
		"IMPORT_FAILED_UNSUPPORTED_FILE_TYPE",
		"DELETE_IN_PROGRESS",
		"DELETE_COMPLETE",
		"DELETE_FAILED",
		"DELETE_FAILED_LIMIT_EXCEEDED",
		"INTERNAL_ERROR",
	}
}

type ImportTaskFilterName string

// Enum values for ImportTaskFilterName
const (
	ImportTaskFilterNameImportTaskId       ImportTaskFilterName = "IMPORT_TASK_ID"
	ImportTaskFilterNameStatus             ImportTaskFilterName = "STATUS"
	ImportTaskFilterNameName               ImportTaskFilterName = "NAME"
	ImportTaskFilterNameFileClassification ImportTaskFilterName = "FILE_CLASSIFICATION"
)

// Values returns all known values for ImportTaskFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ImportTaskFilterName) Values() []ImportTaskFilterName {
	return []ImportTaskFilterName{
		"IMPORT_TASK_ID",
		"STATUS",
		"NAME",
		"FILE_CLASSIFICATION",
	}
}

type OfferingClass string

// Enum values for OfferingClass
const (
	OfferingClassStandard    OfferingClass = "STANDARD"
	OfferingClassConvertible OfferingClass = "CONVERTIBLE"
)

// Values returns all known values for OfferingClass. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OfferingClass) Values() []OfferingClass {
	return []OfferingClass{
		"STANDARD",
		"CONVERTIBLE",
	}
}

type OrderString string

// Enum values for OrderString
const (
	OrderStringAsc  OrderString = "ASC"
	OrderStringDesc OrderString = "DESC"
)

// Values returns all known values for OrderString. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OrderString) Values() []OrderString {
	return []OrderString{
		"ASC",
		"DESC",
	}
}

type PurchasingOption string

// Enum values for PurchasingOption
const (
	PurchasingOptionAllUpfront     PurchasingOption = "ALL_UPFRONT"
	PurchasingOptionPartialUpfront PurchasingOption = "PARTIAL_UPFRONT"
	PurchasingOptionNoUpfront      PurchasingOption = "NO_UPFRONT"
)

// Values returns all known values for PurchasingOption. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PurchasingOption) Values() []PurchasingOption {
	return []PurchasingOption{
		"ALL_UPFRONT",
		"PARTIAL_UPFRONT",
		"NO_UPFRONT",
	}
}

type Tenancy string

// Enum values for Tenancy
const (
	TenancyDedicated Tenancy = "DEDICATED"
	TenancyShared    Tenancy = "SHARED"
)

// Values returns all known values for Tenancy. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Tenancy) Values() []Tenancy {
	return []Tenancy{
		"DEDICATED",
		"SHARED",
	}
}

type TermLength string

// Enum values for TermLength
const (
	TermLengthOneYear   TermLength = "ONE_YEAR"
	TermLengthThreeYear TermLength = "THREE_YEAR"
)

// Values returns all known values for TermLength. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TermLength) Values() []TermLength {
	return []TermLength{
		"ONE_YEAR",
		"THREE_YEAR",
	}
}
