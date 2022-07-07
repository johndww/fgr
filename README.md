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
- [ ] website explanation

### Authenticated Home Page
- [x] Show all events you are linked to
- [ ] Create Event button
- [ ] Logout

### Create Event
- [x] capture event details
- [ ] add members
    - [ ] by email 
    - [ ] or by previous association

### Show Event
- [x] display all members & gifts
- [x] ability to claim and unclaim gifts
- [x] add/remove your gift requests
- [ ] event owner controls
  - [ ] delete event (with confirmation)
  - [ ] update event details (& membership)

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
    - eventId
    - userId

* gift_requests
    - userId
    - eventId
    - name
    - description
    - claimedByUserId