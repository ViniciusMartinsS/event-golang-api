# Event Rest API
Event Manager is a simple application that allows you to **create**, **update**, **list** & **delete** an event.

# Stack
- GoLang

# Getting Started


- Clone Project
- In the project root run the following command: `go build`
- Run the application by executing the following command: `./event-golang-api`
- Server runs on **8080**

# EndPoints

## Events

All possible endpoints for the event service:

- `GET /events/:id?`
- `POST /events`
- `PATCH /events/id`
- `DELETE /events/id`

# Disclaimer
This API was built to **practice** GoLang skills. <br/>
I have followed this [**ARTICLE**](https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da) and changed the code as I thought it would be better to arrange it. <br/>
Still missing some error handling. As I get experience on it, I will improve the code to follow patterns and make it better.
