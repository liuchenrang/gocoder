package contractService

import(
     {{range $element := .imports}}
        "{{$element}}"
     {{end}}
)

type I{{.CreatorName}} interface {

      Add(model models.{{.Name}})  (models.{{.Name}}, error)
      Delete(model models.{{.Name}})  (int, error)
      DeleteById(id uint)  (int, error)

      Update(id uint, attr map[string]interface{})  (bool, error)
      UpdateByModel(model models.{{.Name}})  (bool, error)
      Incr(id uint,attr map[string]interface{}) (bool, error)

      Find(id uint)  (*models.{{.Name}}, error)

      FindAll(page components.Page) (*[]*models.{{.Name}}, error)
      FindWithID(ids []uint) (*[]*models.{{.Name}}, error)

      CountWhereForAdmin(wh string,bind ...interface{}) (int64,  error)

      FindWhereForAdmin(wh string,page components.Page,bind ...interface{}) (*[]*models.{{.Name}}, error)

      Count() (int64, error)
}


