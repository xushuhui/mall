-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        8.0.19 - MySQL Community Server - GPL
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  11.0.0.5919
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- 导出  表 qmplus.sys_authorities 结构
CREATE TABLE IF NOT EXISTS `sys_authorities` (
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  `authority_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '角色名',
  `parent_id` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '父角色ID',
  PRIMARY KEY (`authority_id`) USING BTREE,
  UNIQUE KEY `id` (`authority_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  qmplus.sys_authorities 的数据：~3 rows (大约)
/*!40000 ALTER TABLE `sys_authorities` DISABLE KEYS */;
INSERT INTO `sys_authorities` (`created_at`, `updated_at`, `deleted_at`, `authority_id`, `authority_name`, `parent_id`) VALUES
	('2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, '888', '普通用户', '0'),
	('2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, '8881', '普通用户子角色', '888'),
	('2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, '9528', '测试角色', '0');
/*!40000 ALTER TABLE `sys_authorities` ENABLE KEYS */;

-- 导出  表 qmplus.sys_authority_menus 结构
CREATE TABLE IF NOT EXISTS `sys_authority_menus` (
  `sys_authority_authority_id` varchar(90) COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色ID',
  `sys_base_menu_id` bigint unsigned NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  qmplus.sys_authority_menus 的数据：~53 rows (大约)
/*!40000 ALTER TABLE `sys_authority_menus` DISABLE KEYS */;
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`) VALUES
	('888', 1),
	('888', 2),
	('888', 3),
	('888', 4),
	('888', 5),
	('888', 6),
	('888', 7),
	('888', 8),
	('888', 9),
	('888', 10),
	('888', 11),
	('888', 12),
	('888', 13),
	('888', 14),
	('888', 15),
	('888', 16),
	('888', 17),
	('888', 18),
	('888', 19),
	('888', 20),
	('888', 21),
	('888', 22),
	('888', 23),
	('888', 24),
	('888', 25),
	('888', 26),
	('888', 27),
	('8881', 1),
	('8881', 2),
	('8881', 8),
	('8881', 17),
	('8881', 18),
	('8881', 19),
	('8881', 20),
	('9528', 1),
	('9528', 2),
	('9528', 3),
	('9528', 4),
	('9528', 5),
	('9528', 6),
	('9528', 7),
	('9528', 8),
	('9528', 9),
	('9528', 10),
	('9528', 11),
	('9528', 12),
	('9528', 13),
	('9528', 14),
	('9528', 15),
	('9528', 17),
	('9528', 18),
	('9528', 19),
	('9528', 20);
/*!40000 ALTER TABLE `sys_authority_menus` ENABLE KEYS */;

-- 导出  表 qmplus.sys_base_menus 结构
CREATE TABLE IF NOT EXISTS `sys_base_menus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `menu_level` bigint unsigned DEFAULT NULL,
  `parent_id` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint DEFAULT NULL COMMENT '排序标记',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  qmplus.sys_base_menus 的数据：~27 rows (大约)
/*!40000 ALTER TABLE `sys_base_menus` DISABLE KEYS */;
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`) VALUES
	(1, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '0', 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, 0, 0, '仪表盘', 'setting'),
	(2, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '0', 'about', 'about', 0, 'view/about/index.vue', 7, 0, 0, '关于我们', 'info'),
	(3, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '0', 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 3, 0, 0, '超级管理员', 'user-solid'),
	(4, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '3', 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 1, 0, 0, '角色管理', 's-custom'),
	(5, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '3', 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 2, 1, 0, '菜单管理', 's-order'),
	(6, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '3', 'api', 'api', 0, 'view/superAdmin/api/api.vue', 3, 1, 0, 'api管理', 's-platform'),
	(7, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '3', 'user', 'user', 0, 'view/superAdmin/user/user.vue', 4, 0, 0, '用户管理', 'coordinate'),
	(8, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '0', 'person', 'person', 1, 'view/person/person.vue', 4, 0, 0, '个人信息', 'message-solid'),
	(9, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '0', 'example', 'example', 0, 'view/example/index.vue', 6, 0, 0, '示例文件', 's-management'),
	(10, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '9', 'table', 'table', 0, 'view/example/table/table.vue', 1, 0, 0, '表格示例', 's-order'),
	(11, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '9', 'form', 'form', 0, 'view/example/form/form.vue', 2, 0, 0, '表单示例', 'document'),
	(12, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '9', 'rte', 'rte', 0, 'view/example/rte/rte.vue', 3, 0, 0, '富文本编辑器', 'reading'),
	(13, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '9', 'excel', 'excel', 0, 'view/example/excel/excel.vue', 4, 0, 0, 'excel导入导出', 's-marketing'),
	(14, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '9', 'upload', 'upload', 0, 'view/example/upload/upload.vue', 5, 0, 0, '上传下载', 'upload'),
	(15, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '9', 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', 6, 0, 0, '断点续传', 'upload'),
	(16, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '9', 'customer', 'customer', 0, 'view/example/customer/customer.vue', 7, 0, 0, '客户列表（资源示例）', 's-custom'),
	(17, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '0', 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', 5, 0, 0, '系统工具', 's-cooperation'),
	(18, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '17', 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', 1, 1, 0, '代码生成器', 'cpu'),
	(19, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '17', 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', 2, 1, 0, '表单生成器', 'magic-stick'),
	(20, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '17', 'system', 'system', 0, 'view/systemTools/system/system.vue', 3, 0, 0, '系统配置', 's-operation'),
	(21, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '0', 'iconList', 'iconList', 0, 'view/iconList/index.vue', 2, 0, 0, '图标集合', 'star-on'),
	(22, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '3', 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/sysDictionary.vue', 5, 0, 0, '字典管理', 'notebook-2'),
	(23, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '3', 'dictionaryDetail/:id', 'dictionaryDetail', 1, 'view/superAdmin/dictionary/sysDictionaryDetail.vue', 1, 0, 0, '字典详情', 's-order'),
	(24, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '3', 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 6, 0, 0, '操作历史', 'time'),
	(25, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '9', 'simpleUploader', 'simpleUploader', 0, 'view/example/simpleUploader/simpleUploader', 6, 0, 0, '断点续传（插件版）', 'upload'),
	(26, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '0', 'https://www.gin-vue-admin.com', 'https://www.gin-vue-admin.com', 0, '/', 0, 0, 0, '官方网站', 's-home'),
	(27, '2020-11-09 11:50:44', '2020-11-09 11:50:44', NULL, 0, '0', 'state', 'state', 0, 'view/system/state.vue', 6, 0, 0, '服务器状态', 'cloudy');
/*!40000 ALTER TABLE `sys_base_menus` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
