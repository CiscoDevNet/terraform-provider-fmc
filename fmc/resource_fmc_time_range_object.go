package fmc

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFmcTimeRangeObject() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Time Range Object in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_time_range_object\" \"test\" {\n" +
			"    name        		  = \"test-time-range\"\n" +
			"    description 		  = \"Test time range\"\n" +
			"    effective_start_date = \"2019-09-19T15:53:00\"\n" +
			"    effective_end_date   = \"2019-09-21T17:53:00\"\n" +
			"}\n" +
			"```",
		CreateContext: resourceFmcTimeRangeObjectCreate,
		ReadContext:   resourceFmcTimeRangeObjectRead,
		UpdateContext: resourceFmcTimeRangeObjectUpdate,
		DeleteContext: resourceFmcTimeRangeObjectDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this resource",
			},
			"effective_start_date": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Effective start date for this time range object (time in RFC3339 format)",
			},
			"effective_end_date": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Effective end date for this time range object (time in RFC3339 format)",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of this resource",
				Default:     " ",
				StateFunc: func(val interface{}) string {
					state := val.(string)
					if val == nil || state == "" || state == " " {
						return " "
					}
					return state
				},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// Fix for bug in the FMC API which returns " " for empty description
					if (new == " " && old == "") || (old == " " && new == "") {
						return true
					}
					return old == new
				},
			},
			"recurrence": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start_time": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "Start date for this recurrence (time in RFC3339 format)",
						},
						"end_time": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "End date for this recurrence (time in RFC3339 format)",
						},
						"start_day": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "Start day for this recurrence (time in RFC3339 format)",
						},
						"end_day": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "End day for this recurrence (time in RFC3339 format)",
						},
						"daily_start_time": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "Daily start time for this recurrence (time in RFC3339 format)",
						},
						"daily_end_time": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "Daily end time for this recurrence (time in RFC3339 format)",
						},
						"days": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type:    schema.TypeString,
								Default: "",
							},
						},
						"recurrence_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Type of recurrence. Allowed values: \"DAILY_INTERVAL\", \"RANGE\"",
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								v := strings.ToUpper(val.(string))
								allowedValues := []string{"DAILY_INTERVAL", "RANGE"}
								for _, allowed := range allowedValues {
									if v == allowed {
										return
									}
								}
								errs = append(errs, fmt.Errorf("%q must be in %v, got: %q", key, allowedValues, v))
								return
							},
						},
					},
				},
				Description: "List of URL objects to add",
			},
		},
	}
}

func resourceFmcTimeRangeObjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	var recurrences []TimeRangeRecurrence
	if inputObjs, ok := d.GetOk("recurrence"); ok {
		for _, obj := range inputObjs.([]interface{}) {
			obji := obj.(map[string]interface{})

			days := []string{}
			if obji["days"] != nil {
				for _, day := range obji["days"].([]interface{}) {
					days = append(days, day.(string))
				}
			}

			recurrence := TimeRangeRecurrence{
				StartTime:      obji["start_time"].(string),
				EndTime:        obji["end_time"].(string),
				StartDay:       obji["start_day"].(string),
				EndDay:         obji["end_day"].(string),
				DailyStartTime: obji["daily_start_time"].(string),
				DailyEndTime:   obji["daily_end_time"].(string),
				RecurrenceType: obji["recurrence_type"].(string),
				Days:           days,
			}

			// there is a bug in FMC API: when type is set to DAILY_INTERVAL -
			// start day and end day are not required, but API throws an error
			// if they were not passed, however when passed during creation,
			// start day and end day are not returned by API in GET requests
			if recurrence.RecurrenceType == "DAILY_INTERVAL" {
				// if days array is empty for DAILY_INTERVAL recurrence type
				// - API will throw an error, it's invalid request anyway, but
				// this additional check will prevent provider panic
				if len(days) > 1 {
					recurrence.StartDay = days[0]
					recurrence.EndDay = days[len(days)-1]
				}
			}

			recurrences = append(recurrences, recurrence)
		}
	}

	res, err := c.CreateFmcTimeRangeObject(ctx, &TimeRangeObjectInput{
		Name:               d.Get("name").(string),
		Description:        d.Get("description").(string),
		EffectiveStartDate: d.Get("effective_start_date").(string),
		EffectiveEndDate:   d.Get("effective_end_date").(string),
		RecurrenceList:     recurrences,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create time range object",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourceFmcTimeRangeObjectRead(ctx, d, m)
}

func resourceFmcTimeRangeObjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetFmcTimeRangeObject(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read time range object",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read time range object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("description", item.Description); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read time range object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("effective_start_date", item.EffectiveStartDate); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read time range object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("effective_end_date", item.EffectiveEndDate); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read time range object",
			Detail:   err.Error(),
		})
		return diags
	}

	if item.RecurrenceList != nil {
		recurrenceList := []interface{}{}

		for _, recurrence := range item.RecurrenceList {

			recurrenceList = append(recurrenceList, map[string]interface{}{
				"start_time":       recurrence.StartTime,
				"end_time":         recurrence.EndTime,
				"start_day":        recurrence.StartDay,
				"end_day":          recurrence.EndDay,
				"daily_start_time": recurrence.DailyStartTime,
				"daily_end_time":   recurrence.DailyEndTime,
				"recurrence_type":  recurrence.RecurrenceType,
				"days":             recurrence.Days,
			})
		}

		if err := d.Set("recurrence", recurrenceList); err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to read time range object",
				Detail:   err.Error(),
			})
			return diags
		}
	}

	return diags
}

func resourceFmcTimeRangeObjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("name", "description", "effective_start_date", "effective_end_date", "recurrence") {
		var recurrences []TimeRangeRecurrence
		if inputObjs, ok := d.GetOk("recurrence"); ok {
			for _, obj := range inputObjs.([]interface{}) {
				obji := obj.(map[string]interface{})

				days := []string{}
				if obji["days"] != nil {
					for _, day := range obji["days"].([]interface{}) {
						days = append(days, day.(string))
					}
				}

				recurrence := TimeRangeRecurrence{
					StartTime:      obji["start_time"].(string),
					EndTime:        obji["end_time"].(string),
					StartDay:       obji["start_day"].(string),
					EndDay:         obji["end_day"].(string),
					DailyStartTime: obji["daily_start_time"].(string),
					DailyEndTime:   obji["daily_end_time"].(string),
					RecurrenceType: obji["recurrence_type"].(string),
					Days:           days,
				}

				// there is a bug in FMC API: when type is set to DAILY_INTERVAL -
				// start day and end day are not required, but API throws an error
				// if they were not passed, however when passed during creation,
				// start day and end day are not returned by API in GET requests
				if recurrence.RecurrenceType == "DAILY_INTERVAL" {
					// if days array is empty for DAILY_INTERVAL recurrence type
					// - API will throw an error, it's invalid request anyway, but
					// this additional check will prevent provider panic
					if len(days) > 1 {
						recurrence.StartDay = days[0]
						recurrence.EndDay = days[len(days)-1]
					}
				}

				recurrences = append(recurrences, recurrence)
			}
		}

		_, err := c.UpdateFmcTimeRangeObject(ctx, id, &TimeRangeObject{
			Name:               d.Get("name").(string),
			Description:        d.Get("description").(string),
			EffectiveStartDate: d.Get("effective_start_date").(string),
			EffectiveEndDate:   d.Get("effective_end_date").(string),
			ID:                 id,
			RecurrenceList:     recurrences,
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update time range object",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	return resourceFmcTimeRangeObjectRead(ctx, d, m)
}

func resourceFmcTimeRangeObjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteFmcTimeRangeObject(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete time range object",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
