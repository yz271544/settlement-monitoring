-- ========================================
-- JingWei Guard Module - PostgreSQL 17 Schema
-- 从 Dameng 数据库迁移到此 PostgreSQL 17 建表语句
-- ========================================

-- 启用 PostGIS 扩展（用于 GIS 数据）
CREATE EXTENSION IF NOT EXISTS postgis;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ========================================
-- 人员管理表
-- ========================================

-- 工作人员表 (T_STAFF)
CREATE TABLE t_staff (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    type VARCHAR(10) NOT NULL,
    sex CHAR(1) NOT NULL,
    education_background VARCHAR(10),
    post VARCHAR(20),
    birth_time TIMESTAMP,
    dept VARCHAR(255),
    tel VARCHAR(20) NOT NULL,
    picture VARCHAR(100),
    company VARCHAR(100),
    company_name VARCHAR(255),
    duty VARCHAR(255),
    remark VARCHAR(255),
    personnel_id VARCHAR(100) NOT NULL,
    personnel_num VARCHAR(100),
    politics_status VARCHAR(10),
    ethnicity VARCHAR(255),
    nationality VARCHAR(255),
    birthplace VARCHAR(255),
    address VARCHAR(255),
    belong VARCHAR(100),
    proc_ins_id VARCHAR(32),
    merge_status VARCHAR(1) DEFAULT '1',
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted VARCHAR(1) DEFAULT '0'
);

CREATE INDEX idx_staff_personnel_id ON t_staff(personnel_id);
CREATE INDEX idx_staff_tel ON t_staff(tel);
CREATE INDEX idx_staff_type ON t_staff(type);
CREATE INDEX idx_staff_deleted ON t_staff(deleted);

COMMENT ON TABLE t_staff IS '工作人员表';
COMMENT ON COLUMN t_staff.merge_status IS '数据状态（0已合并  1未合并）';

-- 执勤人员表 (T_WATCHKEEPER)
CREATE TABLE t_watchkeeper (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(200),
    sex CHAR(1),
    birth_time TIMESTAMP,
    politics_status VARCHAR(10),
    education_background VARCHAR(10),
    post VARCHAR(20),
    tel VARCHAR(20),
    health_condition VARCHAR(10),
    picture VARCHAR(100),
    company VARCHAR(255),
    company_name VARCHAR(100),
    job_count INTEGER,
    psychological_result VARCHAR(255),
    duty VARCHAR(255),
    dept VARCHAR(255),
    watchkeeper_height VARCHAR(10),
    type VARCHAR(10),
    political_review VARCHAR(10),
    personnel_id VARCHAR(100),
    belong VARCHAR(100),
    proc_ins_id VARCHAR(32),
    merge_status VARCHAR(1) DEFAULT '1',
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    remark VARCHAR(255),
    deleted VARCHAR(1) DEFAULT '0'
);

CREATE INDEX idx_watchkeeper_type ON t_watchkeeper(type);
CREATE INDEX idx_watchkeeper_tel ON t_watchkeeper(tel);
CREATE INDEX idx_watchkeeper_deleted ON t_watchkeeper(deleted);

COMMENT ON TABLE t_watchkeeper IS '执勤人员表';

-- 重点人员表 (T_TARGETED_INDIVIDUALS)
CREATE TABLE t_targeted_individuals (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(10) NOT NULL,
    sex CHAR(1) NOT NULL,
    birthplace VARCHAR(255),
    birth_time TIMESTAMP,
    alias VARCHAR(255),
    company VARCHAR(255),
    police_station VARCHAR(255),
    tel VARCHAR(20) NOT NULL,
    personnel_id VARCHAR(100) NOT NULL,
    warden_name VARCHAR(255),
    warden_tel VARCHAR(255),
    warden_post VARCHAR(255),
    address VARCHAR(255),
    problem_manifestations VARCHAR(255),
    measures VARCHAR(255),
    picture VARCHAR(100),
    remark VARCHAR(255),
    ethnicity VARCHAR(255),
    post VARCHAR(255),
    belong VARCHAR(100),
    proc_ins_id VARCHAR(32),
    merge_status VARCHAR(1) DEFAULT '1',
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted VARCHAR(1) DEFAULT '0'
);

CREATE INDEX idx_targeted_individuals_type ON t_targeted_individuals(type);
CREATE INDEX idx_targeted_individuals_tel ON t_targeted_individuals(tel);
CREATE INDEX idx_targeted_individuals_deleted ON t_targeted_individuals(deleted);

