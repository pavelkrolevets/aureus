## Open source cryptocurrency and derivatives exchange.

### Architecture
<p align="center"><img width="100%" src="https://oooooo.oss-cn-hangzhou.aliyuncs.com/gitbitex.png?v=1" /></p>

### Dependencies
* MySql (**BINLOG[ROW format]** enabled)
* Kafka
* Redis

### Install
#### Server
* git clone https://github.com/pavelkrolevets/aureus.git
* Create database `aureus` and make sure **BINLOG[ROW format]** enabled
* Execute ddl.sql
* Modify conf.json
* Run go build
* Run ./aureus

#### Web
* git clone https://github.com/pavelkrolevets/aureus-web.git
* Run `npm install`
* Configure settings to 127.0.0.1 to run locally
* Run `npm start`
* Run `npm run build` to build production

#### Referenced and respect
This project is a fork and continue of the original [gitbitex](https://github.com/gitbitex/gitbitex-web) work.