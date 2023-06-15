
### 🌍 A Go library for seamless integration with the Google Translate API.

go-translate provides an intuitive way to translate text between different languages using the powerful Google Translate API. With its simple yet robust design, you can effortlessly incorporate multilingual support into your Go applications. Let's explore some exciting use cases and get started with go-translate!

## Installation
To install go-translate, use go get:

#### ```  go get github.com/sp4reeee/go-translate  ```

## Quick Start

Using go-translate is easy! Here's a simple example to get you started:

```
package main

import (
	"fmt"

	"github.com/sp4reeee/go-translate"
)

func main() {
	resp, err := translate.Translate("hello buddy", "fr")
	if err != nil {
		fmt.Println("Translation error:", err)
		return
	}

	fmt.Println(resp)
}
```
Please note that the country codes used in the above list are based on the ISO 3166-1 alpha-2 codes. These codes are used to uniquely identify each country in international information systems, including the Google Translate API. 😊

The country code for France is "fr" according to the ISO 3166-1 alpha-2 standard. When using the Google Translate API, you can specify this code "fr" to translate text to French or other ISO 3166-1 alpha-2 standard to another language. 🌍

This allows for easy integration with the Google Translate API, enabling you to provide multilingual translation capabilities in your Go applications with remarkable simplicity and accuracy. 🚀

🎉 That's it! With just a few lines of code, you can translate "hello buddy" to French.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

📜 Enjoy translating with go-translate! If you encounter any issues or have suggestions for improvement, please don't hesitate to reach out. Happy coding! 😊

