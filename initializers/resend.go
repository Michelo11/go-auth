package initializers

import (
	"os"

	"github.com/resendlabs/resend-go"
)

var Resend *resend.Client

func ResendMail() {
	apiKey := os.Getenv("RESEND_API_KEY")

	Resend = resend.NewClient(apiKey)
}
