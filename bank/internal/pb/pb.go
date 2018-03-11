package pb

//go:generate protoc -I ../../ --go_out=plugins=grpc,import_path=pb:. ../../bank.proto
