package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fogleman/gg"
)

func main() {
	// Define flags
	imageFlag := flag.String("image", "input.jpg", "Path to the input image")
	fontFlag := flag.String("font", "font.ttf", "Path to the font file")
	paddingFlag := flag.Int("padding", 50, "Padding above the image")
	textHeightFlag := flag.Int("textHeight", 72, "Height of the caption text")

	// Parse flags
	flag.Parse()

	// Load the image
	img, err := gg.LoadImage(*imageFlag)
	if err != nil {
		fmt.Println("Error loading image:", err)
		return
	}

	// Load the font
	if _, err := os.Stat(*fontFlag); os.IsNotExist(err) {
		fmt.Println("Font file does not exist:", err)
		return
	}

	// Calculate the total padding
	totalPadding := *paddingFlag + *textHeightFlag

	// Create a new drawing context with the increased height
	dc := gg.NewContext(img.Bounds().Max.X, img.Bounds().Max.Y+totalPadding)

	// Draw a white rectangle to fill the padding area
	dc.SetRGB(1, 1, 1) // Set color to white
	dc.DrawRectangle(0, 0, float64(img.Bounds().Max.X), float64(totalPadding))
	dc.Fill()

	// Set font face and size
	if err := dc.LoadFontFace(*fontFlag, float64(*textHeightFlag)); err != nil {
		fmt.Println("Error loading font:", err)
		return
	}

	// Caption text
	caption := "Your caption here"

	// Measure the text dimensions
	tw, _ := dc.MeasureString(caption)

	// Calculate the position to draw the caption
	x := (float64(img.Bounds().Size().X) - tw) / 2 // Center horizontally
	y := float64(*paddingFlag) + float64(*textHeightFlag)/2   // Center vertically within the padding

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

