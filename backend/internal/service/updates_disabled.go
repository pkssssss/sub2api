//go:build !unit

package service

// updatesDisabledOnShenle 自建补丁版（sub2api-shenle）专用：
// 生产构建（无 unit tag）禁用更新检查，不访问 GitHub。
func updatesDisabledOnShenle() bool { return true }
