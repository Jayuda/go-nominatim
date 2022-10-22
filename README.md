# go-nominatim

simple golang gogpslocation (gps) calculations

```
package main

import(
    "https://github.com/Jayuda/go-nominatim"
    "fmt"
)

func main() {
    ReverseData := gonominatim.GPSToAddress("https://nominatim.server.com", "-6.09182031", "101.09128301")
    Address := ReverseData.Address
    fmt.Println(Address)
}
```