COMMENT ON TABLE t_targeted_individuals IS '重点人员表';

-- 训犬员表 (T_DOG_TRAINER)
CREATE TABLE t_dog_trainer (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    type VARCHAR(10) NOT NULL,
    sex CHAR(1) NOT NULL,
    education_background VARCHAR(10),
    post VARCHAR(20),
    birth_time TIMESTAMP,
    dept VARCHAR(255),
    tel VARCHAR(20) NOT NULL,
    picture VARCHAR(100),
    company VARCHAR(255),
    company_name VARCHAR(255),
    duty VARCHAR(255),
    remark VARCHAR(255),
    personnel_id VARCHAR(100) NOT NULL,
    politics_status VARCHAR(10),
    ethnicity VARCHAR(255),
    nationality VARCHAR(255),
    birthplace VARCHAR(255),
    address VARCHAR(255),
    belong VARCHAR(100),
    proc_ins_id VARCHAR(32),
    merge_status VARCHAR(1) DEFAULT '1',
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted VARCHAR(1) DEFAULT '0'
);

CREATE INDEX idx_dog_trainer_type ON t_dog_trainer(type);
CREATE INDEX idx_dog_trainer_tel ON t_dog_trainer(tel);
CREATE INDEX idx_dog_trainer_deleted ON t_dog_trainer(deleted);

COMMENT ON TABLE t_dog_trainer IS '训犬员表';

-- 警犬表 (T_POLICE_DOG)
CREATE TABLE t_police_dog (
    id VARCHAR(32) PRIMARY KEY,
    personnel_id VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(10) NOT NULL,
    dog_age VARCHAR(10) NOT NULL,
    post VARCHAR(100),
    training_experience VARCHAR(255),
    task_execution_status VARCHAR(255),
    remark VARCHAR(255),
    tel VARCHAR(20),
    picture VARCHAR(100),
    company VARCHAR(255),
    company_name VARCHAR(255),
    belong VARCHAR(100),
    proc_ins_id VARCHAR(32),
    merge_status VARCHAR(1) DEFAULT '1',
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted VARCHAR(1) DEFAULT '0'
);

CREATE INDEX idx_police_dog_trainer_id ON t_police_dog(personnel_id);
CREATE INDEX idx_police_dog_type ON t_police_dog(type);
CREATE INDEX idx_police_dog_deleted ON t_police_dog(deleted);

COMMENT ON TABLE t_police_dog IS '警犬表';

-- ========================================
-- 设备装备表
-- ========================================

-- 车辆表 (T_CAR)
CREATE TABLE t_car (
    id VARCHAR(32) PRIMARY KEY,
    car_number VARCHAR(100) NOT NULL,
    type VARCHAR(10) NOT NULL,
    remark VARCHAR(255),
    company VARCHAR(255),
    company_name VARCHAR(255),
    belong VARCHAR(100),
    proc_ins_id VARCHAR(32),
    merge_status VARCHAR(1) DEFAULT '1',
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted VARCHAR(1) DEFAULT '0'
);

CREATE INDEX idx_car_type ON t_car(type);
CREATE INDEX idx_car_number ON t_car(car_number);
CREATE INDEX idx_car_deleted ON t_car(deleted);

COMMENT ON TABLE t_car IS '车辆表';

-- 枪支表 (T_FIREARM)
CREATE TABLE t_firearm (
    id VARCHAR(32) PRIMARY KEY,
    equipment_id VARCHAR(100) NOT NULL,
    type VARCHAR(10) NOT NULL,
    picture VARCHAR(100),
    company VARCHAR(255),
    company_name VARCHAR(255),
    remark VARCHAR(255),
    belong VARCHAR(100),
    proc_ins_id VARCHAR(32),
    merge_status VARCHAR(1) DEFAULT '1',
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted VARCHAR(1) DEFAULT '0'
);

CREATE INDEX idx_firearm_type ON t_firearm(type);
CREATE INDEX idx_firearm_equipment_id ON t_firearm(equipment_id);
CREATE INDEX idx_firearm_deleted ON t_firearm(deleted);

COMMENT ON TABLE t_firearm IS '枪支表';

-- 对讲机表 (T_INTERPHONE)
CREATE TABLE t_interphone (
    id VARCHAR(32) PRIMARY KEY,
    equipment_id VARCHAR(100) NOT NULL,
    interphone_name VARCHAR(255) NOT NULL,
    type VARCHAR(10) NOT NULL,
    remark VARCHAR(255),
    company VARCHAR(255),
    company_name VARCHAR(255),
    belong VARCHAR(100),
    proc_ins_id VARCHAR(32),
    merge_status VARCHAR(1) DEFAULT '1',
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted VARCHAR(1) DEFAULT '0'
);

