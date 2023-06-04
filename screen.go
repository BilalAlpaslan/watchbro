package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/kbinani/screenshot"
)

func takeScreenshot() {
	bounds := image.Rect(-500, -100, 2000, 1200) // Ekran boyutu

	img, err := screenshot.CaptureRect(bounds) // Ekranı yakala
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ekran boyutuu:", bounds.Dx(), "x", bounds.Dy())

	file, err := os.Create("screenshot.png") // Dosyaya kaydet
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("Ekran görüntüsü kaydediliyor...")

	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ekran görüntüsü başarıyla kaydedildi.")
}
