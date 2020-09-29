package writer

import (
	"encoding/json"
	"nn/data"
	"os"
	"path/filepath"
)

type FileWriter struct {
	Folder string
}

func (f *FileWriter) Write(records []*data.DayRecord) error {
	if len(records) == 0 {
		return nil
	}
	bys, err := json.Marshal(&records)
	if err != nil {
		return err
	}
	record := records[0]
	file, err := os.OpenFile(filepath.Join(f.Folder, record.Code+".json"), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(bys)
	return err
}
