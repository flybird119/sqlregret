package mysql

import (
	"fmt"
	"testing"
	"tsg/client"
)

func getNameFromIndex(FieldNames map[string]int, index int) string {
	for k, v := range FieldNames {
		if v == index {
			return k
		}
	}

	return ""
}

func TestQuery(t *testing.T) {
	connector := client.NewMysqlConnection("192.168.0.135", "logreader", "123456", "test", 3306, 123)
	connector.Connect()
	if rst, err := connector.Query("desc testa"); nil != err {
		return
	} else {
		fmt.Println("Field num:", len(rst.Fields))
		for k, v := range rst.FieldNames {
			fmt.Print("key:", k)
			fmt.Println("\tvalue:", v)
		}
		fmt.Println("--------------------------------------------------------------")

		fieldMetas := make([]*client.FieldMeta, 0)

		for index := range rst.Values {
			fieldMeta := new(client.FieldMeta)
			var result string
			for index2 := range rst.Values[index] {
				switch value := rst.Values[index][index2].(type) {
				case []uint8:
					result = string(value)
				default:
					result = ""
				}

				if index2 == 0 {
					fieldMeta.ColumnName = result
				} else if index2 == 1 {
					fieldMeta.ColumnType = result
				} else if index2 == 2 {
					fieldMeta.IsNullable = result
				} else if index2 == 3 {
					fieldMeta.IsKey = result
				} else if index2 == 4 {
					fieldMeta.DefaultValue = result
				} else if index2 == 5 {
					fieldMeta.Extra = result
				}
			}
			fieldMetas = append(fieldMetas, fieldMeta)
		}
		tableMeta := client.NewTableMeta("testa", fieldMetas)

		fmt.Println(tableMeta)
	}
}
