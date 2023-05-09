
# Kavenegar-OTP 

Kavenegar OTP in Golang


## Installation


```bash
go get -u github.com/rootiens/kavenegar-otp
```
    

## Usage/Examples


```go
package main

import (
    kotp "github.com/rootiens/kavenegar-otp"
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

    if err := kotp.Send(otp); err != nil {
        // log errors here
    }
}

```





