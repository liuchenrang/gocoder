package library

import (
	"bufio"
	"coder/config"
	"coder/generator/form"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/polaris1119/goutils"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
	"os"
	"reflect"
	"regexp"
	"strings"
	"unicode"
)

type ParseModel struct {
	fields []form.Field
}

func (p *ParseModel) parseTagSetting(tagQuery string) map[string]string {
	setting := map[string]string{}
	tagRaws := []string{tagQuery}
	for _, str := range tagRaws {
		if str == "" {
			continue
		}
		tags := strings.Split(str, ";")
		for _, value := range tags {
			v := strings.Split(value, ":")
			k := strings.TrimSpace(strings.ToUpper(v[0]))
			if len(v) >= 2 {
				setting[k] = strings.Join(v[1:], ":")
			} else {
				setting[k] = k
			}
		}
	}
	return setting
}
func (p *ParseModel) GetFields() ([]form.Field, error) {
	return p.fields, nil
}
func Lcfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToLower(v))
		return u + str[len(u):]
	}
	return ""
}
func LowerWord(str string) string {
	var nw []byte
	var isFind = false
	for _, value := range str {
		if value < 97 && value > 65 && !isFind {
			nw = append(nw, byte(value+32))
			isFind = true
		} else {
			nw = append(nw, byte(value))
		}
	}
	return string(nw)
}

func (p *ParseModel) ParseTable(name string) {
	var engine *xorm.Engine
	var err error
	
	engine, err = xorm.NewEngine(config.Project.DatabaseType, config.Project.Database)
	if err != nil {
		println(err.Error());
		return
	}
	//	tables,_ := engine.Dialect().GetColumns()
	if err != nil {
		println(err.Error())
		return
	}
	table := core.NewEmptyTable()
	table.Name = "sqb_user"
	table.Comment = ""
	table.StoreEngine = "innodb"
	colSeq, cols, err := engine.Dialect().GetColumns(table.Name)
	
	if err != nil {
		println(err.Error())
		return
	}
	for _, name := range colSeq {
		table.AddColumn(cols[name])
	}
	indexes, err := engine.Dialect().GetIndexes(table.Name)
	if err != nil {
		println(err.Error())
		return
	}
	table.Indexes = indexes
	
	for _, index := range indexes {
		for _, name := range index.Cols {
			if col := table.GetColumn(name); col != nil {
				col.Indexes[index.Name] = index.Type
			} else {
				println(fmt.Errorf("Unknown col %s in index %v of table %v, columns %v", name, index.Name, table.Name, table.ColumnsSeq()))
				return
			}
		}
	}
	fmt.Printf("%+v", table)
	
	//engine, err = xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
}
func (p *ParseModel) Parse(filePath string) (error) {
	file, er := os.Open(filePath);
	if er != nil {
		return er
	}
	reg, er := regexp.Compile(".*`(.*gorm.*?`)");
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		
		if er != nil {
			return er
		}
		re := regexp.MustCompile("\\s{2,}")
		result := reg.FindAll([]byte(scanner.Text()), 1)
		for _, value := range result {
			field := form.Field{}
			
			fieldInfo := strings.Trim(string(value), "\t ")
			fieldInfo = re.ReplaceAllString(fieldInfo, " ")
			fieldAttr := strings.Split(fieldInfo, " ")
			
			field.Name = strings.Trim(fieldAttr[0], " ");
			field.Kind = strings.Trim(fieldAttr[1], " ");
			field.DBName = goutils.UnderscoreName(field.Name);
			field.FormSettings = make(map[string]string)
			field.TagSettings = make(map[string]string)
			tagRaw := strings.Join(fieldAttr[2:], " ")
			tag := reflect.StructTag(strings.Trim(tagRaw, "`"))
			
			gormInfo := tag.Get("gorm")
			if len(gormInfo) > 0 {
				field.TagSettings = p.parseTagSetting(gormInfo)
				if val, ok := field.TagSettings["COMMENT"]; ok {
					field.TagSettings["COMMENT"] = strings.Trim(val, "\"'");
				}
			}
			formInfo := tag.Get("form")
			if (len(formInfo) > 0) {
				field.FormSettings = p.parseTagSetting(formInfo)
			}
			field.Render = form.NewSelectElement(field)
			field.FormSettings["COMMENT"] = field.TagSettings["COMMENT"]
			
			p.fields = append(p.fields, field)
		}
	}
	
	return nil
}
