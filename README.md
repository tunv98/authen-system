# Promotion System

This repository contains the codes of the Authentication System and Promotion System.

Require:
+ User login with account and password, account is username or email or phone number
+ Register a profile: full name, phone number, email, user name, password, birthday, latest login and 3 of them can not null
+ Support 100 first login users per campaign when clients register a new account on the system will get a 30% discount voucher 

Feature:
+ Sign in 
  + Store information includes full name, phone number, email, username, password, birthday, latest login
  + Password hashed by bcrypt
+ Log in
  + With one phone number, email, or username and password 
  + Response JWT token and give token in the header to pass dashboard
+ Access resource 
  + With /homepage to use middleware for authentication (*check /api/user.http for request*) 
+ Add voucher for a user if user login the first time and based on campaigns
  + Number of users who received vouchers will be limited by total_vouchers(campaigns)
  + Use latest_login is nil to check user login the first time

Usage:
+ Use a counter to keep a consistent variable counter to avoid dirty reads and control for total participants of campaigns 
+ Use pipeline to pass userid to create voucher, reduce database hits continuous

Plan:
+ Use memcached or redis to store counter data
+ Use another broker queue such as gg pub/sub, Kafka instead of
+ Separate monolith to microservice with authentication service, promotion, and more...
+ Sharding database for users if optimize big queries, reduce hit db
+ Add metrics following parameters as memory, CPU, latency => Set auto scale (HPA) if necessary
+ If voucher delivery is for a large number of users, use Redis to store the data cache and use Bitmap to check users in the list to create vouchers then 
+ Add unit test, benchmark, and load testing
### Sequence 
![sequence](./docs/sequence.jpg)

### Database
![database](./docs/database.jpg)
+ Status is active, used, or expired to get a status voucher of the user and show in the app list voucher user owner with its status if necessary.

### Setup infrastructure
+ Run MySQL by docker-compose with init by file sql/init.sql
  
        docker-compose up -d
+ Add environment to get config from file

        CONFIG_PATH=config/local.yaml
+ Run go project
  
        go build -o authen-system
