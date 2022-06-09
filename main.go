package main

import (
	"fmt"
	"net/http"

	"github.com/stripe/stripe-go/v72"
)

func main() {
	stripe.Key = "sk_test_51ICNuYCJt5NMbUDVfZlTRHfqOvfmq4PcDIZ3F8ISF0tZ2keLcnGho6gY6XioUJ1RcmFJELswrYRvUYwyzt33ogUp0057LXN4he"
	fmt.Println("The stripe key is: ", stripe.Key)

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/create-checkout-session", createCheckoutSession)
	addr := "localhost:4242"
	http.ListenAndServe(addr, nil)
}

func createCheckoutSession(w http.ResponseWriter, r *http.Request) {
	domain := "http://localhost:4242"
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				//provide the exact Price_ID of the product you want to sell.
				Price: stripe.String("price_1L8NiZCJt5NMbUDVvGCaAkeX"),
				Quantity: stripe.Int64(1),
			},
		},
		Mode: stripe.String(string(stripe.))
	}


}
