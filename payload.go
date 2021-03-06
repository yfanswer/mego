package mego

// Request 呈現了一個客戶端所傳送過來的請求內容。
type Request struct {
	// Method 是欲呼叫的方法名稱。
	Method string `codec:"m" msgpack:"m"`
	// Files 是此請求所包含的檔案欄位與其內容。
	Files map[string][]*RawFile `codec:"f" msgpack:"f"`
	// Params 是資料或參數。
	Params interface{} `codec:"p" msgpack:"p"`
	// ID 為本次請求編號，若無則為單次通知廣播不需回應。
	ID int `codec:"i" msgpack:"i"`
}

// Response 呈現了 Mego 將會回應給客戶端的內容。
type Response struct {
	// Event 是欲呼叫的客戶端事件名稱。
	Event string `codec:"v" msgpack:"v"`
	// Result 是正常回應時的資料酬載。
	Result interface{} `codec:"r" msgpack:"r"`
	// Error 是錯誤回應時的資料酬載，與 Result 兩者擇其一，不會同時使用。
	Error ResponseError `codec:"e" msgpack:"e"`
	// ID 是當時發送此請求的編號，用以讓客戶端比對是哪個請求所造成的回應。
	ID int `codec:"i" msgpack:"i"`
}

// ResponseError 是回應錯誤資料建構體。
type ResponseError struct {
	// Code 是錯誤代號。
	Code int `codec:"c" msgpack:"c"`
	// Message 是人類可讀的簡略錯誤訊息。
	Message string `codec:"m" msgpack:"m"`
	// Data 是本次錯誤的詳細資料。
	Data interface{} `codec:"d" msgpack:"d"`
}

// RawFile 是尚未轉化成為可供開發者使用之前的生檔案資料內容。
type RawFile struct {
	// Binary 是檔案的二進制。
	Binary []byte `codec:"b" msgpack:"b"`
	// ID 是由客戶端替此檔案所產生的順序編號，用於區塊組合。
	ID int `codec:"i" msgpack:"i"`
	// Parts 呈現區塊的分塊進度。索引 0 表示總共區塊數，索引 1 則是本區塊編號。
	// 如果這個切片是空的表示此為實體檔案而非區塊。
	Parts []int `codec:"p" msgpack:"p"`
	// Name 是檔案的原始名稱。
	Name string `codec:"n" msgpack:"n"`
}
