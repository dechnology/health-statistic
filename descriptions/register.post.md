This is the ONLY way to create a new user. The registration questionnaire must be filled and the answers must be sent via this endpoint accordingly.

While the initial questionnaire may have 30 questions, some answers are part of the user data like "birth_year", "gender", and so on. Therefore, we will use some of the data as inputs to create the user while other questions are combined and treated like a questionnaire called "registration questionnaire".

Hence, the user registration needs two things: user data and the response to the registration questionnaire. A user can only be register using this endpoint. Note that the ID field needs to be provided and it needs to match the Auth0 ID. This endpoint does 4 things in order: create a new user, respond to the registration questionnaire, assign one MyCard to this user, and send a notification via FCM.

## Request Body

- `response` _`object`_ **Required**
  The response to the registration questionnaire.

  - `questionnaire_id` _`string`_ **Required**
    The ID of the registration questionnaire. This special questionnaire can be obtained via `questionnaires/registration [GET]` endpoint.

  - `answers` _`array`_ **Required**
    The answers of questions in the registration questionnaire. All questions must be answered.

    - `question_id` _`string`_ **Required**
      The ID of the question in the registration questionnaire.

    - `choise_ids` _`string[]`_ **Required**
      The ID of the choise chosen for this question in the registration questionnaire. Note that though all questions in the registration questionnaire are single choise question, you are required to send in array format.

- `user` _`object`_ **Required**
  The user to be created. A user must be created with the following 15 pieces of data provided. Note that the ID field also needs to be provided since we want to match the ID in our database with the one in Auth0.

  - `id` _`string`_ **Required**
    The user's Auth0 ID.

  - `birth_year` _`string`_ **Required**
    Part of the registration questionnaire.

  - `height` _`float`_ **Required**
    Part of the registration questionnaire.

  - `weight` _`float`_ **Required**
    Part of the registration questionnaire.

  - `gender` _`string`_ **Required**
    Part of the registration questionnaire.

  - `education_level` _`string`_ **Required**
    Part of the registration questionnaire.

  - `occupation` _`string`_ **Required**
    Part of the registration questionnaire.

  - `marriage` _`string`_ **Required**
    Part of the registration questionnaire.

  - `medical_history` _`string`_ **Required**
    Part of the registration questionnaire.

  - `medication_status` _`string`_ **Required**
    Part of the registration questionnaire.

  - `demented_among_direct_relatives` _`bool`_ **Required**
    Part of the registration questionnaire.

  - `head_injury_experience` _`bool`_ **Required**
    Part of the registration questionnaire.

  - `ear_condition` _`string`_ **Required**
    Part of the registration questionnaire.

  - `eyesight_condition` _`string`_ **Required**
    Part of the registration questionnaire.

  - `smoking_habit` _`string`_ **Required**
    Part of the registration questionnaire.
