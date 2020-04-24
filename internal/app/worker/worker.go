package worker

import (
	"context"
	"os"

	"github.com/pipetail/cloudlint/internal/pkg/check"
	"github.com/pipetail/cloudlint/internal/pkg/checkcompleted"
	"github.com/pipetail/cloudlint/internal/pkg/checkreportstarted"
	log "github.com/sirupsen/logrus"

	"github.com/jedib0t/go-pretty/table"
)

// Handle function
func Handle() check.Result {

	reportID := "awioavaovao"
	outputReport := checkreportstarted.New(reportID)

	result := check.Result{}

	for _, val := range outputReport.Payload.Checks {
		newmsg := check.New(reportID)

		//newmsg.Payload.AWSAuth = rcvdEvent.Payload.AWSAuth

		newmsg.Payload.CheckID = val.ID

		newmsg.Payload.CheckType = val.Type

		//go handler(nil, newmsg)
		res := handler(nil, newmsg)
		result.CheckResult = append(result.CheckResult, res.Payload.Check)
		result.CheckInfo = append(result.CheckInfo, val)
	}

	return result
}

func Print(res check.Result) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Group", "Name", "Type", "Impact $", "Severity"})

	totalImpact := 0
	for i, result := range res.CheckResult {
		//fmt.Printf("%s: %+v %+v\n", i, res.CheckInfo[i].ID, result)
		t.AppendRow([]interface{}{res.CheckInfo[i].ID, res.CheckInfo[i].Group, res.CheckInfo[i].Name, res.CheckInfo[i].Type, result.Impact, checkcompleted.Severity(result.Severity)})
		totalImpact += result.Impact
	}

	t.AppendFooter(table.Row{"", "", "", "Total", totalImpact})
	t.Render()

}

func handler(ctx context.Context, message check.Event) *checkcompleted.Event {

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
		return nil
	}

	log.WithFields(log.Fields{
		"report": outputReport,
	}).Info("report finished")

	return outputReport
}
