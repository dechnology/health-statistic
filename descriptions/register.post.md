This is the ONLY way to create a new user. The registration questionnaire must be filled and the answers must be sent via this endpoint accordingly.

## Request Body

- `response` *`object`* **Required**
    The response to the registration questionnaire.

    - `questionnaire_id` *`string`* **Required**
        The ID of the registration questionnaire. This special questionnaire can be obtained via `questionnaires/registration [GET]` endpoint.

    - `answers` *`array`* **Required**
        The answers of questions in the registration questionnaire. All questions must be answered.

        - `question_id` *`string`* **Required**
        The ID of the question in the registration questionnaire.

        - `body` *`string`* **Required**
        The answer to the question in the registration questionnaire.

- `user` *`object`* **Required**
    The user to be created. A user must be created with the following 15 pieces of data provided. Note that the ID field also needs to be provided since we want to match the ID in our database with the one in Auth0.

    - `id` *`string`* **Required**
        The user's Auth0 ID.

    - `birth_year` *`string`* **Required**
        Part of the registration questionnaire.

    - `height` *`float`* **Required**
        Part of the registration questionnaire.

    - `weight` *`float`* **Required**
        Part of the registration questionnaire.

    - `gender` *`string`* **Required**
        Part of the registration questionnaire.

    - `education_level` *`string`* **Required**
        Part of the registration questionnaire.

    - `occupation` *`string`* **Required**
        Part of the registration questionnaire.

    - `marriage` *`string`* **Required**
        Part of the registration questionnaire.

    - `medical_history` *`string`* **Required**
        Part of the registration questionnaire.

    - `medication_status` *`string`* **Required**
        Part of the registration questionnaire.

    - `demented_among_direct_relatives` *`bool`* **Required**
        Part of the registration questionnaire.

    - `head_injury_experience` *`bool`* **Required**
        Part of the registration questionnaire.

    - `ear_condition` *`string`* **Required**
        Part of the registration questionnaire.

    - `eyesight_condition` *`string`* **Required**
        Part of the registration questionnaire.

    - `smoking_habit` *`string`* **Required**
        Part of the registration questionnaire.

