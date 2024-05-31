## Micro Service for Authorization.

This is a service that will be one part of a 3 part system going together to build the backend of my personal asset and inventory application. I am building this application to take my learning and understanding of the Go programming language to another, or I guess you'd say the next level.

## Tech:

- Go: The programming language the service is built with.
- mySQL: Relational  Database Management chosen for the project. Tried and true.
- GORM: A powerful ORM. It simplifies many database operations and adds a layet of abstraction that should make my life a lot easier..  but it does come with a learning curve. It's the first part of the service I am implementing. I plan to use it across all the services.
- Gin: the web framework I chose to base the API on. Previously I didn't use a framework. [Here](https://github.com/ddcroft73/go-crud-app) I found it a lot easier to build an API without employing an API. This isn't something I'd ever do in a language like Python, or Node. In fact I almost didn't even use one, but I feel I should learn Gin because it is really popular and the way it handles routing and JSON is a bit more inviting than without. Also the middleware looks a lot easier to work with..

## Routes:

This service will be responsible for handling the following parts of the backend.

- /login-acces-token
- /test-token
- /register.. (CreateUSer))
- /verify-email
- /resend-email-verification
- /reset-password
- /recover-password
- /verify-2fa
- /resend-2fa
- /get-user/{id}
- /update-user/{id}
- /delete-user/{id}
- /get-all-users

...and any others I may have forgotten.

The application:

One of the hardest things about programming personally is coming up with an idea for an app that keeps you engaged long enough to finish it. At least fore anyway. Programming is all about solving problems. At present I can't think of any problems I really need to solve, but I want to go deeper with Go so I need to build something to learn these skills. A Personal Asset Inventory system isnt the greatest idea and it certainly won't make me rich. But it should give me enough purchase to learn a good bit about building a backend in Go. It will be user based. It will have admin and well user roles on the system so I will have to build it with that in mind. There will be access only for administration and access for everyday users. I am using g a JWT system with 2fa. Access as well as refresh tokens. Users will be able to inventory any and all of their assets/memorabila/collectables and keep images of all.

Micro Services:

There will be 4 services in all.

- Auth Service: wil handle all onboarding and access to the application. As well as password recovery, email/phone verification. 2FA etc. as well
  as User maintence. Since the API is mostly built around the User model it makes sense to just incorporate a few user CRUD endpoints foruser maintence
- Asset Service: Will handle the management of user assets. It will have its own database to store info about assets as well as images. It will expose API endpoints for creating, updating, deleting of assets. Assets will be searchable, filtration and pagination.
- Notification Service: will be responsible for any emails, or SMS messages that need to be sent to user sfor verification purposes. This service has already been built for use in another backend I have written in python. This is a perfect example of microservices in different languages.
