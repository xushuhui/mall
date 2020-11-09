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

-- 导出  表 mall.activity 结构
CREATE TABLE IF NOT EXISTS `activity` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `start_time` datetime(3) NOT NULL,
  `end_time` datetime(3) NOT NULL,
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `delete_time` datetime(3) DEFAULT NULL,
  `remark` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `online` tinyint unsigned DEFAULT '1',
  `entrance_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `internal_top_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.activity 的数据：~1 rows (大约)
/*!40000 ALTER TABLE `activity` DISABLE KEYS */;
INSERT INTO `activity` (`id`, `title`, `description`, `start_time`, `end_time`, `create_time`, `update_time`, `delete_time`, `remark`, `online`, `entrance_img`, `internal_top_img`, `name`) VALUES
	(2, '夏日好礼送不停', '长夏村墟风日清', '2019-08-03 18:04:52.000', '2030-08-31 18:05:16.000', '2019-08-03 17:59:01.000', '2020-11-04 15:58:03.163', NULL, '限服装、鞋、文具等商品', 1, 'http://i2.sleeve.talelin.com/m14.png', NULL, 'a-2');
/*!40000 ALTER TABLE `activity` ENABLE KEYS */;

-- 导出  表 mall.activity_category 结构
CREATE TABLE IF NOT EXISTS `activity_category` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `category_id` int unsigned NOT NULL,
  `activity_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.activity_category 的数据：~11 rows (大约)
/*!40000 ALTER TABLE `activity_category` DISABLE KEYS */;
INSERT INTO `activity_category` (`id`, `category_id`, `activity_id`) VALUES
	(1, 2, 2),
	(2, 7, 2),
	(6, 4, 2),
	(7, 27, 2),
	(8, 32, 2),
	(9, 27, 1),
	(11, 1, 3),
	(12, 2, 3),
	(13, 1, 4),
	(14, 2, 4),
	(15, 3, 4);
/*!40000 ALTER TABLE `activity_category` ENABLE KEYS */;

-- 导出  表 mall.activity_coupon 结构
CREATE TABLE IF NOT EXISTS `activity_coupon` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `coupon_id` int unsigned NOT NULL,
  `activity_id` int unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.activity_coupon 的数据：~9 rows (大约)
/*!40000 ALTER TABLE `activity_coupon` DISABLE KEYS */;
INSERT INTO `activity_coupon` (`id`, `coupon_id`, `activity_id`) VALUES
	(1, 3, 2),
	(2, 4, 2),
	(3, 5, 2),
	(4, 7, 1),
	(6, 3, 3),
	(7, 4, 3),
	(8, 3, 4),
	(9, 4, 4),
	(10, 5, 4);
/*!40000 ALTER TABLE `activity_coupon` ENABLE KEYS */;

-- 导出  表 mall.activity_spu 结构
CREATE TABLE IF NOT EXISTS `activity_spu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `activity_id` int unsigned NOT NULL,
  `spu_id` int unsigned NOT NULL,
  `participation` tinyint unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.activity_spu 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `activity_spu` DISABLE KEYS */;
/*!40000 ALTER TABLE `activity_spu` ENABLE KEYS */;

-- 导出  表 mall.banner 结构
CREATE TABLE IF NOT EXISTS `banner` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `delete_time` datetime(3) DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '部分banner可能有标题图片',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.banner 的数据：~2 rows (大约)
/*!40000 ALTER TABLE `banner` DISABLE KEYS */;
INSERT INTO `banner` (`id`, `name`, `description`, `create_time`, `update_time`, `delete_time`, `title`, `img`) VALUES
	(1, 'b-1', '首页顶部主banner', '2019-07-28 04:47:15.000', '2019-08-04 01:03:16.000', NULL, NULL, NULL),
	(2, 'b-2', '热销商品banner', '2019-08-01 00:37:47.000', '2020-11-04 16:02:30.006', NULL, NULL, 'http://i2.sleeve.talelin.com/m4.png');
/*!40000 ALTER TABLE `banner` ENABLE KEYS */;

-- 导出  表 mall.banner_item 结构
CREATE TABLE IF NOT EXISTS `banner_item` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `keyword` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `type` smallint unsigned NOT NULL DEFAULT '0',
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `delete_time` datetime(3) DEFAULT NULL,
  `banner_id` int unsigned NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.banner_item 的数据：~6 rows (大约)
/*!40000 ALTER TABLE `banner_item` DISABLE KEYS */;
INSERT INTO `banner_item` (`id`, `img`, `keyword`, `type`, `create_time`, `update_time`, `delete_time`, `banner_id`, `name`) VALUES
	(5, 'http://i2.sleeve.talelin.com/m6.png', '28', 1, '2019-08-01 00:41:41.000', '2020-11-04 16:03:04.821', NULL, 2, 'left'),
	(6, 'http://i2.sleeve.talelin.com/m7.png', '26', 1, '2019-08-01 00:41:41.000', '2020-11-04 16:03:14.214', NULL, 2, 'right-top'),
	(7, 'http://i2.sleeve.talelin.com/m8.png', '27', 1, '2019-08-01 00:41:41.000', '2020-11-04 16:03:21.092', NULL, 2, 'right-bottom'),
	(12, 'http://i2.sleeve.talelin.com/m1.pnghttp://i2.sleeve.talelin.com/m1.png', 't-2', 3, '2019-09-15 17:29:52.000', '2020-11-04 16:03:37.397', NULL, 1, NULL),
	(13, 'http://i1.sleeve.7yue.pro/assets/702f2ce9-5729-4aa4-aeb3-921513327747.png', '23', 1, '2019-07-28 04:39:22.000', '2020-11-04 16:03:51.181', NULL, 1, NULL),
	(14, 'http://i1.sleeve.7yue.pro/assets/b8e510a1-8340-43c2-a4b0-0e56a40256f9.png', '24', 1, '2019-07-28 04:40:10.000', '2020-11-04 16:04:00.388', NULL, 1, NULL);
/*!40000 ALTER TABLE `banner_item` ENABLE KEYS */;

-- 导出  表 mall.brand 结构
CREATE TABLE IF NOT EXISTS `brand` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.brand 的数据：~17 rows (大约)
/*!40000 ALTER TABLE `brand` DISABLE KEYS */;
INSERT INTO `brand` (`id`, `name`, `description`, `create_time`, `update_time`, `delete_time`) VALUES
	(1, 'Ikea', NULL, '2019-07-15 12:42:43', '2019-07-15 12:54:11', NULL),
	(2, 'Theory', NULL, '2019-07-15 12:46:00', '2019-07-15 12:46:00', NULL),
	(3, 'Converse', NULL, '2019-07-15 12:47:00', '2019-07-15 12:47:00', NULL),
	(4, 'Hickies', NULL, '2019-07-15 12:48:13', '2019-07-15 12:48:13', NULL),
	(5, 'Bucketfeet', NULL, '2019-07-15 12:48:55', '2019-07-15 12:48:55', NULL),
	(6, 'ROUJE\nROUJE\n', NULL, '2019-07-15 12:53:21', '2019-07-15 12:53:21', NULL),
	(7, 'Claudie Pierlot', NULL, '2019-07-15 12:53:40', '2019-07-15 12:53:40', NULL),
	(8, 'Lemaire&UNIQLIO', NULL, '2019-07-15 12:53:51', '2019-07-15 12:54:02', NULL),
	(9, 'The Kooples', NULL, '2019-07-15 12:54:39', '2019-07-15 12:54:39', NULL),
	(10, 'Marc Jacobs', NULL, '2019-07-15 12:55:08', '2019-07-15 12:55:08', NULL),
	(11, 'Michael Kors', NULL, '2019-07-15 12:55:17', '2019-07-15 12:55:17', NULL),
	(12, 'Tory Burch', NULL, '2019-07-15 12:55:25', '2019-07-15 12:55:25', NULL),
	(13, 'Kate Spade', NULL, '2019-07-15 12:55:43', '2019-07-15 12:55:43', NULL),
	(14, 'Apple', NULL, '2019-07-17 07:59:59', '2019-07-17 07:59:59', NULL),
	(15, 'Staedtler', NULL, '2019-07-17 08:00:30', '2019-07-17 08:00:30', NULL),
	(16, '林间有风', NULL, '2019-07-17 08:19:36', '2019-07-17 08:19:36', NULL),
	(17, '九州', '请打', '2019-08-26 22:08:20', '2019-08-26 22:08:20', NULL);
/*!40000 ALTER TABLE `brand` ENABLE KEYS */;

