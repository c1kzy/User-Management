# Content Management

This project aims to create a simple Content Management System for managing posts and users. Users can add posts to the system, and other users can browse through all articles from different authors. Additionally, authors have the privilege to edit or delete their articles after providing the required password.
## Installation


```bash
git clone https://github.com/c1kzy/User-Management.git
cd rest-api
go run cmd/main.go
```

## Usage

Adding a New Entity:

* Choose between post and user when adding a new entity.
  Specify the relevant details for the chosen entity type.

## User Interaction:

* Users add, update, delete options available.
## Browsing Articles:

* All users can browse through a list of posts.
## Author Actions:

* Post Interaction.
  User can create, update and delete posts(if permission available).
  Status Tracking:

Use created_at, updated_at, and status to determine the current status of a user or article.

## Documentation

Documentation available at
```
http://localhost:8080/swagger/index.html
```
