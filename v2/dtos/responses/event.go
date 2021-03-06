//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
)

// EventCountResponse defines the Response Content for GET event count DTO.
// This object and its properties correspond to the EventCountResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/EventCountResponse
type EventCountResponse struct {
	common.BaseResponse `json:",inline"`
	Count               uint32
	DeviceName          string `json:"deviceName"`
}

// EventResponse defines the Response Content for GET event DTOs.
// This object and its properties correspond to the EventResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/EventResponse
type EventResponse struct {
	common.BaseResponse `json:",inline"`
	Event               dtos.Event `json:"event"`
}

// MultiEventsResponse defines the Response Content for GET multiple event DTOs.
// This object and its properties correspond to the MultiEventsResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/MultiEventsResponse
type MultiEventsResponse struct {
	common.BaseResponse `json:",inline"`
	Events              []dtos.Event `json:"events"`
}

// UpdateEventPushedByChecksumResponse defines the Response Content for PUT event as pushed DTO.
// This object and its properties correspond to the UpdateEventPushedByChecksumResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/UpdateEventPushedByChecksumResponse
type UpdateEventPushedByChecksumResponse struct {
	common.BaseResponse `json:",inline"`
	Checksum            string `json:"checksum"`
}

func NewEventCountResponseNoMessage(requestId string, statusCode int, count uint32, deviceId string) EventCountResponse {
	return NewEventCountResponse(requestId, "", statusCode, count, deviceId)
}

func NewEventCountResponse(requestId string, message string, statusCode int, count uint32, deviceName string) EventCountResponse {
	return EventCountResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Count:        count,
		DeviceName:   deviceName,
	}
}

func NewEventResponseNoMessage(requestId string, statusCode int, event dtos.Event) EventResponse {
	return NewEventResponse(requestId, "", statusCode, event)
}

func NewEventResponse(requestId string, message string, statusCode int, event dtos.Event) EventResponse {
	return EventResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Event:        event,
	}
}

func NewMultiEventsResponse(requestId string, message string, statusCode int, events []dtos.Event) MultiEventsResponse {
	return MultiEventsResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Events:       events,
	}
}

func NewUpdateEventPushedByChecksumResponseNoMessage(requestId string, statusCode int, checksum string) UpdateEventPushedByChecksumResponse {
	return NewUpdateEventPushedByChecksumResponse(requestId, "", statusCode, checksum)
}

func NewUpdateEventPushedByChecksumResponse(requestId string, message string, statusCode int, checksum string) UpdateEventPushedByChecksumResponse {
	return UpdateEventPushedByChecksumResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Checksum:     checksum,
	}
}
