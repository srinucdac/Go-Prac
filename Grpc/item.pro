syntax = "proto3";

package main;

message Item {
    string id = 1;
    string name = 2;
}

message ItemRequest {
    Item item = 1;
}

message ItemResponse {
    repeated Item items = 1;
}

service ItemService {
    rpc GetItems (google.protobuf.Empty) returns (ItemResponse);
    rpc CreateItem (ItemRequest) returns (Item);
}
