package worker

import (
	"context"
	"fmt"

	"github.com/pipetail/cloudlint/internal/pkg/check"
	"github.com/pipetail/cloudlint/internal/pkg/checkcompleted"
	"github.com/pipetail/cloudlint/internal/pkg/checkreportstarted"
	log "github.com/sirupsen/logrus"
)

func Printhello() {
	fmt.Println("helloaaa")
}

func Handle() {

	reportId := "awioavaovao"
	outputReport := checkreportstarted.New(reportId)

	for _, val := range outputReport.Payload.Checks {
		newmsg := check.New(reportId)

		//newmsg.Payload.AWSAuth = rcvdEvent.Payload.AWSAuth

		newmsg.Payload.CheckID = val.ID

		newmsg.Payload.CheckType = val.Type
		//msg, _ := json.Marshal(newmsg)
		handler(nil, newmsg)
	}

}

func handler(ctx context.Context, message check.Event) (string, error) {

	//message := check.Event{}

	var err error

	// initialize CheckCompleted event
	var outputReport *checkcompleted.Event

	// do the checks here
	switch message.Payload.CheckType {
	case check.GetChecks()[0].Type:
		outputReport, err = billingInfo(message)
	case check.GetChecks()[1].Type:
		outputReport, err = dmsUnused(message)
	case check.GetChecks()[2].Type:
		outputReport, err = ebsunused(message)
	case check.GetChecks()[3].Type:
		outputReport, err = paidSupport(message)
	case check.GetChecks()[4].Type:
		outputReport, err = ebsSnapshotsOld(message)
	case check.GetChecks()[5].Type:
		outputReport, err = elbUnused(message)
	case check.GetChecks()[6].Type:
		outputReport, err = vpcendpoints(message)
	case check.GetChecks()[7].Type:
		outputReport, err = natgwunused(message)
	case check.GetChecks()[8].Type:
		outputReport, err = eipunused(message)
	case check.GetChecks()[9].Type:
		outputReport, err = amiOld(message)
	case check.GetChecks()[10].Type:
		outputReport, err = ebsopt(message)
	default:
		log.WithFields(log.Fields{
			"err": err,
		}).Error("received new check request")
	}
	if err != nil {
		return "", err
	}

	// send report to SNS
	//err = sendReport(outputReport, "arn:aws:sns:eu-central-1:680177765279:result")
	if err != nil {
		return "", err
	}

	log.WithFields(log.Fields{
		"report": outputReport,
	}).Info("report finished")

	return "ok", nil

}
