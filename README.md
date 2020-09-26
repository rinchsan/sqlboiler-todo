# sqlboiler-todo

Sample Todo Application written in Go

```shell script
make up
docker-compose exec mysql bash -c "mysql -uroot -ppassword db < var/tmp/ddl.sql"

curl "localhost:8080"
curl "localhost:8080/users"
curl "localhost:8080/users" -X POST -d '{"username":"rinchsan"}'
curl "localhost:8080/users" -X PUT -d '{"id":3, "username":"rinchsan"}'
curl "localhost:8080/todos"
curl "localhost:8080/todos" -X POST -d '{"title":"todo title", "detail":"todo detail", "due_date":"2020-08-20T00:00:00Z", "author_user_id":3, "assignee_user_ids":[2, 3]}'
curl "localhost:8080/todos" -X PUT -d '{"id":1, "title":"new title", "detail":"new detail", "due_date":"2020-08-31T00:00:00Z", "author_user_id":3, "assignee_user_ids":[2, 3]}'
```
