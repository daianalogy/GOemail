package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"strings"
)

func main() {

	envVars := []string{"SET_SMTP", "SET_EMAIL", "SET_PASS"}
	var decodedValues []string
	var encodedValues []string

	for _, envVar := range envVars {
		value, exists := os.LookupEnv(envVar)
		if !exists {
			fmt.Printf("%s is undefined\nDo you want to set a value to %s: (y/n): ", envVar, envVar)
			var answer string
			fmt.Scan(&answer)
			if strings.ToLower(answer) == "y" {
				fmt.Printf("Enter value to %s: ", envVar)
				var value string
				fmt.Scan(&value)
				encodedValue := base64.StdEncoding.EncodeToString([]byte(value))
				encodedValues = append(encodedValues, encodedValue)
				decodedValues = append(decodedValues, value)

			} else {
				fmt.Printf("Program will be terminated, $%s still undefined. ", envVar)
				os.Exit(1)
			}

		} else {
			decodedValue, err := base64.StdEncoding.DecodeString(value)
			if err != nil {
				// Si el valor no se puede decodificar, se asume que aún no ha sido codificado
				encodedValue := base64.StdEncoding.EncodeToString([]byte(value))
				os.Setenv(envVar, encodedValue)
				decodedValues = append(decodedValues, value)
			} else {
				// Si el valor se puede decodificar, se asume que ya está codificado y se añade a la lista.
				decodedValues = append(decodedValues, string(decodedValue))
			}
		}

	}

	if len(encodedValues) != 0 {
		fmt.Println("Please, export the following environment variables and restart the programm: ")
	}

	for i, envVar := range envVars {
		if len(encodedValues) != 0 {
			envar, val := envVar, encodedValues[i]
			fmt.Printf("%s=%s ", envar, val)

		} else {
			continue
		}
	}

	if len(decodedValues) < 3 {
		log.Fatal("Vars SET_SMTP, SET_EMAIL y SET_PASS not found in environment variables,  please set value a export")

	}

	// Establecer la configuración del servidor SMTP
	smtpServer := strings.TrimSpace(decodedValues[0])
	email := strings.TrimSpace(decodedValues[1])
	password := strings.TrimSpace(decodedValues[2])
	auth := smtp.PlainAuth("", email, password, smtpServer)
	//fmt.Printf("%v", decodedValues)

	file, err1 := os.Open("email/emails.txt")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer file.Close()

	// Leer el archivo de texto plano
	scanner := bufio.NewScanner(file)
	var to []string
	for scanner.Scan() {
		email := strings.TrimSpace(scanner.Text())
		to = append(to, email)
	}

	// Leer el archivo HTML
	htmlBytes, err2 := ioutil.ReadFile("email/index.html")
	if err2 != nil {
		log.Fatal("Error while reading HTML:", err2)
	}

	// Crear el correo electrónico

	subject := "Test, auto mails with Go"
	htmlBody := string(htmlBytes)
	msg := []byte("To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"utf-8\"\r\n" +
		"\r\n" +
		htmlBody)
	
		// Enviar el correo electrónico
	err := smtp.SendMail(smtpServer+":587", auth, email, to, msg)
	if err != nil {
		log.Fatal(" Ops! There was an error while attempting to send [ERROR] ", err)
	}
	log.Println("Yay! Email sent succesfully")
}
