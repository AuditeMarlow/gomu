package template

var ProtoFNC = `syntax = "proto3";

package {{dehyphen .Alias}};

option go_package = "./proto;{{dehyphen .Alias}}";

service {{title .Alias}} {
	rpc Call(CallRequest) returns (CallResponse) {}
}

message CallRequest {
	string name = 1;
}

message CallResponse {
	string msg = 1;
}
`

var ProtoSRV = `syntax = "proto3";

package {{dehyphen .Alias}};

option go_package = "./proto;{{dehyphen .Alias}}";

service {{title .Alias}} {
	rpc Call(CallRequest) returns (CallResponse) {}
	rpc ServerStream(ServerStreamRequest) returns (stream ServerStreamResponse) {}
	rpc BidiStream(stream BidiStreamRequest) returns (stream BidiStreamResponse) {}
}

message CallRequest {
	string name = 1;
}

message CallResponse {
	string msg = 1;
}

message ServerStreamRequest {
	int64 count = 1;
}

message ServerStreamResponse {
	int64 count = 1;
}

message BidiStreamRequest {
	int64 stroke = 1;
}

message BidiStreamResponse {
	int64 stroke = 1;
}
`
