package main

type model struct {
	Order_uid               string   `json:"order_uid"`
	Track_number            string   `json:"track_number"`
	Entry                   string   `json:"entry"`
	Delivery                delivery `json:"delivery"`
	Payment                 payment  `json:"payment"`
	Items                   []items  `json:"items"`
	Locale                  string   `json:"locale"`
	International_signature string   `json:"international_signature"`
	Customer_ID             string   `json:"customer_id"`
	Delivery_service        string   `json:"delivery_service"`
	Shard_Key               string   `json:"shardkey"`
	Sm_ID                   int      `json:"sm_id"`
	Date_Created            string   `json:"date_created"`
	Oof_Shard               string   `json:"oof_shard"`
}

type delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type payment struct {
	Transaction   string `json:"transaction"`
	Request_ID    string `json:"request_id"`
	Currency      string `json:"currency"`
	Provider      string `json:"provider"`
	Amount        int    `json:"amount"`
	Payment_dt    int    `json:"payment_dt"`
	Bank          string `json:"bank"`
	Delivery_cost int    `json:"delivery_cost"`
	Goods_total   int    `json:"goods_total"`
	Custom_fee    int    `json:"custom_fee"`
}

type items struct {
	Chrt_ID      int    `json:"chrt_id"`
	Track_number string `json:"track_number"`
	Price        int    `json:"price"`
	Rid          string `json:"rid"`
	Name         string `json:"name"`
	Sale         int    `json:"sale"`
	Size         string `json:"size"`
	Total_price  int    `json:"total_price"`
	NM_ID        int64  `json:"nm_id"`
	Brand        string `json:"brand"`
	Status       int    `json:"status"`
}
