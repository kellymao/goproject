-- MySQL dump 10.13  Distrib 5.7.9, for osx10.9 (x86_64)
--
-- Host: localhost    Database: yyserver
-- ------------------------------------------------------
-- Server version	5.7.9

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `apis`
--

DROP TABLE IF EXISTS `apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `apis` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `group` varchar(255) DEFAULT NULL,
  `relatemenuid` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `apis`
--

LOCK TABLES `apis` WRITE;
/*!40000 ALTER TABLE `apis` DISABLE KEYS */;
INSERT INTO `apis` VALUES (1,'2020-03-13 06:45:22','2020-03-13 06:45:22',NULL,'/menu/getmenu','获取菜单','menu',NULL),(2,'2020-03-13 06:46:40','2020-03-13 06:46:40',NULL,'/role/getrolelist','获取角色','role',NULL),(3,'2020-03-13 06:47:56','2020-03-13 06:47:56',NULL,'/role_menu_tree/getmenutree','获取角色资源','role_menu_tree',NULL),(4,'2020-03-13 06:48:30','2020-03-13 06:48:30',NULL,'/role_menu_tree/savemenutree','修改角色资源','role_menu_tree',NULL),(5,'2020-03-13 06:49:14','2020-03-13 06:49:14',NULL,'/user/getuserlist','获取用户','user',NULL),(6,'2020-03-13 06:49:47','2020-03-13 06:49:47',NULL,'/api/getapilist','获取权限','api',NULL),(7,'2020-04-17 17:22:57','2020-04-17 17:22:57',NULL,'/api/getallapis','获取所有权限','api',NULL),(8,'2020-04-17 17:27:28','2020-04-17 17:27:28',NULL,'/role_to_api/getapilist','获取角色的权限','role_to_api',NULL);
/*!40000 ALTER TABLE `apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `authorities`
--

