package types

type (
	Test struct {
		Text string
	}

	SekaidKey struct {
		Address string `yaml:"address"`
	}

	AddressPermissions struct {
		BlackList []int `json:"blacklist"`
		WhiteList []int `json:"whitelist"`
	}
)
