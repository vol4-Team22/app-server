# Mikke-Server
> [!CAUTION]
> 実運用する際は鍵を新たに生成して、秘密鍵は絶対に後悔しないでください。
## Start the server
Create a Docker image in advance.
```zsh
$ make build-local
```
Use Docker Compose to start each service.
```zsh
$ make up
```
Perform a migration to MySQL.
```zsh
$ make migrate
```