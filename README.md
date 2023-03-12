# Vehicle Rental

Vehicle Rental is a two-sided marketplace app for renting vehicles. Its primary feature is that consumers can rent vehicles while renters can offer their vehicles for rent. This Restful API was built using Golang and SQL, GORM for ORM and dependency injection as the chosen modularization strategy.

## 🔥 Showcase

- [Web Client with ReactJS](https://github.com/rfauzi44/vehicle-rental-web)
- Web Screenshot
<p float="left">
<img src="app1.png" alt="Alt text" height="500">
<img src="app2.png" alt="Alt text" height="500">
</p>

- Database Schema

<img src="db-vehicle-rental.png" alt="Alt text" height="250">

- [Docker Image](https://hub.docker.com/r/rfauzi/vehicle-rental-backend)
- [Postman Docs](https://documenter.getpostman.com/view/25042327/2s93JtQPPz)

## ⚡ Feature

- CRUD for all modules
- GORM Query
- Authentication and Authorization
- Sorting, Filtering, Searching
- Database Migration

## 💻 Built with

- [Gorilla MUX](https://github.com/gorilla/mux): for handling HTTP requests and responses
- [GORM](https://github.com/go-gorm/gorm): for ORM library
- [JWT](https://github.com/golang-jwt/jwt): for authentication and authorization
- [Postgres](https://github.com/postgres/postgres): for DBMS
- [Sendinblue](https://github.com/sendinblue/APIv3-go-library): for sending emails