-- 导出  表 mall.category 结构
CREATE TABLE IF NOT EXISTS `category` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `delete_time` datetime(3) DEFAULT NULL,
  `is_root` tinyint unsigned NOT NULL DEFAULT '0',
  `parent_id` int unsigned DEFAULT NULL,
  `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `index` int unsigned DEFAULT NULL,
  `online` int unsigned DEFAULT '1',
  `level` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.category 的数据：~39 rows (大约)
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` (`id`, `name`, `description`, `create_time`, `update_time`, `delete_time`, `is_root`, `parent_id`, `img`, `index`, `online`, `level`) VALUES
	(1, '鞋', NULL, '2019-07-15 08:51:19.000', '2020-11-04 16:05:16.647', NULL, 1, NULL, 'http://i1.sleeve.7yue.pro/grid/shoes.png', 3, 1, NULL),
	(2, '服装', '', '2019-07-15 08:51:28.000', '2020-11-04 16:05:01.687', NULL, 1, NULL, 'http://i1.sleeve.7yue.pro/grid/clothing.png', 2, 1, NULL),
	(3, '包包', NULL, '2019-07-15 08:51:35.000', '2020-11-04 16:05:09.062', NULL, 1, NULL, 'http://i1.sleeve.7yue.pro/grid/bag.png', 1, 1, NULL),
	(4, '居家', NULL, '2019-07-15 08:51:42.000', '2020-11-04 16:05:41.200', NULL, 1, NULL, 'http://i1.sleeve.7yue.pro/grid/furniture.png', 5, 1, NULL),
	(5, '饰品', NULL, '2019-07-15 08:51:49.000', '2020-11-04 16:05:29.375', NULL, 1, NULL, 'http://i1.sleeve.7yue.pro/grid/jewelry.png', 4, 1, NULL),
	(6, '平底鞋', NULL, '2019-07-15 08:51:55.000', '2019-09-21 23:32:00.822', NULL, 0, 1, NULL, NULL, 1, NULL),
	(7, '凉鞋', NULL, '2019-07-15 08:52:47.000', '2019-09-21 23:32:04.346', NULL, 0, 1, NULL, NULL, 1, NULL),
	(8, '拖鞋', NULL, '2019-07-15 08:53:04.000', '2019-09-21 23:37:12.047', NULL, 0, 1, NULL, NULL, 1, NULL),
	(9, '运动鞋', NULL, '2019-07-15 08:53:23.000', '2019-09-21 23:37:18.707', NULL, 0, 1, NULL, NULL, 1, NULL),
	(10, '高跟鞋', NULL, '2019-07-15 08:53:51.000', '2019-09-21 23:37:26.999', NULL, 0, 1, NULL, NULL, 1, NULL),
	(11, '衬衫', NULL, '2019-07-15 08:54:41.000', '2019-09-21 23:37:40.172', NULL, 0, 2, NULL, NULL, 1, NULL),
	(12, 'T恤', NULL, '2019-07-15 08:55:11.000', '2019-09-21 23:37:47.560', NULL, 0, 2, NULL, NULL, 1, NULL),
	(13, '牛仔裤', NULL, '2019-07-15 08:55:58.000', '2019-09-21 23:37:50.095', NULL, 0, 2, NULL, NULL, 1, NULL),
	(14, '针织衫', NULL, '2019-07-15 08:56:27.000', '2019-09-21 23:37:55.138', NULL, 0, 2, NULL, NULL, 1, NULL),
	(15, '连衣裙', NULL, '2019-07-15 08:56:43.000', '2019-09-21 23:38:08.140', NULL, 0, 2, NULL, NULL, 1, NULL),
	(16, '风衣', NULL, '2019-07-15 08:57:38.000', '2019-09-21 23:38:10.184', NULL, 0, 2, NULL, NULL, 1, NULL),
	(17, '手包', NULL, '2019-07-15 08:58:12.000', '2019-09-21 23:38:17.871', NULL, 0, 3, NULL, NULL, 1, NULL),
	(18, '旅行包', NULL, '2019-07-15 08:58:38.000', '2019-09-21 23:38:24.079', NULL, 0, 3, NULL, NULL, 1, NULL),
	(19, '单肩包', NULL, '2019-07-15 08:58:51.000', '2019-09-21 23:38:28.196', NULL, 0, 3, NULL, NULL, 1, NULL),
	(20, '收纳', NULL, '2019-07-15 09:00:19.000', '2019-09-21 23:31:27.823', NULL, 0, 4, NULL, NULL, 1, NULL),
	(21, '毛巾', NULL, '2019-07-15 09:01:38.000', '2019-09-21 23:31:32.074', NULL, 0, 4, NULL, NULL, 1, NULL),
	(22, '四件套', NULL, '2019-07-15 09:04:52.000', '2019-09-21 23:31:35.949', NULL, 0, 4, NULL, NULL, 1, NULL),
	(23, '台灯', NULL, '2019-07-15 14:18:40.000', '2019-09-21 23:31:45.456', NULL, 0, 4, NULL, NULL, 1, NULL),
	(24, '工艺', NULL, '2019-07-15 14:20:05.000', '2020-11-04 16:06:09.730', NULL, 1, NULL, 'http://i1.sleeve.7yue.pro/grid/book.pnghttp://i1.sleeve.7yue.pro/grid/book.png', 1, 1, NULL),
	(25, '玻璃杯', NULL, '2019-07-15 14:23:01.000', '2019-09-21 23:40:47.698', NULL, 0, 27, NULL, NULL, 1, NULL),
	(26, '桌布', NULL, '2019-07-15 14:26:25.000', '2019-07-15 14:32:14.000', NULL, 0, 27, NULL, NULL, 1, NULL),
	(27, '餐边', NULL, '2019-07-15 14:32:01.000', '2019-09-11 14:35:36.000', NULL, 1, NULL, NULL, 6, 0, NULL),
	(28, '盘和碗', NULL, '2019-07-15 14:40:23.000', '2019-07-15 14:43:47.000', NULL, 0, 27, NULL, NULL, 1, NULL),
	(29, '数码', NULL, '2019-07-17 07:51:07.000', '2019-09-11 14:36:54.000', NULL, 1, NULL, NULL, 7, 0, NULL),
	(30, '笔电', NULL, '2019-07-17 07:51:57.000', '2019-07-17 07:52:25.000', NULL, 0, 29, NULL, NULL, 1, NULL),
	(31, '手机', NULL, '2019-07-17 08:01:59.000', '2019-07-17 08:02:02.000', NULL, 0, 29, NULL, NULL, 1, NULL),
	(32, '文具', NULL, '2019-07-30 01:29:19.000', '2019-09-22 00:14:40.026', NULL, 0, 24, NULL, NULL, 1, NULL),
	(33, '女妆', NULL, '2019-08-02 02:54:11.000', '2019-09-11 14:35:27.000', NULL, 1, NULL, NULL, 8, 0, NULL),
	(34, '香水', NULL, '2019-08-02 02:54:30.000', '2019-08-02 02:54:30.000', NULL, 0, 33, NULL, NULL, 1, NULL),
	(35, '家具', NULL, '2019-08-02 19:44:24.000', '2019-09-21 23:40:27.664', NULL, 0, 4, NULL, 9, 1, NULL),
	(36, '酷玩', NULL, '2019-08-07 22:49:22.000', '2019-09-21 23:38:50.402', NULL, 0, 5, NULL, NULL, 1, NULL),
	(37, '测试数据', NULL, '2019-08-25 19:06:24.000', '2019-09-21 23:39:21.756', NULL, 1, NULL, NULL, NULL, 1, NULL),
	(38, '测试数据', NULL, '2019-08-25 19:07:53.000', '2019-09-21 23:38:52.170', NULL, 0, 37, NULL, NULL, 1, NULL),
	(39, '束带', NULL, '2019-09-07 16:43:29.000', '2019-09-21 23:39:47.444', NULL, 0, 5, NULL, NULL, 1, NULL);
/*!40000 ALTER TABLE `category` ENABLE KEYS */;

-- 导出  表 mall.coupon 结构
CREATE TABLE IF NOT EXISTS `coupon` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `start_time` datetime DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `full_money` decimal(10,2) DEFAULT NULL,
  `minus` decimal(10,2) DEFAULT NULL,
  `rate` decimal(10,2) DEFAULT NULL COMMENT '国内多是打折，国外多是百分比 off',
  `type` smallint NOT NULL COMMENT '1. 满减券 2.折扣券 3.无门槛券 4.满金额折扣券',
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `delete_time` datetime(3) DEFAULT NULL,
  `valitiy` int unsigned DEFAULT NULL,
  `activity_id` int unsigned DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `whole_store` tinyint unsigned DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.coupon 的数据：~5 rows (大约)
/*!40000 ALTER TABLE `coupon` DISABLE KEYS */;
INSERT INTO `coupon` (`id`, `title`, `start_time`, `end_time`, `description`, `full_money`, `minus`, `rate`, `type`, `create_time`, `update_time`, `delete_time`, `valitiy`, `activity_id`, `remark`, `whole_store`) VALUES
	(3, '无门槛减0.1券', '2019-08-05 06:11:42', '2031-08-05 06:11:48', NULL, NULL, 0.10, NULL, 3, '2019-08-03 08:22:06.000', '2019-08-26 14:50:55.000', NULL, NULL, 2, '全场无门槛立减', 1),
	(4, '满500减100券', '2019-08-05 06:11:42', '2030-08-05 06:11:48', NULL, 500.00, 100.00, NULL, 1, '2019-08-03 08:19:34.000', '2019-09-15 21:44:53.000', NULL, NULL, 2, '限服装、居家、文具等商品', 0),
	(7, '满1000减230券', '2019-08-05 06:10:48', '2030-03-05 06:11:17', NULL, 1000.00, 230.00, NULL, 1, '2019-08-03 08:18:36.000', '2019-09-15 21:44:53.000', NULL, NULL, 2, '限服装、家具、文具等商品', 0),
	(10, '满1000打8折', '2019-08-23 09:07:29', '2030-08-23 09:07:36', NULL, NULL, NULL, NULL, 2, '2019-08-23 09:07:57.000', '2019-08-23 09:07:57.000', NULL, NULL, NULL, NULL, NULL),
	(11, '满100打9.9折', '2019-08-28 03:49:55', '2030-08-28 03:49:59', NULL, 100.00, NULL, 0.99, 2, '2019-08-28 03:49:11.000', '2019-09-05 19:14:56.000', NULL, NULL, 2, '全场通用券', 1);
/*!40000 ALTER TABLE `coupon` ENABLE KEYS */;

-- 导出  表 mall.coupon_category 结构
CREATE TABLE IF NOT EXISTS `coupon_category` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `category_id` int unsigned NOT NULL,
  `coupon_id` int unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.coupon_category 的数据：~6 rows (大约)
/*!40000 ALTER TABLE `coupon_category` DISABLE KEYS */;
INSERT INTO `coupon_category` (`id`, `category_id`, `coupon_id`) VALUES
	(1, 15, 4),
	(2, 32, 4),
	(6, 35, 4),
	(7, 15, 7),
	(8, 35, 7),
	(9, 32, 7);
/*!40000 ALTER TABLE `coupon_category` ENABLE KEYS */;

-- 导出  表 mall.coupon_template 结构
CREATE TABLE IF NOT EXISTS `coupon_template` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `full_money` decimal(10,2) DEFAULT NULL,
  `minus` decimal(10,2) DEFAULT NULL,
  `discount` decimal(10,2) DEFAULT NULL COMMENT '国内多是打折，国外多是百分比 off',
  `type` smallint NOT NULL COMMENT '1. 满减券 2.折扣券 3.无门槛券 4.满金额折扣券',
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `delete_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.coupon_template 的数据：~7 rows (大约)
/*!40000 ALTER TABLE `coupon_template` DISABLE KEYS */;
INSERT INTO `coupon_template` (`id`, `title`, `description`, `full_money`, `minus`, `discount`, `type`, `create_time`, `update_time`, `delete_time`) VALUES
	(3, '满1000-230', NULL, 1000.00, 230.00, NULL, 1, '2019-08-03 08:18:36.000', '2019-08-03 08:20:18.000', NULL),
	(4, '满500-100', NULL, 500.00, 100.00, NULL, 1, '2019-08-03 08:19:34.000', '2019-08-03 08:19:34.000', NULL),
	(5, '满2000-500', NULL, 2000.00, 500.00, NULL, 1, '2019-08-03 08:20:03.000', '2019-08-03 08:20:23.000', NULL),
	(6, '满5000-1500', NULL, 5000.00, 1500.00, NULL, 1, '2019-08-03 08:21:15.000', '2019-08-03 08:21:15.000', NULL),
	(7, '无门槛-20', NULL, NULL, NULL, NULL, 3, '2019-08-03 08:22:06.000', '2019-08-03 08:22:06.000', NULL),
	(8, '满减折扣', NULL, 800.00, NULL, 0.90, 4, '2019-08-03 08:23:54.000', '2019-08-03 08:25:02.000', NULL),
	(9, '无门槛折扣', NULL, NULL, NULL, 0.95, 2, '2019-08-03 08:25:54.000', '2019-08-03 08:25:54.000', NULL);
/*!40000 ALTER TABLE `coupon_template` ENABLE KEYS */;

-- 导出  表 mall.coupon_type 结构
CREATE TABLE IF NOT EXISTS `coupon_type` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `code` int NOT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.coupon_type 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `coupon_type` DISABLE KEYS */;
/*!40000 ALTER TABLE `coupon_type` ENABLE KEYS */;

-- 导出  表 mall.grid_category 结构
CREATE TABLE IF NOT EXISTS `grid_category` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `delete_time` datetime(3) DEFAULT NULL,
  `category_id` int DEFAULT NULL,
  `root_category_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.grid_category 的数据：~6 rows (大约)
