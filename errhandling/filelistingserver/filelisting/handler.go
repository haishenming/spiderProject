package filelisting

import (
	"net/http"
	"os"
	"io/ioutil"
)




func HandleFileList(writer http.ResponseWriter, requests *http.Request) error {
		path := requests.URL.Path[len("/list/"):]
		file,err := os.Open(path)
		if err != nil {
			http.Error(writer,
				err.Error(),
				http.StatusInternalServerError)
			return err
		}
		defer file.Close()

		all, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		writer.Write(all)
		return nil
	}
