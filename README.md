service-notification-api

A repository for the  notification service api being developed 
for ant investors

### How do I update the definitions? ###

* The api definition is defined in the proto file notification.proto
* To update the proto service you need to run the command :
    
    `protoc -I ./ ./notification.proto --go_out=./`
    `protoc -I ./ ./notification.proto --go-grpc_out=./`

    with that in place update the implementation appropriately
