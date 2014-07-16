package main

import (
	"./util"
	"os"
	//"strings"
	"text/template"
)

func Gen(langPath string, msgs []Message, outPath string) {
	files := util.GetAllFiles(langPath, ".tmpl")
	for _, file := range files {
		//tmpl := template.New("")
		//tmpl = tmpl.Funcs(template.FuncMap{"Title": strings.Title})
		tmpl, err := template.ParseFiles(file)
		if err != nil {
			println(err.Error())
			return
		}

		outFile := file[0 : len(file)-len(".tmpl")]
		of, _ := os.Create(outFile)
		err = tmpl.Execute(of, msgs)
		if err != nil {
			print(err.Error())
			return
		}
	}
	print("gen ok.")
}
