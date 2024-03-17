rd /s /q   .\common\rpc\clients\
mkdir .\common\rpc\clients\

goctl.exe rpc protoc --go-grpc_out=.\common\rpc\pb --go_out=.\common\rpc\pb --zrpc_out=apps\admin  apps\admin\proto\admin.proto
MOVE .\apps\admin\adminclient .\common\rpc\clients\


protoc-go-inject-tag -input=".\common\rpc\pb\*\*.pb.go"