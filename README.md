# libadrsir-go
golang製 ビットトレードワン赤外線送受信機 ADRSIR 用のライブラリ

i2cのライブラリはperiphのものを使用することを想定している

# How to use

see more: https://github.com/on0z/libadrsir-go/blob/main/cmd/libadrsir-sample/main.go

```golang
package main

import (
	"log"

	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	host "periph.io/x/host/v3"

	libadrsir "github.com/on0z/libadrsir-go"
)

func main() {
  // setup periph.io host
	_, err := host.Init()
	if err != nil {
		log.Fatalf("failed to initialize periph: %v", err)
	}
  
  // Use i2creg I²C bus registry to find the first available I²C bus.
	b, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()

	// Dev is a valid conn.Conn.
	d := &i2c.Dev{Addr: uint16(libadrsir.ADDR), Bus: b}

	adrsir := libadrsir.NewADRSIR(d)
	adrsir.Send("00002800D00029003900160038001600120016001300160012001700120016001300160012001700380016001200170012001600130016001200170012001600130016003800160013001600380016001300160012001700120016001300160012001700120016001300160012001700120016001300160012001700120016001300160012001700120016003900160012001600390016003800160012001600390016003800160011004205")
}
```


# 参考
https://github.com/tokieng/adrsirlib/blob/master/adrsirlib.py

https://bit-trade-one.co.jp/blog/2017121302/
