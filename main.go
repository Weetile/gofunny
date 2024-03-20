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

	// Create a new drawing context
	dc := gg.NewContextForImage(img)

	// Set font face and size
	if err := dc.LoadFontFace(fontPath, 72); err != nil { // Change 24 to your desired font size
		fmt.Println("Error loading font:", err)
		return
	}

	// Caption text
	caption := "Your caption here"

	// Measure the text dimensions
	tw, th := dc.MeasureString(caption)

	// Calculate the position to draw the caption
	imgWidth := float64(img.Bounds().Size().X)
	x := (imgWidth - tw) / 2 // Center horizontally
	y := th                  // Place above the image with some padding

	// Draw the caption text in black color
	dc.SetRGB(0, 0, 0) // Set text color to black
	dc.DrawString(caption, x, y)

	// Save or display the resulting image
	if err := dc.SavePNG("output.png"); err != nil {
		fmt.Println("Error saving image:", err)
		return
	}
	fmt.Println("Image with caption saved as output.png")
}