DROP TABLE IF EXISTS `authorities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `authorities` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `authority_id` varchar(255) NOT NULL,
  `authority_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `authority_id` (`authority_id`),
  KEY `idx_authorities_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `authorities`
--

LOCK TABLES `authorities` WRITE;
/*!40000 ALTER TABLE `authorities` DISABLE KEYS */;
INSERT INTO `authorities` VALUES (2,'2019-09-08 16:18:45','2019-09-08 16:18:45',NULL,'888','普通用户'),(6,'2019-09-18 22:23:33','2019-09-18 22:23:33',NULL,'9528','测试角色'),(12,'2019-10-09 23:04:18','2019-10-09 23:04:18',NULL,'999','封禁');
/*!40000 ALTER TABLE `authorities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menus`
--

DROP TABLE IF EXISTS `menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `menus` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `menu_level` int(10) unsigned DEFAULT NULL,
  `parent_id` varchar(255) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `hidden` tinyint(1) DEFAULT NULL,
  `component` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `icon` varchar(255) DEFAULT NULL,
  `nick_name` varchar(255) DEFAULT NULL,
  `menuid` varchar(255) DEFAULT NULL,
  `authority_id` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menus`
--

LOCK TABLES `menus` WRITE;
/*!40000 ALTER TABLE `menus` DISABLE KEYS */;
INSERT INTO `menus` VALUES (1,'2020-02-24 10:59:26','2020-02-24 10:59:45',NULL,NULL,'0','/404','404',1,'404.vue',NULL,'',NULL,'1','888'),(2,'2020-02-24 10:59:26','2020-02-24 10:59:45',NULL,NULL,'0','/','系统管理',0,'Home.vue',NULL,'ios-home',NULL,'2','888'),(3,'2020-02-24 10:59:26','2020-02-24 10:59:45',NULL,NULL,'0','/','数据库管理',0,'Home.vue',NULL,'ios-apps',NULL,'3','888'),(4,'2020-02-24 10:59:26','2020-02-24 10:59:45',NULL,NULL,'0','/','导航三',0,'Home.vue',NULL,'ios-paw',NULL,'4','888'),(5,'2020-02-24 10:59:26','2020-02-24 10:59:45',NULL,NULL,'0','/','导航四',0,'Home.vue',NULL,'ios-cog',NULL,'5','888'),(6,'2020-02-24 11:06:56','2020-02-24 11:06:56',NULL,NULL,'2','/main','主页',1,'Main.vue',NULL,'ios-paw',NULL,'6','888'),(7,'2020-02-24 11:06:56','2020-02-24 11:06:56',NULL,NULL,'2','/table','表格',0,'nav1/Table.vue',NULL,'ios-paw',NULL,'7','888'),(8,'2020-02-24 11:06:56','2020-02-24 11:06:56',NULL,NULL,'2','/form','表单',0,'nav1/Form.vue',NULL,'ios-compass',NULL,'8','888'),(9,'2020-02-24 11:06:56','2020-02-24 11:06:56',NULL,NULL,'2','/user','用户管理',0,'nav1/user.vue',NULL,'ios-flask',NULL,'9','888'),(10,'2020-02-29 22:20:52','2020-02-29 22:20:52',NULL,NULL,'2','/role','角色管理',0,'nav1/role.vue',NULL,'ios-map',NULL,'10','888'),(11,'2020-03-05 08:13:53','2020-03-05 08:13:53',NULL,NULL,'2','/*/edit','编辑',1,'edit.vue',NULL,'ios-paw',NULL,'11','888'),(12,'2020-03-13 06:39:56','2020-03-13 06:39:56',NULL,NULL,'2','/api','权限管理',0,'nav1/Api.vue',NULL,'ios-cog',NULL,'12','888'),(13,'2020-03-17 11:34:09','2020-03-17 11:34:09',NULL,NULL,'3','/dbmanage','数据库监控',0,'nav2/Dbmonitor.vue',NULL,'ios-compass',NULL,'13',NULL),(14,'2020-04-17 17:33:26','2020-04-17 17:33:26',NULL,NULL,'5','/echarts','echart',0,'charts/echarts.vue',NULL,'ios-cog',NULL,'14',NULL),(15,'2020-04-17 17:35:20','2020-04-17 17:35:20',NULL,NULL,'5','/treechat','treechart',0,'charts/charts.vue',NULL,'ios-map',NULL,'15',NULL);
/*!40000 ALTER TABLE `menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_to_apis`
--

DROP TABLE IF EXISTS `role_to_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role_to_apis` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `roleid` varchar(255) DEFAULT NULL,
  `apiid` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_role_to_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_to_apis`
--

