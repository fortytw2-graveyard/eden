echo "preparing database for test"
export DB_USER="eden"
export DB_PASSWORD="password=eden"
export DB_NAME="eden"
export DB_EXTRA="sslmode=disable"
export REDIS_URL="localhost:6379"

$GOPATH/bin/goose down
$GOPATH/bin/goose up
redis-cli flushall

go test -v ./...

echo "clearing database after tests"

redis-cli flushall
$GOPATH/bin/goose down
$GOPATH/bin/goose up
