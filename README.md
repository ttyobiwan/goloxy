# Goloxy

Simple service for handling Auth0 authentication. There are three routes in place:

- `/login?redirect_url=<str>` redirects to Auth0 login page
- `/callback?code=<str>&redirect_url=<str>` gets the access token and redirects to `redirect_url?access_token=<str>`
- `/logout?redirect_url=<str>` redirects to Auth logout page

Server uses Go default net/http package. There is only one dependency installed.

## Usage

Build Docker images

```bash
make build
```

Start Docker containers

```bash
make up
```

All three routes should be available at `http://localhost:8080`.
