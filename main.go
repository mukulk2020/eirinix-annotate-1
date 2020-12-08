package main

import (
	"fmt"
	"os"

	eirinix "github.com/SUSE/eirinix"
	annotate "github.com/kansal-mukul/eirinix-annotate/annotate"
)

func main() {
	fmt.Println("Running kansal-mukul/eirinix-annotate...")
	options := eirinix.ManagerOptions{
		Namespace:           os.Getenv("POD_NAMESPACE"),
		Host:                "0.0.0.0",
		Port:                4545,
		ServiceName:         os.Getenv("WEBHOOK_SERVICE_NAME"),
		WebhookNamespace:    os.Getenv("WEBHOOK_NAMESPACE"),
		OperatorFingerprint: "eirini-x-kansal-mukul-annotate"
	}
	fmt.Printf("--> %#v\n", options)
	x := eirinix.NewManager(options)
	x.AddExtension(&annotate.Extension{})
	x.Start()
}
