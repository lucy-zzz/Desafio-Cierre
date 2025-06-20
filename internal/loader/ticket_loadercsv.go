package loader

import (
	"app/internal"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func NewLoaderTicketCSV(filePath string) *LoaderTicketCSV {
	return &LoaderTicketCSV{
		filePath: filePath,
	}
}

type LoaderTicketCSV struct {
	filePath string
}

func (tl *LoaderTicketCSV) Load() (t map[int]internal.TicketAttributes, err error) {
	f, err := os.Open(tl.filePath)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return t, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	t = make(map[int]internal.TicketAttributes)
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			err = fmt.Errorf("error reading record: %v", err)

			return t, err
		}

		idStr := record[0]
		id, _ := strconv.Atoi(idStr)
		price, _ := strconv.ParseFloat(record[5], 64)
		ticket := internal.TicketAttributes{
			Name:    record[1],
			Email:   record[2],
			Country: record[3],
			Hour:    record[4],
			Price:   price,
		}

		t[id] = ticket
	}

	return
}
