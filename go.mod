module partyApp

go 1.12

require (
	cloud.google.com/go v0.44.0
	github.com/golang/protobuf v1.3.2
	github.com/julienschmidt/httprouter v1.2.0
	github.com/spf13/viper v1.4.0
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 // indirect
	golang.org/x/net v0.0.0 // indirect
	golang.org/x/sys v0.0.0-20190804053845-51ab0e2deafa // indirect
	golang.org/x/tools v0.0.0-20190809145639-6d4652c779c4 // indirect
)

replace golang.org/x/net v0.0.0 => github.com/golang/net v0.0.0-20190724013045-ca1201d0de80
