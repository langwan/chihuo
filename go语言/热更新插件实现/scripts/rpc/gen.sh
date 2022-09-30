echo "生成 rpc server 代码"

OUT=../../pb
protoc \
--go_out=${OUT} \
--go-grpc_out=${OUT} \
--go-grpc_opt=require_unimplemented_servers=false \
language.proto






