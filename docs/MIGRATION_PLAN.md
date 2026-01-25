# JingWei Guard 模块迁移计划

## 项目信息

### 源项目
- **框架**: Spring Boot 2.7.11 + Vue 2
- **数据库**: Dameng (MySQL 兼容模式)
- **模块**: jingwei-ztuo-guard, jingwei-system (guard 相关)
- **路径**: `/data/lyndon/iProject/javapath/jingwei-xserver-plus`

### 目标项目
- **框架**: gin-admin (Golang) + React
- **数据库**: PostgreSQL 17
- **项目路径**: `/lyndon/iProject/gopath/src/github.com/settlement-monitoring`
- **Vue项目**: `/data/lyndon/iProject/webpath/scmap-shanxi`（参考）

---

## 迁移范围

### 后端 API 迁移

| 来源 | 目标 | 优先级 |
|------|------|--------|
| 人员管理 (5个表) | guard/person API | P0 |
| 设备装备 (6个表) | guard/equipment API | P0 |
| 位置地标 (3个表) | guard/location API | P1 |
| 场景管理 (13个表) | guard/scene API | P0 |
| 流程相关 | guard/workflow API | P2 |

### 前端页面迁移

| 来源 | 目标 | 优先级 |
|------|------|--------|
| 工作人员管理 | pages/guard/staff | P0 |
| 执勤人员管理 | pages/guard/watchkeeper | P0 |
| 重点人员管理 | pages/guard/targeted-individuals | P0 |
| 设备管理 | pages/guard/equipment | P0 |
| 场景管理 | pages/guard/scene | P0 |
| 标绘管理 | pages/guard/plotting | P1 |

---

## 迁移阶段

### 阶段一：后端基础设施 (当前阶段)

**目标**: 搭建 Golang 后端开发环境

#### 任务列表
- [ ] 1.1 创建 guard 模块目录结构
- [ ] 1.2 定义数据库模型 (schema)
- [ ] 1.3 实现数据访问层 (dal)
- [ ] 1.4 实现业务逻辑层 (biz)
- [ ] 1.5 实现 API 控制器 (api)
- [ ] 1.6 配置路由和权限 (wire/casbin)
- [ ] 1.7 编写单元测试

#### 预计时间: 2-3 天

---

### 阶段二：后端核心 API 开发

**目标**: 实现核心 CRUD API

#### 任务列表
- [ ] 2.1 工作人员 CRUD API
- [ ] 2.2 执勤人员 CRUD API
- [ ] 2.3 重点人员 CRUD API
- [ ] 2.4 训犬员 CRUD API
- [ ] 2.5 警犬 CRUD API
- [ ] 2.6 车辆 CRUD API
- [ ] 2.7 枪支 CRUD API
- [ ] 2.8 对讲机 CRUD API
- [ ] 2.9 无人机反制设备 CRUD API
- [ ] 2.10 防暴安检设备 CRUD API
- [ ] 2.11 执法记录仪 CRUD API
- [ ] 2.12 制高点 CRUD API
- [ ] 2.13 重点部位 CRUD API
- [ ] 2.14 危险部位 CRUD API
- [ ] 2.15 场景 CRUD API
- [ ] 2.16 标绘 CRUD API

#### 预计时间: 5-7 天

---

### 阶段三：后端高级功能

**目标**: 实现业务特色功能

#### 任务列表
- [ ] 3.1 Excel 导入导出
- [ ] 3.2 数据统计 API
- [ ] 3.3 场景合并/拆分
- [ ] 3.4 标绘统计分析
- [ ] 3.5 GIS 空间查询
- [ ] 3.6 工作流集成（Flowable）
- [ ] 3.7 文件上传下载
- [ ] 3.8 数据验证和校验

#### 预计时间: 3-5 天

---

### 阶段四：前端基础设施

**目标**: 搭建 React 前端开发环境

#### 任务列表
- [ ] 4.1 创建 guard 模块目录结构
- [ ] 4.2 定义 API 服务层 (services)
- [ ] 4.3 创建通用组件
- [ ] 4.4 配置路由
- [ ] 4.5 配置权限控制
- [ ] 4.6 配置国际化 (i18n)

#### 预计时间: 1-2 天

---

### 阶段五：前端核心页面开发

**目标**: 实现核心业务页面

