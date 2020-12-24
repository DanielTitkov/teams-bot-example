# Business Automation Robot for Teams

This is an example service providing business automation functions via Teams with possibily to expand to other messagers. 

Requirements: 
* Golang
* Postgres/Docker
* Registered Teams bot
* ngrok (to connect local version to Teams) 

To run

`make db` to start db in Docker
`make run` to start service
`make expose` to expose service endpoint via ngrok