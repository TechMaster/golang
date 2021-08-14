package email

import (
	"demofiber/db"

	"github.com/TechMaster/eris"
)

func SendEmail(to string, subject string, content string) error {
	// Try to connect email server but failed
	return eris.Panic("Failed to email server") //Lỗi trả về có Stack Trace
}

func Foo() error {
	if err := db.Bar(); err != nil {
		return err
	}
	return nil
}

func GotMail() error {
	if err := db.Query2(); err != nil {
		return eris.Wrap(err, "Failed to Query")
	}
	return nil
}
