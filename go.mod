module github.com/pipetail/cloudlint

go 1.14

require (
	github.com/aws/aws-sdk-go v1.30.13
	github.com/google/uuid v1.1.1
	github.com/jedib0t/go-pretty/v6 v6.0.4
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.0.0-20200226121028-0de0cce0169b // indirect
	golang.org/x/sys v0.0.0-20200122134326-e047566fdf82 // indirect
	golang.org/x/text v0.3.2 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace github.com/pipetail/cloudlint/pkg/check v0.0.0 => ../pkg/check

replace github.com/pipetail/cloudlint/internal/app/worker v0.0.0 => ../internal/app/worker

replace github.com/pipetail/cloudlint/pkg/aws v0.0.0 => ../pkg/aws
