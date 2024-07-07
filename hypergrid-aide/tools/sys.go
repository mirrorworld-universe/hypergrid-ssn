package tools

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const AIDE_CONFIG_FILE = "/home/ubuntu/.hypergrid-ssn/last_slot.txt"

func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}
func GetLastSentSlot() (uint64, error) {
	if CheckFileExist(AIDE_CONFIG_FILE) {
		var f *os.File
		var err error
		f, err = os.Open(AIDE_CONFIG_FILE)
		if err != nil {
			log.Fatal(err)
			return 0, err
		}
		defer f.Close()

		fd, err := io.ReadAll(f)
		if err != nil {
			fmt.Println("read to fd fail", err)
			return 0, err
		}
		last_sent_slot, err := strconv.ParseUint(string(fd), 10, 64) // 将fd从[]byte转换为string，然后转换为int
		if err != nil {
			fmt.Println("convert fd to int fail", err)
			return 0, err

		}
		return last_sent_slot, nil
	}
	return 0, nil
}

func SetLastSentSlot(slot uint64) (bool, error) {
	if !CheckFileExist(AIDE_CONFIG_FILE) {

		f, err := os.Create(AIDE_CONFIG_FILE)
		if err != nil {
			log.Fatal(err)
			return false, err
		}
		defer f.Close()
		_, err = f.WriteString(strconv.FormatUint(slot, 10))
		if err != nil {
			log.Fatal(err)
			return false, err
		}
		return true, nil
	} else {
		f, err := os.OpenFile(AIDE_CONFIG_FILE, os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
			return false, err
		}
		defer f.Close()
		_, err = f.WriteString(strconv.FormatUint(slot, 10))
		if err != nil {
			log.Fatal(err)
			return false, err
		}
		return true, nil
	}
}
