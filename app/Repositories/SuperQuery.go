package Repository

import (
	"database/sql"
)

func SuperQuery() {

}

// 获取rows数据和字段名称
func SuperQueryRowsToArrMap(rows *sql.Rows) *map[string]interface{} {
	// 列名字
	columns, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	length := len(columns)
	result := make([]map[string]interface{}, 0)

	// 行数据
	for rows.Next() {
		current := makeResultReceiver(length)
		if err := rows.Scan(current...); err != nil {
			panic(err)
		}
		// logger.AppendDebug("asas", current)
		value := make(map[string]interface{})
		for i := 0; i < length; i++ {
			k := columns[i]
			v := current[i].(*interface{})
			// value[k] = v
			// 解决字符类型的字段显示为base64
			arr, found := (*v).([]byte)
			if found {
				value[k] = string(arr[:])
			} else {
				value[k] = *v
			}
		}
		result = append(result, value)
	}

	// 返回数据
	data := make(map[string]interface{})
	data["columns"] = columns
	data["values"] = &result

	return &data
}

// then the result is what you want, keep in mind that the value of the map is always pointer, since the value may be sql null
func makeResultReceiver(length int) []interface{} {
	result := make([]interface{}, 0, length)
	for i := 0; i < length; i++ {
		var current interface{}
		// current = struct{}{}
		result = append(result, &current)
	}
	return result
}
