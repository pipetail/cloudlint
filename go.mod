module github.com/pipetail/cloudlint

go 1.12

require (
	github.com/Azure/azure-sdk-for-go v41.3.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.10.0 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.8.3 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/aws/aws-lambda-go v1.13.3
	github.com/aws/aws-sdk-go v1.28.9
	github.com/evanphx/json-patch v4.5.0+incompatible // indirect
	github.com/getlantern/deepcopy v0.0.0-20160317154340-7f45deb8130a // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/golang/protobuf v1.4.0 // indirect
	github.com/google/uuid v1.1.1
	github.com/hashicorp/go-hclog v0.12.2 // indirect
	github.com/hashicorp/go-plugin v1.2.2 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/motemen/go-quickfix v0.0.0-20200118031250-2a6e54e79a50 // indirect
	github.com/motemen/gore v0.5.0 // indirect
	github.com/peterh/liner v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.5.1 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v1.0.0 // indirect
	github.com/vmware-tanzu/velero v1.3.2 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/tools v0.0.0-20200422205258-72e4a01eba43 // indirect
	google.golang.org/grpc v1.29.0 // indirect
	k8s.io/apiextensions-apiserver v0.18.2 // indirect
	k8s.io/cli-runtime v0.18.2 // indirect
	k8s.io/client-go v11.0.0+incompatible // indirect
)

replace github.com/pipetail/cloudlint/internal/pkg/check v0.0.0 => ../../internal/pkg/check

// replace github.com/pipetail/cloudlint/internal/pkg/checkcompleted v0.0.0 => ./internal/pkg/checkcompleted

// replace github.com/pipetail/cloudlint/internal/pkg/checkreportstarted v0.0.0 => ./internal/pkg/checkreportstarted

// replace github.com/pipetail/cloudlint/internal/pkg/checkreport v0.0.0 => ./internal/pkg/checkreport

// replace github.com/pipetail/cloudlint/internal/pkg/checkawsintegration v0.0.0 => ./internal/pkg/checkawsintegration

// replace github.com/pipetail/cloudlint/internal/pkg/checkawsintegrationcompleted v0.0.0 => ./internal/pkg/checkawsintegrationcompleted

// replace github.com/pipetail/cloudlint/internal/pkg/awsregions v0.0.0 => ./internal/pkg/awsregions

replace github.com/pipetail/cloudlint/internal/app/worker v0.0.0 => ../../internal/app/worker
