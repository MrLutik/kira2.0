package types

type Test struct {
	Text string
}

type SekaidKey struct {
	Address string `yaml:"address"`
}

type AddressPermissions struct {
	BlackList []int `json:"blacklist"`
	WhiteList []int `json:"whitelist"`
}


