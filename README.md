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
- Clone this repository and open the "Uptime-Monitoring-Service":
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
- Clone this repository and open the "Uptime-Monitoring-Service":
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

### 3. Without Docker, Without Build
- Clone this repository and open the "Uptime-Monitoring-Service":
```
cd Uptime-Monitoring-Service
```
Run:
```
./main
```
