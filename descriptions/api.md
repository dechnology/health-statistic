This is the API documentation for 「健康資料公鏈」

# Usage

## Flow Diagram
<!-- ![flow_diagram](https://cdn.discordapp.com/attachments/874556062815100940/1132920083174408222/App-.drawio.png) -->

## API Endpoints for Frontend Application

Full descriptions of all endpoints are described below; however, more common endpoints for a client application are described here. For detailed types of params, body, or response, please refer to the full descriptiions below. 

### Server Health Check - `/health_check` [**GET**]

This endpoint checks if the server is alive and the database is connected. It returns a code of 200 if ture.

---

### Get Registration Questionnaire - `/questionnaires/registration` [**GET**]

For the sake of scalibility, we make the questionnaire system a general one. That is, we can later create more questionnaires in our application. Upon initialization of our database, the registration questionnaire is created and assigned a fixed UUID. Before the user register, we need to have them fill in the registration questionnaire. This questionniare can be fetch using this endpoint. Note that this questionnaire correspond to question 16 to 30 since the answers to question 1 to 15 are part of the user data.

---

### User Registration - `/register` [**POST**]

While the initial questionnaire may have 30 questions, some answers are part of the user data like "birth_year", "gender", and so on. Therefore, we will use some of the data as inputs to create the user while other questions are combined and treated like a questionnaire called "registration questionnaire".

Hence, the user registration needs two things: user data and the response to the registration questionnaire. A user can only be register using this endpoint. Note that the ID field needs to be provided and it needs to match the Auth0 ID. This endpoint does 4 things in order: create a new user, respond to the registration questionnaire, assign one MyCard to this user, and send a notification via FCM.

---

### After User Registration

#### Get user data - `/users/{id}` [**GET**]

#### Get user's history notifications - `/users/{id}/notifications` [**GET**]

#### Get user's prices - `/users/{id}/prices` [**GET**]
Get all prices the user has won. Practically, a user may only have 1 or no prices. But for the sake of scalibility, we allow multiple prices to be owned by an user.

#### Get user's MyCards - `/users/{id}/mycards` [**GET**]
Note that the ID of this data is the card number. Currently, all users will have only one MyCard assigned to them. However, for the sake of scalibility, multiple MyCards can be assigned to an user.

#### Get all prices - `/prices` [**GET**]
Get all prices from the database, this is merely for demonstration purposes. The prices will be assigned randomly after the event is finished.