package localfs

type bootClient struct {
	data struct {
		Matchers []struct {
			Serial string
			Asset string
			UUID string
			Mac string
	}
	BootScript string
	Name string
}