CREATE INDEX idx_interphone_type ON t_interphone(type);
CREATE INDEX idx_interphone_equipment_id ON t_interphone(equipment_id);
CREATE INDEX idx_interphone_deleted ON t_interphone(deleted);

COMMENT ON TABLE t_interphone IS '对讲机表';

-- 无人机反制设备表 (T_DRONE_COUNTER_EQUIPMENT)
CREATE TABLE t_drone_counter_equipment (
    id VARCHAR(32) NOT NULL,
    equipment_id VARCHAR(100) NOT NULL,
    equipment_type VARCHAR(100) NOT NULL,
    type VARCHAR(10) NOT NULL,
    equipment_params VARCHAR(255) NOT NULL,
    remark VARCHAR(255),
    picture VARCHAR(100),
    company VARCHAR(255),
    company_name VARCHAR(255),
    belong VARCHAR(100),
    proc_ins_id VARCHAR(32),
    merge_status VARCHAR(1) DEFAULT '1',
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted VARCHAR(1) DEFAULT '0',
    PRIMARY KEY (id, equipment_id)
);

CREATE INDEX idx_drone_counter_type ON t_drone_counter_equipment(type);
CREATE INDEX idx_drone_counter_deleted ON t_drone_counter_equipment(deleted);

COMMENT ON TABLE t_drone_counter_equipment IS '无人机反制设备表';

-- 防爆安检设备表 (T_SECURITY_SCREENING_EQUIPMENT)
CREATE TABLE t_security_screening_equipment (
    id VARCHAR(32) PRIMARY KEY,
    equipment_id VARCHAR(100) NOT NULL,
    equipment_params VARCHAR(255),
    company VARCHAR(255) NOT NULL,
    company_name VARCHAR(255),
    remark VARCHAR(255),
    picture VARCHAR(100),
    equipment_name VARCHAR(255) NOT NULL,
    type VARCHAR(10) NOT NULL,
    belong VARCHAR(100),
    proc_ins_id VARCHAR(32),
    merge_status VARCHAR(1) DEFAULT '1',
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted VARCHAR(1) DEFAULT '0'
);

CREATE INDEX idx_security_screening_type ON t_security_screening_equipment(type);
CREATE INDEX idx_security_screening_equipment_id ON t_security_screening_equipment(equipment_id);
CREATE INDEX idx_security_screening_deleted ON t_security_screening_equipment(deleted);

COMMENT ON TABLE t_security_screening_equipment IS '防爆安检设备表';

-- 执法记录仪表 (T_POLICE_RECORDER)
CREATE TABLE t_police_recorder (
    id VARCHAR(32) PRIMARY KEY,
    equipment_id VARCHAR(100) NOT NULL,
    type VARCHAR(10) NOT NULL,
    picture VARCHAR(100),
    company VARCHAR(255),
    company_name VARCHAR(255),
    remark VARCHAR(255),
    belong VARCHAR(100),
    proc_ins_id VARCHAR(32),
    merge_status VARCHAR(1) DEFAULT '1',
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted VARCHAR(1) DEFAULT '0'
);

CREATE INDEX idx_police_recorder_type ON t_police_recorder(type);
CREATE INDEX idx_police_recorder_equipment_id ON t_police_recorder(equipment_id);
CREATE INDEX idx_police_recorder_deleted ON t_police_recorder(deleted);

COMMENT ON TABLE t_police_recorder IS '执法记录仪表';

-- ========================================
-- 位置地标表
-- ========================================

-- 哨位表 (T_VANTAGE_POINT)
CREATE TABLE t_vantage_point (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    vantage_point_type VARCHAR(100),
    position VARCHAR(255),
    line VARCHAR(100),
    height VARCHAR(100),
    range VARCHAR(100),
    management_unit VARCHAR(255),
    targeted_individuals_num INTEGER,
    belong VARCHAR(100),
    proc_ins_id VARCHAR(32),
    merge_status VARCHAR(1) DEFAULT '1',
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    remark VARCHAR(255),
    deleted VARCHAR(1) DEFAULT '0'
);

CREATE INDEX idx_vantage_point_type ON t_vantage_point(vantage_point_type);
CREATE INDEX idx_vantage_point_deleted ON t_vantage_point(deleted);

COMMENT ON TABLE t_vantage_point IS '哨位表';

