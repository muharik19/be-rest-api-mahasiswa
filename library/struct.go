package library

type RespUpload struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type RespMahasiswaAll struct {
	ID      int    `json:"id"`
	Nim     string `json:"nim"`
	Nama    string `json:"nama"`
	Jurusan string `json:"jurusan"`
	Hp      string `json:"hp"`
	Photo   string `json:"photo"`
}

type RespResult struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Results []RespMahasiswaAll `json:"results"`
}

type checkNim struct {
	nim string
}