/*!40000 ALTER TABLE `grid_category` DISABLE KEYS */;
INSERT INTO `grid_category` (`id`, `title`, `img`, `name`, `create_time`, `update_time`, `delete_time`, `category_id`, `root_category_id`) VALUES
	(1, '服装', NULL, NULL, '2019-08-02 19:41:55.000', '2019-09-18 21:27:20.349', NULL, NULL, 2),
	(2, '包包', NULL, NULL, '2019-08-02 19:42:11.000', '2019-09-18 21:27:20.349', NULL, NULL, 3),
	(3, '鞋', NULL, NULL, '2019-08-02 19:42:57.000', '2019-09-18 21:27:20.349', NULL, NULL, 1),
	(4, '饰品', NULL, NULL, '2019-08-02 19:43:34.000', '2019-09-18 21:27:20.349', NULL, NULL, 5),
	(5, '居家', NULL, NULL, '2019-08-02 19:44:34.000', '2019-09-18 21:27:20.349', NULL, NULL, 4),
	(6, '工艺', NULL, NULL, '2019-08-02 19:46:53.000', '2019-09-18 21:27:20.349', NULL, NULL, 24);
/*!40000 ALTER TABLE `grid_category` ENABLE KEYS */;

-- 导出  表 mall.order 结构
CREATE TABLE IF NOT EXISTS `order` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `order_no` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `user_id` int unsigned DEFAULT NULL COMMENT 'user表外键',
  `total_price` decimal(10,2) DEFAULT '0.00',
  `total_count` int unsigned DEFAULT '0',
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `delete_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `snap_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `snap_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `snap_items` json DEFAULT NULL,
  `snap_address` json DEFAULT NULL,
  `prepay_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `final_total_price` decimal(10,2) DEFAULT NULL,
  `status` tinyint unsigned DEFAULT '1',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_order_no` (`order_no`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.order 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `order` DISABLE KEYS */;
/*!40000 ALTER TABLE `order` ENABLE KEYS */;

-- 导出  表 mall.sale_explain 结构
CREATE TABLE IF NOT EXISTS `sale_explain` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `fixed` tinyint unsigned DEFAULT '0',
  `text` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `spu_id` int DEFAULT NULL,
  `index` int unsigned DEFAULT NULL,
  `replace_id` int unsigned DEFAULT NULL,
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `delete_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.sale_explain 的数据：~4 rows (大约)
/*!40000 ALTER TABLE `sale_explain` DISABLE KEYS */;
INSERT INTO `sale_explain` (`id`, `fixed`, `text`, `spu_id`, `index`, `replace_id`, `create_time`, `delete_time`, `update_time`) VALUES
	(1, 1, '发货地：上海', NULL, 1, NULL, '2019-08-17 04:59:44.000', NULL, '2019-08-17 04:59:44.000'),
	(2, 1, '物流：顺丰', NULL, 2, NULL, '2019-08-17 04:59:44.000', NULL, '2019-08-17 05:00:27.000'),
	(3, 1, '发货时间：七个工作日', NULL, 3, NULL, '2019-08-17 04:59:44.000', NULL, '2019-08-17 05:00:29.000'),
	(4, 1, '售后：不支持7天无理由退货', NULL, 4, NULL, '2019-08-17 04:59:44.000', NULL, '2019-09-10 12:17:04.000');
/*!40000 ALTER TABLE `sale_explain` ENABLE KEYS */;

