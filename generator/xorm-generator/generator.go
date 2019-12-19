package xorm_generator

import (
	"bytes"
	"coder/config"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"github.com/lunny/log"
	"github.com/polaris1119/goutils"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

func GetTable(tableName string) (*core.Table,error)  {
	var engine *xorm.Engine
	var err error
	
	engine, err = xorm.NewEngine(config.Project.DatabaseType, config.Project.Database)
	if err != nil {
		return nil,err
	}
	//	tables,_ := engine.Dialect().GetColumns()
	if err != nil {
		return nil,err
	}
	table := core.NewEmptyTable()
	table.Name = tableName
	table.Comment = ""
	table.StoreEngine = "innodb"
	colSeq, cols, err := engine.Dialect().GetColumns(table.Name)
	
	if err != nil {
		return nil,err
	}
	for _, name := range colSeq {
		table.AddColumn(cols[name])
	}
	indexes, err := engine.Dialect().GetIndexes(table.Name)
	if err != nil {
		return nil,err
	}
	table.Indexes = indexes
	
	for _, index := range indexes {
		for _, name := range index.Cols {
			if col := table.GetColumn(name); col != nil {
				col.Indexes[index.Name] = index.Type
			} else {
				return nil,fmt.Errorf("Unknown col %s in index %v of table %v, columns %v", name, index.Name, table.Name, table.ColumnsSeq())

			}
		}
	}
	return table,nil
}
func  Render(templateDir string,genDir string,modelName string) {
	var langTmpl LangTmpl
	var ok bool
	var lang string = "go"
	var prefix string = "" //[SWH|+]
	
	genDir = strings.Replace(genDir, "\\", "/", -1)
	models := path.Base(genDir)
	cfgPath := path.Join(templateDir, "config")
	info, err := os.Stat(cfgPath)
	var configs map[string]string
	if err == nil && !info.IsDir() {
		configs = loadConfig(cfgPath)
		if l, ok := configs["lang"]; ok {
			lang = l
		}
		if j, ok := configs["genJson"]; ok {
			genJson, err = strconv.ParseBool(j)
		}
		
		//[SWH|+]
		if j, ok := configs["prefix"]; ok {
			prefix = j
		}
		
		if j, ok := configs["ignoreColumnsJSON"]; ok {
			ignoreColumnsJSON = strings.Split(j, ",")
		}
		
		if j, ok := configs["created"]; ok {
			created = strings.Split(j, ",")
		}
		
		if j, ok := configs["updated"]; ok {
			updated = strings.Split(j, ",")
		}
		
		if j, ok := configs["deleted"]; ok {
			deleted = strings.Split(j, ",")
		}
		
	}
	
	if langTmpl, ok = langTmpls[lang]; !ok {
		fmt.Println("Unsupported programing language", lang)
		return
	}
	
	os.MkdirAll(genDir, os.ModePerm)
	
	supportComment = true
	var tables []*core.Table
	table,err := GetTable(goutils.UnderscoreName(modelName))
	tables = append(tables, table)
	
	filepath.Walk(templateDir, func (f string, info os.FileInfo, err error) error{
		err = WorkFile(f, info, err,langTmpl, tables, prefix,genDir, models)
		return err
	})
	
	
}

func WorkFile(f string, info os.FileInfo, err error, langTmpl LangTmpl, tables []*core.Table, prefix string, genDir string,models string) error {
	if info.IsDir() {
		return nil
	}
	
	if info.Name() == "config" {
		return nil
	}
	
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		log.Errorf("%v", err)
		return err
	}
	
	t := template.New(f)
	t.Funcs(langTmpl.Funcs)
	
	tmpl, err := t.Parse(string(bs))
	if err != nil {
		log.Errorf("%v", err)
		return err
	}
	
	fileName := info.Name()
	newFileName := fileName[:len(fileName)-4]
	ext := path.Ext(newFileName)
	
	
	for _, table := range tables {
		//[SWH|+]
		if prefix != "" {
			table.Name = strings.TrimPrefix(table.Name, prefix)
		}
		// imports
		tbs := []*core.Table{table}
		imports := langTmpl.GenImports(tbs)
		
		w, err := os.Create(path.Join(genDir, table.Name+ext))
		if err != nil {
			log.Errorf("%v", err)
			return err
		}
		defer w.Close()
		
		newbytes := bytes.NewBufferString("")
		
		t := &Tmpl{Tables: tbs, Imports: imports, Models: models}
		err = tmpl.Execute(newbytes, t)
		if err != nil {
			log.Errorf("%v", err)
			return err
		}
		
		tplcontent, err := ioutil.ReadAll(newbytes)
		if err != nil {
			log.Errorf("%v", err)
			return err
		}
		var source string
		if langTmpl.Formater != nil {
			source, err = langTmpl.Formater(string(tplcontent))
			if err != nil {
				log.Errorf("%v-%v", err, string(tplcontent))
				return err
			}
		} else {
			source = string(tplcontent)
		}
		
		w.WriteString(source)
		w.Close()
	}
	
	return nil
}