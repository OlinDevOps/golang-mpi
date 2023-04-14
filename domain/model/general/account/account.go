package account

type Registration struct {
	PartyName         string `json:"party_name"` //as outlet name
	Alamat            string `json:"alamat"`
	PartySiteName     string `json:"party_site_name"`     //as outlet name
	ShortCode         string `json:"short_code"`          //as branch name
	Attribute2        string `json:"attribute_1"`         //as name npwp
	Attribute3        string `json:"attribute_3"`         //as alamat npwp
	Attribute4        string `json:"attribute_4"`         //as no npwp
	Attribute5        string `json:"attribute_5"`         //as Nama pemilik
	Attribute6        string `json:"attribute_6"`         //as apoteker
	Attribute12       string `json:"attribute_12"`        //as nomor sia
	Attribute13       string `json:"attribute_13"`        //as no ktp
	Attribute14       string `json:"attribute_14"`        //as expired sia
	Attribute18       string `json:"attribute_18"`        //as file sia
	Attribute19       string `json:"attribute_19"`        //as expired sipa
	AttributeNumber12 string `json:"attribute_number_12"` // as status olin
}