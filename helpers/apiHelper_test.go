package helpers

import (
	"log"
	"testing"
)

func TestSendOTpOnMail(t *testing.T) {
	mail, err := SendOTpOnMail("devendrapohekar.siliconithub@gmail.com", "Devendra Pohekar")
	if err != nil {
		log.Print(err)
	}
	log.Print(mail, " ", "successfully work")
}
