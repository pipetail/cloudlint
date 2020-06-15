package worker

import (
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/pipetail/cloudlint/pkg/check"
	"github.com/pipetail/cloudlint/pkg/checkcompleted"
	"github.com/pipetail/cloudlint/pkg/checkreportstarted"
	log "github.com/sirupsen/logrus"

	"github.com/jedib0t/go-pretty/table"
)

// Handle function
func Handle(filterChecks []string) check.Result {

	reportID := uuid.New().String()
	outputReport := checkreportstarted.New(reportID).WithFilter(filterChecks)

	result := check.Result{}

	var wg sync.WaitGroup // create waitgroup (empty struct)

	queue := make(chan checkcompleted.Check, len(outputReport.Payload.Checks))

	log.WithFields(log.Fields{
		"report": outputReport,
	}).Debug("report started")

	for _, val := range outputReport.Payload.Checks {

		wg.Add(1)

		newmsg := check.New(reportID)

		//newmsg.Payload.AWSAuth = rcvdEvent.Payload.AWSAuth

		newmsg.Payload.CheckID = val.ID

		newmsg.Payload.CheckType = val.Type

		go concurrentHandler(newmsg, &wg, queue)
		res := checkcompleted.Check{
			ID:       val.ID,
			Severity: checkcompleted.INFO,
			Impact:   0,
		}
		//res := handler(newmsg)
		result.CheckResult = append(result.CheckResult, res)
		result.CheckInfo = append(result.CheckInfo, val)
	}

	wg.Wait() // blocks here

	for range result.CheckResult {
		item := <-queue

		idx := indexOf(item.ID, result.CheckResult)
		if idx != -1 {
			result.CheckResult[idx].Impact = item.Impact
			result.CheckResult[idx].Severity = item.Severity
		}
	}

	return result
}

func indexOf(element string, data []checkcompleted.Check) int {
	for k, v := range data {
		if element == v.ID {
			return k
		}
	}
	return -1 //not found.
}

// Print prints the report in a pretty table output
func Print(res check.Result) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Group", "Name", "Impact $", "Severity"})

	totalImpact := 0
	billIdx := -1

	for i, result := range res.CheckResult {
		//fmt.Printf("%s: %+v %+v\n", i, res.CheckInfo[i].ID, result)
		if res.CheckInfo[i].Type == "aws_monthly_bill" {
			billIdx = i
			continue
		}
		t.AppendRow([]interface{}{res.CheckInfo[i].Type, res.CheckInfo[i].Group, res.CheckInfo[i].Name, result.Impact, checkcompleted.Severity(result.Severity)})
		totalImpact += result.Impact
	}

	t.AppendFooter(table.Row{"", "", "Total impact", totalImpact})
	if billIdx >= 0 {
		t.AppendFooter(table.Row{"", "", res.CheckInfo[billIdx].Name, res.CheckResult[billIdx].Impact})
	}
	t.Render()

}

func concurrentHandler(message check.Event, wg *sync.WaitGroup, c chan<- checkcompleted.Check) {

	outputReport := Handler(message)

	if outputReport != nil {
		c <- outputReport.Payload.Check
	} else {
		log.WithFields(log.Fields{
			"message": message,
		}).Error("report finished with error")
		c <- checkcompleted.New(message.Payload.CheckID).Payload.Check
	}

	log.WithFields(log.Fields{
		"report": outputReport,
	}).Info("report finished")

	wg.Done()
}

// Handler handles the one incoming check and returns the result
func Handler(message check.Event) *checkcompleted.Event {

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
	case check.GetChecks()[11].Type:
		outputReport, err = datatransferhuge(message)
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