-- 导出  表 mall.sku 结构
CREATE TABLE IF NOT EXISTS `sku` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `price` decimal(10,2) unsigned NOT NULL,
  `discount_price` decimal(10,2) unsigned DEFAULT NULL,
  `online` tinyint unsigned NOT NULL DEFAULT '1',
  `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `spu_id` int unsigned NOT NULL,
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `delete_time` datetime(3) DEFAULT NULL,
  `specs` json DEFAULT NULL,
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `stock` int unsigned NOT NULL DEFAULT '0',
  `category_id` int unsigned DEFAULT NULL,
  `root_category_id` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=52 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.sku 的数据：~51 rows (大约)
/*!40000 ALTER TABLE `sku` DISABLE KEYS */;
INSERT INTO `sku` (`id`, `price`, `discount_price`, `online`, `img`, `title`, `spu_id`, `create_time`, `update_time`, `delete_time`, `specs`, `code`, `stock`, `category_id`, `root_category_id`) VALUES
	(1, 13.80, 11.10, 1, NULL, '青峰·7英寸', 1, '2019-07-16 13:14:26.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "青蓝色", "key_id": 1, "value_id": 1}, {"key": "尺寸", "value": "7英寸", "key_id": 2, "value_id": 5}]', '1$1-1#2-5', 99, 28, 27),
	(2, 77.76, NULL, 1, NULL, '金属灰·七龙珠', 2, '2019-07-18 13:11:07.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "金属灰", "key_id": 1, "value_id": 45}, {"key": "图案", "value": "七龙珠", "key_id": 3, "value_id": 9}, {"key": "尺码", "value": "小号 S", "key_id": 4, "value_id": 14}]', '2$1-45#3-9#4-14', 5, 17, 3),
	(3, 66.00, 59.00, 1, NULL, '青芒色·灌篮高手', 2, '2019-07-18 13:11:11.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "青芒色", "key_id": 1, "value_id": 42}, {"key": "图案", "value": "灌篮高手", "key_id": 3, "value_id": 10}, {"key": "尺码", "value": "中号 M", "key_id": 4, "value_id": 15}]', '2$1-42#3-10#4-15', 999, 17, 3),
	(4, 88.00, NULL, 1, NULL, '青芒色·圣斗士', 2, '2019-07-18 13:11:13.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "青芒色", "key_id": 1, "value_id": 42}, {"key": "图案", "value": "圣斗士", "key_id": 3, "value_id": 11}, {"key": "尺码", "value": "大号  L", "key_id": 4, "value_id": 16}]', '2$1-42#3-11#4-16', 8, 17, 3),
	(5, 77.00, 59.00, 1, NULL, '橘黄色·七龙珠', 2, '2019-07-18 13:11:16.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "橘黄色", "key_id": 1, "value_id": 44}, {"key": "图案", "value": "七龙珠", "key_id": 3, "value_id": 9}, {"key": "尺码", "value": "小号 S", "key_id": 4, "value_id": 14}]', '2$1-44#3-9#4-14', 7, 17, 3),
	(6, 29.99, 19.99, 1, NULL, '黑色', 3, '2019-07-30 02:26:25.000', '2020-02-25 15:07:01.442', '2019-09-14 06:43:46.000', '[{"key": "颜色", "value": "黑色", "key_id": 1, "value_id": 12}]', '3$1-12', 2, 32, 24),
	(7, 32.00, NULL, 1, NULL, '银色', 3, '2019-07-30 02:27:39.000', '2020-02-25 15:07:01.442', '2019-09-14 06:43:49.000', '[{"key": "颜色", "value": "银色", "key_id": 1, "value_id": 18}]', '3$1-18', 3, 32, 24),
	(8, 49.70, 37.70, 1, NULL, '金色', 3, '2019-07-30 02:27:39.000', '2020-02-25 15:07:01.442', '2019-09-14 06:43:51.000', '[{"key": "颜色", "value": "金色", "key_id": 1, "value_id": 19}]', '3$1-19', 0, 32, 24),
	(9, 16.60, NULL, 1, NULL, '白羽·4.3英寸', 1, '2019-07-30 02:36:26.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "深白色", "key_id": 1, "value_id": 3}, {"key": "尺寸", "value": "4.3英寸", "key_id": 2, "value_id": 7}]', '1$1-3#2-7', 0, 28, 27),
	(10, 78.00, NULL, 1, NULL, '桌旗 130 cm', 4, '2019-07-30 06:46:56.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色规格", "value": "桌旗 30x 100 cm", "key_id": 5, "value_id": 20}]', '4$5-20', 665, 26, 27),
	(11, 128.00, NULL, 1, NULL, '桌布 140x 360 cm', 4, '2019-07-30 06:52:40.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色规格", "value": "桌布 140x 360 cm", "key_id": 5, "value_id": 21}]', '4$5-21', 555, 26, 27),
	(12, 72.00, NULL, 1, NULL, '桌旗 30x 220 cm', 4, '2019-07-30 06:52:40.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色规格", "value": "桌旗 30x 220 cm", "key_id": 5, "value_id": 22}]', '4$5-22', 556, 26, 27),
	(13, 188.00, NULL, 1, NULL, '桌布 160x 330 cm', 4, '2019-07-30 06:52:40.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色规格", "value": "桌布 160x 330 cm", "key_id": 5, "value_id": 23}]', '4$5-23', 555, 26, 27),
	(14, 77.00, NULL, 1, NULL, '飞行员墨镜（阳光橙）', 11, '2019-08-07 22:53:15.000', '2020-02-25 15:07:01.442', NULL, NULL, NULL, 331, 36, 5),
	(15, 0.20, NULL, 1, NULL, '林白的测试数据', 12, '2019-08-26 01:01:35.000', '2020-02-25 15:07:01.442', NULL, '[]', '12$', 942, 38, 37),
	(16, 72.00, 68.00, 1, NULL, '碳素墨水', 10, '2019-08-28 01:25:41.000', '2020-02-25 15:07:01.442', NULL, NULL, NULL, 71, 32, 24),
	(17, 297.00, 236.00, 1, NULL, '爱丽丝Alisia束带（黑色）', 13, '2019-08-28 01:25:41.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "黑色", "key_id": 1, "value_id": 12}]', '13$1-12', 71, 39, 5),
	(18, 297.00, 236.00, 1, NULL, '爱丽丝Alisia束带（棕色）', 13, '2019-08-28 01:25:41.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "棕色", "key_id": 1, "value_id": 24}]', '13$1-24', 71, 39, 5),
	(19, 297.00, 236.00, 1, NULL, '爱丽丝Alisia束带  （咖色）', 13, '2019-08-28 01:25:41.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "咖色", "key_id": 1, "value_id": 25}]', '13$1-25', 71, 39, 5),
	(20, 19.90, NULL, 1, NULL, 'Decend Mini小夹子  （橙色）', 14, '2019-08-28 01:25:41.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "橙色", "key_id": 1, "value_id": 27}]', '14$1-27', 71, 20, 4),
	(21, 19.90, NULL, 1, NULL, 'Decend Mini小夹子  （红色）', 14, '2019-08-28 01:25:41.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "红色", "key_id": 1, "value_id": 26}]', '14$1-26', 71, 20, 4),
	(22, 19.90, NULL, 1, NULL, 'Decend Mini小夹子   （金色）', 14, '2019-08-28 01:25:41.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "金色", "key_id": 1, "value_id": 19}]', '14$1-19', 71, 20, 4),
	(23, 24.00, 19.90, 1, NULL, '多彩别针、回形针 Mini （金色）一盒30个', 15, '2019-09-10 02:05:48.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "金色", "key_id": 1, "value_id": 19}, {"key": "数量", "value": "一盒30个", "key_id": 6, "value_id": 28}]', '15$1-19#6-28', 71, 32, 24),
	(24, 24.00, 19.90, 1, NULL, '多彩别针、回形针 Mini （银色）一盒30个', 15, '2019-09-10 02:05:48.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "银色", "key_id": 1, "value_id": 18}, {"key": "数量", "value": "一盒30个", "key_id": 6, "value_id": 28}]', '15$1-18#6-28', 71, 32, 24),
	(25, 24.00, 19.90, 1, NULL, '多彩别针、回形针 Mini （黑色）一盒30个', 15, '2019-09-10 02:05:48.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "黑色", "key_id": 1, "value_id": 12}, {"key": "数量", "value": "一盒30个", "key_id": 6, "value_id": 28}]', '15$1-12#6-28', 71, 32, 24),
	(26, 32.00, 24.00, 1, NULL, '多彩别针、回形针 Mini （黑色）一盒50个', 15, '2019-09-10 02:05:48.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "黑色", "key_id": 1, "value_id": 12}, {"key": "数量", "value": "一盒50个", "key_id": 6, "value_id": 29}]', '15$1-12#6-29', 71, 32, 24),
	(27, 1799.00, NULL, 1, NULL, '风袖说 Sleeven牛仔 ', 6, '2019-09-10 02:05:48.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "黑色", "key_id": 1, "value_id": 12}, {"key": "尺码", "value": "小号 S", "key_id": 4, "value_id": 14}]', '6$1-12#4-14', 70, 14, 2),
	(28, 1799.00, NULL, 1, NULL, '风袖说 Sleeve牛仔 ', 6, '2019-09-10 02:05:48.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "黑色", "key_id": 1, "value_id": 12}, {"key": "尺码", "value": "中号 M", "key_id": 4, "value_id": 15}]', '6$1-12#4-15', 70, 14, 2),
	(29, 1799.00, NULL, 1, NULL, '风袖说 Sleeve牛仔 ', 6, '2019-09-10 02:05:48.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "黑色", "key_id": 1, "value_id": 12}, {"key": "尺码", "value": "中号 L", "key_id": 4, "value_id": 16}]', '6$1-12#4-16', 70, 14, 2),
	(30, 1399.00, NULL, 1, NULL, '双色百褶裙（棕色）', 23, '2019-09-12 22:42:33.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "棕色", "key_id": 1, "value_id": 24}, {"key": "尺码", "value": "小号 S", "key_id": 4, "value_id": 14}]', '23$1-24#4-14', 69, 15, 2),
	(31, 1399.00, NULL, 1, NULL, '双色百褶裙（鹅暖青）', 23, '2019-09-14 01:44:49.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "鹅暖青", "key_id": 1, "value_id": 30}, {"key": "尺码", "value": "中号 M", "key_id": 4, "value_id": 15}]', '23$1-30#4-15', 7, 15, 2),
	(32, 1399.00, NULL, 1, NULL, '双色百褶裙（鹅暖青 小号）', 23, '2019-09-14 01:46:06.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "鹅暖青", "key_id": 1, "value_id": 30}, {"key": "尺码", "value": "小号 S", "key_id": 4, "value_id": 14}]', '23$1-30#4-14', 8, 15, 2),
	(33, 2799.00, 1799.00, 1, NULL, '秋冬新款驼色大衣', 24, '2019-09-14 02:17:15.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "驼色", "key_id": 1, "value_id": 31}, {"key": "尺码", "value": "小号 S", "key_id": 4, "value_id": 14}]', '24$1-31#4-14', 70, 16, 2),
	(34, 2699.00, 1799.00, 1, NULL, '秋冬新款驼色大衣（驼色 M号）', 24, '2019-09-14 02:19:39.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "驼色", "key_id": 1, "value_id": 31}, {"key": "尺码", "value": "中号 M", "key_id": 4, "value_id": 15}]', '24$1-31#4-15', 70, 16, 2),
	(35, 2999.00, 1699.00, 1, NULL, '秋冬新款驼色大衣 （L号）', 24, '2019-09-14 02:21:02.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "驼色", "key_id": 1, "value_id": 31}, {"key": "尺码", "value": "大号  L", "key_id": 4, "value_id": 16}]', '24$1-31#4-16', 47, 16, 2),
	(36, 3999.00, NULL, 1, NULL, '复古双色沙发（藏青色）', 25, '2019-09-14 02:42:50.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "藏青色", "key_id": 1, "value_id": 2}, {"key": "双色沙发尺寸（非标）", "value": "1.5米 x 1米", "key_id": 7, "value_id": 32}]', '25$1-2#7-32', 87, 35, NULL),
	(37, 3999.00, NULL, 1, NULL, '复古双色沙发 (米黄色）', 25, '2019-09-14 02:43:47.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "米黄色", "key_id": 1, "value_id": 35}, {"key": "双色沙发尺寸（非标）", "value": "2米 x 1米", "key_id": 7, "value_id": 33}]', '25$1-35#7-33', 56, 35, NULL),
	(38, 4799.00, 4299.00, 1, NULL, 'SemiConer柔质沙发（长款）', 26, '2019-09-14 05:44:45.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "金色", "key_id": 1, "value_id": 27}, {"key": "双色沙发尺寸（非标）", "value": "2米 x 1米", "key_id": 7, "value_id": 33}]', '26$1-27#7-33', 7, 35, 4),
	(39, 4799.00, NULL, 1, NULL, 'SemiConer柔质沙发 （L型）', 26, '2019-09-14 05:45:44.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "金色", "key_id": 1, "value_id": 27}, {"key": "双色沙发尺寸（非标）", "value": "L型 2米 + 0.8米", "key_id": 7, "value_id": 34}]', '26$1-27#7-34', 7, 35, 4),
	(40, 4799.00, 4299.00, 1, NULL, 'SemiConer柔质沙发（短款）', 26, '2019-09-14 05:46:29.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "金色", "key_id": 1, "value_id": 27}, {"key": "双色沙发尺寸（非标）", "value": "1.5米 x 1米", "key_id": 7, "value_id": 32}]', '26$1-27#7-32', 6, 35, 4),
	(41, 1399.00, NULL, 1, NULL, 'Mier双色靠椅（海蓝色）', 27, '2019-09-14 06:11:11.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "海蓝色", "key_id": 1, "value_id": 36}]', '27$1-36', 67, 35, 4),
	(42, 1399.00, NULL, 1, NULL, 'Mier双色靠椅 （象牙白）', 27, '2019-09-14 06:11:50.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "象牙白", "key_id": 1, "value_id": 37}]', '27$1-37', 13, 35, 4),
	(43, 999.00, NULL, 1, NULL, 'Ins复古金色落地灯（落地灯）', 28, '2019-09-14 06:25:46.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "金色", "key_id": 1, "value_id": 27}, {"key": "台灯高低", "value": "落地灯", "key_id": 8, "value_id": 38}]', '28$1-27#8-38', 19, 23, 4),
	(44, 999.00, NULL, 1, NULL, 'Ins复古金色落地灯 （台灯）', 28, '2019-09-14 06:26:25.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "金色", "key_id": 1, "value_id": 27}, {"key": "台灯高低", "value": "台灯", "key_id": 8, "value_id": 39}]', '28$1-27#8-39', 19, 23, 4),
	(45, 22.00, 17.00, 1, NULL, '抹茶小橡皮 （一盒30个）', 3, '2019-09-14 06:36:02.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "抹茶绿", "key_id": 1, "value_id": 40}, {"key": "数量", "value": "一盒30个", "key_id": 6, "value_id": 28}]', '3$1-40#6-28', 120, 32, 24),
	(46, 22.00, 20.00, 1, NULL, '抹茶小橡皮（一盒50个）', 3, '2019-09-14 06:40:00.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "抹茶绿", "key_id": 1, "value_id": 40}, {"key": "数量", "value": "一盒50个", "key_id": 6, "value_id": 29}]', '3$1-40#6-29', 8, 32, 24),
	(47, 29.00, 20.00, 1, NULL, 'ins复古印花NoteBook （一盒30）', 8, '2019-09-14 06:51:47.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "青草绿", "key_id": 1, "value_id": 41}, {"key": "数量", "value": "一盒30个", "key_id": 6, "value_id": 28}]', '8$1-41#6-28', 19, 32, 24),
	(48, 29.00, 27.00, 1, NULL, 'ins复古印花NoteBook （一盒50个）', 8, '2019-09-14 06:55:14.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "青草绿", "key_id": 1, "value_id": 41}, {"key": "数量", "value": "一盒50个", "key_id": 6, "value_id": 29}]', '8$1-41#6-29', 89, 32, 24),
	(49, 379.00, 279.00, 1, NULL, '七色针织衫（米黄）', 5, '2019-09-14 07:08:06.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "米黄色", "key_id": 1, "value_id": 35}, {"key": "尺码", "value": "中号 M", "key_id": 4, "value_id": 15}]', '5$1-35#4-15', 7, 14, 2),
	(50, 349.00, 279.00, 1, NULL, '七色针织衫（白色）', 5, '2019-09-14 07:09:31.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "象牙白", "key_id": 1, "value_id": 37}, {"key": "尺码", "value": "小号 S", "key_id": 4, "value_id": 14}]', '5$1-37#4-14', 60, 14, 2),
	(51, 349.00, 279.00, 1, NULL, '七色针织衫（蓝色）', 5, '2019-09-14 07:12:13.000', '2020-02-25 15:07:01.442', NULL, '[{"key": "颜色", "value": "青蓝色", "key_id": 1, "value_id": 1}, {"key": "尺码", "value": "中号 M", "key_id": 4, "value_id": 15}]', '5$1-1#4-15', 60, 14, 2);
/*!40000 ALTER TABLE `sku` ENABLE KEYS */;

-- 导出  表 mall.sku_spec 结构
CREATE TABLE IF NOT EXISTS `sku_spec` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `spu_id` int unsigned NOT NULL,
  `sku_id` int unsigned NOT NULL,
  `key_id` int unsigned NOT NULL,
  `value_id` int unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=90 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.sku_spec 的数据：~60 rows (大约)
/*!40000 ALTER TABLE `sku_spec` DISABLE KEYS */;
INSERT INTO `sku_spec` (`id`, `spu_id`, `sku_id`, `key_id`, `value_id`) VALUES
	(5, 23, 2728, 1, 25),
	(6, 23, 2728, 4, 14),
	(7, 23, 2729, 1, 1),
	(8, 23, 2729, 4, 14),
	(9, 23, 2730, 1, 24),
	(10, 23, 2730, 4, 14),
	(11, 23, 30, 1, 24),
	(12, 23, 30, 4, 14),
	(13, 23, 31, 1, 30),
	(14, 23, 31, 4, 15),
	(15, 23, 32, 1, 30),
	(16, 23, 32, 4, 14),
	(23, 25, 36, 1, 2),
	(24, 25, 36, 7, 32),
	(25, 25, 37, 1, 35),
	(26, 25, 37, 7, 33),
	(27, 26, 38, 1, 27),
	(28, 26, 38, 7, 33),
	(29, 26, 39, 1, 27),
	(30, 26, 39, 7, 34),
	(31, 26, 40, 1, 27),
	(32, 26, 40, 7, 32),
	(33, 27, 41, 1, 36),
	(34, 27, 42, 1, 37),
	(35, 28, 43, 1, 27),
	(36, 28, 43, 8, 38),
	(37, 28, 44, 1, 27),
	(38, 28, 44, 8, 39),
	(41, 3, 45, 1, 40),
	(42, 3, 45, 6, 28),
	(43, 3, 46, 1, 40),
	(44, 3, 46, 6, 29),
	(47, 8, 47, 1, 41),
	(48, 8, 47, 6, 28),
	(49, 8, 48, 1, 41),
	(50, 8, 48, 6, 29),
	(51, 5, 49, 1, 35),
	(52, 5, 49, 4, 15),
	(53, 5, 50, 1, 37),
	(54, 5, 50, 4, 14),
	(55, 5, 51, 1, 1),
	(56, 5, 51, 4, 15),
	(57, 24, 34, 1, 31),
	(58, 24, 34, 4, 15),
	(59, 24, 35, 1, 31),
	(60, 24, 35, 4, 16),
	(61, 24, 33, 1, 31),
	(62, 24, 33, 4, 14),
	(78, 2, 5, 1, 43),
	(79, 2, 5, 3, 9),
	(80, 2, 5, 4, 14),
	(81, 2, 3, 1, 42),
	(82, 2, 3, 3, 10),
	(83, 2, 3, 4, 15),
	(84, 2, 2, 1, 3),
	(85, 2, 2, 3, 9),
	(86, 2, 2, 4, 14),
	(87, 2, 4, 1, 42),
	(88, 2, 4, 3, 11),
	(89, 2, 4, 4, 16);
/*!40000 ALTER TABLE `sku_spec` ENABLE KEYS */;

-- 导出  表 mall.spec_key 结构
CREATE TABLE IF NOT EXISTS `spec_key` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `unit` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `standard` tinyint unsigned NOT NULL DEFAULT '0',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.spec_key 的数据：~8 rows (大约)
/*!40000 ALTER TABLE `spec_key` DISABLE KEYS */;
INSERT INTO `spec_key` (`id`, `name`, `unit`, `standard`, `create_time`, `update_time`, `delete_time`, `description`) VALUES
	(1, '颜色', NULL, 1, '2019-07-15 15:33:18', '2019-07-15 15:33:18', NULL, NULL),
	(2, '尺寸', '英寸', 0, '2019-07-15 15:48:52', '2019-07-16 13:02:44', NULL, NULL),
	(3, '图案', NULL, 0, '2019-07-17 08:21:42', '2019-07-17 08:21:42', NULL, NULL),
	(4, '尺码', NULL, 1, '2019-07-17 08:24:40', '2019-07-17 08:24:47', NULL, NULL),
	(5, '颜色与规格', NULL, 0, '2019-07-30 06:39:27', '2019-07-30 06:39:27', NULL, NULL),
	(6, '数量', '个', 0, '2019-09-10 02:13:11', '2019-09-10 02:13:11', NULL, NULL),
	(7, '双色沙发尺寸（非标）', '米', 0, '2019-09-14 02:32:05', '2019-09-14 02:32:05', NULL, ''),
	(8, '台灯高低', '', 0, '2019-09-14 03:28:00', '2019-09-14 03:28:00', NULL, '');
/*!40000 ALTER TABLE `spec_key` ENABLE KEYS */;

-- 导出  表 mall.spec_value 结构
CREATE TABLE IF NOT EXISTS `spec_value` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `spec_id` int unsigned NOT NULL,
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `delete_time` datetime(3) DEFAULT NULL,
  `extend` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.spec_value 的数据：~45 rows (大约)
/*!40000 ALTER TABLE `spec_value` DISABLE KEYS */;
INSERT INTO `spec_value` (`id`, `value`, `spec_id`, `create_time`, `update_time`, `delete_time`, `extend`) VALUES
	(1, '青蓝色', 1, '2019-07-16 12:52:36.000', '2019-07-16 12:52:36.000', NULL, NULL),
	(2, '藏青色', 1, '2019-07-16 12:55:45.000', '2019-07-16 12:55:45.000', NULL, NULL),
	(3, '深白色', 1, '2019-07-16 12:56:33.000', '2019-07-16 12:56:33.000', NULL, NULL),
	(4, '少女粉', 1, '2019-07-16 12:57:01.000', '2019-07-16 12:57:01.000', NULL, NULL),
	(5, '7英寸', 2, '2019-07-16 12:58:47.000', '2019-07-16 12:58:47.000', NULL, NULL),
	(6, '5英寸', 2, '2019-07-16 12:58:56.000', '2019-07-16 12:58:56.000', NULL, NULL),
	(7, '4.3英寸', 2, '2019-07-16 12:59:12.000', '2019-07-16 12:59:12.000', NULL, NULL),
	(8, '10英寸', 2, '2019-07-16 12:59:23.000', '2019-07-16 12:59:23.000', NULL, NULL),
	(9, '七龙珠', 3, '2019-07-17 08:22:05.000', '2019-07-17 08:22:05.000', NULL, NULL),
	(10, '灌篮高手', 3, '2019-07-17 08:22:15.000', '2019-07-17 08:22:15.000', NULL, NULL),
	(11, '圣斗士', 3, '2019-07-17 08:22:24.000', '2019-07-17 08:22:24.000', NULL, NULL),
	(12, '黑色', 1, '2019-07-17 08:22:42.000', '2019-07-17 08:22:42.000', NULL, NULL),
	(13, '墨绿色', 1, '2019-07-17 08:25:04.000', '2019-07-17 08:25:04.000', NULL, NULL),
	(14, '小号 S', 4, '2019-07-17 08:25:34.000', '2019-09-10 15:47:36.000', NULL, NULL),
	(15, '中号 M', 4, '2019-07-17 08:25:44.000', '2019-09-10 15:47:40.000', NULL, NULL),
	(16, '大号  L', 4, '2019-07-17 08:25:50.000', '2019-09-10 15:47:45.000', NULL, NULL),
	(17, '加大 XL', 4, '2019-07-17 08:26:03.000', '2019-09-10 15:47:50.000', NULL, NULL),
	(18, '银色', 1, '2019-07-30 02:23:10.000', '2019-07-30 02:23:10.000', NULL, NULL),
	(19, '金色', 1, '2019-07-30 02:23:21.000', '2019-07-30 02:23:21.000', NULL, NULL),
	(20, '桌旗 30x 100 cm', 5, '2019-07-30 06:40:26.000', '2019-07-30 06:43:19.000', NULL, NULL),
	(21, '桌旗 30x 220 cm', 5, '2019-07-30 06:42:16.000', '2019-07-30 06:43:22.000', NULL, NULL),
	(22, '桌布 140x 360 cm', 5, '2019-07-30 06:42:48.000', '2019-07-30 06:43:27.000', NULL, NULL),
	(23, '桌布 160x 330 cm', 5, '2019-07-30 06:43:16.000', '2019-07-30 06:43:29.000', NULL, NULL),
	(24, '棕色', 1, '2019-07-30 06:43:16.000', '2019-09-07 19:22:05.000', NULL, NULL),
	(25, '咖色', 1, '2019-07-30 06:43:16.000', '2019-09-07 19:22:08.000', NULL, NULL),
	(26, '红色', 1, '2019-09-07 19:22:01.000', '2019-09-07 19:22:01.000', NULL, NULL),
	(27, '金色', 1, '2019-09-07 19:22:33.000', '2019-09-07 19:22:33.000', NULL, NULL),
	(28, '一盒30个', 6, '2019-09-10 02:13:42.000', '2019-09-10 02:13:42.000', NULL, NULL),
	(29, '一盒50个', 6, '2019-09-10 02:13:42.000', '2019-09-10 02:13:42.000', NULL, NULL),
	(30, '鹅暖青', 1, '2019-09-14 01:42:00.000', '2019-09-14 01:42:00.000', NULL, ''),
	(31, '驼色', 1, '2019-09-14 02:15:59.000', '2019-09-14 02:15:59.000', NULL, ''),
	(32, '1.5米 x 1米', 7, '2019-09-14 02:32:57.000', '2019-09-14 02:32:57.000', NULL, ''),
	(33, '2米 x 1米', 7, '2019-09-14 02:33:39.000', '2019-09-14 02:33:39.000', NULL, ''),
	(34, 'L型 2米 + 0.8米', 7, '2019-09-14 02:34:12.000', '2019-09-14 02:34:12.000', NULL, ''),
	(35, '米黄色', 1, '2019-09-14 02:40:45.000', '2019-09-14 02:40:45.000', NULL, ''),
	(36, '海蓝色', 1, '2019-09-14 05:50:31.000', '2019-09-14 05:50:31.000', NULL, ''),
	(37, '象牙白', 1, '2019-09-14 05:50:42.000', '2019-09-14 05:50:42.000', NULL, ''),
	(38, '落地灯', 8, '2019-09-14 06:24:15.000', '2019-09-14 06:24:15.000', NULL, ''),
	(39, '台灯', 8, '2019-09-14 06:24:27.000', '2019-09-14 06:24:27.000', NULL, ''),
	(40, '抹茶绿', 1, '2019-09-14 06:37:13.000', '2019-09-14 06:37:13.000', NULL, ''),
	(41, '青草绿', 1, '2019-09-14 06:53:08.000', '2019-09-14 06:53:08.000', NULL, ''),
	(42, '青芒色', 1, '2019-09-15 18:02:01.000', '2019-09-15 18:02:01.000', NULL, ''),
	(43, '粉色', 1, '2019-09-15 18:02:16.000', '2019-09-15 18:02:16.000', NULL, ''),
	(44, '橘黄色', 1, '2019-09-18 03:32:43.000', '2019-09-18 03:32:43.000', NULL, ''),
	(45, '金属灰', 1, '2019-09-18 03:32:59.000', '2019-09-18 03:32:59.000', NULL, '');
/*!40000 ALTER TABLE `spec_value` ENABLE KEYS */;

-- 导出  表 mall.spu 结构
CREATE TABLE IF NOT EXISTS `spu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `subtitle` varchar(800) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `category_id` int unsigned NOT NULL,
  `root_category_id` int DEFAULT NULL,
  `online` tinyint unsigned NOT NULL DEFAULT '1',
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `delete_time` datetime(3) DEFAULT NULL,
  `price` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文本型价格，有时候SPU需要展示的是一个范围，或者自定义平均价格',
  `sketch_spec_id` int unsigned DEFAULT NULL COMMENT '某种规格可以直接附加单品图片',
  `default_sku_id` int DEFAULT NULL COMMENT '默认选中的sku',
  `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `discount_price` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `tags` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `is_test` tinyint unsigned DEFAULT '0',
  `spu_theme_img` json DEFAULT NULL,
  `for_theme_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.spu 的数据：~19 rows (大约)
