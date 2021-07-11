# TaxiFare by @jerryalfs

Please run command below, before run this project

```bash
go mod tidy
```
try to do
```bash
go build
```
to make sure project can run well
and before that I want to explain, why I put name in parameter
not only time and mileage, because I want to categorize user
and also store it in redis as a key prefix
after try go build and success
you can run binary with this command
```bash
./taxiFare
```
or
```bash
go run app.go
```
but dont forget must on specific path directory same path with app.go
and i do using redis as database in cache because it's fast and easy to use
easy to install or uninstall, easy to check also.

can use this command to install redis
```bash
wget http://download.redis.io/redis-stable.tar.gz
tar xvzf redis-stable.tar.gz
cd redis-stable
make
```
and type
```bash
redis-cli
```
you can try to hit this API to see the result, since running on localhost
you can hit to localhost:9000 or 127.0.0.1:9000.
for port it's changeable in app.go
```bash
localhost:9000/taxifare/detail
```
can use postman or curl with post method
this sample request body
```bash
{
    "timemileage" : "00:12:00.000 1000",
    "name" : "jerry"
}
```
sample curl
```bash
curl -H "Content-Type: application/json" --data '{"timemileage" : "00:12:00.000 1000","name" : "jerry"}' localhost:9000/taxifare/detail
```