-- 重点部位表 (T_KEY_PART)
CREATE TABLE t_key_part (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    key_part_type VARCHAR(100),
    position VARCHAR(255),
    line VARCHAR(100),
    management_unit VARCHAR(255),
    belong VARCHAR(100),
    proc_ins_id VARCHAR(32),
    merge_status VARCHAR(1) DEFAULT '1',
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    remark VARCHAR(255),
    deleted VARCHAR(1) DEFAULT '0'
);

CREATE INDEX idx_key_part_type ON t_key_part(key_part_type);
CREATE INDEX idx_key_part_deleted ON t_key_part(deleted);

COMMENT ON TABLE t_key_part IS '重点部位表';

-- 危险部位表 (T_DANGEROUS_PART)
CREATE TABLE t_dangerous_part (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    dangerous_part_type VARCHAR(100),
    position VARCHAR(255),
    line VARCHAR(100),
    management_unit VARCHAR(255),
    belong VARCHAR(100),
    proc_ins_id VARCHAR(32),
    merge_status VARCHAR(1) DEFAULT '1',
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    remark VARCHAR(255),
    deleted VARCHAR(1) DEFAULT '0'
);

CREATE INDEX idx_dangerous_part_type ON t_dangerous_part(dangerous_part_type);
CREATE INDEX idx_dangerous_part_deleted ON t_dangerous_part(deleted);

COMMENT ON TABLE t_dangerous_part IS '危险部位表';

-- ========================================
-- 流程表 (T_FLOW)
-- ========================================
CREATE TABLE t_flow (
    id VARCHAR(32) PRIMARY KEY,
    type VARCHAR(10) NOT NULL,
    remark VARCHAR(255),
    company VARCHAR(255),
    company_name VARCHAR(255),
    belong VARCHAR(100),
    proc_ins_id VARCHAR(32),
    merge_status VARCHAR(1) DEFAULT '1',
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted VARCHAR(1) DEFAULT '0'
);

CREATE INDEX idx_flow_type ON t_flow(type);
CREATE INDEX idx_flow_deleted ON t_flow(deleted);

COMMENT ON TABLE t_flow IS '流程表';

-- ========================================
-- GIS 区域表 (T_GIS_REGION)
-- ========================================
CREATE TABLE t_gis_region (
    id VARCHAR(32) PRIMARY KEY,
    region_name VARCHAR(255),
    region_code VARCHAR(20),
    region_shape TEXT,
    center_point VARCHAR(50),
    level INTEGER,
    parent_id VARCHAR(32),
    province_code VARCHAR(20),
    province_name VARCHAR(100),
    city_code VARCHAR(20),
    city_name VARCHAR(100),
    county_code VARCHAR(20),
    county_name VARCHAR(100),
    township_code VARCHAR(20),
    township_name VARCHAR(100),
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted VARCHAR(1) DEFAULT '0'
);

CREATE INDEX idx_gis_region_code ON t_gis_region(region_code);
CREATE INDEX idx_gis_region_level ON t_gis_region(level);
CREATE INDEX idx_gis_region_parent ON t_gis_region(parent_id);
CREATE INDEX idx_gis_region_deleted ON t_gis_region(deleted);

COMMENT ON TABLE t_gis_region IS 'GIS 区域表';
COMMENT ON COLUMN t_gis_region.region_shape IS '区域形状（GeoJSON 格式）';

-- ========================================
-- 场景相关表 (jingwei-ztuo-guard 模块)
-- ========================================

-- 场景表 (t_scene)
CREATE TABLE t_scene (
    id VARCHAR(32) PRIMARY KEY,
    scene_no VARCHAR(50),
    scene_name VARCHAR(200),
    org_id VARCHAR(32),
    org_name VARCHAR(255),
    contact_name VARCHAR(100),
    contact_phone VARCHAR(50),
    scene_description TEXT,
    scene_cover VARCHAR(500),
    scene_type VARCHAR(10),
    province_code VARCHAR(20),
    province_name VARCHAR(100),
    city_code VARCHAR(20),
    city_name VARCHAR(100),
    county_code VARCHAR(20),
    county_name VARCHAR(100),
    township_code VARCHAR(20),
    township_name VARCHAR(100),
    region_shape TEXT,
    kml_name VARCHAR(255),
    service_coordinate VARCHAR(255),
    status VARCHAR(20),
    center_point VARCHAR(50),
    envelope TEXT,
    type VARCHAR(10),
    neighborhood_situation TEXT,
    del_flag BOOLEAN DEFAULT FALSE,
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    remark VARCHAR(255)
);

