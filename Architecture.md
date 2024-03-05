# Use cases
User is able to create, update and delete his profile. Passwords are mandatory

[![](https://mermaid.ink/img/pako:eNp90MEKwjAMBuBXCTkpzBfYQXAOEfQg6ECwHsqa6qBbR9ciYn13awvDXcylIfngD31hrQVhjlLpR33nxsKpZB2EWl2qgcwVFoulPza3DlzvoZgdjJaNIqgNcUtinnARWdWLMIM-Eb8dsYuLKS5J0Q8-j1jExRSvYxr0erAeNrNDeP9eEN0uuX_h0e2Tm-Zihi2Zljci_M7rO2No79QSwzy0giR3yjJk3TtQ7qw-Prsac2scZZgyy4bfDG8xl1wN9P4ASpZ4EA?type=png)](https://mermaid.live/edit#pako:eNp90MEKwjAMBuBXCTkpzBfYQXAOEfQg6ECwHsqa6qBbR9ciYn13awvDXcylIfngD31hrQVhjlLpR33nxsKpZB2EWl2qgcwVFoulPza3DlzvoZgdjJaNIqgNcUtinnARWdWLMIM-Eb8dsYuLKS5J0Q8-j1jExRSvYxr0erAeNrNDeP9eEN0uuX_h0e2Tm-Zihi2Zljci_M7rO2No79QSwzy0giR3yjJk3TtQ7qw-Prsac2scZZgyy4bfDG8xl1wN9P4ASpZ4EA)
# DB diagram
[![](https://mermaid.ink/img/pako:eNqVkbFuxDAIhl_FYr57gajjLR06dY10QjZprCZ2BFjV6Zp3L05OctqqQ70Anw3-gTv4HAg68BOKXCK-Mc59cnY24ooQ33dQTwzNHyKLuoQzNWY5PxHNGKcWjigjhetixT8yH8pxng5ZngnV3qE2Vpbwi-XnSwtEUYvs8bqbqt89fZ7PbsmidndsbkN_dFcTrzE4poGYkqeN_F_gUdM3wSucYCa26QSb_6aiBx3JZgeduQH5vYc-1XdYNL_ekodOudAJ9o8e64JuwEmMUoia-eWx0GrWLw0qj5U?type=png)](https://mermaid.live/edit#pako:eNqVkbFuxDAIhl_FYr57gajjLR06dY10QjZprCZ2BFjV6Zp3L05OctqqQ70Anw3-gTv4HAg68BOKXCK-Mc59cnY24ooQ33dQTwzNHyKLuoQzNWY5PxHNGKcWjigjhetixT8yH8pxng5ZngnV3qE2Vpbwi-XnSwtEUYvs8bqbqt89fZ7PbsmidndsbkN_dFcTrzE4poGYkqeN_F_gUdM3wSucYCa26QSb_6aiBx3JZgeduQH5vYc-1XdYNL_ekodOudAJ9o8e64JuwEmMUoia-eWx0GrWLw0qj5U)# Layers

* .env
* .gitignore
# docs
* swagger
# cmd
* main.go

This layer is our entry point to the app

## internal
* infrastructure
  * inputports
    * http
      * handlers.go
      * models.go
* domain
  * models.go
  * repository.go
* app
  * queries
  * commands
  * auth
  * cache
* pkg
  * logger
  * database
* tools 
  * mockgen.go

**Infrastructure** directory handles our communication, http subdirectory for http connection and handlers logic(created subdirectory in case if connection will change)

**Models** models for CREATE, UPDATE, DELETE handlers, without created_at, deleted_at, status fields

**App** directory handles main business logic. E.g. profile create, get token, password hash

**Queries** directory handles request operations with data(GET requests)

**Commands** directory handles any operations with data that make changes(CREATE, UPDATE, DELETE)

**Cache** may be used to cache auth tokens, passwords etc.

**Auth** handles user authorization

**Domain** contain interfaces with CRUD operations and models for user struct, no dependencies

**Logger** directory for custom logger

**Database** directory handles a connection to database





