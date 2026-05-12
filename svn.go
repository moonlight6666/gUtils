package gUtils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// SvnCheckAndUpdate
// 作用：检测SVN目录是否需要更新，需要则自动更新
// 参数：dir - SVN工作副本目录
// 返回：更新失败返回error，成功返回nil
func SvnCheckAndUpdate(dir string) error {
	// 1. 校验目录是否存在
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("目录不存在: %s", dir)
	}

	// 2. 校验是否为合法SVN工作副本
	infoCmd := exec.Command("svn", "info", dir)
	if _, err := infoCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("不是合法的SVN目录: %s, err: %v", dir, err)
	}

	// 3. 检测是否需要更新（svn status -u 检查远程版本，包含 * 代表需要更新）
	statusCmd := exec.Command("svn", "status", "-u", dir)
	output, err := statusCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("检测SVN状态失败: %v", err)
	}

	// 判断输出中是否包含 *（* 表示本地版本落后远程）
	needUpdate := strings.Contains(string(output), "*")
	if !needUpdate {
		fmt.Printf("[SVN] 目录 %s 已是最新版本，无需更新\n", dir)
		return nil
	}

	fmt.Printf("[SVN] 目录 %s 检测到新版本，开始执行更新\n", dir)

	// 4. 执行SVN更新
	updateCmd := exec.Command("svn", "update", dir)
	if _, err := updateCmd.CombinedOutput(); err != nil {
		fmt.Printf("[SVN] 目录 %s 更新失败: %v\n", dir, err)
		return err
	}
	fmt.Printf("[SVN] 目录 %s 更新成功\n", dir)
	return nil

	//// ====================== 更新失败：自动重试逻辑 ======================
	//fmt.Printf("[SVN] 目录 %s 更新失败，3秒后清理锁并重试\n", dir)
	//time.Sleep(3 * time.Second)
	//
	//// 执行 svn cleanup 清理锁
	//cleanupCmd := exec.Command("svn", "cleanup", dir)
	//if _, err := cleanupCmd.CombinedOutput(); err != nil {
	//	return fmt.Errorf("svn cleanup失败: %v", err)
	//}
	//fmt.Printf("[SVN] 目录 %s 清理锁成功\n", dir)
	//
	//// 重试更新
	//retryCmd := exec.Command("svn", "update", dir)
	//if _, err := retryCmd.CombinedOutput(); err != nil {
	//	return fmt.Errorf("重试更新失败: %v", err)
	//}
	//
	//fmt.Printf("[SVN] 目录 %s 重试更新成功\n", dir)
	//return nil
}
