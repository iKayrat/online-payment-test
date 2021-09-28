# Online-payment-test
It is a simple go application, which includes logic of generating signature for FONDY.


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



```go
// Set signature convert struct to map and 
func (c *CheckouRequest) SetSignature(password string) {
	params := structs.Map(c)
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	values := []string{}
	for _, key := range keys {
		value := params[key].(string)
		if value == "" {
			continue
		}
		values = append(values, value)
	}

	c.Signature = generateSignature(values, password)
}
```
