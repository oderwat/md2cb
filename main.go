package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework AppKit
#import <Foundation/Foundation.h>
#import <AppKit/AppKit.h>

void copyToClipboard(const char* html) {
    @autoreleasepool {
        NSString *htmlString = [NSString stringWithUTF8String:html];
        NSData *htmlData = [htmlString dataUsingEncoding:NSUTF8StringEncoding];
        
        NSDictionary *options = @{
            NSDocumentTypeDocumentAttribute: NSHTMLTextDocumentType,
            NSCharacterEncodingDocumentAttribute: @(NSUTF8StringEncoding)
        };
        
        NSAttributedString *attributedString = [[NSAttributedString alloc]
            initWithData:htmlData
            options:options
            documentAttributes:nil
            error:nil];
        
        NSPasteboard *pasteboard = [NSPasteboard generalPasteboard];
        [pasteboard clearContents];
        [pasteboard writeObjects:@[attributedString]];
    }
}
*/
import "C"
import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"unsafe"
)

func convertMarkdownToHTML(markdown string) (string, error) {
	cmd := exec.Command("pandoc", "-f", "markdown", "-t", "html")
	cmd.Stdin = bytes.NewBufferString(markdown)
	
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error running pandoc: %v - %s", err, stderr.String())
	}

	return out.String(), nil
}

func copyToClipboardGo(html string) error {
	// Format as complete HTML document
	doc := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<style>
table { border-collapse: collapse; }
th, td { border: 1px solid black; padding: 8px; }
</style>
</head>
<body>
%s
</body>
</html>`, html)

	cstr := C.CString(doc)
	defer C.free(unsafe.Pointer(cstr))
	C.copyToClipboard(cstr)
	return nil
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	if len(input) == 0 {
		fmt.Fprintf(os.Stderr, "No input received on stdin\n")
		os.Exit(1)
	}

	html, err := convertMarkdownToHTML(string(input))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error converting markdown: %v\n", err)
		os.Exit(1)
	}

	if err := copyToClipboardGo(html); err != nil {
		fmt.Fprintf(os.Stderr, "Error copying to clipboard: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Content copied to clipboard successfully!\n")
}