/*!40000 ALTER TABLE `spu` DISABLE KEYS */;
INSERT INTO `spu` (`id`, `title`, `subtitle`, `category_id`, `root_category_id`, `online`, `create_time`, `update_time`, `delete_time`, `price`, `sketch_spec_id`, `default_sku_id`, `img`, `discount_price`, `description`, `tags`, `is_test`, `spu_theme_img`, `for_theme_img`) VALUES
	(1, '青锋大碗', '大碗主要用来盛宽面，凡凡倾情推荐', 28, 27, 1, '2019-07-15 14:47:11.000', '2020-11-04 16:09:37.485', NULL, '12.99', 1, NULL, 'http://i2.sleeve.talelin.com/n9.pnghttp://i2.sleeve.talelin.com/n9.png', '11.11', NULL, '林白推荐', 1, NULL, NULL),
	(2, '林间有风自营针织衫', '秋日冬款，浪漫满屋', 12, 2, 1, '2019-07-31 08:19:24.000', '2020-11-04 16:09:06.787', NULL, '77.00', 1, 2, 'http://i1.sleeve.7yue.pro/assets/ecf8d824-19d4-4db2-a5da-872ab014fecd.png', '62.00', NULL, '秋日冬款$浪漫满屋', 1, NULL, 'https://gitee.com/lrelia7/sleeve-static/raw/master/theme/spu1.png'),
	(3, '抹茶小橡皮', '小作文写错了，用它擦一擦', 32, 24, 1, '2019-09-16 09:55:51.000', '2020-11-04 16:14:56.341', NULL, '29.99', 1, NULL, 'http://i2.sleeve.talelin.com/m17.png', '17.00', NULL, '一飞推荐', 0, NULL, 'https://gitee.com/lrelia7/sleeve-static/raw/master/theme/spu2.png'),
	(4, '印花桌布', '生活需要仪式感，吃饭也一样。桌旗+桌布给你绚烂的生命色彩', 26, 27, 1, '2019-07-30 06:26:33.000', '2020-11-04 16:09:19.274', NULL, '119.00', 5, NULL, 'http://i2.sleeve.talelin.com/n10.png', '97.00', NULL, '风袖臻选', 1, NULL, NULL),
	(5, '七色针织衫', '女朋友不给你洗衣服？没关系，每天换一件。', 14, 2, 0, '2019-07-16 14:47:11.000', '2020-02-25 15:07:01.451', NULL, '349', 1, NULL, NULL, '279', NULL, 'pedro推荐', 1, NULL, NULL),
	(6, 'Sleeve风袖说牛仔系列', 'Sleeve风袖说当季经典款', 14, 2, 1, '2019-08-01 08:19:24.000', '2020-11-04 16:10:15.715', NULL, '1799', 1, 5, 'http://i2.sleeve.talelin.com/n14.png', '', NULL, '包邮$热门', 1, NULL, NULL),
	(8, 'Ins复古碎花NoteBook', '林白默默的掏出小本本，将她说的话一次不漏的记了下来。', 32, 24, 1, '2019-09-15 05:00:21.000', '2020-11-04 16:13:51.845', NULL, '29.99', 1, NULL, 'http://i2.sleeve.talelin.com/m19.png', '27.8', NULL, '林白推荐', 0, NULL, 'http://i1.sleeve.7yue.pro/assets/b6442702-4810-46cb-bb0b-f4602d38e4ff.png'),
	(10, '碳素墨水', '虽然我们早已不再使用钢笔书写，但钢笔在纸上划过的笔触永远是键盘无法替代的。一只钢笔+一瓶墨水在一个安静的午后，写写内心的故事。', 32, 24, 1, '2019-09-16 09:57:15.000', '2020-11-04 16:15:30.078', NULL, '80.00', NULL, NULL, 'http://i2.sleeve.talelin.com/m24.png', '69.00', NULL, '', 1, NULL, NULL),
	(11, '飞行员墨镜', '戴起来像小李子', 36, 5, 1, '2019-08-07 22:47:05.000', '2020-11-04 16:10:41.063', NULL, '77.00', NULL, NULL, 'http://i2.sleeve.talelin.com/n2.png', NULL, NULL, NULL, 1, NULL, NULL),
	(12, '林间有风测试商品', '测试商品，可购买体验完整支付和订单流程', 38, 37, 1, '2019-08-25 19:03:03.000', '2020-11-04 16:11:09.151', NULL, '0.2', NULL, NULL, 'http://i2.sleeve.talelin.com/x1.png', NULL, NULL, '测试数据$可支付', 0, NULL, NULL),
	(13, '基克的聚合束带', '三色可选，加攻、加防、还能加血', 39, 5, 0, '2019-09-07 16:06:47.000', '2020-02-25 15:07:01.451', NULL, '279', NULL, NULL, NULL, NULL, NULL, NULL, 1, NULL, NULL),
	(14, 'Ins 复古小夹子（Mini)', '静静的，享受时光的流逝', 32, 24, 1, '2019-09-16 09:54:47.000', '2020-11-04 16:14:12.845', NULL, '19.9', NULL, NULL, 'http://i2.sleeve.talelin.com/m23.png', NULL, NULL, '三色可选', 1, NULL, NULL),
	(15, '多彩别针、回形针', '每盒70个，可爱多彩', 32, 24, 1, '2019-09-16 09:55:27.000', '2020-11-04 16:14:25.965', NULL, '24', 1, 25, 'http://i2.sleeve.talelin.com/m26.png', '19.9', NULL, '三色可选', 1, NULL, NULL),
	(23, '双色牛仔裤', '秋冬新款，做一个Cool Boy', 15, 2, 1, '2019-09-16 10:26:04.000', '2020-11-04 16:16:01.365', NULL, '1399', 1, NULL, 'http://i2.sleeve.talelin.com/n11.png', NULL, NULL, '', 1, NULL, 'http://i1.sleeve.7yue.pro/assets/702f2ce9-5729-4aa4-aeb3-921513327747.png'),
	(24, '秋冬新款驼色大衣', '2020新款，暖暖过秋冬', 16, 2, 1, '2019-09-14 02:13:20.000', '2020-11-04 16:12:11.413', NULL, '2999', 1, NULL, 'http://i2.sleeve.talelin.com/n3.png', '2699', NULL, '经典单色', 1, NULL, 'http://i1.sleeve.7yue.pro/assets/b8e510a1-8340-43c2-a4b0-0e56a40256f9.png'),
	(25, '复古双色沙发', '双色可选，经典青黄两色', 35, 4, 1, '2019-09-14 02:30:23.000', '2020-11-04 16:12:24.718', NULL, '3999', 1, NULL, 'http://i2.sleeve.talelin.com/h3.pngv', NULL, NULL, '复刻经典$双色可选', 1, NULL, NULL),
	(26, 'SemiConer柔质沙发', '窝在沙发上，一杯红酒配电影', 35, 4, 1, '2019-09-14 05:43:19.000', '2020-11-04 16:12:34.996', NULL, '4799', 1, NULL, 'http://i2.sleeve.talelin.com/3.png', '4200', NULL, '', 1, NULL, NULL),
	(27, 'Mier双色靠椅', '安静的午后，一杯清茶，追忆似水年华。看清风浮动，看落日余晖', 35, 4, 1, '2019-09-09 02:26:12.000', '2020-11-04 16:11:45.508', NULL, '1299', 1, NULL, 'http://i2.sleeve.talelin.com/h1.png', NULL, NULL, '', 1, NULL, 'http://i1.sleeve.7yue.pro/assets/f6c9fce8-626f-44c0-a709-3f6ef9f3fbef.png'),
	(28, 'Ins复古金色落地灯', 'Instagram复古台灯，就在此刻回到过去，找寻逝去的记忆', 23, 4, 1, '2019-09-14 06:19:12.000', '2020-11-04 16:12:44.182', NULL, '999', 8, NULL, 'http://i2.sleeve.talelin.com/a9.png', NULL, NULL, 'Ins$复古流行', 1, NULL, NULL);
