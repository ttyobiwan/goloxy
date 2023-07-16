# Goloxy

Simple service for handling Auth0 authentication. There are three routes in place:

- `/login?redirect_url=<str>` redirects to Auth0 login page
- `/callback?code=<str>&redirect_url=<str>` gets the access token and redirects to `redirect_url?access_token=<str>`
- `/logout?redirect_url=<str>` redirects to Auth logout page

Server uses Go default net/http package. There is only one dependency installed.

This is for my personal usage, to be able to just copy & paste the config and get the token, without adding these endpoints into every single backend app, or setting up some insane JS apps with million dependencies. If you want to use it in any different way: go ahead.

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
