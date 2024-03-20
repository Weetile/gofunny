package main

import (
	"fmt"
	"os"

	"github.com/fogleman/gg"
)

func main() {
	// Load your image
	imgPath := "input.jpg" // Change this to your image file path
	img, err := gg.LoadImage(imgPath)
	if err != nil {
		fmt.Println("Error loading image:", err)
		return
	}

	// Load your font
	fontPath := "font.ttf" // Change this to your font file path
	if _, err := os.Stat(fontPath); os.IsNotExist(err) {
		fmt.Println("Font file does not exist:", err)
		return
	}

	// Calculate the desired padding and text height
	padding := 50    // Adjust this value for desired padding
	textHeight := 72 // Adjust this value based on the font size
	totalPadding := padding + textHeight

	// Create a new drawing context with the increased height
	dc := gg.NewContext(img.Bounds().Max.X, img.Bounds().Max.Y+totalPadding)

	// Draw a white rectangle to fill the padding area
	dc.SetRGB(1, 1, 1) // Set color to white
	dc.DrawRectangle(0, 0, float64(img.Bounds().Max.X), float64(totalPadding))
	dc.Fill()

	// Set font face and size
	if err := dc.LoadFontFace(fontPath, float64(textHeight)); err != nil {
		fmt.Println("Error loading font:", err)
		return
	}

	// Caption text
	caption := "Your caption here"

	// Measure the text dimensions
	tw, _ := dc.MeasureString(caption)

	// Calculate the position to draw the caption
	x := (float64(img.Bounds().Size().X) - tw) / 2 // Center horizontally
	y := float64(padding) + float64(textHeight)/2  // Center vertically within the padding

	// Draw the caption text in black color
	dc.SetRGB(0, 0, 0) // Set text color to black
	dc.DrawString(caption, x, y)

	// Draw the original image onto the context
	dc.DrawImage(img, 0, totalPadding)

	// Save or display the resulting image
	if err := dc.SavePNG("output.png"); err != nil {
		fmt.Println("Error saving image:", err)
		return
	}
	fmt.Println("Image with caption saved as output.png")
}
