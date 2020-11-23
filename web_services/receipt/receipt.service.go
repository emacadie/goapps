package receipt

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"os"
	"shelfunit.info/golang/inventoryservice/cors"
	"strconv"
	"strings"
	
)

const receiptPath = "receipts"

const fileNameService = "in receipt.service."

func handleReceipts( w http.ResponseWriter, r *http.Request ) {
	var funcName = fileNameService + "handleReceipts: "
	log.Println( funcName + "starting switch" )
	switch r.Method {
	case http.MethodGet:
		receiptList, err := GetReceipts()
		if err != nil {
			w.WriteHeader( http.StatusInternalServerError )
			return
		}
		j, err := json.Marshal( receiptList )
		if err != nil {
			log.Fatal( err )
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		r.ParseMultipartForm( 5 << 20 ) // limit to 5 MB
		file, handler, err := r.FormFile( "receipt" )
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer file.Close()
		f, err := os.OpenFile( filepath.Join( ReceiptDirectory, handler.Filename ), os.O_WRONLY|os.O_CREATE, 0666 )
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		w.WriteHeader(http.StatusCreated)
		
	case http.MethodOptions:
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
		
	} // switch r.Method
} // handleReceipts

func handleDownload(w http.ResponseWriter, r *http.Request) {
	var funcName = fileNameService + "handleDownload: "
	log.Println( funcName + "starting function" )
	urlPathSegments := strings.Split( r.URL.Path, fmt.Sprintf( "%s/", receiptPath ) )
	if len( urlPathSegments[ 1: ] ) > 1 {
		w.WriteHeader( http.StatusBadRequest )
		return
	}
	fileName := urlPathSegments[ 1: ][ 0 ]
	file, err := os.Open( filepath.Join( ReceiptDirectory, fileName ) )
	defer file.Close()
	if err != nil {
		w.WriteHeader( http.StatusNotFound )
		return
	}
	fHeader := make( []byte, 512 )
	file.Read( fHeader )
	fContentType := http.DetectContentType( fHeader )

	stat, err := file.Stat()
	if err != nil {
		w.WriteHeader( http.StatusInternalServerError )
		return
	}
	fSize := strconv.FormatInt( stat.Size(), 10 )
	w.Header().Set( "Content-Disposition", "attachment; filename=" + fileName )
	w.Header().Set( "Content-Type", fContentType )
	w.Header().Set( "Content-Length", fSize )
	file.Seek( 0, 0 )
	io.Copy( w, file )
} // handleDownload

// SetupRoutes
func SetupRoutes( apiBasePath string ) {
	var funcName = fileNameService + "SetupRoutes: "
	log.Println( funcName + "starting route setup" )
	receiptHandler  := http.HandlerFunc( handleReceipts )
	downloadHandler := http.HandlerFunc( handleDownload )
	http.Handle( fmt.Sprintf( "%s/%s",  apiBasePath, receiptPath ), cors.Middleware( receiptHandler ) )
	http.Handle( fmt.Sprintf( "%s/%s/", apiBasePath, receiptPath ), cors.Middleware( downloadHandler ) )
	log.Printf( "%s done w/route setup w/ base path %s and receipt path %s \n", funcName, apiBasePath, receiptPath )
}

