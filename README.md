# fgr

basic gift registry

## Dev Instructions

### Stack
- UI: Vue
- Backend: Golang HTTP server
- DB: Postgres run on Droplet
- Deployment: Digital Ocean Droplet

### Starting UI
- npm install
- npm install -g @vue/cli
- npm run dev

### Backend

login to be able to see claimed gifts & to claim gifts

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
  - 
### TODO
- [ ] Add versioning to endpoints
- [ ] Support invited user getting claimed
- [x] CSRF protection
- [ ] Invited user email is not the one they authed with
- [x] Session ID table
- [ ] Restrict firewall on digital ocean (remove postgres)
- [ ] Secret storage on digital ocean?
- [ ] Admin select user login (deprecated public select user) and create user and users
- [ ] Prevent multiple clicks doing the same thing

## DB Schema

* event 
    - id
    - ownerUserId
    - name

* users
    - id
    - name
    - email

* membership
    - id
    - eventId
    - userId

* gift_requests
    - id
    - userId
    - eventId
    - name
    - assignedUserId