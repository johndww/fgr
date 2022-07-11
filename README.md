# fgr

basic gift registry

## Dev Instructions

### Stack
- UI: Vue
- Backend: Golang HTTP server
- DB: ???
- Deployment: ??? (digital ocean droplet & postgres likely)

### Starting UI
- npm install
- npm install -g @vue/cli
- npm run serve

### Backend

login to be able to see claimed gifts & to claim gifts

## Pages

### Unauthenticated Home Page
- [ ] Login/Create Account
- [x] Welcome page (unauthenticated)

### Navigation
- [ ] General Navigation
- [ ] Back button (breadcrumbs?)

### Authenticated Home Page
- [x] Show all events you are linked to
- [x] Create Event button
- [ ] Logout

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
    - description
    - assignedUserId