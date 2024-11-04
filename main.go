package main

import (
	"fmt"
	"math/rand"
	"os"
)


func main() {
    var amount int
    var prefix string
    var name string
    var email string
    var expirationDate string
    var value float64
    
    fmt.Print("amount of vouchers :")
    fmt.Scanln(&amount)
    
    fmt.Print("voucher values : ")
    fmt.Scanln(&value)

    fmt.Print("prefix :")
    fmt.Scanln(&prefix)
    
    fmt.Print("name :")
    fmt.Scanln(&name)
    
    fmt.Print("email :")
    fmt.Scanln(&email)

    fmt.Print("expirationDate :")
    fmt.Scanln(&expirationDate)

    
    codes := make([]string, amount)
    for i := 0; i < amount; i++ {
        codes[i] = generateCode(prefix)
    }

    var csvCode string = printCsvCode(codes, value, name, email, expirationDate)
    
    os.WriteFile(prefix+".csv", []byte(csvCode), 0644)
}


func printCsvCode(codes []string, value float64, name string, email string, expirationDate string) string {
    var csv string = "codes,value,expirationDate,name,email/phone\n"
    
    for i := 0; i < len(codes); i++ {
        csv += codes[i] + ","
        csv += fmt.Sprint(value) + ","
        csv += expirationDate + ","
        csv += name + ","
        csv += email + "\n"
    }
    
    return csv
}

func generateCode(prefix string) string {
    var code string = randStringBytes(10)
    return prefix + "-" + code
}

func randStringBytes(length int) string { 
    var runes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    
    bytes := make([]byte, length)

    for index := range bytes {
        bytes[index] = runes[rand.Intn(len(runes))]
    }

    return string(bytes)
}
