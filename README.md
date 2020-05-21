<p align="center"><img width="40%" src="https://getbitex.oss-cn-beijing.aliyuncs.com/projects/image/logo.svg" /></p>

[![Go Report Card](https://goreportcard.com/badge/github.com/pavelkrolevets/gitbitex-spot)](https://goreportcard.com/report/github.com/pavelkrolevets/gitbitex-spot)

GitBitEx is an open source cryptocurrency and derivatives exchange.

## Architecture
<p align="center"><img width="100%" src="https://oooooo.oss-cn-hangzhou.aliyuncs.com/gitbitex.png?v=1" /></p>

## Dependencies
* MySql (**BINLOG[ROW format]** enabled)
* Kafka
* Redis

## Install
### Server
* git clone https://github.com/pavelkrolevets/gitbitex-spot.git
* Create database and make sure **BINLOG[ROW format]** enabled
* Execute ddl.sql
* Modify conf.json
* Run go build
* Run ./gitbitex-spot

### Web
* git clone https://github.com/gitbitex/gitbitex-web.git
* Run `npm install`
* Configure settings to 127.0.0.1 to run locally
* Run `npm start`
* Run `npm run build` to build production

#### Referenced and respect
This project is continue of the original gitbitex work.