curl http://localhost:8080/ -X POST -d '{}'
echo "\n"
curl http://localhost:8080/ -X POST -d '{"name":"gevin"}'
echo '\n'
curl http://localhost:8080/ -X POST -d '{"name":"gevin", "age":20}'