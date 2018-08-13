# Quack: Tamuhack's application/registration system.
## Quackity quack quack.

### Backend dev-environment instructions:
- Set up a standard go environment with something like this: https://medium.com/@AkyunaAkish/setting-up-a-golang-development-environment-mac-os-x-d58e5a7ea24f (should be easy to find instructions for PC)
- Run `go get github.com/tamuhack-org-quack`
- `cd` into that ^ directory and run:
  - `go get github.com/oxequa/realize`
  - `go get -d -v golang.org/x/net/html`
  - `go get -d -v github.com/gorilla/handlers`
  - `go get -d -v github.com/gorilla/mux`
- Finally, in a seperate terminal window, but in the same directory, run:
  -`realize start --run --server --no-config`
- This should start a server on `http://localhost:8080`

### Frontend dev-environment instructions:
- Make sure you have the following dependecies/versions:
  - `node -v`: 10+
  - `yarn -v`: 1.9+
- `cd` into the `frontend/` directory
- run `yarn install`
- run `yarn dev`
- This should start a server on `http://localhost:3000`

### Actually contributing code:
- For debugging the project locally, visit `http://localhost:8080`. In short, the server serves an html file which references the frontend dev server. You won't have to visit the frontend (webpack) server at all.
- Submit a PR for most changes.
