module github.com/YoungsoonLee/practical-go/ch08/user-service/client-json

go 1.17

replace github.com/YoungsoonLee/practical-go/ch08/user-service/service => ../service

require (
	github.com/YoungsoonLee/practical-go/ch08/user-service/service v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/golang/protobuf v1.5.0 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)
