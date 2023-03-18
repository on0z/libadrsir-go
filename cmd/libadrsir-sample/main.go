package main

import (
	"fmt"
	"log"

	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	host "periph.io/x/host/v3"

	libadrsir "github.com/on0z/libadrsir-go"
)

func main() {
	// host.Init() registers all the periph-provided host driver automatically,
	// so it is preferable to use than periph.Init().
	//
	// You can also use periph.io/x/extra/hostextra.Init() for additional drivers
	// that depends on cgo and/or third party packages.

	// state, err := host.Init()
	_, err := host.Init()
	if err != nil {
		log.Fatalf("failed to initialize periph: %v", err)
	}

	/*
		periph.io example

		// Prints the loaded driver.
		fmt.Printf("Using drivers:\n")
		for _, driver := range state.Loaded {
			fmt.Printf("- %s\n", driver)
		}

		// Prints the driver that were skipped as irrelevant on the platform.
		fmt.Printf("Drivers skipped:\n")
		for _, failure := range state.Skipped {
			fmt.Printf("- %s: %s\n", failure.D, failure.Err)
		}

		// Having drivers failing to load may not require process termination. It
		// is possible to continue to run in partial failure mode.
		fmt.Printf("Drivers failed to load:\n")
		for _, failure := range state.Failed {
			fmt.Printf("- %s: %v\n", failure.D, failure.Err)
		}
	*/

	// Use pins, buses, devices, etc.

	// Use i2creg I²C bus registry to find the first available I²C bus.
	b, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()

	// Dev is a valid conn.Conn.
	d := &i2c.Dev{Addr: uint16(libadrsir.ADDR), Bus: b}

	adrsir := libadrsir.NewADRSIR(d)
	errCh := make(chan error)
	go func() {
		errCh <- adrsir.Send("00002800D00029003900160038001600120016001300160012001700120016001300160012001700380016001200170012001600130016001200170012001600130016003800160013001600380016001300160012001700120016001300160012001700120016001300160012001700120016001300160012001700120016001300160012001700120016003900160012001600390016003800160012001600390016003800160011004205")
	}()
	if err := <-errCh; err != nil {
		fmt.Println(err)
	}
}
