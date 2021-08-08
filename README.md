# MVC model
- controllers
  - validate input request
  - NOT ANY business logic
- service
  - handle the business logic
  - stateless
  - singloten
- model/domain/DAOs
  - define the structure of domain object
  - persistence layer

# Package organizing
- mvc
  - controllers
  - services
  - domain
  - app
  - utils
