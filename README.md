# SLS

```
package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

func main() {
	// Пример использования
	data := [][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5},
		{6.6, 7.7, 8.8, 9.9},
	}

	// Сохраняем данные в файл
	err := SaveSliceToFile("data.bin", data)
	if err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}

	// Загружаем данные из файла
	loadedData, err := LoadSliceFromFile("data.bin")
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
		return
	}
}
```


Формат файла:

    Первые 4 байта: количество строк (int32)

    Для каждой строки:

        4 байта: количество элементов (int32)

        N × 8 байт: элементы float64

Этот формат эффективен для хранения числовых данных и сохраняет точность значений с плавающей запятой.
