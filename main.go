package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fogleman/gg"
)

func main() {
	// Define flags
	imageFlag := flag.String("image", "input.jpg", "Path to the image file")
	fontFlag := flag.String("font", "font.ttf", "Path to the font file")
	captionFlag := flag.String("caption", "Hello, World!", "Caption text")
	paddingFlag := flag.Int("padding", 50, "Padding value")
	flag.Parse()

	// Load your image
	imgPath := *imageFlag
	img, err := gg.LoadImage(imgPath)
	if err != nil {
		fmt.Println("Error loading image:", err)
		return
	}

	// Load your font
	fontPath := *fontFlag
	if _, err := os.Stat(fontPath); os.IsNotExist(err) {
		fmt.Println("Font file does not exist:", err)
		return
	}

	// Calculate the desired text height proportional to 1/8 of the image height
	textHeight := int(float64(img.Bounds().Dy()) / 8)

	// Use the padding from the flag
	padding := *paddingFlag

	// Calculate the total padding
	totalPadding := padding + textHeight

	// Define the caption
	caption := *captionFlag

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

	// Measure the text dimensions
	tw, th := dc.MeasureString(caption)

	// Calculate the position to draw the caption
	x := (float64(img.Bounds().Size().X) - tw) / 2 // Center horizontally
	// Adjust y coordinate to center vertically within the padding area
	y := float64(totalPadding)/2 + float64(th)/2

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

