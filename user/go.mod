module micros/user

go 1.13

require (
	github.com/DataDog/datadog-go v0.0.0-20190315133836-a5d50a065561 // indirect
	github.com/SAP/go-hdb v0.14.1 // indirect
	github.com/StackExchange/wmi v0.0.0-20181212234831-e0a55b97c705 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20190320094055-98404a78c009 // indirect
	github.com/coredns/coredns v1.4.0 // indirect
	github.com/envoyproxy/go-control-plane v0.6.9 // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/gocql/gocql v0.0.0-20190319151216-6bdac5e86117 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/gogo/googleapis v1.1.0 // indirect
	github.com/golang/protobuf v1.3.1
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/hashicorp/go-discover v0.0.0-20190319153616-61771d82ff54 // indirect
	github.com/hashicorp/hil v0.0.0-20190212132231-97b3a9cdfa93 // indirect
	github.com/hashicorp/net-rpc-msgpackrpc v0.0.0-20151116020338-a14192a58a69 // indirect
	github.com/hashicorp/raft-boltdb v0.0.0-20171010151810-6e5ba93211ea // indirect
	github.com/hashicorp/vault v1.1.0 // indirect
	github.com/hashicorp/vault-plugin-auth-alicloud v0.0.0-20190320211238-36e70c54375f // indirect
	github.com/hashicorp/vault-plugin-auth-azure v0.0.0-20190320211138-f34b96803f04 // indirect
	github.com/hashicorp/vault-plugin-auth-centrify v0.0.0-20190320211357-44eb061bdfd8 // indirect
	github.com/hashicorp/vault-plugin-auth-gcp v0.0.0-20190320214413-e8308b5e41c9 // indirect
	github.com/hashicorp/vault-plugin-auth-jwt v0.0.0-20190320205321-c6ec63d2528c // indirect
	github.com/hashicorp/vault-plugin-auth-kubernetes v0.0.0-20190320210228-426b5188498e // indirect
	github.com/hashicorp/vault-plugin-secrets-ad v0.0.0-20190320211735-127b63e898e6 // indirect
	github.com/hashicorp/vault-plugin-secrets-alicloud v0.0.0-20190320213517-3307bdf683cb // indirect
	github.com/hashicorp/vault-plugin-secrets-azure v0.0.0-20190320211922-2dc8a8a5e490 // indirect
	github.com/hashicorp/vault-plugin-secrets-gcp v0.0.0-20190320211452-71903323ecb4 // indirect
	github.com/hashicorp/vault-plugin-secrets-gcpkms v0.0.0-20190320213325-9e326a9e802d // indirect
	github.com/hashicorp/vault-plugin-secrets-kv v0.0.0-20190320211621-3ccc8684cf25 // indirect
	github.com/jinzhu/gorm v1.9.11
	github.com/lyft/protoc-gen-validate v0.0.14 // indirect
	github.com/micro/go-config v1.1.0
	github.com/micro/go-grpc v1.0.0
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.0.0
	github.com/mitchellh/pointerstructure v0.0.0-20170205204203-f2329fcfa9e2 // indirect
	github.com/shirou/gopsutil v2.18.12+incompatible // indirect
	github.com/shirou/w32 v0.0.0-20160930032740-bb4de0191aa4 // indirect
	github.com/softlayer/softlayer-go v0.0.0-20180806151055-260589d94c7d
	github.com/ugorji/go/codec v0.0.0-20190320090025-2dc34c0b8780 // indirect
	golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c
	gopkg.in/ini.v1 v1.42.0 // indirect
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 // indirect
	micros/common v0.0.0-00010101000000-000000000000 // indirect
	sigs.k8s.io/structured-merge-diff v0.0.0-20190302045857-e85c7b244fd2 // indirect
)

replace (
	micros/common => ../common
	micros/product => ../product
	micros/user => ../user
)

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.9

replace github.com/golang/lint => golang.org/x/lint v0.0.0-20190930215403-16217165b5de
