package main

import (
    "fmt"
    "errors"
)

func intro_errors(){
    r, e := fe(100)
    fmt.Println("Result :", r)
    fmt.Println("Error :", e)

    r, e = fe(10)
    fmt.Println("Result :", r)
    fmt.Println("Error :", e)

    r, e = fae(100)
    fmt.Println("Result :", r)
    fmt.Println("Error :", e)


}

func fe(x int) (int, error){
    if x == 100 {
        return -1, errors.New("Can not take 100")
    }

    return x, nil
}

type argError struct {
    arg int
    prob string
}

func (ae *argError) Error() string {
    return fmt.Sprintf("Error %d - %s", ae.arg, ae.prob)
}

func fae(x int) (int, error){
    if x  == 100 {
        return -1, &argError{x, "can not take this argument"}
    }
    return x, nil
}

