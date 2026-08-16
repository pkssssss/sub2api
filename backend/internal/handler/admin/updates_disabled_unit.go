//go:build unit

package admin

// updatesDisabledOnShenle 在 unit 测试构建中不拦截更新，保持上游测试语义。
func updatesDisabledOnShenle() bool { return false }
