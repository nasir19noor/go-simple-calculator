package main

import (
    "fmt"
    "html/template"
    "net/http"
    "strconv"
)

var tmpl = template.Must(template.New("calculator").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>Simple Calculator</title>
</head>
<body>
    <h1>Simple Calculator</h1>
    <form method="post">
        <input type="number" name="num1" step="any" required>
        <select name="operation">
            <option value="add">+</option>
            <option value="subtract">-</option>
            <option value="multiply">*</option>
            <option value="divide">/</option>
        </select>
        <input type="number" name="num2" step="any" required>
        <button type="submit">Calculate</button>
    </form>
    {{if .Result}}
    <h2>Result: {{.Result}}</h2>
    {{end}}
</body>
</html>
`))

func calculatorHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        num1, err1 := strconv.ParseFloat(r.FormValue("num1"), 64)
        num2, err2 := strconv.ParseFloat(r.FormValue("num2"), 64)
        operation := r.FormValue("operation")

        if err1 == nil && err2 == nil {
            var result float64
            switch operation {
            case "add":
                result = num1 + num2
            case "subtract":
                result = num1 - num2
            case "multiply":
                result = num1 * num2
            case "divide":
                if num2 != 0 {
                    result = num1 / num2
                } else {
                    result = 0
                }
            }
            tmpl.Execute(w, struct{ Result float64 }{Result: result})
            return
        }
    }
    tmpl.Execute(w, nil)
}

func main() {
    http.HandleFunc("/", calculatorHandler)
    fmt.Println("Starting server on :8080")
    http.ListenAndServe(":8080", nil)
}
