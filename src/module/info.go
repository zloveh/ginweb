package module

import (
	"ginweb/src/conf"
)

func GetInfo() ([]string, error) {
	var tp string
	sold_type := []string{}
	sql := "select distinct order_type from sold_info"

	db := conf.Riskdb.DBPool.GetConn()
	defer conf.Riskdb.DBPool.Release(db)

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&tp)
		if err != nil {
			return nil, err
		}
		sold_type = append(sold_type, tp)
	}
	return sold_type, nil
}
