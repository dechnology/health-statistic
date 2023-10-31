# Description

Upload HealthKit data for a user.

## Request Body

- `data` _`array`_ **Required**
  The data is an array of arrays. Each array represents a unit of healthkit data. Each unit of healthkit data is an array of 6 elements. The 6 elements are:
  
  - `type` _`string`_ **Required**
    The type of the healthkit data. For example, `HKQuantityTypeIdentifierHeartRate`, `HKQuantityTypeIdentifierStepCount`, and etc.
  
  - `count` _`string`_ **Required**
    The count of the healthkit data. For example, if the type is `HKQuantityTypeIdentifierHeartRate`, then the count is the heart rate.
  
  - `uuid` _`string`_ **Required**
    The UUID of the healthkit data.
  
  - `start_time` _`string`_ **Required**
    The start time of the healthkit data.
  
  - `end_time` _`string`_ **Required**
    The end time of the healthkit data.
  
  - `timezone_id` _`string`_ **Required**
    The timezone ID of the healthkit data.

- `start_time` _`datetime`_ **Required**
  The start time of this batch of healthkit data.

- `end_time` _`datetime`_ **Required**
  The end time of this batch of healthkit data.
