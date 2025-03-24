package sls

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

// SaveSliceToFile сохраняет слайс [][]float64 в файл
func SaveSliceToFile(filename string, data [][]float64) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Записываем количество строк
	rows := int32(len(data))
	if err := binary.Write(file, binary.LittleEndian, rows); err != nil {
		return err
	}

	for _, row := range data {
		// Записываем количество элементов в строке
		cols := int32(len(row))
		if err := binary.Write(file, binary.LittleEndian, cols); err != nil {
			return err
		}

		// Записываем сами элементы как float64
		for _, num := range row {
			if err := binary.Write(file, binary.LittleEndian, num); err != nil {
				return err
			}
		}
	}

	return nil
}

// LoadSliceFromFile загружает слайс [][]float64 из файла
func LoadSliceFromFile(filename string) ([][]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rows int32
	if err := binary.Read(file, binary.LittleEndian, &rows); err != nil {
		return nil, err
	}

	if rows < 0 {
		return nil, errors.New("invalid row count")
	}

	data := make([][]float64, rows)

	for i := 0; i < int(rows); i++ {
		var cols int32
		if err := binary.Read(file, binary.LittleEndian, &cols); err != nil {
			return nil, err
		}

		if cols < 0 {
			return nil, errors.New("invalid column count")
		}

		row := make([]float64, cols)
		for j := 0; j < int(cols); j++ {
			var num float64
			if err := binary.Read(file, binary.LittleEndian, &num); err != nil {
				return nil, err
			}
			row[j] = num
		}
		data[i] = row
	}

	return data, nil
}
