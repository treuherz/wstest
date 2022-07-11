# Testing WebSocket proxying 

First, `go run .` to start the "backend" server.

Run `npm start` to start the dev server and proxy. If the `react-dev-utils` dep is pointed to `5.0.1` you should see something like

```
WEB REQUEST ORIGIN: http://localhost:8080
WEB REQUEST HOST: localhost:8080
WS REQUEST ORIGIN: http://localhost:3000
WS REQUEST HOST: localhost:8080
```

in the output of the Go server. Note that the Web request Origin header has been changed to match the proxied host but the WS request hasn't been.

If `react-dev-utils` points to `https://gitpkg.now.sh/treuherz/create-react-app/packages/react-dev-utils?main` then you should see

```
WS REQUEST ORIGIN: http://localhost:8080
WS REQUEST HOST: localhost:8080
WEB REQUEST ORIGIN: http://localhost:8080
WEB REQUEST HOST: localhost:8080
```

where both Origin headers have been correctly changed.
