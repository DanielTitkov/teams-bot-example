# Business Automation Robot for Teams

This is an example service providing business automation functions via Teams with possibily to expand to other messagers. 

Requirements: 
* Golang
* Postgres/Docker
* Registered Teams bot
* ngrok (to connect local version to Teams) 

To run

1. Create dev.yml in /configs with relevant data (use default.yml as an example)
2. `make db` to start db in Docker
3. `make run` to start service
4. `make expose` to expose service endpoint via ngrok