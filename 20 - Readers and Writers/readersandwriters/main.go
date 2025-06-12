package main

import (
	"fmt"
	"io"
	"strings"
)

func processData(reader io.Reader, Writer io.Writer) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			Writer.Write(b[0:count])
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}
func CopingData(reader io.Reader, writer io.Writer) {
	count, err := io.Copy(writer, reader)
	if err == nil {
		Printfln("Read : %v bytes", count)
	} else {
		Printfln("Read : %v", err.Error())
	}
}
func scanFormReader(reader io.Reader, template string, vals ...interface{}) (int, error) {
	return fmt.Fscanf(reader, template, vals...)
}
func ScanData() {
	reader := strings.NewReader("Kayak Watersports $279.00")
	var name, category string
	var price float64
	scanTemplate := " %s , %s,%f"
	_, err := scanFormReader(reader, scanTemplate, &name, &category, &price)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Name: %v", name)
		Printfln("Category: %v", category)
		Printfln("Price: %.2f", price)
	}
}
func main() {
	//r := strings.NewReader("Kayak")
	//var builder strings.Builder
	//processData(r, &builder)
	//CopingData(r, &builder)
	//Printfln("String builder contents: %s", builder.String())
	//pipeReader, pipeWriter := io.Pipe()
	//go func() {
	//	GenanrateData(pipeWriter)
	//	pipeWriter.Close()
	//}()
	//ConsumeData(pipeReader)
	ScanData()
}
