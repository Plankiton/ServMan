package main

import (
    "encoding/json"
    "../pdf"
)

type J struct {
    Name  string
    Coisa string
}

func main() {
    m := map[string]interface{}{}

    jj := J {
        Name: "joao",
        Coisa: "maria",
    }

    j, _ := json.Marshal(jj)
    _ = json.Unmarshal(j, &m)
    html := pdf.ToHtml(m)

    print(html)
    pdf.RenderPdf(`
<style>
* {
    background: white;
    text-align: center;
    font-size: 24;
}
</style>
        `+html, "joao.pdf")
}
