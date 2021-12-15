module github.com/YoungsoonLee/practical-go/ch08/multiple-services/server

go 1.17

require (
	github.com/YoungsoonLee/practical-go/ch08/multiple-services/service v0.0.0
	google.golang.org/grpc v1.37.0
)

require (
	github.com/golang/protobuf v1.4.3 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/YoungsoonLee/practical-go/ch08/multiple-services/service v0.0.0 => ../service