CREATE INDEX idx_scene_no ON t_scene(scene_no);
CREATE INDEX idx_scene_type ON t_scene(type);
CREATE INDEX idx_scene_del_flag ON t_scene(del_flag);

COMMENT ON TABLE t_scene IS '场景表';
COMMENT ON COLUMN t_scene.region_shape IS '行政区划范围（GeoJSON 格式）';

-- 驻地表 (t_residence)
CREATE TABLE t_residence (
    id VARCHAR(32) PRIMARY KEY,
    scene_id VARCHAR(32) NOT NULL,
    name VARCHAR(200),
    unit_property VARCHAR(100),
    unit_leader VARCHAR(100),
    phone VARCHAR(50),
    id_card VARCHAR(50),
    address VARCHAR(500),
    star_level VARCHAR(20),
    person_number VARCHAR(20),
    foreigner_number VARCHAR(20),
    region_number VARCHAR(20),
    building VARCHAR(200),
    height VARCHAR(50),
    area VARCHAR(50),
    room_number VARCHAR(20),
    internal_person VARCHAR(100),
    internal_phone VARCHAR(50),
    internal_id_card VARCHAR(50),
    police_station VARCHAR(255),
    police_person VARCHAR(100),
    police_phone VARCHAR(50),
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    remark VARCHAR(255)
);

CREATE INDEX idx_residence_scene_id ON t_residence(scene_id);

COMMENT ON TABLE t_residence IS '驻地表';