LOCK TABLES `role_to_apis` WRITE;
/*!40000 ALTER TABLE `role_to_apis` DISABLE KEYS */;
INSERT INTO `role_to_apis` VALUES (7,'2020-03-15 22:53:33','2020-03-15 22:53:33','2020-03-15 22:54:15','888','1'),(8,'2020-03-15 22:54:15','2020-03-15 22:54:15',NULL,'888','1'),(9,'2020-03-15 22:54:15','2020-03-15 22:54:15',NULL,'888','2'),(10,'2020-03-15 22:54:15','2020-03-15 22:54:15',NULL,'888','3'),(11,'2020-03-15 22:54:15','2020-03-15 22:54:15',NULL,'888','4'),(12,'2020-03-15 22:54:15','2020-03-15 22:54:15',NULL,'888','5'),(13,'2020-03-15 22:54:15','2020-03-15 22:54:15',NULL,'888','6'),(14,'2020-03-15 23:02:37','2020-03-15 23:02:37',NULL,'9528','1'),(15,'2020-03-15 23:02:37','2020-03-15 23:02:37',NULL,'9528','2'),(16,'2020-03-15 23:02:37','2020-03-15 23:02:37',NULL,'9528','3'),(17,'2020-03-15 23:02:37','2020-03-15 23:02:37',NULL,'9528','4'),(18,'2020-03-16 08:00:48','2020-03-16 08:00:48',NULL,'999','1'),(19,'2020-04-17 17:23:57','2020-04-17 17:23:57',NULL,'888','7'),(20,'2020-04-17 17:27:34','2020-04-17 17:27:34',NULL,'888','8');
/*!40000 ALTER TABLE `role_to_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_to_menus`
--

DROP TABLE IF EXISTS `role_to_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role_to_menus` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `roleid` varchar(255) DEFAULT NULL,
  `menuid` varchar(255) DEFAULT NULL,
  `parentid` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_role_to_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=273 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_to_menus`
--

LOCK TABLES `role_to_menus` WRITE;
/*!40000 ALTER TABLE `role_to_menus` DISABLE KEYS */;
INSERT INTO `role_to_menus` VALUES (158,'2020-03-11 09:15:21','2020-03-11 09:15:21','2020-03-12 21:05:55','888','3','0'),(159,'2020-03-11 09:15:21','2020-03-11 09:15:21','2020-03-12 21:05:55','888','4','0'),(160,'2020-03-11 09:15:21','2020-03-11 09:15:21','2020-03-12 21:05:55','888','5','0'),(164,'2020-03-11 09:19:58','2020-03-11 09:19:58','2020-03-11 09:22:02','9528','2','0'),(167,'2020-03-11 09:19:58','2020-03-11 09:19:58','2020-03-11 09:22:02','9528','4','0'),(168,'2020-03-11 09:22:02','2020-03-11 09:22:02','2020-03-11 09:22:30','9528','2','0'),(169,'2020-03-11 09:22:02','2020-03-11 09:22:02','2020-03-11 09:22:30','9528','7','2'),(170,'2020-03-11 09:22:02','2020-03-11 09:22:02','2020-03-11 09:22:30','9528','8','2'),(171,'2020-03-11 09:22:02','2020-03-11 09:22:02','2020-03-11 09:22:30','9528','4','0'),(172,'2020-03-11 09:22:30','2020-03-11 09:22:30',NULL,'9528','4','0'),(173,'2020-03-11 09:22:30','2020-03-11 09:22:30',NULL,'9528','1','0'),(174,'2020-03-11 09:26:41','2020-03-11 09:26:41','2020-03-11 09:27:08','999','1','0'),(175,'2020-03-11 09:26:41','2020-03-11 09:26:41','2020-03-11 09:27:08','999','4','0'),(176,'2020-03-11 09:27:08','2020-03-11 09:27:08','2020-03-11 09:27:24','999','2','0'),(177,'2020-03-11 09:27:08','2020-03-11 09:27:08','2020-03-11 09:27:24','999','7','2'),(178,'2020-03-11 09:27:08','2020-03-11 09:27:08','2020-03-11 09:27:24','999','3','0'),(179,'2020-03-11 09:27:24','2020-03-11 09:27:24','2020-03-14 22:14:00','999','3','0'),(180,'2020-03-12 21:05:55','2020-03-12 21:05:55','2020-03-15 22:50:04','888','2','0'),(181,'2020-03-12 21:05:55','2020-03-12 21:05:55','2020-03-15 22:50:04','888','7','2'),(182,'2020-03-12 21:05:55','2020-03-12 21:05:55','2020-03-15 22:50:04','888','8','2'),(183,'2020-03-12 21:05:55','2020-03-12 21:05:55','2020-03-15 22:50:04','888','3','0'),(184,'2020-03-12 21:05:55','2020-03-12 21:05:55','2020-03-15 22:50:04','888','4','0'),(185,'2020-03-12 21:05:55','2020-03-12 21:05:55','2020-03-15 22:50:04','888','5','0'),(186,'2020-03-14 22:14:00','2020-03-14 22:14:00','2020-03-15 23:33:32','999','3','0'),(187,'2020-03-14 22:14:00','2020-03-14 22:14:00','2020-03-15 23:33:32','999','4','0'),(188,'2020-03-15 22:50:04','2020-03-15 22:50:04','2020-03-17 11:37:50','888','1','0'),(189,'2020-03-15 22:50:04','2020-03-15 22:50:04','2020-03-17 11:37:50','888','2','0'),(190,'2020-03-15 22:50:04','2020-03-15 22:50:04','2020-03-17 11:37:50','888','7','2'),(191,'2020-03-15 22:50:04','2020-03-15 22:50:04','2020-03-17 11:37:50','888','8','2'),(192,'2020-03-15 22:50:04','2020-03-15 22:50:04','2020-03-17 11:37:50','888','3','0'),(193,'2020-03-15 22:50:04','2020-03-15 22:50:04','2020-03-17 11:37:50','888','4','0'),(194,'2020-03-15 22:50:04','2020-03-15 22:50:04','2020-03-17 11:37:50','888','5','0'),(195,'2020-03-15 23:33:32','2020-03-15 23:33:32','2020-03-16 08:00:05','999','1','0'),(196,'2020-03-15 23:33:32','2020-03-15 23:33:32','2020-03-16 08:00:05','999','3','0'),(197,'2020-03-15 23:33:32','2020-03-15 23:33:32','2020-03-16 08:00:05','999','4','0'),(198,'2020-03-16 08:00:05','2020-03-16 08:00:05','2020-03-16 08:04:01','999','1','0'),(199,'2020-03-16 08:00:05','2020-03-16 08:00:05','2020-03-16 08:04:01','999','2','0'),(200,'2020-03-16 08:00:05','2020-03-16 08:00:05','2020-03-16 08:04:01','999','9','2'),(201,'2020-03-16 08:00:05','2020-03-16 08:00:05','2020-03-16 08:04:01','999','10','2'),(202,'2020-03-16 08:00:05','2020-03-16 08:00:05','2020-03-16 08:04:01','999','11','2'),(203,'2020-03-16 08:00:05','2020-03-16 08:00:05','2020-03-16 08:04:01','999','12','2'),(204,'2020-03-16 08:00:05','2020-03-16 08:00:05','2020-03-16 08:04:01','999','3','0'),(205,'2020-03-16 08:00:05','2020-03-16 08:00:05','2020-03-16 08:04:01','999','4','0'),(206,'2020-03-16 08:04:01','2020-03-16 08:04:01',NULL,'999','1','0'),(207,'2020-03-16 08:04:01','2020-03-16 08:04:01',NULL,'999','2','0'),(208,'2020-03-16 08:04:01','2020-03-16 08:04:01',NULL,'999','6','2'),(209,'2020-03-16 08:04:01','2020-03-16 08:04:01',NULL,'999','9','2'),(210,'2020-03-16 08:04:01','2020-03-16 08:04:01',NULL,'999','10','2'),(211,'2020-03-16 08:04:01','2020-03-16 08:04:01',NULL,'999','11','2'),(212,'2020-03-16 08:04:01','2020-03-16 08:04:01',NULL,'999','12','2'),(213,'2020-03-16 08:04:01','2020-03-16 08:04:01',NULL,'999','3','0'),(214,'2020-03-16 08:04:01','2020-03-16 08:04:01',NULL,'999','4','0'),(215,'2020-03-17 11:37:50','2020-03-17 11:37:50','2020-04-17 17:09:20','888','1','0'),(216,'2020-03-17 11:37:50','2020-03-17 11:37:50','2020-04-17 17:09:20','888','2','0'),(217,'2020-03-17 11:37:50','2020-03-17 11:37:50','2020-04-17 17:09:20','888','7','2'),(218,'2020-03-17 11:37:50','2020-03-17 11:37:50','2020-04-17 17:09:20','888','10','2'),(219,'2020-03-17 11:37:50','2020-03-17 11:37:50','2020-04-17 17:09:20','888','3','0'),(220,'2020-03-17 11:37:50','2020-03-17 11:37:50','2020-04-17 17:09:20','888','13','3'),(221,'2020-03-17 11:37:50','2020-03-17 11:37:50','2020-04-17 17:09:20','888','4','0'),(222,'2020-03-17 11:37:50','2020-03-17 11:37:50','2020-04-17 17:09:20','888','5','0'),(223,'2020-04-17 17:09:20','2020-04-17 17:09:20','2020-04-17 17:28:25','888','1','0'),(224,'2020-04-17 17:09:20','2020-04-17 17:09:20','2020-04-17 17:28:25','888','2','0'),(225,'2020-04-17 17:09:20','2020-04-17 17:09:20','2020-04-17 17:28:25','888','7','2'),(226,'2020-04-17 17:09:20','2020-04-17 17:09:20','2020-04-17 17:28:25','888','10','2'),(227,'2020-04-17 17:09:20','2020-04-17 17:09:20','2020-04-17 17:28:25','888','3','0'),(228,'2020-04-17 17:09:20','2020-04-17 17:09:20','2020-04-17 17:28:25','888','13','3'),(229,'2020-04-17 17:09:20','2020-04-17 17:09:20','2020-04-17 17:28:25','888','4','0'),(230,'2020-04-17 17:09:20','2020-04-17 17:09:20','2020-04-17 17:28:25','888','5','0'),(231,'2020-04-17 17:28:25','2020-04-17 17:28:25','2020-04-17 17:34:02','888','1','0'),(232,'2020-04-17 17:28:25','2020-04-17 17:28:25','2020-04-17 17:34:02','888','2','0'),(233,'2020-04-17 17:28:25','2020-04-17 17:28:25','2020-04-17 17:34:02','888','6','2'),(234,'2020-04-17 17:28:25','2020-04-17 17:28:25','2020-04-17 17:34:02','888','7','2'),(235,'2020-04-17 17:28:25','2020-04-17 17:28:25','2020-04-17 17:34:02','888','8','2'),(236,'2020-04-17 17:28:25','2020-04-17 17:28:25','2020-04-17 17:34:02','888','9','2'),(237,'2020-04-17 17:28:25','2020-04-17 17:28:25','2020-04-17 17:34:02','888','10','2'),(238,'2020-04-17 17:28:25','2020-04-17 17:28:25','2020-04-17 17:34:02','888','11','2'),(239,'2020-04-17 17:28:25','2020-04-17 17:28:25','2020-04-17 17:34:02','888','12','2'),(240,'2020-04-17 17:28:25','2020-04-17 17:28:25','2020-04-17 17:34:02','888','3','0'),(241,'2020-04-17 17:28:25','2020-04-17 17:28:25','2020-04-17 17:34:02','888','13','3'),(242,'2020-04-17 17:28:25','2020-04-17 17:28:25','2020-04-17 17:34:02','888','4','0'),(243,'2020-04-17 17:28:25','2020-04-17 17:28:25','2020-04-17 17:34:02','888','5','0'),(244,'2020-04-17 17:34:02','2020-04-17 17:34:02','2020-04-17 17:36:29','888','1','0'),(245,'2020-04-17 17:34:02','2020-04-17 17:34:02','2020-04-17 17:36:29','888','2','0'),(246,'2020-04-17 17:34:02','2020-04-17 17:34:02','2020-04-17 17:36:29','888','6','2'),(247,'2020-04-17 17:34:02','2020-04-17 17:34:02','2020-04-17 17:36:29','888','7','2'),(248,'2020-04-17 17:34:02','2020-04-17 17:34:02','2020-04-17 17:36:29','888','8','2'),(249,'2020-04-17 17:34:02','2020-04-17 17:34:02','2020-04-17 17:36:29','888','9','2'),(250,'2020-04-17 17:34:02','2020-04-17 17:34:02','2020-04-17 17:36:29','888','10','2'),(251,'2020-04-17 17:34:02','2020-04-17 17:34:02','2020-04-17 17:36:29','888','11','2'),(252,'2020-04-17 17:34:02','2020-04-17 17:34:02','2020-04-17 17:36:29','888','12','2'),(253,'2020-04-17 17:34:02','2020-04-17 17:34:02','2020-04-17 17:36:29','888','3','0'),(254,'2020-04-17 17:34:02','2020-04-17 17:34:02','2020-04-17 17:36:29','888','13','3'),(255,'2020-04-17 17:34:02','2020-04-17 17:34:02','2020-04-17 17:36:29','888','4','0'),(256,'2020-04-17 17:34:02','2020-04-17 17:34:02','2020-04-17 17:36:29','888','5','0'),(257,'2020-04-17 17:34:02','2020-04-17 17:34:02','2020-04-17 17:36:29','888','14','5'),(258,'2020-04-17 17:36:29','2020-04-17 17:36:29',NULL,'888','1','0'),(259,'2020-04-17 17:36:29','2020-04-17 17:36:29',NULL,'888','2','0'),(260,'2020-04-17 17:36:29','2020-04-17 17:36:29',NULL,'888','6','2'),(261,'2020-04-17 17:36:29','2020-04-17 17:36:29',NULL,'888','7','2'),(262,'2020-04-17 17:36:29','2020-04-17 17:36:29',NULL,'888','8','2'),(263,'2020-04-17 17:36:29','2020-04-17 17:36:29',NULL,'888','9','2'),(264,'2020-04-17 17:36:29','2020-04-17 17:36:29',NULL,'888','10','2'),(265,'2020-04-17 17:36:29','2020-04-17 17:36:29',NULL,'888','11','2'),(266,'2020-04-17 17:36:29','2020-04-17 17:36:29',NULL,'888','12','2'),(267,'2020-04-17 17:36:29','2020-04-17 17:36:29',NULL,'888','3','0'),(268,'2020-04-17 17:36:29','2020-04-17 17:36:29',NULL,'888','13','3'),(269,'2020-04-17 17:36:29','2020-04-17 17:36:29',NULL,'888','4','0'),(270,'2020-04-17 17:36:29','2020-04-17 17:36:29',NULL,'888','5','0'),(271,'2020-04-17 17:36:29','2020-04-17 17:36:29',NULL,'888','14','5'),(272,'2020-04-17 17:36:29','2020-04-17 17:36:29',NULL,'888','15','5');
/*!40000 ALTER TABLE `role_to_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `t1`
--

DROP TABLE IF EXISTS `t1`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `t1` (
  `id` int(11) NOT NULL,
  `createtime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `t1`
--

LOCK TABLES `t1` WRITE;
/*!40000 ALTER TABLE `t1` DISABLE KEYS */;
INSERT INTO `t1` VALUES (0,'2020-03-12 09:48:13');
/*!40000 ALTER TABLE `t1` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `uuid` varbinary(255) DEFAULT NULL,
  `user_name` varchar(255) DEFAULT NULL,
  `pass_word` varchar(255) DEFAULT NULL,
  `nick_name` varchar(255) DEFAULT 'QMPlusUser',
  `header_img` varchar(255) DEFAULT 'http://www.henrongyi.top/avatar/lufu.jpg',
  `authority_id` double DEFAULT '888',
  `authority_name` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_users_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (10,'2019-09-13 09:23:46','2019-10-21 03:16:03',NULL,'ce0d6685-c15f-4126-a5b4-890bc9d2356d',NULL,NULL,'超级管理员','http://qmplusimg.henrongyi.top/1571627762timg.jpg',888,NULL,'admin','e10adc3949ba59abbe56e057f20f883e'),(11,'2019-09-13 09:27:29','2019-09-13 09:27:29',NULL,'fd6ef79b-944c-4888-8377-abe2d2608858',NULL,NULL,'QMPlusUser','http://www.henrongyi.top/avatar/lufu.jpg',888,NULL,'a303176530','3ec063004a6f31642261936a379fde3d'),(12,'2019-09-13 09:28:56','2019-09-13 09:28:56',NULL,'e799cec6-4c7f-438c-8647-7d5c3397451e',NULL,NULL,'QMPlusUser','http://www.henrongyi.top/avatar/lufu.jpg',888,NULL,'a30317465','3ec063004a6f31642261936a379fde3d'),(13,'2019-09-13 09:29:26','2019-09-13 09:29:26',NULL,'ee7d5922-2331-41bc-93c2-a6e6da0de420',NULL,NULL,'QMPlusUser','http://www.henrongyi.top/avatar/lufu.jpg',888,NULL,'a30317465','3ec063004a6f31642261936a379fde3d'),(14,'2019-09-13 09:29:28','2019-09-13 09:29:28',NULL,'5b4d34a2-42f5-47c5-a92d-a67ae6d4ad34',NULL,NULL,'QMPlusUser','http://www.henrongyi.top/avatar/lufu.jpg',888,NULL,'a30317465','3ec063004a6f31642261936a379fde3d'),(15,'2019-09-13 09:31:16','2019-10-09 15:04:28',NULL,'40d7946a-6728-4e6f-9d40-1425efe81602',NULL,NULL,'QMPlusUser','http://www.henrongyi.top/avatar/lufu.jpg',999,NULL,'user','e10adc3949ba59abbe56e057f20f883e');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-04-20 13:32:39
