package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func CheckError(err error, msg string) {
	if err != nil {
		log.Println(msg, err)
	}
}

func FailOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	}
}

func DownloadFile(filepath string, url string) error {
	// Cria o arquivo
	out, err := os.Create(filepath)
	if err != nil  {
	  return err
	}
	defer out.Close()
	// Baixa o arquivo
	resp, err := http.Get(url)
	if err != nil {
	  return err
	}
	defer resp.Body.Close()
	// Verifica se baixou corretamente
	if resp.StatusCode != http.StatusOK {
	  return fmt.Errorf("bad status: %s", resp.Status)
	}
	// Armazena o arquvbo
	_, err = io.Copy(out, resp.Body)
	if err != nil  {
	  return err
	}
	return nil
  }