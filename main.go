package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
)

func main() {
	text := flag.String("text", "", "Текст или vCard для генерации QR-кода")
	output := flag.String("output", "pixqry.png", "Имя выходного файла PNG")
	logoPath := flag.String("logo", "logo.png", "Путь к PNG логотипу")
	size := flag.Int("size", 512, "Размер QR-кода (пиксели)")
	flag.Parse()

	if *text == "" {
		log.Fatal("Ошибка: необходимо указать --text для генерации QR-кода")
	}

	err := generateQRCodeWithLogo(*text, *output, *logoPath, *size)
	if err != nil {
		log.Fatalf("Ошибка при генерации PixQRy: %v", err)
	}

	fmt.Println("✅ PixQRy успешно создан:", *output)
}

func generateQRCodeWithLogo(content, outPath, logoPath string, size int) error {
	// Создаём QR-код с уровнем коррекции Medium (15%)
	qr, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return err
	}
	qr.DisableBorder = true

	qrImg := qr.Image(size)

	// Открываем логотип
	logoFile, err := os.Open(logoPath)
	if err != nil {
		return fmt.Errorf("не удалось открыть логотип: %w", err)
	}
	defer logoFile.Close()

	logoImg, err := png.Decode(logoFile)
	if err != nil {
		return fmt.Errorf("не удалось декодировать логотип PNG: %w", err)
	}

	// Масштабируем логотип — 20% от размера QR
	logoSize := size / 5
	logoResized := resize.Resize(uint(logoSize), 0, logoImg, resize.Lanczos3)

	// Вычисляем позицию для центрирования логотипа
	offset := image.Pt(
		(qrImg.Bounds().Dx()-logoResized.Bounds().Dx())/2,
		(qrImg.Bounds().Dy()-logoResized.Bounds().Dy())/2,
	)

	// Создаём новое изображение и копируем QR
	finalImg := image.NewRGBA(qrImg.Bounds())
	draw.Draw(finalImg, qrImg.Bounds(), qrImg, image.Point{0, 0}, draw.Src)

	// Накладываем логотип сверху с альфа-композицией
	draw.Draw(finalImg, logoResized.Bounds().Add(offset), logoResized, image.Point{0, 0}, draw.Over)

	// Сохраняем результат в PNG
	outFile, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return png.Encode(outFile, finalImg)
} 