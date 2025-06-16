# 🧙‍♂️ Проект **PixQRy** — оффлайн генератор QR-кодов с логотипом

---

## Структура проекта

```
pixqry/
├── .gitignore
├── go.mod
├── main.go
├── logo.png         # Ваш логотип (квадратный PNG)
```

---

## 1. `.gitignore`

```gitignore
# Go артефакты и временные файлы
bin/
*.exe
*.test
*.out

# Зависимости
vendor/

# IDE
.idea/
.vscode/
*.swp

# Системные файлы
.DS_Store
Thumbs.db

# Выходные файлы QR-кодов
qr_*.png
*.svg
*.pdf
```

---

## 2. `go.mod`

```go
module pixqry

go 1.21

require (
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/skip2/go-qrcode v1.2.0
)
```

Запустить после создания:

```bash
go mod tidy
```

---

## 3. `main.go`

```go
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
```

---

## 4. Использование

### Шаг 1: Подготовь логотип

* Помести квадратный PNG-файл `logo.png` в корень проекта.
* Рекомендуемый размер логотипа — 256×256 px.

---

### Шаг 2: Запусти сборку и генерацию

```bash
go mod tidy
go run main.go --text "https://pixqry.com" --output pixqry.png --logo logo.png --size 512
```

---

### Шаг 3: Используй с vCard

Пример для визитки (vCard):

```bash
go run main.go --text "BEGIN:VCARD
VERSION:3.0
N:Иванов;Иван
FN:Иван Иванов
ORG:PixQRy Inc.
TITLE:Разработчик
TEL:+1234567890
EMAIL:ivan@pixqry.example.com
END:VCARD" --output ivan_pixqry.png --logo logo.png
```

---

## 5. Опции CLI

| Параметр   | Описание                             | По умолчанию |
| ---------- | ------------------------------------ | ------------ |
| `--text`   | Текст для кодирования (обязательный) |              |
| `--output` | Имя выходного файла PNG              | `pixqry.png` |
| `--logo`   | Путь к PNG логотипу                  | `logo.png`   |
| `--size`   | Размер QR в пикселях                 | `512`        |

---

## 6. Что внутри PixQRy

* Создаёт QR-код с уровнем коррекции Medium (чтобы логотип не ломал считывание)
* Вставляет логотип в центр (20% от размера QR)
* Выходной файл — PNG с высоким разрешением
* Работает полностью оффлайн, не требует интернета и API