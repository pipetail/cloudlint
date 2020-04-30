package worker

import (
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/pipetail/cloudlint/internal/pkg/check"
	"github.com/pipetail/cloudlint/internal/pkg/checkcompleted"
	"github.com/pipetail/cloudlint/internal/pkg/checkreportstarted"
	log "github.com/sirupsen/logrus"

	"github.com/jedib0t/go-pretty/table"
)

// Handle function
func Handle() check.Result {

	reportID := uuid.New().String()
	outputReport := checkreportstarted.New(reportID)

	result := check.Result{}

	var wg sync.WaitGroup // create waitgroup (empty struct)

	queue := make(chan checkcompleted.Check, len(outputReport.Payload.Checks))

	for _, val := range outputReport.Payload.Checks {

		wg.Add(1)

		newmsg := check.New(reportID)

		//newmsg.Payload.AWSAuth = rcvdEvent.Payload.AWSAuth

		newmsg.Payload.CheckID = val.ID

		newmsg.Payload.CheckType = val.Type

		go concurrentHandler(newmsg, &wg, queue)
		res := checkcompleted.Check{
			ID:       val.ID,
			Severity: 1,
			Impact:   123,
		}
		//res := handler(newmsg)
		result.CheckResult = append(result.CheckResult, res)
		result.CheckInfo = append(result.CheckInfo, val)
	}

	wg.Wait() // blocks here

	for i := range result.CheckResult {
		item := <-queue
		result.CheckResult[i].Impact = item.Impact
		result.CheckResult[i].Severity = item.Severity
		//fmt.Println("[main]: msg %v %v ", item, res)
	}

	return result
}

// Print prints the report in a pretty table output
func Print(res check.Result) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Group", "Name", "Impact $", "Severity"})

	totalImpact := 0
	for i, result := range res.CheckResult {
		//fmt.Printf("%s: %+v %+v\n", i, res.CheckInfo[i].ID, result)
		t.AppendRow([]interface{}{res.CheckInfo[i].Type, res.CheckInfo[i].Group, res.CheckInfo[i].Name, result.Impact, checkcompleted.Severity(result.Severity)})
		totalImpact += result.Impact
	}

	t.AppendFooter(table.Row{"", "", "Total", totalImpact})
	t.Render()

}

func concurrentHandler(message check.Event, wg *sync.WaitGroup, c chan<- checkcompleted.Check) {

	outputReport := handler(message)

	c <- outputReport.Payload.Check

	log.WithFields(log.Fields{
		"report": outputReport,
	}).Info("report finished")

	wg.Done()
}

func handler(message check.Event) *checkcompleted.Event {

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