#### 任务列表
- [ ] 5.1 工作人员管理页面
- [ ] 5.2 执勤人员管理页面
- [ ] 5.3 重点人员管理页面
- [ ] 5.4 训犬员管理页面
- [ ] 5.5 警犬管理页面
- [ ] 5.6 车辆管理页面
- [ ] 5.7 装备管理页面
- [ ] 5.8 位置地标管理页面
- [ ] 5.9 场景管理页面
- [ ] 5.10 标绘管理页面

#### 预计时间: 5-7 天

---

### 阶段六：前端高级功能

**目标**: 实现复杂交互和业务功能

#### 任务列表
- [ ] 6.1 Excel 导入导出组件
- [ ] 6.2 数据统计图表
- [ ] 6.3 地图标绘组件
- [ ] 6.4 场景地图展示
- [ ] 6.5 图片上传组件
- [ ] 6.6 表单验证
- [ ] 6.7 批量操作

#### 预计时间: 3-5 天

---

### 阶段七：数据迁移

**目标**: 从 Dameng 迁移数据到 PostgreSQL

#### 任务列表
- [ ] 7.1 评估 Dameng 数据量和结构
- [ ] 7.2 编写数据转换脚本
- [ ] 7.3 迁移字典数据
- [ ] 7.4 迁移业务数据（分批）
- [ ] 7.5 数据验证和校对
- [ ] 7.6 性能测试

#### 预计时间: 2-3 天

---

### 阶段八：测试和部署

**目标**: 全面测试和上线部署

#### 任务列表
- [ ] 8.1 单元测试
- [ ] 8.2 集成测试
- [ ] 8.3 功能测试
- [ ] 8.4 性能测试
- [ ] 8.5 安全测试
- [ ] 8.6 用户验收测试
- [ ] 8.7 部署文档
- [ ] 8.8 正式上线

#### 预计时间: 3-5 天

---

## 总体时间表

| 阶段 | 任务 | 预计天数 | 开始日期 |
|------|------|----------|----------|
| 1 | 后端基础设施 | 2-3 | 2025-01-25 |
| 2 | 后端核心 API | 5-7 | 2025-01-28 |
| 3 | 后端高级功能 | 3-5 | 2025-02-04 |
| 4 | 前端基础设施 | 1-2 | 2025-02-09 |
| 5 | 前端核心页面 | 5-7 | 2025-02-11 |
| 6 | 前端高级功能 | 3-5 | 2025-02-18 |
| 7 | 数据迁移 | 2-3 | 2025-02-23 |
| 8 | 测试和部署 | 3-5 | 2025-02-26 |

**总计**: 24-37 天（约 5-7 周）

---

## 当前进度

- [x] 0.1 数据库结构设计和创建
- [x] 0.2 PostgreSQL 17 数据库搭建
- [x] 0.3 Redis 连接测试
- [ ] 0.4 Golang 开发环境验证
- [ ] 0.5 React 开发环境验证

---

## 下一步行动

### 立即开始：阶段一 - 后端基础设施

1. 创建 guard 模块目录结构
2. 定义数据库模型
3. 实现第一个表的 CRUD API（t_staff）

---

## 风险和注意事项

### 技术风险
- **GIS 数据处理**: PostGIS 空间数据查询需要专业知识
- **工作流集成**: 可能需要保留 Flowable 或寻找 Go 替代方案
- **数据类型转换**: Dameng 到 PostgreSQL 的类型差异

### 业务风险
- **功能完整性**: 确保所有业务功能都正确迁移
- **数据一致性**: 迁移过程数据不要丢失
- **用户体验**: React 和 Vue 的交互差异

### 缓解措施
- 分阶段验证，确保每个阶段质量
- 保留原系统并行运行一段时间
- 编写详细的数据迁移回滚方案

---

## 资源和参考资料

### 技术文档
- [gin-admin 文档](https://github.com/LyricTian/gin-admin)
- [gin-admin-frontend 文档](https://github.com/gin-admin/gin-admin-frontend)
- [PostgreSQL 17 文档](https://www.postgresql.org/docs/17/)
- [PostGIS 文档](https://postgis.net/documentation/)

### 源代码参考
- Java Controller: `jingwei-admin/src/main/java/com/jingwei/web/controller/guard/`
- Java Service: `jingwei-system/src/main/java/com/jingwei/guard/service/`
- Entity 定义: `jingwei-system/src/main/java/com/jingwei/guard/domain/`
- Vue API: `scmap-shanxi/src/api/`
- Vue Components: 参考 scmap-shanxi 相关页面

---

**创建时间**: 2026-01-25
**最后更新**: 2026-01-25
**状态**: 🚀 开始执行阶段一