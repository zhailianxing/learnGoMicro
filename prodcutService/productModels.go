package prodcutService

import "strconv"

type ProductInfo struct {
	ProductId int `json:"product_id"`
	ProductName string `json:"product_name"`
}


func NewProducts(n int) []*ProductInfo {
	ret := make([]*ProductInfo, 0)
	for i := 0; i < n; i++ {
		ret = append(ret, &ProductInfo{
			ProductId: 100+i,
			ProductName: "product_name_"+strconv.Itoa(i),
		})
	}
	return ret
}

