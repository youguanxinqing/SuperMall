package model

type Good struct {
	ID       string `json:"id"`
	Image    string `json:"imgae"`
	Link     string `json:"link"`
	Title    string `json:"title"`
	Fav      uint32 `json:"fav"`
	OrgPrice uint16 `json:"orgPrice"`
	Price    uint16 `json:"price"`
}

type Goods struct {
	List []Good `json:"list"`
}

// f, _ := os.Open("config/Goods.json")
// 	defer f.Close()
// 	data, _ := ioutil.ReadAll(f)
// 	var gds model.Goods
// 	gds.UnmarshalJSON(data)
// 	fmt.Println(gds)
// 	for _, i := range gds.List {
// 		fmt.Println(i.Image, i.ID, i.OrgPrice)
// 	}
