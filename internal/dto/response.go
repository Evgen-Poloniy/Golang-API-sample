package dto

// Response represents a data record from database
// @Description Data record structure returned from the database
type Data struct {
}

type DataResponse struct {
    Data interface{} `json:"data"`
}
