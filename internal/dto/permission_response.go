package dto

// GroupedPermissions 分组的权限（按资源分组，方便前端使用）
type GroupedPermissions struct {
	Livestock   []string `json:"livestock"`   // 养殖管理权限
	Environment []string `json:"environment"` // 环境监控权限
	Device      []string `json:"device"`      // 设备控制权限
	User        []string `json:"user"`        // 用户管理权限
	Report      []string `json:"report"`      // 报表权限
}
