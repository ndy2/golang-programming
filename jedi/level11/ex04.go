package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error : %v, %v, %v", se.lat, se.long, se.err)
}

func main() {

	_, err := sqrt(-10.0)
	if err != nil {
		log.Println(err) // "2009/11/10 23:00:00 math error : , , f is negative"
	}

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := errors.New("f is negative")
		return 0, sqrtError{"", "", e}
	}

	return 42, nil
}
