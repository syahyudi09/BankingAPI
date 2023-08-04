package model


type RegistrasiCustomerModel struct {
	ID             string  `json:"id"`
	NamaCustomer   string  `json:"nama_customer" validate:"required"`
	Alamat         string  `json:"alamat" validate:"required"`
	NomorTelepon   string  `json:"nomor_telepon" validate:"required"`
	Email          string  `json:"alamat_email" validate:"required,email"`
	Password       string  `json:"password" validate:"required,min=8"`
	TanggalLahir   string  `json:"tanggal_lahir" validate:"required,datetime=2006-01-02"`
	Status         string  `json:"status_perkawinan" validate:"required"`
	Pekerjaan      string  `json:"pekerjaan" validate:"required"`
	Pendapatan     float64 `json:"pendapatan" validate:"min=0"`
	Kewarganegaraan string  `json:"kewarganegaraan" validate:"required,eq=WNI|eq=WNA"`
	Foto           string  `json:"foto"`
}

type CustomerLogin struct {
	ID       string  `json:"id"`
	Email    string `jsnton:"email"`
	Password string `json:"password"`
}

type GetAllCustomer struct{
	ID             string  `json:"id"`
	NamaCustomer   string  `json:"nama_customer" validate:"required,regex=^[A-Z][a-zA-Z ]*$"`
	Alamat         string  `json:"alamat" validate:"required"`
	NomorTelepon   string  `json:"nomor_telepon" validate:"required"`
	Email          string  `json:"alamat_email" validate:"required,email"`
	TanggalLahir   string  `json:"tanggal_lahir" validate:"required,datetime=2006-01-02"`
	Status         string  `json:"status_perkawinan" validate:"required"`
	Pekerjaan      string  `json:"pekerjaan" validate:"required"`
	Pendapatan     float64 `json:"pendapatan" validate:"min=0"`
	Kewarganegaraan string  `json:"kewarganegaraan" validate:"required,eq=wni|eq=wna"`
	Foto           string  `json:"foto"`
}

