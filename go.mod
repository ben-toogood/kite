module github.com/ben-toogood/kite

go 1.15

require (
	cloud.google.com/go/storage v1.6.0
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef // indirect
	github.com/cloudflare/cfssl v1.5.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/form3tech-oss/jwt-go v3.2.2+incompatible
	github.com/friendsofgo/graphiql v0.2.2
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/gofrs/uuid v3.3.0+incompatible // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.3 // indirect
	github.com/google/martian v2.1.1-0.20190517191504-25dcb96d9e51+incompatible // indirect
	github.com/google/uuid v1.1.2
	github.com/graph-gophers/graphql-go v0.0.0-20201113091052-beb923fada29
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/lib/pq v1.8.0 // indirect
	github.com/lileio/lile/v2 v2.2.14-0.20200617145758-67284c38c56e
	github.com/lileio/pubsub v0.0.0-20180730130251-70c350806efc
	github.com/lileio/pubsub/v2 v2.3.1
	github.com/maxbrunsfeld/counterfeiter/v6 v6.3.0 // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/common v0.6.0 // indirect
	github.com/segmentio/ksuid v1.0.3
	github.com/sendgrid/rest v2.6.2+incompatible
	github.com/sendgrid/sendgrid-go v3.7.2+incompatible
	github.com/sirupsen/logrus v1.7.0
	github.com/stretchr/testify v1.6.1
	github.com/tj/assert v0.0.3
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9 // indirect
	golang.org/x/sys v0.0.0-20201130171929-760e229fe7c5 // indirect
	golang.org/x/text v0.3.4 // indirect
	google.golang.org/api v0.29.0 // indirect
	google.golang.org/genproto v0.0.0-20201119123407-9b1e624d6bc4 // indirect
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.1-0.20201020201750-d3470999428b
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.7
	honnef.co/go/tools v0.0.1-2020.1.4 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.27.1
