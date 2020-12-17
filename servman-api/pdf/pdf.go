package pdf

import (
    "strings"
    h2pdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
    "log"
)

func ToHtml(item map[string]interface{}) string {
    div := " "
    div += "<div class='row'>"

    for k, v := range(item) {
        div += "<div class='column'>"
        div += "<span class='key'>"+ k +"</span>:"
        div += "<span class='value'>"+ v.(string) +"</span>"
        div += "</div>"
    }

    div += "</div>"

    return div
}

/*
j, _ := json.Marshal(jj)
_ = json.Unmarshal(j, &m)
html := pdf.ToHtml(m)
*/

func RenderPdf(htmlStr string, filename string) {
    pdfg, err :=  h2pdf.NewPDFGenerator()
    if err != nil{
        return
    }

    pdfg.AddPage(h2pdf.NewPageReader(strings.NewReader(htmlStr)))


    // Create PDF document in internal buffer
    err = pdfg.Create()
    if err != nil {
        log.Fatal(err)
    }

    //Your Pdf Name
    err = pdfg.WriteFile(filename)
    if err != nil {
        log.Fatal(err)
    }
}
