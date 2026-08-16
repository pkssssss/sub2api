//go:build !unit

package admin

// updatesDisabledOnShenle 自建补丁版（sub2api-shenle）专用：
// 生产构建（无 unit tag）禁用在线更新，防止补丁被官方版本覆盖。
func updatesDisabledOnShenle() bool { return true }
