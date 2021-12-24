package main

import (
	"fmt"
	"github.com/Crec0/Commotion/noise"
)	

func main() {
	xoroshiro := noise.Xoroshiro(928371928731, 138179328713892);
	for i := 0; i < 100; i++ {
		fmt.Println(noise.NextLong(&xoroshiro));
	}
}

