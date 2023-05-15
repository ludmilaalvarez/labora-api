package models

type Item struct {
	Id           int     `json:"id"`
	CustomerName string  `json:"name"`
	OrderDate    string  `json:"date"`
	Product      string  `json:"product"`
	Quantity     int     `json:"quantity"`
	Price        float64 `json:"price"`
	ItemDetails  string  `json:"itemdetails"`
	TotalPrice   float64 `json:"totalprice"`
}


func (i *Item) PrecioTotal(){
	(*i).TotalPrice= (*i).Price * float64((*i).Quantity)
} 
 

/* 
func PrecioTotal(item Item) float64{
	 return (item.Price * float64(item.Quantity))
} */ 