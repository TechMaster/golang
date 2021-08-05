package main

import (
	"fmt"

	"demofiber/eris"

	"github.com/gofiber/fiber/v2"
)

// Chuyên xử lý các err mà controller trả về
func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	var statusCode = 500

	if e, ok := err.(*eris.Error); ok {
		handleEris(e)
		if e.Code > 0 { // Mặc định là 500, nếu e.Code > 0 thì gán vào statusCode
			statusCode = e.Code
		}
	} else if e, ok := err.(*fiber.Error); ok { //Thử kiểm tra xem có phải là kiểu fiber.Error không
		statusCode = e.Code
		fmt.Println(err.Error())
	} else {
		fmt.Println(err.Error())
	}

	if err = ctx.Render("error/error", fiber.Map{
		"ErrorMessage": err.Error(),
		"StatusCode":   statusCode,
	}); err != nil {
		return ctx.Status(500).SendString("Internal Server Error")
	}

	return nil
}

//Hàm chuyên xử lý Eris Error có Stack Trace
func handleEris(err *eris.Error) {
	formattedStr := eris.ToCustomString(err, eris.StringFormat{
		Options: eris.FormatOptions{
			InvertOutput: true, // flag that inverts the error output (wrap errors shown first)
			WithTrace:    true, // flag that enables stack trace output
			InvertTrace:  true, // flag that inverts the stack trace output (top of call stack shown first)
			Skip:         3,    // Bỏ qua 3 dòng lệnh cuối cùng trong Stack
		},
		MsgStackSep:  "\n",  // separator between error messages and stack frame data
		PreStackSep:  "\t",  // separator at the beginning of each stack frame
		StackElemSep: " | ", // separator between elements of each stack frame
		ErrorSep:     "\n",  // separator between each error in the chain
	})

	colorReset := string("\033[0m")
	colorRed := string("\033[31m")
	//Chỗ này log ra console
	if err.IsPanic() {
		fmt.Println(colorRed, formattedStr, colorReset)
		//Lỗi Panic và Error nhất thiết phải ghi vào file !
	} else {
		fmt.Println(formattedStr)
	}
}
