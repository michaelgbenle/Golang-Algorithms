package main

//
//import (
//	"bufio"
//	"fmt"
//	"io"
//	"net"
//	"os"
//	"strings"
//	"testing"
//)
//
//
////TEST READING FROM SERVER
//func TestRead(t *testing.T) {
//	expected := "network is not reading data"
//	var connection net.Conn
//	if Read(connection) != expected {
//		t.Errorf("Expected %v this to equal %s", connection, expected)
//	}
//}
////TEST WRITING TO SERVER
//func TestWrite(t *testing.T) {
//	expected := "network not writing message on platform"
//	var connection net.Conn
//	var clientName string
//	if Write(connection, clientName) != expected {
//		t.Errorf("Expected %v this to equal %s", connection, expected)
//	}
//}
//
//
//func Read(connection net.Conn) string { //read from connection and display in terminal
//	confrmRd := ""
//	if connection != nil {
//		confrmRd += "network is reading data"
//		for {
//			reader := bufio.NewReader(connection)
//			message, err := reader.ReadString('\n')
//			if err == io.EOF {
//				connection.Close()
//				fmt.Println("Connection Closed....")
//				os.Exit(0)
//			}
//			fmt.Println(message)
//			fmt.Println("-----------------------------------")
//		}
//	} else {
//		confrmRd += "network is not reading data"
//	}
//	return confrmRd
//}
//func Write(connection net.Conn, clientName string) string { //write message from the terminal to the connection
//	// the plan is to display message entered by the user to the terminal
//	confrmWrt := ""
//	if connection != nil || clientName != "" {
//		confrmWrt += "network writing message on platform"
//		for {
//			reader := bufio.NewReader(os.Stdin)
//			message, err := reader.ReadString('\n')
//			if err != nil {
//				break
//			}
//			//formatting the string
//			message = fmt.Sprintf("%s:-%s\n", clientName, strings.Trim(message, " \r\n"))
//			connection.Write([]byte(message))
//		}
//	} else {
//		confrmWrt += "network not writing message on platform"
//	}
//	return confrmWrt
//}
