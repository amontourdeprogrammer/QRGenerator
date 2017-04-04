package main 

import (
	"testing"
	"net"
	"bufio"
)


func Testconnection(t *testing.T) {
	go main()
	conn, err := net.Dial("tcp", ":8080")
	if err != nil{
		t.Errorf("No server connection : ", err)
	}	
	statusLine, err := bufio.NewReader(conn).ReadString('\n')
	if statusLine != "HTTP1.0 200 OK" {
		t.Errorf("No content on page : ", statusLine)
	}
}

// todo
// Generate QR code
////Data analysis : Numeric ? alphanumeric ? bytes ? --> Choose a mode / Check if input corresponds to mode
////Data Encoding (http://www.thonky.com/qr-code-tutorial/data-encoding)
//////Step 1: Choose the Error Correction Level
//////Step 2: Determine the Smallest Version for the Data
//////Step 3: Add the Mode Indicator
//////Step 4: Add the Character Count Indicator
//////Step 5: Encode Using the Selected Mode
//////Step 6: Break Up into 8-bit Codewords and Add Pad Bytes if Necessary
////Error Correction Coding
//////Step 1: Break Data Codewords into Blocks if Necessary
//////Step 2: Understand Polynomial Long Division
//////Step 3: Understand The Galois Field
//////Step 4: Understand Galois Field Arithmetic
//////Step 5: Generate Powers of 2 using Byte-Wise Modulo 100011101
//////Step 6: Understand Multiplication with Logs and Antilogs
//////Step 7: Understanding The Generator Polynomial
//////Step 8: Generating Error Correction Codewords
//////Step 9: Divide the Message Polynomial by the Generator Polynomial
////Structure Final Message
//////Step 1: Determine How Many Blocks and Error Correction Codewords are Required
//////Step 2: Interleave the Blocks
//////Step 3: Convert to Binary
//////Step 4: Add Remainder Bits if Necessary
/////Module Placement in Matrix
//////Step 1: Add the Finder Patterns
//////Step 2: Add the Separators
//////Step 3: Add the Alignment Patterns
//////Step 4: Add the Timing Patterns
////Data Masking
////Step 1: Determining the Best Mask
////Step 2: Choose the Lowest Penalty Score for the Eight Mask Patterns
////Add Format and Version Information
//////Step 1: Format String
//////Step 2: Output the Final Matrix

// Set up server
// if get method with data, responds with qr code


