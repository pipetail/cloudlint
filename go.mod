module github.com/pipetail/cloudlint

go 1.12

require (
	github.com/aws/aws-sdk-go v1.30.13
	github.com/getlantern/deepcopy v0.0.0-20160317154340-7f45deb8130a // indirect
	github.com/go-openapi/strfmt v0.19.3 // indirect
	github.com/google/go-cmp v0.4.0 // indirect
	github.com/google/uuid v1.1.1
	github.com/jedib0t/go-pretty v4.3.0+incompatible
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5 // indirect
	go.mongodb.org/mongo-driver v1.1.2 // indirect
	golang.org/x/net v0.0.0-20200226121028-0de0cce0169b // indirect
	golang.org/x/sys v0.0.0-20200122134326-e047566fdf82 // indirect
	golang.org/x/text v0.3.2 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace github.com/pipetail/cloudlint/pkg/check v0.0.0 => ../pkg/check

replace github.com/pipetail/cloudlint/internal/app/worker v0.0.0 => ../internal/app/worker

replace github.com/pipetail/cloudlint/pkg/aws v0.0.0 => ../pkg/aws
