# Snippetbox
We’ll be building a web application called Snippetbox, which lets people paste and share snippets of text — a bit like Pastebin or GitHub’s Gists. This is a Golang training to learn how to build production-ready web applications with Go, built while reading Alex Edwards "Let’s Go" book.

A user is able save and view snippets via the app. Here, we learn how to structure a project, routing requests, working with a database, processing forms and displaying dynamic data safely.Then later we’ll add user accounts, and restrict the application so that only registered users can create snippets. This will take us through more advanced topics like configuring a HTTPS server, session management, user authentication and middleware.

Touch Points:
- # All the fundamentals  — How to start a server, create handlers, send responses, route requests and serve static files.
- Structure and organization — How to create an idiomatic and scalable structure for your web application.
- Using Modules — How to use Go's Module functionality to manage and version control your dependencies.
- Managing configuration — How to use command-line flags and dependency injection to manage your application settings.
- Logging and Error Handling — How to implement leveled logging and centralized error handling.
- SQL databases — How to design a database model, set up a connection pool, and execute statements and queries.
- HTML templating — How to cache your templates, display dynamic data, create custom functions and handle runtime errors.
- Middleware — How to create your own middleware to perform common actions (like logging requests and recovering panics).
- RESTful routing — How to create a modern request routing structure that follows the principles of REST.
-Form validation — How to implement reusable and user-friendly pattern for validating forms and displaying errors.
-Session management — How to use and configure sessions to persist data between requests.
- Using HTTPS — How to correctly setup a HTTPS server and configure it for improved performance and security.
- Prevent common vulnerabilities — How to prevent SQL injection, CSRF, XSS, clickjacking and slow-client attacks.
- Authentication and authorization — How to safely encrypt user passwords and add signup, login and logout functionality.
- Request context — How to use Go's context.Context to pass data between your middleware and handlers.
- Testing — How to create unit tests, integration tests and end-to-end tests, mock dependencies and measure test coverage.
And most importantly… How to put it together in a fully-functioning application!


Below a list of used libraries:

- [zerolog](https://github.com/rs/zerolog): fast and structured logging
- [lumberjack](https://github.com/natefinch/lumberjack): log rolling files
- [cobra](https://github.com/spf13/cobra): command like support
- [viper](https://github.com/spf13/viper): flags and file configuration
- [goland-migrate](https://github.com/golang-migrate/migrate): sql database migration
