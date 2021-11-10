### Bubble

##### 启动 mysql server

```bash
docker run --name mysql8019 -p 13306:3306 -e MYSQL_ROOT_PASSWORD=root1234 -v /root/GolandProjects/bubble/mysql:/var/lib/mysql -d mysql:8.0.19
```


##### 启动 mysql client

```bash
docker run -it --network host --rm mysql mysql -h127.0.0.1 -P13306 --default-character-set=utf8mb4 -uroot -proot1234

goreleaser init
goreleaser --snapshot --skip-publish --rm-dist
```
