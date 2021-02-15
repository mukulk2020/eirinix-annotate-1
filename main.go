package main

import (
	"fmt"
	"os"

	eirinix "github.com/SUSE/eirinix"
	annotate "eirinix-annotate/annotate"
)

func main() {
	fmt.Println("Running eirinix-annotate...")
	options := eirinix.ManagerOptions{
		Namespace:           os.Getenv("POD_NAMESPACE"),
		Host:                "0.0.0.0",
		Port:                4545,
		ServiceName:         os.Getenv("WEBHOOK_SERVICE_NAME"),
		WebhookNamespace:    os.Getenv("WEBHOOK_NAMESPACE"),
	}
	fmt.Printf("--> %#v\n", options)
	x := eirinix.NewManager(options)
	x.AddExtension(&annotate.Extension{})
	x.Start()
}
