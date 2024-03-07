# Promotion System

This repository contains the codes of the Authentication System and Promotion System.

Feature:
+ Sign in 
  + store information includes full name, phone number, email, username, pass word, birthday, latest login
  + pass word hashed by bcrypt
+ Log in
  + with one of phone number, email or username and password 
  + response jwt token and give token in header to pass dashboard
+ Ping with /homepage to use middleware for accessing (*check /api/user.http for request*) 
+ Add voucher for user if user login the first time and based on campaigns
  + number of users received voucher will be limited by total_vouchers(campaigns)
  + use latest_login is nil to check user login the first time

Usage:
+ Use counter to keep consistent variable counter avoid dirty read, and control for total participants of campaigns 
+ Use pipeline to pass userid to create voucher, reduce database hits continuous

Plan:
+ Use memcached or redis to store counter data
+ Use other broker queue such as gg pub/sub, kafka instead of
+ Separate monolith to microservice with authentication service, promotion, and more...
### Database
![database](./docs/database.jpg)
+ Status is in active, used, expired to get status voucher of user and show in app list voucher user owner with its status if necessary.

### Sequence 


### Setup infrastructure
+ Run MySQL by docker-compose with init by file sql/init.sql
  
        docker-compose up -d
+ Add environment to get config from file

        CONFIG_PATH=config/local.yaml
+ Run go project
  
        go build -o authen-system