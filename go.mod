module gin-mongo-demo

go 1.14

replace github.com/coreos/bbolt v1.3.5 => go.etcd.io/bbolt v1.3.5

replace go.etcd.io/bbolt v1.3.5 => github.com/coreos/bbolt v1.3.5

replace google.golang.org/grpc v1.28.0 => google.golang.org/grpc v1.26.0

replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.14.1

require (
	github.com/coreos/bbolt v1.3.5 // indirect
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-log/log v0.2.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/mock v1.4.4 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0 // indirect
	github.com/jonboulle/clockwork v0.2.1 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/lucas-clemente/quic-go v0.18.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/miekg/dns v1.1.31 // indirect
	github.com/nats-io/jwt v1.0.1 // indirect
	github.com/nats-io/nats-server/v2 v2.1.8 // indirect
	github.com/nats-io/nkeys v0.2.0 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200427203606-3cfed13b9966 // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	go.mongodb.org/mongo-driver v1.4.1
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/sys v0.0.0-20200917073148-efd3b9a0ff20 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.org/x/tools v0.0.0-20200917050209-655488c8ae71 // indirect
	google.golang.org/genproto v0.0.0-20200122232147-0452cf42e150 // indirect
	google.golang.org/grpc v1.26.0 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	honnef.co/go/tools v0.0.1-2020.1.5 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)
