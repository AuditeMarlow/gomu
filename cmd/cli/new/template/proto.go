package template

var ProtoFNC = `syntax = "proto3";

package {{dehyphen .Alias}};

option go_package = "./proto;{{dehyphen .Alias}}";

service {{title .Alias}} {
	rpc Call(Request) returns (Response) {}
}

message Request {
	string name = 1;
}

message Response {
	string msg = 1;
}
`

var ProtoSRV = `syntax = "proto3";

package {{dehyphen .Alias}};

option go_package = "./proto;{{dehyphen .Alias}}";

service {{title .Alias}} {
	rpc Call(Request) returns (Response) {}
	rpc Stream(StreamRequest) returns (stream StreamResponse) {}
}

message Request {
	string name = 1;
}

message Response {
	string msg = 1;
}

message StreamRequest {
	int64 count = 1;
}

message StreamResponse {
	int64 count = 1;
}
`
