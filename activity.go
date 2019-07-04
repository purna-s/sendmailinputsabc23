package sendmailinputsabc23

import (

	"fmt"
	"log"
	"strings"
	"net/smtp"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-sendmailinputsabc23")

// MyActivity is a stub for your Activity implementation
type sendmailinputsabc23 struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &sendmailinputsabc23{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *sendmailinputsabc23) Metadata() *activity.Metadata {
	return a.metadata
}


// Eval implements activity.Activity.Eval
func (a *sendmailinputsabc23) Eval(ctx activity.Context) (done bool, err error) {

	arcpnt := ctx.GetInput("Arcpnt").(string)
	bsub := ctx.GetInput("Bsub").(string)
	cbody := ctx.GetInput("Cbody").(string)
	wsHost := context.GetInput("server").(string)
	
	
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"sendalertsforq@gmail.com",
		"ptcxejoylzgtrfmh",
		"smtp.gmail.com",
	)
	
	t := []string{"To:", arcpnt}
	strings.Join(t, " ")
	
	s := []string{"Subject:", bsub}
	strings.Join(s, " ")
	
	
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	
	to := []string{arcpnt}
	msg := []byte(strings.Join(t, " ") + "\r\n" + strings.Join(s, " ") + "\r\n" + cbody + "\r\n")
	
	err = smtp.SendMail("smtp.gmail.com:587", auth, "sendalertsforq@gmail.com", to, msg)
	if err != nil {
		activityLog.Debugf("Error ", err)
		fmt.Println("error: ", err)
		return
	}
	
	fmt.Println("Mail Sent")
	log.Println("Mail Sent")


	// Set the output as part of the context
	activityLog.Debugf("Activity has sent the mail Successfully")
	fmt.Println("Activity has sent the mail Successfully")

	ctx.SetOutput("output", "Mail_Sent_Successfully")

	return true, nil
}
