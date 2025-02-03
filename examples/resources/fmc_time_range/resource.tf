resource "fmc_time_range" "example" {
  name        = "time_range_1"
  description = "My time range object"
  start_time  = "2025-01-07T20:20"
  end_time    = "2025-01-22T09:20"
  recurrence_list = [
    {
      recurrence_type  = "RANGE"
      range_start_time = "09:00"
      range_end_time   = "17:00"
      range_start_day  = "MON"
      range_end_day    = "FRI"
    }
  ]
}