-- 标绘表 (t_plotting)
CREATE TABLE t_plotting (
    id VARCHAR(32) PRIMARY KEY,
    scene_id VARCHAR(32),
    plotting_type INTEGER,
    coverage_id VARCHAR(32),
    coverage_code VARCHAR(50),
    basics_properties_json TEXT,
    extend_properties_json TEXT,
    name VARCHAR(255),
    license_plate_number VARCHAR(100),
    longitude_latitude VARCHAR(100),
    shape TEXT,
    style_flag INTEGER DEFAULT 0,
    layer_flag INTEGER DEFAULT 0,
    style_info_json TEXT,
    reality_images_one VARCHAR(500),
    reality_images_two VARCHAR(500),
    range_images VARCHAR(500),
    images VARCHAR(500),
    sort INTEGER DEFAULT 0,
    watch_person VARCHAR(500),
    equipment VARCHAR(500),
    del_flag BOOLEAN DEFAULT FALSE,
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_plotting_scene_id ON t_plotting(scene_id);
CREATE INDEX idx_plotting_coverage_id ON t_plotting(coverage_id);
CREATE INDEX idx_plotting_type ON t_plotting(plotting_type);
CREATE INDEX idx_plotting_del_flag ON t_plotting(del_flag);

COMMENT ON TABLE t_plotting IS '标绘表';
COMMENT ON COLUMN t_plotting.shape IS '坐标信息（GeoJSON 格式）';

-- 摄像头表 (t_camera)
CREATE TABLE t_camera (
    id VARCHAR(32) PRIMARY KEY,
    category VARCHAR(10),
    camera_type VARCHAR(50),
    name VARCHAR(255),
    status VARCHAR(10) DEFAULT '0',
    camera_index_code VARCHAR(100),
    longitude VARCHAR(50),
    latitude VARCHAR(50),
    altitude VARCHAR(50),
    install_place VARCHAR(255),
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    remark VARCHAR(255),
    del_flag BIGINT DEFAULT 0
);

CREATE INDEX idx_camera_category ON t_camera(category);
CREATE INDEX idx_camera_status ON t_camera(status);
CREATE INDEX idx_camera_del_flag ON t_camera(del_flag);

COMMENT ON TABLE t_camera IS '摄像头表';

-- 装备信息表 (t_equipment)
CREATE TABLE t_equipment (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(255),
    type VARCHAR(10),
    serial_number VARCHAR(100),
    model VARCHAR(100),
    effective_range VARCHAR(100),
    region VARCHAR(50),
    unit VARCHAR(255),
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    remark VARCHAR(255)
);

CREATE INDEX idx_equipment_type ON t_equipment(type);
CREATE INDEX idx_equipment_region ON t_equipment(region);

COMMENT ON TABLE t_equipment IS '装备信息表';

-- 人员表 (t_person)
CREATE TABLE t_person (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(100),
    sex VARCHAR(10),
    nationality VARCHAR(50),
    nation VARCHAR(50),
    native_place VARCHAR(255),
    birth_time TIMESTAMP,
    politics_status VARCHAR(50),
    education_background VARCHAR(50),
    post VARCHAR(100),
    tel VARCHAR(50),
    driving_type VARCHAR(50),
    driving_age INTEGER,
    heath_condition VARCHAR(50),
    picture VARCHAR(500),
    company VARCHAR(255),
    ever_post VARCHAR(255),
    job_resume VARCHAR(255),
    training_record VARCHAR(255),
    service_condition VARCHAR(255),
    recommend_unit_opinion VARCHAR(255),
    recommend_unit_opinion_date TIMESTAMP,
    political_dept_opinion VARCHAR(255),
    political_dept_opinion_date TIMESTAMP,
    police_dept_opinion VARCHAR(255),
    police_dept_opinion_date TIMESTAMP,
    lead_name VARCHAR(100),
    lead_tel VARCHAR(50),
    person_type VARCHAR(10),
    nature VARCHAR(100),
    scene_id VARCHAR(32),
    alias VARCHAR(255),
    card VARCHAR(50),
    address VARCHAR(500),
    physical_result VARCHAR(255),
    psychological_result VARCHAR(255),
    job_count INTEGER,
    person_attribute VARCHAR(255),
    dept VARCHAR(255),
    duty VARCHAR(255),
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_person_type ON t_person(person_type);
CREATE INDEX idx_person_scene_id ON t_person(scene_id);

COMMENT ON TABLE t_person IS '人员维护表';

-- 三维模型表 (t_three_dimension_model)
CREATE TABLE t_three_dimension_model (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(255),
    model_type VARCHAR(50),
    model_url VARCHAR(500),
    thumbnail_url VARCHAR(500),
    position VARCHAR(100),
    rotation VARCHAR(100),
    scale VARCHAR(100),
    scene_id VARCHAR(32),
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    remark VARCHAR(255)
);

CREATE INDEX idx_3d_model_scene_id ON t_three_dimension_model(scene_id);

COMMENT ON TABLE t_three_dimension_model IS '三维模型表';

-- VR 信息表 (t_vr_info)
CREATE TABLE t_vr_info (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(255),
    vr_type VARCHAR(50),
    vr_url VARCHAR(500),
    thumbnail_url VARCHAR(500),
    scene_id VARCHAR(32),
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    remark VARCHAR(255)
);

CREATE INDEX idx_vr_info_scene_id ON t_vr_info(scene_id);

COMMENT ON TABLE t_vr_info IS 'VR 信息表';

-- 图例表 (t_legend)
CREATE TABLE t_legend (
    id VARCHAR(32) PRIMARY KEY,
    legend_code VARCHAR(50) NOT NULL,
    legend_name VARCHAR(255),
    legend_type VARCHAR(50),
    icon_url VARCHAR(500),
    color VARCHAR(50),
    layer_name VARCHAR(255),
    sort INTEGER DEFAULT 0,
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    remark VARCHAR(255)
);

CREATE INDEX idx_legend_code ON t_legend(legend_code);

COMMENT ON TABLE t_legend IS '图例表';

-- 覆盖类型树表 (t_coverage_tree)
CREATE TABLE t_coverage_tree (
    id VARCHAR(32) PRIMARY KEY,
    parent_id VARCHAR(32),
    code VARCHAR(50) NOT NULL,
    name VARCHAR(255),
    type VARCHAR(50),
    sort INTEGER DEFAULT 0,
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    remark VARCHAR(255)
);

CREATE INDEX idx_coverage_tree_parent ON t_coverage_tree(parent_id);
CREATE INDEX idx_coverage_tree_code ON t_coverage_tree(code);

COMMENT ON TABLE t_coverage_tree IS '覆盖类型树表';

-- 场景地图表 (t_scene_map)
CREATE TABLE t_scene_map (
    id VARCHAR(32) PRIMARY KEY,
    scene_id VARCHAR(32),
    map_name VARCHAR(255),
    map_type VARCHAR(50),
    map_url VARCHAR(500),
    map_config_json TEXT,
    sort INTEGER DEFAULT 0,
    is_default BOOLEAN DEFAULT FALSE,
    load_flag INTEGER DEFAULT 1,
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    remark VARCHAR(255)
);

CREATE INDEX idx_scene_map_scene_id ON t_scene_map(scene_id);

COMMENT ON TABLE t_scene_map IS '场景地图表';

-- 演练路线表 (t_simulation_route)
CREATE TABLE t_simulation_route (
    id VARCHAR(32) PRIMARY KEY,
    route_name VARCHAR(255),
    route_type VARCHAR(50),
    route_shape TEXT,
    route_points_json TEXT,
    scene_id VARCHAR(32),
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    remark VARCHAR(255)
);

CREATE INDEX idx_simulation_route_scene_id ON t_simulation_route(scene_id);

COMMENT ON TABLE t_simulation_route IS '演练路线表';
COMMENT ON COLUMN t_simulation_route.route_shape IS '路线形状（GeoJSON 格式）';

-- ========================================
-- 系统字典数据迁移
-- ========================================

-- 字典类型表 (sys_dict_type)
CREATE TABLE IF NOT EXISTS sys_dict_type (
    dict_id VARCHAR(32) PRIMARY KEY,
    tenant_id VARCHAR(20) DEFAULT '000000',
    dict_name VARCHAR(100) DEFAULT '',
    dict_type VARCHAR(100) DEFAULT '',
    status CHAR(1) DEFAULT '0',
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    remark VARCHAR(500)
);

INSERT INTO sys_dict_type (dict_id, dict_name, dict_type, status) VALUES
('1001', '人员类型', 'guard_person_type', '0'),
('1002', '车辆类型', 'guard_car_type', '0'),
('1003', '装备类型', 'guard_equipment_type', '0'),
('1004', '设备类型', 'guard_device_type', '0'),
('1005', '场景类型', 'guard_scene_type', '0'),
('1006', '标绘类型', 'guard_plotting_type', '0'),
('1007', '安检设备类型', 'guard_screening_type', '0');

-- 字典数据表 (sys_dict_data)
CREATE TABLE IF NOT EXISTS sys_dict_data (
    dict_code VARCHAR(32) PRIMARY KEY,
    tenant_id VARCHAR(20) DEFAULT '000000',
    dict_sort INTEGER DEFAULT 0,
    dict_label VARCHAR(100) DEFAULT '',
    dict_value VARCHAR(100) DEFAULT '',
    dict_type VARCHAR(100) DEFAULT '',
    css_class VARCHAR(100),
    list_class VARCHAR(100),
    is_default CHAR(1) DEFAULT 'N',
    status CHAR(1) DEFAULT '0',
    create_by VARCHAR(64) DEFAULT '',
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_by VARCHAR(64) DEFAULT '',
    update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    remark VARCHAR(500)
);

-- 人员类型字典数据
INSERT INTO sys_dict_data (dict_code, dict_sort, dict_label, dict_value, dict_type, status) VALUES
('100101', 1, '执勤人员', 'watchkeeper', 'guard_person_type', '0'),
('100102', 2, '工作人员', 'staff', 'guard_person_type', '0'),
('100103', 3, '训犬员', 'dog_trainer', 'guard_person_type', '0'),
('100104', 4, '重点人员', 'targeted_individuals', 'guard_person_type', '0'),
('100105', 5, '警犬', 'police_dog', 'guard_person_type', '0');

-- 车辆类型字典数据
INSERT INTO sys_dict_data (dict_code, dict_sort, dict_label, dict_value, dict_type, status) VALUES
('100201', 1, '小汽车', '01', 'guard_car_type', '0'),
('100202', 2, '卡车', '02', 'guard_car_type', '0'),
('100203', 3, '消防车', '03', 'guard_car_type', '0'),
('100204', 4, '民警车', '04', 'guard_car_type', '0'),
('100205', 5, '巡警车', '05', 'guard_car_type', '0'),
('100206', 6, '武警车', '06', 'guard_car_type', '0'),
('100207', 7, '特警车', '07', 'guard_car_type', '0'),
('100208', 8, '救护车', '08', 'guard_car_type', '0'),
('100209', 9, '摩托车', '09', 'guard_car_type', '0');

-- 场景类型字典数据
INSERT INTO sys_dict_data (dict_code, dict_sort, dict_label, dict_value, dict_type, status) VALUES
('100501', 1, '行政边界', '01', 'guard_scene_type', '0'),
('100502', 2, 'KML文件', '02', 'guard_scene_type', '0'),
('100503', 3, '驻地', '04', 'guard_scene_type', '0');

-- 标绘类型字典数据
INSERT INTO sys_dict_data (dict_code, dict_sort, dict_label, dict_value, dict_type, status) VALUES
('100601', 1, '点', '0', 'guard_plotting_type', '0'),
('100602', 2, '线', '1', 'guard_plotting_type', '0'),
('100603', 3, '面', '2', 'guard_plotting_type', '0');

-- ========================================
-- 创建触发器：自动更新 update_time
-- ========================================

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.update_time = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 为所有需要更新的表创建触发器
CREATE TRIGGER t_staff_update_time BEFORE UPDATE ON t_staff
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_watchkeeper_update_time BEFORE UPDATE ON t_watchkeeper
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_targeted_individuals_update_time BEFORE UPDATE ON t_targeted_individuals
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_dog_trainer_update_time BEFORE UPDATE ON t_dog_trainer
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_police_dog_update_time BEFORE UPDATE ON t_police_dog
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_car_update_time BEFORE UPDATE ON t_car
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_firearm_update_time BEFORE UPDATE ON t_firearm
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_interphone_update_time BEFORE UPDATE ON t_interphone
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_drone_counter_equipment_update_time BEFORE UPDATE ON t_drone_counter_equipment
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_security_screening_equipment_update_time BEFORE UPDATE ON t_security_screening_equipment
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_police_recorder_update_time BEFORE UPDATE ON t_police_recorder
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_vantage_point_update_time BEFORE UPDATE ON t_vantage_point
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_key_part_update_time BEFORE UPDATE ON t_key_part
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_dangerous_part_update_time BEFORE UPDATE ON t_dangerous_part
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_flow_update_time BEFORE UPDATE ON t_flow
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_gis_region_update_time BEFORE UPDATE ON t_gis_region
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_scene_update_time BEFORE UPDATE ON t_scene
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_residence_update_time BEFORE UPDATE ON t_residence
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_plotting_update_time BEFORE UPDATE ON t_plotting
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_camera_update_time BEFORE UPDATE ON t_camera
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_equipment_update_time BEFORE UPDATE ON t_equipment
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_person_update_time BEFORE UPDATE ON t_person
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_three_dimension_model_update_time BEFORE UPDATE ON t_three_dimension_model
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_vr_info_update_time BEFORE UPDATE ON t_vr_info
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_legend_update_time BEFORE UPDATE ON t_legend
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_coverage_tree_update_time BEFORE UPDATE ON t_coverage_tree
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_scene_map_update_time BEFORE UPDATE ON t_scene_map
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER t_simulation_route_update_time BEFORE UPDATE ON t_simulation_route
    FOR EACH ROW EXECUTE FUNCTION update_timestamp();

-- ========================================
-- 创建视图：常用查询
-- ========================================

-- 场景统计视图
CREATE OR REPLACE VIEW v_scene_statistics AS
SELECT
    s.id AS scene_id,
    s.scene_name,
    s.scene_no,
    s.type AS scene_type,
    COUNT(DISTINCT p.id) AS plotting_count,
    COUNT(DISTINCT ps.id) AS person_count,
    s.create_time
FROM t_scene s
LEFT JOIN t_plotting p ON s.id = p.scene_id AND p.del_flag = FALSE
LEFT JOIN t_person ps ON s.id = ps.scene_id
WHERE s.del_flag = FALSE
GROUP BY s.id, s.scene_name, s.scene_no, s.type, s.create_time;

COMMENT ON VIEW v_scene_statistics IS '场景统计视图';

-- 按场景类型统计标绘数量
CREATE OR REPLACE VIEW v_plotting_by_scene_type AS
SELECT
    s.type AS scene_type,
    COALESCE(p.plotting_type, 0) AS plotting_type,
    COUNT(*) AS count
FROM t_scene s
LEFT JOIN t_plotting p ON s.id = p.scene_id AND p.del_flag = FALSE
WHERE s.del_flag = FALSE
GROUP BY s.type, p.plotting_type
ORDER BY s.type, p.plotting_type;

COMMENT ON VIEW v_plotting_by_scene_type IS '按场景类型统计标绘数量视图';

-- 人员统计视图
CREATE OR REPLACE VIEW v_person_statistics AS
SELECT
    't_staff' AS person_type,
    COUNT(*) AS count
FROM t_staff WHERE deleted = '0'
UNION ALL
SELECT
    't_watchkeeper' AS person_type,
    COUNT(*) AS count
FROM t_watchkeeper WHERE deleted = '0'
UNION ALL
SELECT
    't_targeted_individuals' AS person_type,
    COUNT(*) AS count
FROM t_targeted_individuals WHERE deleted = '0'
UNION ALL
SELECT
    't_dog_trainer' AS person_type,
    COUNT(*) AS count
FROM t_dog_trainer WHERE deleted = '0';

COMMENT ON VIEW v_person_statistics IS '人员统计视图';

-- ========================================
-- 脚本完成
-- ========================================