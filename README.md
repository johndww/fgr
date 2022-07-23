# fgr

basic gift registry

## Dev Instructions

### Stack
- UI: Vue
- Backend: Golang HTTP server
- DB: Postgres run on Droplet
- Deployment: Digital Ocean Droplet

### Setup UI
- cd ui
- npm install
- npm run dev

### Setup Backend
- create a .env.local file and fill out the fields based off of .env
- cd server
- go mod vendor
- go run main.go

## Pages

### Unauthenticated Home Page
- [x] Login/Create Account
- [x] Welcome page (unauthenticated)

### Navigation
- [x] General Navigation
- [x] Back button to events

### Authenticated Home Page
- [x] Show all events you are linked to
- [x] Create Event button
- [x] Logout
- [x] Google Auth

### Create Event
- [x] capture event details
- [x] add members
    - [x] by email 
    - [x] or by previous association

### Show Event
- [x] display all members & gifts
- [x] ability to claim and unclaim gifts
- [x] add/remove your gift requests
- [ ] event owner controls
  - [ ] delete event (with confirmation)
  - [x] update event details (& membership)
### TODO
- [x] Add versioning to endpoints
- [x] CSRF protection
- [x] Session ID table
- [x] Secret storage on digital ocean (.env file)
- [x] Admin select user login (deprecated public select user) and create user and users
- [ ] Restrict firewall on digital ocean (remove postgres)
- [ ] Prevent multiple clicks doing the same thing
- [ ] Support invited user getting claimed
- [ ] Invited user email is not the one they authed with
