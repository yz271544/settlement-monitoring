# PostgreSQL 17 数据库搭建完成报告

## 连接信息

```bash
# PostgreSQL
Host: localhost
Port: 35432
Database: settlement
Username: postgres
Password: STmt0125

# Redis
Host: localhost
Port: 36379
Password: (无)
```

## 已创建的数据库对象

### 数据表（30张）

#### 人员管理表（5张）
- `t_staff` - 工作人员表
- `t_watchkeeper` - 执勤人员表
- `t_targeted_individuals` - 重点人员表
- `t_dog_trainer` - 训犬员表
- `t_police_dog` - 警犬表

#### 设备装备表（6张）
- `t_car` - 车辆表
- `t_firearm` - 枪支表
- `t_interphone` - 对讲机表
- `t_drone_counter_equipment` - 无人机反制设备表
- `t_security_screening_equipment` - 防爆安检设备表
- `t_police_recorder` - 执法记录仪表

#### 位置地标表（3张）
- `t_vantage_point` - 哨位表
- `t_key_part` - 重点部位表
- `t_dangerous_part` - 危险部位表

#### 场景相关表（13张）
- `t_scene` - 场景表
- `t_residence` - 驻地表
- `t_plotting` - 标绘表
- `t_camera` - 摄像头表
- `t_equipment` - 装备信息表
- `t_person` - 人员维护表
- `t_three_dimension_model` - 三维模型表
- `t_vr_info` - VR信息表
- `t_legend` - 图例表
- `t_coverage_tree` - 覆盖类型树表
- `t_scene_map` - 场景地图表
- `t_simulation_route` - 演练路线表
- `t_gis_region` - GIS区域表

#### 系统支持表（3张）
- `sys_dict_type` - 字典类型表
- `sys_dict_data` - 字典数据表
- `t_flow` - 流程表

### 视图（3个）
- `v_scene_statistics` - 场景统计视图
- `v_plotting_by_scene_type` - 按场景类型统计标绘数量视图
- `v_person_statistics` - 人员统计视图

### 触发器（30个）
所有数据表都创建了 `update_time` 触发器，自动更新 `update_time` 字段。

### 索引
已为所有表的关键字段创建索引，包括：
- 主键索引
- 业务字段索引（tel, type, personnel_id 等）
- 逻辑删除字段索引（deleted, del_flag）
- 外键索引（scene_id 等）

## 字典数据

已预置的字典类型：
- `guard_person_type` - 人员类型（执勤人员、工作人员、训犬员、重点人员、警犬）
- `guard_car_type` - 车辆类型（小汽车、卡车、消防车、民警车、巡警车、武警车、特警车、救护车、摩托车）
- `guard_equipment_type` - 装备类型
- `guard_device_type` - 设备类型
- `guard_scene_type` - 场景类型（行政边界、KML文件、驻地）
- `guard_plotting_type` - 标绘类型（点、线、面）
- `guard_screening_type` - 安检设备类型

## 数据类型映射

| Dameng (DM) | PostgreSQL 17 |
|-------------|---------------|
| VARCHAR(32) | VARCHAR(32) |
| VARCHAR(100) | VARCHAR(100) |
| VARCHAR(255) | VARCHAR(255) |
| CHAR(1) | CHAR(1) |
| TIMESTAMP(0) | TIMESTAMP |
| INT | INTEGER |
| CLOB | TEXT |
| BLOB | BYTEA |

## 特殊说明

### PostGIS 扩展
已启用 PostGIS 扩展用于存储和查询 GIS 空间数据：
- `region_shape` 字段存储 GeoJSON 格式的区域形状
- `shape` 字段在 `t_plotting` 和 `t_simulation_route` 中存储坐标信息

### 逻辑删除
所有支持逻辑删除的表使用以下字段：
- `deleted` 字段：'0' 表示未删除，'2' 表示已删除
- `del_flag` 字段：false 表示未删除，true 表示已删除

### 时间戳管理
- `create_time`：创建时间，默认当前时间戳
- `update_time`：更新时间，通过触发器自动更新
- `update_timestamp()` 函数：自动更新时间戳的触发器函数

## 数据库验证

```sql
-- 查看所有表
\dt

-- 查看所有视图
\dv

-- 查看所有索引
\di

-- 查看所有触发器
SELECT trigger_name FROM information_schema.triggers
WHERE trigger_schema = 'public';

-- 验证表结构
\d t_staff
\d t_scene
\d t_plotting

-- 验证字典数据
SELECT * FROM sys_dict_type;
SELECT * FROM sys_dict_data;
```

## 下一步操作

### 1. 填充测试数据
```sql
-- 插入测试场景
INSERT INTO t_scene (id, scene_no, scene_name, scene_type, type)
VALUES (gen_random_uuid()::varchar(32), 'SC001', '测试场景', '01', '04');

-- 插入测试工作人员
INSERT INTO t_staff (id, name, type, sex, tel, personnel_id)
VALUES (gen_random_uuid()::varchar(32), '张三', 'staff', '1', '13800138000', '110101199001011234');
```

### 2. 配置 Golang 后端
更新 `backend/configs/config.yaml` 中的数据库连接：
```yaml
database:
  default:
    driver: postgres
    host: localhost
    port: 35432
    database: settlement
    username: postgres
    password: STmt0125
```

### 3. 配置 Redis 连接
```yaml
redis:
  default:
    address: localhost:36379
    password: ""
```

### 4. 数据迁移
从 Dameng 数据库迁移数据：
```bash
# 导出 Dameng 数据
dexp userid=SYSDBA/password FILE=dameng_data.dmp LOGS=dm_export.log FULL=Y

# 转换格式并导入 PostgreSQL
# 需要编写转换脚本处理 Dameng 到 PostgreSQL 的数据类型转换
```

## 性能优化建议

1. **索引优化**：根据实际查询频率调整索引
2. **分区表**：对于大数据量表考虑按时间或区域分区
3. **查询优化**：合理使用视图和物化视图
4. **连接池配置**：根据应用负载调整连接池大小

## 备份与恢复

```bash
# 备份数据库
docker exec settlement-monitoring-pg pg_dump -U postgres settlement > settlement_backup.sql

# 恢复数据库
docker exec -i settlement-monitoring-pg psql -U postgres settlement < settlement_backup.sql
```

---

**创建时间**: 2026-01-25
**数据库版本**: PostgreSQL 17.6
**执行状态**: ✅ 成功