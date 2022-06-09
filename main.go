package main

import (
	"fmt"

	"github.com/stripe/stripe-go/v72"
)

func main() {
	stripe.Key = "sk_test_51ICNuYCJt5NMbUDVfZlTRHfqOvfmq4PcDIZ3F8ISF0tZ2keLcnGho6gY6XioUJ1RcmFJELswrYRvUYwyzt33ogUp0057LXN4he"
	fmt.Println("The stripe key is: ", stripe.Key)
}
