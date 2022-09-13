# user-service
test code simple microservice, you can see how I write code in go
service to create user, that can be use to create product from product-service (need login first from auth-service)
# Architecture
Controller -> Service -> Repository
# Trigger Jenkins
Triggered when every changed on master branch, jenkins will check every change every minutes(depending on cron job's setting)