/*!40000 ALTER TABLE `spu` ENABLE KEYS */;

-- 导出  表 mall.spu_detail_img 结构
CREATE TABLE IF NOT EXISTS `spu_detail_img` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `spu_id` int unsigned DEFAULT NULL,
  `index` int unsigned NOT NULL,
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `delete_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.spu_detail_img 的数据：~4 rows (大约)
/*!40000 ALTER TABLE `spu_detail_img` DISABLE KEYS */;
INSERT INTO `spu_detail_img` (`id`, `img`, `spu_id`, `index`, `create_time`, `update_time`, `delete_time`) VALUES
	(24, '', 2, 1, '2019-09-18 04:50:21.313', '2020-02-25 15:07:01.456', NULL),
	(25, '', 2, 3, '2019-09-18 04:50:21.313', '2020-02-25 15:07:01.456', NULL),
	(26, '', 2, 4, '2019-09-18 04:50:21.313', '2020-02-25 15:07:01.456', NULL),
	(27, '', 2, 2, '2019-09-18 04:50:21.313', '2020-02-25 15:07:01.456', NULL);
/*!40000 ALTER TABLE `spu_detail_img` ENABLE KEYS */;

-- 导出  表 mall.spu_img 结构
CREATE TABLE IF NOT EXISTS `spu_img` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `spu_id` int unsigned DEFAULT NULL,
  `delete_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=192 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.spu_img 的数据：~41 rows (大约)
/*!40000 ALTER TABLE `spu_img` DISABLE KEYS */;
INSERT INTO `spu_img` (`id`, `img`, `spu_id`, `delete_time`, `update_time`, `create_time`) VALUES
	(11, NULL, 11, NULL, '2020-02-25 15:07:01.461', '2019-08-08 03:36:56.000'),
	(15, NULL, 13, NULL, '2020-02-25 15:07:01.461', '2019-09-07 16:33:59.000'),
	(16, NULL, 13, NULL, '2020-02-25 15:07:01.461', '2019-09-07 16:33:59.000'),
	(17, NULL, 13, NULL, '2020-02-25 15:07:01.461', '2019-09-07 16:33:59.000'),
	(18, NULL, 13, NULL, '2020-02-25 15:07:01.461', '2019-09-07 16:33:59.000'),
	(19, NULL, 14, NULL, '2020-02-25 15:07:01.461', '2019-09-07 16:33:59.000'),
	(20, NULL, 14, NULL, '2020-02-25 15:07:01.461', '2019-09-07 16:33:59.000'),
	(30, NULL, 15, NULL, '2020-02-25 15:07:01.461', '2019-09-07 16:33:59.000'),
	(31, NULL, 15, NULL, '2020-02-25 15:07:01.461', '2019-09-07 16:33:59.000'),
	(41, NULL, 16, NULL, '2020-02-25 15:07:01.461', '2019-09-11 21:16:57.000'),
	(42, NULL, 16, NULL, '2020-02-25 15:07:01.461', '2019-09-11 21:16:57.000'),
	(43, NULL, 6, NULL, '2020-02-25 15:07:01.461', '2019-09-12 18:06:59.000'),
	(44, NULL, 6, NULL, '2020-02-25 15:07:01.461', '2019-09-12 18:06:59.000'),
	(88, NULL, 25, NULL, '2020-02-25 15:07:01.461', '2019-09-14 02:37:42.000'),
	(89, NULL, 25, NULL, '2020-02-25 15:07:01.461', '2019-09-14 02:37:42.000'),
	(90, NULL, 25, NULL, '2020-02-25 15:07:01.461', '2019-09-14 02:37:42.000'),
	(107, NULL, 3, NULL, '2020-02-25 15:07:01.461', '2019-09-14 06:34:25.000'),
	(108, NULL, 3, NULL, '2020-02-25 15:07:01.461', '2019-09-14 06:34:25.000'),
	(124, NULL, 12, NULL, '2020-02-25 15:07:01.461', '2019-09-14 07:24:14.000'),
	(126, NULL, 4, NULL, '2020-02-25 15:07:01.461', '2019-09-14 19:15:48.000'),
	(138, NULL, 1, NULL, '2020-02-25 15:07:01.461', '2019-09-14 23:01:57.000'),
	(139, NULL, 1, NULL, '2020-02-25 15:07:01.461', '2019-09-14 23:01:57.000'),
	(141, NULL, 28, NULL, '2020-02-25 15:07:01.461', '2019-09-14 23:06:56.000'),
	(150, NULL, 26, NULL, '2020-02-25 15:07:01.461', '2019-09-14 23:07:16.000'),
	(151, NULL, 5, NULL, '2020-02-25 15:07:01.461', '2019-09-15 04:53:29.000'),
	(152, NULL, 5, NULL, '2020-02-25 15:07:01.461', '2019-09-15 04:53:29.000'),
	(153, NULL, 5, NULL, '2020-02-25 15:07:01.461', '2019-09-15 04:53:29.000'),
	(154, NULL, 5, NULL, '2020-02-25 15:07:01.461', '2019-09-15 04:53:29.000'),
	(155, NULL, 5, NULL, '2020-02-25 15:07:01.461', '2019-09-15 04:53:29.000'),
	(161, NULL, 10, NULL, '2020-02-25 15:07:01.461', '2019-09-15 17:50:36.000'),
	(165, NULL, 2, NULL, '2020-02-25 15:07:01.461', '2019-09-15 17:57:23.000'),
	(166, NULL, 2, NULL, '2020-02-25 15:07:01.461', '2019-09-15 17:57:23.000'),
	(167, NULL, 2, NULL, '2020-02-25 15:07:01.461', '2019-09-15 17:57:23.000'),
	(172, NULL, 24, NULL, '2020-02-25 15:07:01.461', '2019-09-16 01:02:58.000'),
	(177, NULL, 8, NULL, '2020-02-25 15:07:01.461', '2019-09-16 01:03:14.000'),
	(179, NULL, 8, NULL, '2020-02-25 15:07:01.461', '2019-09-16 01:03:14.000'),
	(186, NULL, 23, NULL, '2020-02-25 15:07:01.461', '2019-09-16 01:05:21.000'),
	(187, NULL, 23, NULL, '2020-02-25 15:07:01.461', '2019-09-16 01:05:21.000'),
	(189, NULL, 23, NULL, '2020-02-25 15:07:01.461', '2019-09-16 01:05:21.000'),
	(190, NULL, 27, NULL, '2020-02-25 15:07:01.461', '2019-09-16 01:05:46.000'),
	(191, NULL, 10, NULL, '2020-02-25 15:07:01.461', '2019-09-15 17:50:36.000');
