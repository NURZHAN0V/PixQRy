# 🎨 PixQRy — Генератор QR-кодов с логотипом

Привет! 👋 PixQRy — это простое и удобное приложение для создания красивых QR-кодов с вашим логотипом. Работает полностью оффлайн, не требует интернета и API.

## ✨ Возможности

- 🎯 Создание QR-кодов с встроенным логотипом
- 🎨 Поддержка PNG логотипов
- 📱 Настраиваемый размер QR-кода
- 💼 Поддержка vCard для визиток
- 🔒 Работает полностью оффлайн
- 🎯 Высокое качество выходного изображения

## 🚀 Быстрый старт

### 1. Установка Go

Сначала установите Go с [официального сайта](https://golang.org/dl/).

### 2. Установка PixQRy

```bash
# Клонируйте репозиторий
git clone https://github.com/yourusername/pixqry.git
cd pixqry

# Установите зависимости
go mod tidy
```

### 3. Подготовка логотипа

1. Подготовьте квадратный PNG-файл с вашим логотипом
2. Назовите его `logo.png`
3. Поместите в корень проекта
4. Рекомендуемый размер: 256×256 пикселей

### 4. Создание QR-кода

```bash
# Простой QR-код со ссылкой
go run main.go --text "https://your-website.com" --output my_qr.png

# QR-код с визиткой (vCard)
go run main.go --text "BEGIN:VCARD
VERSION:3.0
N:Иванов;Иван
FN:Иван Иванов
ORG:Компания
TITLE:Должность
TEL:+1234567890
EMAIL:ivan@example.com
END:VCARD" --output business_card.png
```

## ⚙️ Настройка

| Параметр   | Описание                             | По умолчанию |
|------------|--------------------------------------|--------------|
| `--text`   | Текст для кодирования (обязательно)  | -            |
| `--output` | Имя выходного файла PNG              | pixqry.png   |
| `--logo`   | Путь к PNG логотипу                  | logo.png     |
| `--size`   | Размер QR-кода в пикселях            | 512          |

## 🎯 Примеры использования

### 1. Создание QR-кода для сайта
```bash
go run main.go --text "https://example.com" --output website_qr.png
```

### 2. Создание QR-кода для визитки
```bash
go run main.go --text "BEGIN:VCARD
VERSION:3.0
N:Иванов;Иван
FN:Иван Иванов
ORG:Компания
TITLE:Разработчик
TEL:+1234567890
EMAIL:ivan@example.com
END:VCARD" --output business_card.png
```

### 3. Создание QR-кода с кастомным размером
```bash
go run main.go --text "https://example.com" --size 1024 --output large_qr.png
```

## 🛠️ Технические детали

- Использует уровень коррекции ошибок Medium (15%)
- Логотип автоматически масштабируется до 20% от размера QR-кода
- Поддерживает прозрачность PNG
- Оптимизирован для печати

## 🤝 Вклад в проект

Мы приветствуем ваши предложения по улучшению! Если вы хотите внести свой вклад:

1. Форкните репозиторий
2. Создайте ветку для ваших изменений
3. Отправьте pull request

## 📝 Лицензия

MIT License — используйте свободно!

## 💡 Подсказки

- Используйте высококачественные логотипы для лучшего результата
- Для печати рекомендуется использовать размер не менее 512 пикселей
- Тестируйте QR-код перед массовой печатью
- Сохраняйте оригинальный логотип в хорошем качестве

## 🆘 Поддержка

Если у вас возникли вопросы или проблемы:
- Создайте issue в репозитории
- Напишите на email: your-email@example.com

---

Сделано с ❤️ для сообщества разработчиков