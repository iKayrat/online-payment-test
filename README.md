# Online-payment-test
It is a simple go application, which includes logic of generating transaction with signature for FONDY.


```go
// request interface
checkouReq := &CheckouRequest{
		OrderID:    id.String(),
		MerchantId: merchantId,
		// tourism tickets
		OrderDesc: "Expedition Tour",
		Amount:    "5000",
		Currency:  currencyEUR,
		// https://
		ServeCallbackUrl: "",
	}
```
