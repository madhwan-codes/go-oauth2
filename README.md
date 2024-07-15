# OAuth2 Implementation in Go

This repository demonstrates the implementation of OAuth2 authentication using Google, Facebook, and GitHub in a Go application. The application uses the Echo framework for handling HTTP requests.

## Getting Started

### Prerequisites

Make sure you have the following installed:

- [Go](https://golang.org/dl/) 1.16 or higher
- [Git](https://git-scm.com/)
- [Echo](https://echo.labstack.com/) v4
- OAuth2 credentials from Google, Facebook, and GitHub

### Installing

Clone the repository:

```sh
git clone https://github.com/yourusername/go-oauth2-example.git
cd go-oauth2-example
```

Install the dependencies:

```sh
go mod tidy
```

Running the Application
Run the application with:
```sh
go run main.go
```


Routes
- GET /: Health check endpoint
- GET /google_login: Initiates Google OAuth2 login
- GET /google_callback: Google OAuth2 callback
- GET /github_login: Initiates GitHub OAuth2 login
- GET /github_callback: GitHub OAuth2 callback
- GET /facebook_login: Initiates Facebook OAuth2 login
- GET /facebook_callback: Facebook OAuth2 callback
