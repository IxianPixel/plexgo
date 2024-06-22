// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetResourcesStatisticsRequest struct {
	// The timespan to retrieve statistics for
	// the exact meaning of this parameter is not known
	//
	Timespan *int64 `queryParam:"style=form,explode=true,name=Timespan"`
}

func (o *GetResourcesStatisticsRequest) GetTimespan() *int64 {
	if o == nil {
		return nil
	}
	return o.Timespan
}

type StatisticsResources struct {
	Timespan                 *int64   `json:"timespan,omitempty"`
	At                       *int64   `json:"at,omitempty"`
	HostCPUUtilization       *float32 `json:"hostCpuUtilization,omitempty"`
	ProcessCPUUtilization    *float32 `json:"processCpuUtilization,omitempty"`
	HostMemoryUtilization    *float32 `json:"hostMemoryUtilization,omitempty"`
	ProcessMemoryUtilization *float32 `json:"processMemoryUtilization,omitempty"`
}

func (o *StatisticsResources) GetTimespan() *int64 {
	if o == nil {
		return nil
	}
	return o.Timespan
}

func (o *StatisticsResources) GetAt() *int64 {
	if o == nil {
		return nil
	}
	return o.At
}

func (o *StatisticsResources) GetHostCPUUtilization() *float32 {
	if o == nil {
		return nil
	}
	return o.HostCPUUtilization
}

func (o *StatisticsResources) GetProcessCPUUtilization() *float32 {
	if o == nil {
		return nil
	}
	return o.ProcessCPUUtilization
}

func (o *StatisticsResources) GetHostMemoryUtilization() *float32 {
	if o == nil {
		return nil
	}
	return o.HostMemoryUtilization
}

func (o *StatisticsResources) GetProcessMemoryUtilization() *float32 {
	if o == nil {
		return nil
	}
	return o.ProcessMemoryUtilization
}

type GetResourcesStatisticsMediaContainer struct {
	Size                *int                  `json:"size,omitempty"`
	StatisticsResources []StatisticsResources `json:"StatisticsResources,omitempty"`
}

func (o *GetResourcesStatisticsMediaContainer) GetSize() *int {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetResourcesStatisticsMediaContainer) GetStatisticsResources() []StatisticsResources {
	if o == nil {
		return nil
	}
	return o.StatisticsResources
}

// GetResourcesStatisticsResponseBody - Resource Statistics
type GetResourcesStatisticsResponseBody struct {
	MediaContainer *GetResourcesStatisticsMediaContainer `json:"MediaContainer,omitempty"`
}

func (o *GetResourcesStatisticsResponseBody) GetMediaContainer() *GetResourcesStatisticsMediaContainer {
	if o == nil {
		return nil
	}
	return o.MediaContainer
}

type GetResourcesStatisticsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource Statistics
	Object *GetResourcesStatisticsResponseBody
}

func (o *GetResourcesStatisticsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetResourcesStatisticsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetResourcesStatisticsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetResourcesStatisticsResponse) GetObject() *GetResourcesStatisticsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
