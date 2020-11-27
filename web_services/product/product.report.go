package product

import (
	"bytes"
	"net/http"
	"encoding/json"
	"log"
	"path"
	"html/template"
	"time"
)

type ProductReportFilter struct {
	NameFilter         string `json:"productName"`
	ManufacturerFilter string `json:"manufacturer"`
	SKUFilter          string `json:"sku"`
}


func handleProductReport( w http.ResponseWriter, r *http.Request ) {
	var funcName = fileNameData + "handleProductReport: "
	switch r.Method {
		case http.MethodPost:
		var productFilter ProductReportFilter
		err := json.NewDecoder( r.Body ).Decode( &productFilter )
		if err != nil {
			log.Printf( "%s err: %s \n", funcName, err )
			w.WriteHeader( http.StatusBadRequest )
			return
		}
		products, err := searchForProductData( productFilter )
		if err != nil {
			log.Printf( "%s err: %s \n", funcName, err )
			w.WriteHeader( http.StatusInternalServerError )
			return
		}
		
		t := template.New( "report.gotmpl" )
		t, err = t.ParseFiles( path.Join( "templates", "report.gotmpl" ) )
		if err != nil {
			log.Printf( "%s err: %s \n", funcName, err )
			w.WriteHeader( http.StatusInternalServerError )
			return
		}
		var tmpl bytes.Buffer
		var product Product
		if len(products) > 0 {
			product = products[ 0 ]
			err = t.Execute( &tmpl, product )
		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		rdr := bytes.NewReader(tmpl.Bytes())
		w.Header().Set("Content-Disposition", "Attachment")
		http.ServeContent( w, r, "report.html", time.Now(), rdr )
		
		case http.MethodOptions:
		return
		default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
		
	} // switch

} // handleProductReport

