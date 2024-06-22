// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type GetBandwidthStatisticsErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetBandwidthStatisticsErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetBandwidthStatisticsErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetBandwidthStatisticsErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetBandwidthStatisticsResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetBandwidthStatisticsResponseBody struct {
	Errors []GetBandwidthStatisticsErrors `json:"errors,omitempty"`
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GetBandwidthStatisticsResponseBody{}

func (e *GetBandwidthStatisticsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
