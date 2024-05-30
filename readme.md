## Auth micro service in GO

This is the Auth API for a Backend written in Go. This project is going to be a Personal inventory organizer. It will consist of micro services all written in Go. So far the code is mostlu boilerplate. You can add a user and hash their password. That's about it. Still have to build out all the routes.

## Tech Stack:

- GIN
- GORM
- MySQL

It's pretty simple but Learning the way Gin Handles the routes seems a bit more complicated than not using a framework. GORM has been interesting. I'm still figuring it out but the database so far seems to be up to speed.

## Routes:

/login-acces-token
/test-token
/register
/verify-email
/resend-email-verification
/reset-password
/recover-password
/verify-2fa
/resend-2fa

## Other Services:
 
User assets
Still working it out. So far I just have a goal for the app in mind, and Im currently building the service to handling all the authorizations.
