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
- vite (or vite build for prod)
  - for dev login, ```/devlogin``` with ```e36afe1d-ef84-42a8-af38-0e6d60745e9e```

### Setup Backend (Local)
- create a .env.local file and fill out the fields based off of .env
- install https://magefile.org/ (brew install mage)
- cd server
- go mod vendor
- (ensure DB is running - postgres in Docker)
- go run main.go

### SSL (yearly renewal)
- Namecheap SSL
- Create new CSR in Droplet (SSH in, ~/certs using openssl from https://www.namecheap.com/support/knowledgebase/article.aspx/9446/2290/generating-csr-on-apache-opensslmodsslnginx-heroku/)
- Use CSR in Namecheap website for PositiveSSL manage
- Use CNAME verification - add CNAME in Digital Ocean UI and validate in Namecheap
- Cert issued to email (.crt and .bundle)
- Follow Cert install instructions for NGINX https://www.sectigo.com/knowledge-base/detail/Certificate-Installation-NGINX-1527076083655/kA01N000000zFJQ
- Use these locations (replace old): ```    ssl_certificate /home/fgr/certs/simplegift_app.chained.crt;
    ssl_certificate_key /home/fgr/certs/simplegift_app.key;```
- Reboot FGR (pulls new private key): ``` sudo systemctl restart fgr```
- Reboot nginx (pulls new chained cert): ```sudo systemctl restart nginx```

### Server Configuration / Operation
- nginx config: ```/etc/nginx/sites-enabled/```
- view server logs: ``` sudo journalctl -u fgr```
- start/stop app: ``` sudo systemctl stop fgr
   sudo systemctl start fgr```
- verify startup ```sudo systemctl status fgr```
- Update UI:
- - upload UI assets
- - ```./uiReplace.sh```

### Update Backend on Server
- ```~/code/fgr/server```
- ```BEHAVIOR=prod mage -v build```
- ```cd build```
- ``` sftp fgr@164.92.73.41```
- ```put fgr```
- on server in ~: ```mv fgr build```
- ```sudo systemctl stop fgr```
- ```sudo systemctl start fgr```
- verify: ```sudo systemctl status fgr```

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
- [x] event owner controls
  - [x] delete event (with confirmation)
  - [x] update event details (& membership)
### TODO
- [x] Add versioning to endpoints
- [x] CSRF protection
- [x] Session ID table
- [x] Secret storage on digital ocean (.env file)
- [x] Admin select user login (deprecated public select user) and create user and users
- [x] HTTPS and secure cookie
- [ ] Collate loading and errors on pages that have multiple async sources
- [x] Restrict firewall on digital ocean (remove postgres)
- [x] Prevent multiple clicks doing the same thing (gift claim/release)
- [x] Support invited user getting claimed
- [x] Google ad support
- [ ] Invited user email is not the one they authed with
