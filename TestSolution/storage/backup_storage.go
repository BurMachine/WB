package storage

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"os"
)

func (s *Storage) BackupCreate() error {

	jsn, err := json.Marshal(s.data)
	if err != nil {
		return err
	}

	compressedData, err := gZipData(jsn)
	if err != nil {
		return err
	}
	err = os.WriteFile("dump.gz", compressedData, 0777)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) BackupLoad() error {
	file, err := os.Open("dump.gz")
	a := err.Error()
	println(a)
	if err != nil {
		if err.Error() != "open dump.gz: no such file or directory" {
			return err
		} else {
			return nil
		}
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	resData, err := gUnzipData(data)
	if err != nil {
		return err
	}
	m := make(map[string]*UserGrade)
	err = json.Unmarshal(resData, &m)
	s.data = m
	return nil
}
func gUnzipData(data []byte) (resData []byte, err error) {
	b := bytes.NewBuffer(data)

	var r io.Reader
	r, err = gzip.NewReader(b)
	if err != nil {
		return
	}

	var resB bytes.Buffer
	_, err = resB.ReadFrom(r)
	if err != nil {
		return
	}

	resData = resB.Bytes()

	return
}

func gZipData(data []byte) (compressedData []byte, err error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)

	_, err = gz.Write(data)
	if err != nil {
		return
	}

	if err = gz.Flush(); err != nil {
		return
	}

	if err = gz.Close(); err != nil {
		return
	}

	compressedData = b.Bytes()

	return
}
