package assignmentbiz

import (
	"context"
	"fmt"
	"os"
	assignmentrecognizeprovider "server/libs/assignment_recognize_provider"

	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

type recognizeAssignmentBiz struct {
	provider assignmentrecognizeprovider.Provider
}

func NewRecognizeAssignmentBiz(provider assignmentrecognizeprovider.Provider) *recognizeAssignmentBiz {
	return &recognizeAssignmentBiz{provider: provider}
}

func (biz *recognizeAssignmentBiz) RecognizeAssignment(context context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	err := outputPdfText("./dethi.pdf")
	if err != nil {
		fmt.Println(err)
	}

	return biz.provider.RecognizeAssignment(context, data)
}

// outputPdfText prints out contents of PDF file to stdout.
func outputPdfText(inputPath string) error {
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	defer f.Close()

	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		return err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	fmt.Printf("--------------------\n")
	fmt.Printf("PDF to text extraction:\n")
	fmt.Printf("--------------------\n")
	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return err
		}

		ex, err := extractor.New(page)
		if err != nil {
			return err
		}

		text, err := ex.ExtractText()
		if err != nil {
			return err
		}

		fmt.Println("------------------------------")
		fmt.Printf("Page %d:\n", pageNum)
		fmt.Printf("\"%s\"\n", text)
		fmt.Println("------------------------------")
	}

	return nil
}
