package ipgeo

type Domain struct {
	IP          string
	City        string
	CountryName string
	RegionName  string
}

type Repository interface {
	GetLocationByIP(ip string) (Domain, error)
}
