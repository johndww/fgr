# fgr

basic gift registry

## Dev Instructions

### Stack
- UI: Vue
- Backend: Golang HTTP server
- DB: ???
- Deployment: ??? (digitial ocean droplet & postgres likely)

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
- [ ] Show all events you are linked to
- [ ] Create Event button
- [ ] Logout

### Create Event
- [ ] capture event name and description
- [ ] add members
    - [ ] by email 
    - [ ] or by previous association

### Show Event
- [ ] display all members & gifts
- [ ] ability to claim and unclaim gifts
- [ ] add/remove your gift requests
- [ ] event owner controls
  - [ ] delete event (with confirmation)
  - [ ] update event details (& membership)

## DB Schema

* event 
    - id
    - owner
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