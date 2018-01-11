package graphql


type MiniCompany struct {
	Id             string
	CommercialName string
	Email          string
	Password       string
	Ruc string
	SocialReason string

}

var CompanyByEmail = `query ($email: String!){Company (email: $email) {id commercialName ruc socialReason email password}}`
var CompanyBySocialReason = `query ($social: String!){Company (socialReason: $social) {id commercialName ruc socialReason email password}}`
var CompanyByRuc = `query ($ruc: String!){Company (ruc: $ruc) {id commercialName ruc socialReason email password}}`