/*!40000 ALTER TABLE `spu_img` ENABLE KEYS */;

-- 导出  表 mall.spu_key 结构
CREATE TABLE IF NOT EXISTS `spu_key` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `spu_id` int unsigned NOT NULL,
  `spec_key_id` int unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.spu_key 的数据：~26 rows (大约)
/*!40000 ALTER TABLE `spu_key` DISABLE KEYS */;
INSERT INTO `spu_key` (`id`, `spu_id`, `spec_key_id`) VALUES
	(1, 3, 1),
	(3, 2, 1),
	(7, 1, 1),
	(8, 1, 2),
	(9, 15, 1),
	(10, 15, 2),
	(11, 16, 2),
	(13, 16, 1),
	(14, 23, 1),
	(15, 23, 4),
	(16, 24, 1),
	(17, 24, 4),
	(18, 25, 1),
	(19, 25, 7),
	(20, 26, 1),
	(21, 26, 7),
	(22, 27, 1),
	(23, 28, 1),
	(24, 28, 8),
	(25, 3, 6),
	(26, 8, 1),
	(27, 8, 6),
	(28, 5, 1),
	(29, 5, 4),
	(30, 2, 3),
	(31, 2, 4);
/*!40000 ALTER TABLE `spu_key` ENABLE KEYS */;

-- 导出  表 mall.spu_tag 结构
CREATE TABLE IF NOT EXISTS `spu_tag` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `spu_id` int unsigned NOT NULL,
  `tag_id` int unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.spu_tag 的数据：~5 rows (大约)
/*!40000 ALTER TABLE `spu_tag` DISABLE KEYS */;
INSERT INTO `spu_tag` (`id`, `spu_id`, `tag_id`) VALUES
	(1, 2, 5),
	(2, 2, 1),
	(3, 2, 12),
	(4, 12, 13),
	(5, 12, 14);
/*!40000 ALTER TABLE `spu_tag` ENABLE KEYS */;

-- 导出  表 mall.tag 结构
CREATE TABLE IF NOT EXISTS `tag` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '中文限制6个，英文限制12个，由逻辑层控制',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `delete_time` datetime(3) DEFAULT NULL,
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `highlight` tinyint unsigned DEFAULT '0',
  `type` tinyint unsigned DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.tag 的数据：~18 rows (大约)
/*!40000 ALTER TABLE `tag` DISABLE KEYS */;
INSERT INTO `tag` (`id`, `title`, `description`, `update_time`, `delete_time`, `create_time`, `highlight`, `type`) VALUES
	(1, 'Sleeve', '', '2019-09-14 07:35:17.000', NULL, '2019-07-30 04:10:10.000', 1, 1),
	(3, '椅', NULL, '2019-09-14 07:37:52.000', NULL, '2019-07-30 04:11:06.000', NULL, 1),
	(4, '非卖品', NULL, '2019-08-01 18:45:52.000', NULL, '2019-08-01 18:45:52.000', NULL, 1),
	(5, '测试数据', NULL, '2019-09-16 02:19:38.000', NULL, '2019-08-01 18:46:01.000', 1, 1),
	(10, '江浙沪不包邮', NULL, '2019-09-16 02:42:58.000', NULL, '2019-08-01 18:49:02.000', NULL, 1),
	(11, '百褶裙', NULL, '2019-09-14 07:37:52.000', NULL, '2019-08-01 18:50:26.000', 1, 1),
	(15, '秋冬', NULL, '2019-09-14 07:36:43.000', NULL, '2019-09-12 18:06:59.000', 1, 1),
	(17, '落地灯', NULL, '2019-09-14 07:35:17.000', NULL, '2019-09-14 02:13:20.000', NULL, 1),
	(18, '复刻', NULL, '2019-09-16 02:43:03.000', NULL, '2019-09-14 02:37:42.000', NULL, 1),
	(19, '双色可选', NULL, '2019-09-14 02:37:42.000', NULL, '2019-09-14 02:37:42.000', NULL, 1),
	(20, 'Ins', NULL, '2019-09-14 06:19:12.000', NULL, '2019-09-14 06:19:12.000', NULL, 1),
	(21, '复古', NULL, '2019-09-14 07:37:52.000', NULL, '2019-09-14 06:19:12.000', 1, 1),
	(22, '林间有风', NULL, '2019-09-16 02:42:44.000', NULL, '2019-09-14 06:34:25.000', 1, 1),
	(23, '灯', NULL, '2019-09-14 07:36:43.000', NULL, '2019-09-14 06:50:11.000', NULL, 1),
	(25, '可支付', NULL, '2019-09-14 07:24:14.000', NULL, '2019-09-14 07:24:14.000', NULL, 1),
	(26, '风袖臻选', NULL, '2019-09-16 02:36:35.000', NULL, '2019-09-14 19:15:48.000', 0, 1),
	(28, '林白', NULL, '2019-09-16 02:44:26.000', NULL, '2019-09-14 23:01:25.000', NULL, 1),
	(30, '包邮', NULL, '2019-09-15 17:55:04.000', NULL, '2019-09-15 17:55:04.000', NULL, 1);
/*!40000 ALTER TABLE `tag` ENABLE KEYS */;

-- 导出  表 mall.theme 结构
CREATE TABLE IF NOT EXISTS `theme` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `tpl_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `delete_time` datetime(3) DEFAULT NULL,
  `entrance_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `extend` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `internal_top_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `title_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `online` tinyint unsigned DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.theme 的数据：~6 rows (大约)
/*!40000 ALTER TABLE `theme` DISABLE KEYS */;
INSERT INTO `theme` (`id`, `title`, `description`, `name`, `create_time`, `tpl_name`, `update_time`, `delete_time`, `entrance_img`, `extend`, `internal_top_img`, `title_img`, `online`) VALUES
	(1, '清凉一夏，折扣季', '秋天是金色的。麦穗是金色的，秋阳是金色的。虽然冬快至，但宜人的温度总是让我们心情愉快#我们为您精选了一系列秋冬折扣商品，慢慢挑选吧~', 't-1', '2019-07-18 07:10:59.000', 'janna', '2020-11-04 15:59:36.731', NULL, 'http://i2.sleeve.talelin.com/m2.png', NULL, 'http://i2.sleeve.talelin.com/m33.png', NULL, 1),
	(4, '每周上新', '风袖`每周上新`活动#每周都有一款折扣商品，每周都有适合你的唯一专属#有Ins复古风装饰；金属CD唱片夹；每周来逛逛，找到你所喜爱的东西', 't-2', '2019-07-30 00:00:14.000', NULL, '2020-11-04 16:00:32.532', NULL, '', NULL, 'http://i2.sleeve.talelin.com/m1.png', 'http://i2.sleeve.talelin.com/m3.png', 1),
	(5, '风袖甄选', '甄选', 't-3', '2019-07-30 17:20:23.000', 'diana', '2020-11-04 16:01:00.628', NULL, 'http://i2.sleeve.talelin.com/m9.png', NULL, 'http://i2.sleeve.talelin.com/m11.png', NULL, 1),
	(6, '时尚穿搭', '帅点才有女朋友', 't-4', '2019-08-01 02:43:18.000', 'irelia', '2020-11-04 16:18:52.575', NULL, 'http://i2.sleeve.talelin.com/m10.png', NULL, 'http://i2.sleeve.talelin.com/m1.png', 'http://i2.sleeve.talelin.com/m3.png', 1),
	(7, '热卖好评', '林白选的那一定是最好的', 't-5', '2019-08-09 07:19:37.000', 'camille', '2020-11-04 16:19:25.183', NULL, 'http://i2.sleeve.talelin.com/m9.png', NULL, 'http://i2.sleeve.talelin.com/m11.png', 'http://i2.sleeve.talelin.com/m11.png', 1),
	(8, '热门推荐', '林白选的那一定是最好的', 't-6', '2019-09-10 11:43:06.000', 'camille', '2020-11-04 16:19:41.415', NULL, 'http://i2.sleeve.talelin.com/m10.png', NULL, 'http://i2.sleeve.talelin.com/m12.png', NULL, 1);
/*!40000 ALTER TABLE `theme` ENABLE KEYS */;

