# RunExportReportTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportParameters** | Pointer to **map[string]string** |  | [optional] 
**Format** | Pointer to [**ExportFormat**](ExportFormat.md) |  | [optional] 
**PagesCount** | Pointer to **int32** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewRunExportReportTaskVM

`func NewRunExportReportTaskVM(t string, ) *RunExportReportTaskVM`

NewRunExportReportTaskVM instantiates a new RunExportReportTaskVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunExportReportTaskVMWithDefaults

`func NewRunExportReportTaskVMWithDefaults() *RunExportReportTaskVM`

NewRunExportReportTaskVMWithDefaults instantiates a new RunExportReportTaskVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportParameters

`func (o *RunExportReportTaskVM) GetExportParameters() map[string]string`

GetExportParameters returns the ExportParameters field if non-nil, zero value otherwise.

### GetExportParametersOk

`func (o *RunExportReportTaskVM) GetExportParametersOk() (*map[string]string, bool)`

GetExportParametersOk returns a tuple with the ExportParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportParameters

`func (o *RunExportReportTaskVM) SetExportParameters(v map[string]string)`

SetExportParameters sets ExportParameters field to given value.

### HasExportParameters

`func (o *RunExportReportTaskVM) HasExportParameters() bool`

HasExportParameters returns a boolean if a field has been set.

### SetExportParametersNil

`func (o *RunExportReportTaskVM) SetExportParametersNil(b bool)`

 SetExportParametersNil sets the value for ExportParameters to be an explicit nil

### UnsetExportParameters
`func (o *RunExportReportTaskVM) UnsetExportParameters()`

UnsetExportParameters ensures that no value is present for ExportParameters, not even an explicit nil
### GetFormat

`func (o *RunExportReportTaskVM) GetFormat() ExportFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *RunExportReportTaskVM) GetFormatOk() (*ExportFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *RunExportReportTaskVM) SetFormat(v ExportFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *RunExportReportTaskVM) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetPagesCount

`func (o *RunExportReportTaskVM) GetPagesCount() int32`

GetPagesCount returns the PagesCount field if non-nil, zero value otherwise.

### GetPagesCountOk

`func (o *RunExportReportTaskVM) GetPagesCountOk() (*int32, bool)`

GetPagesCountOk returns a tuple with the PagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesCount

`func (o *RunExportReportTaskVM) SetPagesCount(v int32)`

SetPagesCount sets PagesCount field to given value.

### HasPagesCount

`func (o *RunExportReportTaskVM) HasPagesCount() bool`

HasPagesCount returns a boolean if a field has been set.

### GetT

`func (o *RunExportReportTaskVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *RunExportReportTaskVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *RunExportReportTaskVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


