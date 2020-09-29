# Uptime Monitoring Service
It is a system which checks whether the requested URL is up or not.

## Tech Stack Used:
- Golang - gin (microframework)
- Mysql
  - Gorm as orm library
- Docker

## Configuration
- Install MySQL on your machine.
- MySQL Connection should be there, having:
```
Hostname: localhost
Port: 3306
Username: root
Password: rootroot
```
- Create a Database named as "UptimeMonitoringService". Use the following command to create it:
```
CREATE DATABASE UptimeMonitoringService;
```

#### OR
- Clone this repository.
- Edit the varialble present in __.env__ file.
- Then follow the instructions in __2.__ or __3.__ in the __Installation__ section below.

## Installation
### 1. Using Docker, Pulling The Image From Dockerhub
- Pull the image first from dockerhub by using the following command:
```
docker image pull sarthakshekhawat/uptime-monitoring-service
```
- Then run it, using this command:
```
docker run -p 8080:8080 sarthakshekhawat/uptime-monitoring-service
```

### 2. Using Docker, Without Pulling The Image From Dockerhub
- Clone this repository and open __Uptime-Monitoring-Service__ directory:
```
cd Uptime-Monitoring-Service
```
- Build:
```
docker build -t uptime-monitoring-service .
```
- Run:
```
docker run -p 8080:8080 uptime-monitoring-service
```

### 3. Without Docker, With Build
- Clone this repository and open __Uptime-Monitoring-Service__ directory:
```
cd Uptime-Monitoring-Service
```
- Then run following commands in the terminal:
```
go mod download
go build .
```
Run:
```
./Uptime-Monitoring-Service
```

### 4. Without Docker, Without Build
- Clone this repository and open __Uptime-Monitoring-Service__ directory:
```
cd Uptime-Monitoring-Service
```
Run:
```
./main
```

## API
### Base URL
```
http://localhost:8080
```

### Add a URL to Monitor:
__POST /urls/__

Request:
```
{
    "url":                 "http://www.abc.com",
    "crawl_timeout":       3,
    "frequency":           5,
    "failure_threshold":   3
}
```
Response:
```
{
    "crawl_timeout":       3,
    "failure_count":       0,
    "failure_threshold":   3,
    "frequency":           5,
    "id":                  "11912758-bbf5-4890-9421-47bcbc7daf40",
    "status":              "active",
    "url":                 "http://www.abc.com"
}
```

### Get URL Information:
__GET /urls/:id__
Response:
```
{
    "crawl_timeout":       3,
    "failure_count":       0,
    "failure_threshold":   3,
    "frequency":           5,
    "id":                  "11912758-bbf5-4890-9421-47bcbc7daf40",
    "status":              "active",
    "url":                 "http://www.abc.com"
}
```

### Update URL Parameters:
__PATCH /urls/:id__

Request:
```
{
    "crawl_timeout":       1,
    "frequency":           10,
    "failure_threshold":   5
}
```
Response:
```
{
    "crawl_timeout":       1,
    "failure_count":       0,
    "failure_threshold":   5,
    "frequency":           10,
    "id":                  "11912758-bbf5-4890-9421-47bcbc7daf40",
    "status":              "active",
    "url":                 "http://www.abc.com"
}
```

### Activate URL:
__POST /urls/:id/activate__

Response:
```
{
    "crawl_timeout":       1,
    "failure_count":       0,
    "failure_threshold":   5,
    "frequency":           10,
    "id":                  "11912758-bbf5-4890-9421-47bcbc7daf40",
    "status":              "active",
    "url":                 "http://www.abc.com"
}
```

### Deactivate URL:
__POST /urls/:id/deactivate__

Response:
```
{
    "crawl_timeout":       1,
    "failure_count":       0,
    "failure_threshold":   5,
    "frequency":           10,
    "id":                  "11912758-bbf5-4890-9421-47bcbc7daf40",
    "status":              "inactive",
    "url":                 "http://www.abc.com"
}
```

### Delete URL:
__DELETE /urls/:id__
