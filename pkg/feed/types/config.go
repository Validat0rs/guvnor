package types

type Chain []struct {
	ChainId string `toml:"chain_id"`
	API     string `toml:"api"`
}

type Config struct {
	Chain Chain
}
