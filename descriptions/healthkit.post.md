Create a new HealthKit data assoicated with a given user.

## Request Body

- `user_id` *`string`* **Required**
    The ID of the user to whom this HealthKit data belongs.

- **Arbitrary JSON Data** Optional
    The HealthKit data to be stored. Note that this field is arbitrary and you can store any data you want. However, we recommend you to follow the [HealthKit Data Types](https://developer.apple.com/documentation/healthkit/healthkit_data_types) and [HealthKit Data Types Reference](https://developer.apple.com/documentation/healthkit/healthkit_data_types_reference) to store your data. This will make it easier for you to use the data in the future.