-- 导出  表 mall.theme_spu 结构
CREATE TABLE IF NOT EXISTS `theme_spu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `theme_id` int unsigned NOT NULL,
  `spu_id` int unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.theme_spu 的数据：~47 rows (大约)
/*!40000 ALTER TABLE `theme_spu` DISABLE KEYS */;
INSERT INTO `theme_spu` (`id`, `theme_id`, `spu_id`) VALUES
	(22, 6, 23),
	(23, 6, 24),
	(24, 6, 5),
	(25, 6, 6),
	(39, 7, 8),
	(41, 7, 23),
	(42, 7, 24),
	(43, 7, 28),
	(44, 7, 27),
	(45, 7, 25),
	(46, 7, 13),
	(47, 7, 14),
	(48, 6, 28),
	(49, 6, 27),
	(50, 6, 26),
	(51, 6, 8),
	(54, 8, 27),
	(55, 8, 23),
	(56, 8, 6),
	(57, 8, 25),
	(58, 8, 27),
	(59, 8, 26),
	(60, 1, 25),
	(61, 1, 28),
	(62, 1, 23),
	(63, 1, 27),
	(64, 1, 8),
	(65, 1, 14),
	(66, 1, 3),
	(67, 1, 26),
	(70, 8, 8),
	(71, 4, 25),
	(75, 4, 26),
	(76, 4, 28),
	(83, 4, 27),
	(84, 4, 23),
	(85, 4, 8),
	(86, 4, 24),
	(87, 4, 13),
	(88, 4, 6),
	(89, 4, 5),
	(90, 4, 3),
	(91, 5, 23),
	(92, 5, 8),
	(93, 5, 27),
	(94, 5, 24),
	(100, 8, 28);
/*!40000 ALTER TABLE `theme_spu` ENABLE KEYS */;

-- 导出  表 mall.user 结构
CREATE TABLE IF NOT EXISTS `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `openid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `nickname` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `unify_uid` int DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `mobile` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `wx_profile` json DEFAULT NULL,
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `delete_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_openid` (`openid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.user 的数据：~22 rows (大约)
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` (`id`, `openid`, `nickname`, `unify_uid`, `email`, `password`, `mobile`, `wx_profile`, `create_time`, `update_time`, `delete_time`) VALUES
	(1, NULL, NULL, NULL, '123@qq.com', '45678123123', NULL, 'null', '2019-07-13 08:06:35.000', '2019-07-13 08:06:35.000', NULL),
	(2, NULL, NULL, NULL, '1236@qq.com', '45678123123', NULL, 'null', '2019-07-13 08:06:55.000', '2019-07-13 08:06:55.000', NULL),
	(3, NULL, NULL, NULL, '12367@qq.com', '0e40a584386c3f52bee84639fa2608ce', NULL, 'null', '2019-07-13 11:08:20.000', '2019-07-13 11:08:20.000', NULL),
	(4, NULL, NULL, NULL, '123667@qq.com', '0e40a584386c3f52bee84639fa2608ce', NULL, 'null', '2019-07-13 11:20:58.000', '2019-07-13 11:20:58.000', NULL),
	(16, NULL, '7七月', NULL, NULL, NULL, NULL, '{"city": "Haidian", "gender": 1, "country": "China", "language": "zh_CN", "nickName": "7七月", "province": "Beijing", "avatarUrl": "https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTJRgMw7FHDF5Etxb54Qh743fib2ZfO40U7JWyBNcphdDRyaBb1eWXV0NicJtDL59DGGAY8Bf8weqSgg/132"}', '2019-08-19 09:52:31.000', '2020-02-25 15:07:01.469', NULL),
	(17, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '2019-08-21 15:14:23.000', '2020-02-25 15:07:01.469', NULL),
	(18, NULL, NULL, NULL, NULL, NULL, NULL, 'null', '2019-09-11 21:46:46.000', '2020-02-25 15:07:01.469', NULL),
	(19, NULL, NULL, NULL, NULL, NULL, NULL, 'null', '2019-09-11 21:46:47.000', '2020-02-25 15:07:01.469', NULL),
	(20, NULL, 'Jokky💫', NULL, NULL, NULL, NULL, '{"city": "Nantong", "gender": 1, "country": "China", "language": "zh_HK", "nickName": "Jokky💫", "province": "Jiangsu", "avatarUrl": "https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIW4BKAJ072RS11HnwVDR4OEIlt3fr61JyaucribV8KUXWEmYvJ2VUzyYShdibrT5PPP95Y8DKLK2Lg/132"}', '2019-09-11 22:12:25.000', '2020-02-25 15:07:01.469', NULL),
	(21, NULL, 'Colorful3', NULL, NULL, NULL, NULL, '{"city": "", "gender": 1, "country": "Fiji Islands", "language": "zh_CN", "nickName": "Colorful3", "province": "", "avatarUrl": "https://wx.qlogo.cn/mmopen/vi_32/DYAIOgq83eobqRnYzViaMmIt9gAKCIBkib7pF9lcsAmaFfzKFU0MCr5XcLfLzKIckyib3PSFyqMP2ksq3ibFibxb05A/132"}', '2019-09-11 22:12:37.000', '2020-02-25 15:07:01.469', NULL),
	(22, NULL, '一飞同学@Evan', NULL, NULL, NULL, NULL, '{"city": "Suzhou", "gender": 1, "country": "China", "language": "zh_CN", "nickName": "一飞同学@Evan", "province": "Jiangsu", "avatarUrl": "https://wx.qlogo.cn/mmopen/vi_32/DYAIOgq83epwgf48jZfAzOz78zxQOqtWD9SuQNgIUdhficfteoMl8MBcTUOJWTfqTDbYdaEheGWjPp9cSD5Uiclw/132"}', '2019-09-11 22:13:01.000', '2020-02-25 15:07:01.469', NULL),
	(23, NULL, '流乔', NULL, NULL, NULL, NULL, '{"city": "", "gender": 0, "country": "", "language": "zh_CN", "nickName": "流乔", "province": "", "avatarUrl": "https://wx.qlogo.cn/mmopen/vi_32/aN5d8M5StwcQLOic6FkLYrHcpxdaNR7CLfwfBOl9ThCesTjUAVAnR2WyE1sTBficepZ526KAn98bpRpJ35rnGElQ/132"}', '2019-09-11 22:18:23.000', '2020-02-25 15:07:01.469', NULL),
	(24, NULL, 'pedro', NULL, NULL, NULL, NULL, '{"city": "Wuhan", "gender": 1, "country": "China", "language": "zh_CN", "nickName": "pedro", "province": "Hubei", "avatarUrl": "https://wx.qlogo.cn/mmopen/vi_32/So1cw4tZWziadtWHyqALcSSUY2dOjmZ669eARYZoJVMlSSBMzqGFlekHHic0a2MfFCYicURXTghib23mmzjYA2BjLg/132"}', '2019-09-11 22:21:08.000', '2020-02-25 15:07:01.469', NULL),
	(25, NULL, NULL, NULL, NULL, NULL, NULL, 'null', '2019-09-12 05:31:35.000', '2020-02-25 15:07:01.469', NULL),
	(26, NULL, NULL, NULL, NULL, NULL, NULL, 'null', '2019-09-12 05:31:48.000', '2020-02-25 15:07:01.469', NULL),
	(27, NULL, NULL, NULL, NULL, NULL, NULL, 'null', '2019-09-12 05:59:03.000', '2020-02-25 15:07:01.469', NULL),
	(28, NULL, NULL, NULL, NULL, NULL, NULL, 'null', '2019-09-12 06:00:18.000', '2020-02-25 15:07:01.469', NULL),
	(29, NULL, NULL, NULL, NULL, NULL, NULL, 'null', '2019-09-12 06:55:27.000', '2020-02-25 15:07:01.469', NULL),
	(30, NULL, NULL, NULL, NULL, NULL, NULL, 'null', '2019-09-12 17:48:25.000', '2020-02-25 15:07:01.469', NULL),
	(31, NULL, NULL, NULL, NULL, NULL, NULL, 'null', '2019-09-15 18:10:08.000', '2020-02-25 15:07:01.469', NULL),
	(32, NULL, NULL, NULL, NULL, NULL, NULL, 'null', '2019-09-16 03:24:38.000', '2020-02-25 15:07:01.469', NULL),
	(33, NULL, NULL, NULL, NULL, NULL, NULL, 'null', '2019-09-16 04:15:19.000', '2020-02-25 15:07:01.469', NULL);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;

-- 导出  表 mall.user_coupon 结构
CREATE TABLE IF NOT EXISTS `user_coupon` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL,
  `coupon_id` int unsigned NOT NULL,
  `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '1:未使用，2：已使用， 3：已过期',
  `create_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `order_id` int unsigned DEFAULT NULL,
  `update_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_user_coupon` (`user_id`,`coupon_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  mall.user_coupon 的数据：~23 rows (大约)
/*!40000 ALTER TABLE `user_coupon` DISABLE KEYS */;
INSERT INTO `user_coupon` (`id`, `user_id`, `coupon_id`, `status`, `create_time`, `order_id`, `update_time`) VALUES
	(6, 16, 4, 1, '2019-08-20 10:50:20.000', NULL, '2019-09-08 06:06:51.000'),
	(8, 16, 3, 1, '2019-08-20 10:53:06.000', NULL, '2019-09-15 21:14:08.000'),
	(10, 16, 7, 1, '2019-08-25 18:03:54.000', NULL, '2019-09-10 12:21:16.000'),
	(11, 16, 11, 1, '2019-08-28 03:51:07.000', NULL, '2019-09-15 05:53:03.000'),
	(12, 19, 3, 2, '2019-09-11 21:49:55.000', 245, '2019-09-11 21:50:16.000'),
	(13, 19, 4, 1, '2019-09-11 21:51:30.000', NULL, '2019-09-11 21:51:30.000'),
	(14, 19, 7, 1, '2019-09-11 21:51:31.000', NULL, '2019-09-11 21:51:31.000'),
	(15, 19, 11, 1, '2019-09-11 21:51:32.000', NULL, '2019-09-11 21:51:32.000'),
	(16, 22, 11, 1, '2019-09-11 22:14:23.000', NULL, '2019-09-11 22:14:23.000'),
	(17, 21, 3, 1, '2019-09-11 22:14:25.000', NULL, '2019-09-11 22:14:25.000'),
	(18, 22, 3, 2, '2019-09-11 22:14:34.000', 249, '2019-09-11 22:17:22.000'),
	(19, 22, 4, 1, '2019-09-11 22:16:10.000', NULL, '2019-09-11 22:16:10.000'),
	(20, 23, 3, 1, '2019-09-11 22:20:04.000', NULL, '2019-09-11 22:20:04.000'),
	(21, 23, 11, 1, '2019-09-11 22:20:06.000', NULL, '2019-09-11 22:20:06.000'),
	(22, 21, 11, 1, '2019-09-11 22:34:39.000', NULL, '2019-09-11 22:34:39.000'),
	(23, 20, 11, 1, '2019-09-11 22:44:26.000', NULL, '2019-09-11 22:44:26.000'),
	(24, 20, 3, 1, '2019-09-11 22:44:47.000', NULL, '2019-09-11 22:44:47.000'),
	(25, 26, 3, 1, '2019-09-12 05:32:09.000', NULL, '2019-09-12 05:32:09.000'),
	(26, 30, 4, 1, '2019-09-12 17:48:57.000', NULL, '2019-09-12 17:48:57.000'),
	(27, 30, 7, 1, '2019-09-12 17:48:59.000', NULL, '2019-09-12 17:48:59.000'),
	(28, 30, 3, 1, '2019-09-12 17:49:00.000', NULL, '2019-09-12 17:49:00.000'),
	(29, 20, 4, 1, '2019-09-14 08:12:51.000', NULL, '2019-09-14 08:12:51.000'),
	(30, 20, 7, 1, '2019-09-15 22:36:37.000', NULL, '2019-09-15 22:36:37.000');
/*!40000 ALTER TABLE `user_coupon` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
