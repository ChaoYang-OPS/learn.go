package main

type downloadFromNetDisk struct {
	secret   DynamicSecret
	filePath string
}

func (dd *downloadFromNetDisk) DonwloadFile(file string) {
	if err := dd.logincheck(); err != nil {
		// todo 重新登录
	}
	dd.downloadFromAliYun(dd.filePath)
}

func (dd *downloadFromNetDisk) logincheck() error {
	//checkSecret(dd.secret())
	return nil
}

func (dd *downloadFromNetDisk) checkSecret(secret string) {
	// todo 调用阿里云的验证接口去验证密码是否有效
}

func (dd *downloadFromNetDisk) downloadFromAliYun(path string) {

}

func main() {
	dd := &downloadFromNetDisk{
		secret: &mobileTokenDynmic{
			mobileNumber: "131768998978",
		},
		filePath: "file1",
	}

	dd.DonwloadFile(dd.filePath)
}

func dynameicSecret() string {
	return ""
}

type DynamicSecret interface {
	GetSecret() string
}
type mobileTokenDynmic struct {
	mobileNumber string
}

func (d *mobileTokenDynmic) GetSecret() string {
	return "somethhing"
}

// 通常开发的时候，第一个版本叫做happy path
// 剩下的是痛： 它无法应对变更， 简单的变更会带来更痛苦的维护，
