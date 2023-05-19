
# Kavenegar-OTP 

Kavenegar OTP in Golang


## Installation


```bash
go get -u github.com/rootiens/kavenegarotp
```
    

## Usage/Examples


```go
package main

import (
	"fmt"

	kotp "github.com/rootiens/kavenegarotp"
)

func main() {
    otp := kotp.Otp{
        ApiKey: "ASDFGHJKLQWERTYUIOPZXCVBNM",
        Receptor: "0912345678",
        Type: "sms",
        Tokens: []kotp.Token{
            {Key: "token", Value: "12345"},
        },
        Template: "hi",
    }

    response, err := kotp.Send(otp)
    err != nil {
        // log errors here
    }
    fmt.Println(response)
}
```

#### Example response

```json
{
    "return": {
        "status": 200,
        "message": "تایید شد"
    },
    "entries": [
     {
        "messageid": 8792343,
        "message": "ممنون از ثبت نام شما کد تایید عضویت  : 852596	",
        "status": 5,
        "statustext": "ارسال به مخابرات",
        "sender": "10004346",
        "receptor": "09361234567",
        "date": 1356619709,
        "cost": 120
     }
  ]
}
```





