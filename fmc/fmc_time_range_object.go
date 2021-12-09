package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

var timeRangeObjectType = "TimeRange"

type TimeRangeRecurrence struct {
	StartTime      string   `json:"rangeStartTime"`
	EndTime        string   `json:"rangeEndTime"`
	StartDay       string   `json:"rangeStartDay"`
	EndDay         string   `json:"rangeEndDay"`
	Days           []string `json:"days"`
	DailyStartTime string   `json:"dailyStartTime"`
	DailyEndTime   string   `json:"dailyEndTime"`
	RecurrenceType string   `json:"recurrenceType"`
}

type TimeRangeObjectInput struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	Type               string `json:"type"`
	EffectiveStartDate string `json:"effectiveStartDateTime"`
	EffectiveEndDate   string `json:"effectiveEndDateTime"`

	RecurrenceList []TimeRangeRecurrence `json:"recurrenceList"`
}

type TimeRangeObject struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Type               string `json:"type"`
	EffectiveStartDate string `json:"effectiveStartDateTime"`
	EffectiveEndDate   string `json:"effectiveEndDateTime"`

	RecurrenceList []TimeRangeRecurrence `json:"recurrenceList"`
}

func (v *Client) CreateFmcTimeRangeObject(ctx context.Context, object *TimeRangeObjectInput) (*TimeRangeObject, error) {
	if object == nil {
		return nil, fmt.Errorf("input data is empty")
	}
	object.Type = timeRangeObjectType

	url := fmt.Sprintf("%s/object/timeranges", v.domainBaseURL)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("creating time range object: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating time range object: %s - %s", url, err.Error())
	}
	item := &TimeRangeObject{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("getting time range objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) GetFmcTimeRangeObject(ctx context.Context, id string) (*TimeRangeObject, error) {
	url := fmt.Sprintf("%s/object/timeranges/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting time range object: %s - %s", url, err.Error())
	}
	item := &TimeRangeObject{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting time range object: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateFmcTimeRangeObject(ctx context.Context, id string, object *TimeRangeObject) (*TimeRangeObject, error) {
	if object == nil {
		return nil, fmt.Errorf("input data is empty")
	}
	object.Type = timeRangeObjectType

	url := fmt.Sprintf("%s/object/timeranges/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating time range object: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating time range object: %s - %s", url, err.Error())
	}
	item := &TimeRangeObject{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting time range object: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcTimeRangeObject(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/object/timeranges/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting time range object: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
