package utils

type Db struct {
	UsersPath       string `json:"users_path" goint:"users"`
	CompaniesPath       string `json:"companies_path" goint:"companies"`
	PostsPath     string `json:"posts_path" goint:"posts"`

}

type Storage struct {
	Endpoint  string `json:"endpoint"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

type GointConfig struct {
	Db      Db      `json:"db"`
	Storage Storage `json:"storage"`
	GraphQLServer string `json:"graph_ql_server"`
}

