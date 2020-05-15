-- MySQL dump 10.15  Distrib 10.0.32-MariaDB, for debian-linux-gnueabihf (armv8l)
--
-- Host: localhost    Database: homepage
-- ------------------------------------------------------
-- Server version	10.0.32-MariaDB-0+deb8u1

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

DROP DATABASE IF EXISTS homepage;
CREATE DATABASE homepage;
USE homepage;

--
-- Table structure for table `ar_internal_metadata`
--

DROP TABLE IF EXISTS `ar_internal_metadata`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ar_internal_metadata` (
  `key` varchar(255) NOT NULL,
  `value` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ar_internal_metadata`
--

LOCK TABLES `ar_internal_metadata` WRITE;
/*!40000 ALTER TABLE `ar_internal_metadata` DISABLE KEYS */;
INSERT INTO `ar_internal_metadata` VALUES ('environment','production','2018-05-23 15:16:29','2018-05-23 15:16:29');
/*!40000 ALTER TABLE `ar_internal_metadata` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `equipments`
--

DROP TABLE IF EXISTS `equipments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `equipments` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `num` int(11) DEFAULT NULL,
  `note` varchar(255) DEFAULT NULL,
  `tag_id` bigint(20) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_equipments_on_tag_id` (`tag_id`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `equipments`
--

LOCK TABLES `equipments` WRITE;
/*!40000 ALTER TABLE `equipments` DISABLE KEYS */;
INSERT INTO `equipments` VALUES (1,'Surface book',1,'',1,'2018-05-25 16:40:33','2018-05-25 16:40:33'),(2,'Kotlin スタートブック',1,'',4,'2018-05-25 16:41:21','2018-05-25 16:41:21'),(3,'Docker 実践活用ガイド',1,'',4,'2018-05-25 16:41:46','2018-05-25 16:41:46'),(4,'ゼロから作るDeap Learning',1,'',4,'2018-05-25 16:42:56','2018-05-25 16:42:56'),(5,'Ruby on Rails 4 アプリケーションプログラミング',1,'',4,'2018-05-25 16:43:17','2018-05-25 16:43:17'),(6,'実装 ディープラーニング',1,'',4,'2018-05-25 16:49:07','2018-05-25 16:49:07'),(7,'確かな力が身につく Python「超」入門',1,'',4,'2018-06-08 13:03:08','2018-06-08 13:03:08'),(8,'Pythonプログラミングパーフェクトマスター',1,'',4,'2018-06-08 13:03:44','2018-06-08 13:03:44'),(9,'Python プロフェッショナルプログラミング',1,'',4,'2018-06-08 13:04:09','2018-06-08 13:04:09');
INSERT INTO `equipments` VALUES (10,'HTML5 マークアップガイドブック',1,'',4,'2018-06-08 13:07:41','2018-06-08 13:07:41'),(11,'HTML/CSSデザイン講義',1,'',4,'2018-06-08 13:08:08','2018-06-08 13:08:08'),(12,'ホームページ辞典',1,'',4,'2018-06-08 13:09:31','2018-06-08 13:09:31'),(13,'Sassファーストガイド',1,'',4,'2018-06-08 13:10:23','2018-06-08 13:10:23'),(14,'マスタリング TCP/IP',1,'',4,'2018-06-08 13:10:54','2018-06-08 13:10:54'),(15,'ゲーム開発者のためのAI入門',1,'',4,'2018-06-08 13:11:24','2018-06-08 13:11:24'),(16,'OpenGL プログラミングガイド',1,'',4,'2018-06-08 13:11:49','2018-06-08 13:11:49');
INSERT INTO `equipments` VALUES (20,'uGUIではじめる Unity UIデザインの教科書',1,'',4,'2018-06-08 13:17:20','2018-06-08 13:17:20'),(21,'Unity ライブラリ辞典 ランタイム編',1,'',4,'2018-06-08 13:17:47','2018-06-08 13:17:47'),(22,'実践 OpenCV',1,'',4,'2018-06-08 13:18:18','2018-06-08 13:18:18'),(23,'パーフェクトJava',1,'',4,'2018-06-08 13:18:39','2018-06-08 13:18:39'),(24,'一週間でマスターする SQL',1,'',4,'2018-06-08 13:19:51','2018-06-08 13:19:51'),(25,'Linuxサーバ 構築・設定のすべて',1,'',4,'2018-06-08 13:20:35','2018-06-08 13:20:35');
INSERT INTO `equipments` VALUES (26,'Vim テクニックバイブル',1,'',4,'2018-06-08 13:21:07','2018-06-08 13:21:07'),(27,'シェルスクリプト 基本リファレンス',1,'',4,'2018-06-08 13:21:32','2018-06-08 13:21:32'),(28,'逆引きUNIXコマンド',1,'',4,'2018-06-08 13:21:51','2018-06-08 13:21:51'),(29,'UNIX コマンドブック',1,'',4,'2018-06-08 13:22:16','2018-06-08 13:22:16');
INSERT INTO `equipments` VALUES (30,'初めてのPHP, MySQL, JavaScript&CSS',1,'',4,'2018-06-08 13:23:00','2018-06-08 13:23:00'),(31,'Electronではじめるアプリ開発',1,'',4,'2018-06-08 13:23:20','2018-06-08 13:23:20'),(32,'基礎からのWordPress',1,'',4,'2018-06-08 13:23:36','2018-06-08 13:23:36'),(33,'独習Linux専科',1,'',4,'2018-06-08 13:23:59','2018-06-08 13:23:59'),(34,'初めてのRuby',1,'',4,'2018-06-08 13:24:12','2018-06-08 13:24:12'),(35,'Rubyによるクローラー開発技法',1,'',4,'2018-06-08 13:24:31','2018-06-08 13:24:31'),(36,'パーフェクトRuby on Rails',1,'',4,'2018-06-08 13:24:54','2018-06-08 13:24:54'),(37,'基礎 Ruby on Rails',1,'',4,'2018-06-08 13:25:19','2018-06-08 13:25:19'),(38,'たのしいRuby',1,'',4,'2018-06-08 13:25:36','2018-06-08 13:25:36'),(39,'Amazon Web Services',1,'',4,'2018-06-08 13:26:04','2018-06-08 13:26:04');
INSERT INTO `equipments` VALUES (40,'SEOに効くWebライティング',1,'',4,'2018-06-08 13:26:29','2018-06-08 13:26:29'),(41,'GitHub実践入門',1,'',4,'2018-06-08 13:26:52','2018-06-08 13:26:52'),(42,'入門 Git',1,'',4,'2018-06-08 13:27:08','2018-06-08 13:27:08'),(43,'リーダブルコード',1,'一家に一冊常備すべき本',4,'2018-06-08 13:27:53','2018-06-08 13:27:53'),(44,'絶対に挫折しないiPhoneアプリ開発「超」入門',1,'',4,'2018-06-08 13:28:29','2018-06-08 13:28:29'),(45,'確かな力が身につくPHP「超」入門',1,'',4,'2018-06-08 13:28:53','2018-06-08 13:28:53'),(46,'Leap Motionプログラミング',1,'',4,'2018-06-08 13:29:24','2018-06-08 13:29:24'),(47,'Xamarinプログラミング',1,'',4,'2018-06-08 13:29:42','2018-06-08 13:29:42'),(48,'OpenCV3プログラミングブック',1,'',4,'2018-06-08 13:30:04','2018-06-08 13:30:04'),(49,'Unity5入門',1,'',4,'2018-06-08 13:33:02','2018-06-08 13:33:02');
INSERT INTO `equipments` VALUES (17,"猫でもわかる C#プログラミング",1,'',4,'2018-06-08 13:15:48','2018-06-08 13:15:48');
INSERT INTO `equipments` VALUES (18,"独習 C#",1,'',4,'2018-06-08 13:16:10','2018-06-08 13:16:10');
INSERT INTO `equipments` VALUES (19,"Unity5 3Dゲーム開発講座",1,'',4,'2018-06-08 13:16:50','2018-06-08 13:16:50');
INSERT INTO `equipments` VALUES (50,"KINECT for Windows SDKプログラミング C#編",1,'',4,'2018-06-08 13:33:36','2018-06-08 13:33:36');
INSERT INTO `equipments` VALUES (51,"KINECT センサープログラミング",1,'',4,'2018-06-08 13:33:57','2018-06-08 13:33:57');
INSERT INTO `equipments` VALUES (52,"KINECT for Windows SDKプログラミング v2センサー対応版",1,'',4,'2018-06-08 13:34:40','2018-06-08 13:34:40');
INSERT INTO `equipments` VALUES (53,"KINECTセンサー 画像処理プログラミング",1,'',4,'2018-06-08 13:35:07','2018-06-08 13:35:07');
/*!40000 ALTER TABLE `equipments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `introductions`
--

DROP TABLE IF EXISTS `introductions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `introductions` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `department` varchar(255) DEFAULT NULL,
  `grade` int(11) DEFAULT NULL,
  `comments` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_introductions_on_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `introductions`
--

LOCK TABLES `introductions` WRITE;
/*!40000 ALTER TABLE `introductions` DISABLE KEYS */;
INSERT INTO `introductions` VALUES (5,2,'工学部　情報工学科',0,'お布団暖かい','2018-05-24 04:46:38','2019-05-10 04:47:58'),(39,1,'工学部',0,'こっそり誰かを削除してもバレないのでは...?\r\n','2018-06-09 04:18:12','2018-06-09 04:18:12'),(46,20,'工学部',0,'偉大なるゼミ長（仮）','2018-07-18 00:54:47','2018-07-18 00:54:47'),(47,3,'工学部情報工学科',0,'お金稼ぎが好きです。\r\nAndroidアプリをよく作ってます。','2019-05-10 04:48:40','2019-05-10 04:48:40'),(48,4,'',0,'','2019-05-10 04:48:54','2019-05-10 04:48:54'),(49,15,'工学部情報工学科',0,'それに関する最新のデータは持っていません','2019-05-10 04:49:04','2019-05-10 04:49:04'),(50,16,'工学部情報工学科',0,'https://ikutohiraiwa.ml','2019-05-10 04:49:13','2019-05-10 04:49:13'),(51,17,'工学部情報工学科',0,'他の人が抱腹絶倒の超面白いこと書いてるはず','2019-05-10 04:49:23','2019-05-10 04:49:23'),(52,19,'工学部情報工学科',0,'のんほい良いとこ一度はおいで','2019-05-10 04:49:31','2019-05-10 04:49:31'),(53,13,'工学部情報工学科',0,'maimaiして穏やかに眠りたい。','2019-05-10 04:49:44','2019-05-10 04:49:44'),(56,9,'工学部情報工学科',4,'遊戯王大好き','2019-05-10 04:50:12','2019-05-10 04:50:12'),(77,24,'ちょっとわかりません',3,'ﾊﾟﾝﾂﾀﾍﾞﾀｲ','2019-05-10 06:02:43','2019-05-10 06:02:43'),(81,25,'工学部情報工学科',3,'アニメが好きです','2019-05-11 06:02:26','2019-05-11 06:02:26'),(83,22,'工学部情報工学科',3,'筋トレ大好き','2019-05-11 12:48:55','2019-05-11 12:48:55'),(85,23,'工学部',3,'アニメが好き','2019-05-22 15:40:42','2019-05-22 15:40:42'),(86,21,'工学部情報工学科',3,'コーラ大好き','2019-06-27 10:01:13','2019-06-27 10:01:13'),(87,5,'工学部情報工学科',4,'CoCo壱大好き','2019-07-21 11:13:53','2019-07-21 11:13:53'),(93,7,'工学部情報工学科',4,'ゲーム大好き','2019-12-26 17:46:25','2019-12-26 17:46:25'),(94,6,'工学部情報工学科',4,'ジェフリー大好き','2019-12-26 17:46:34','2019-12-26 17:46:34'),(95,10,'工学部情報工学科',4,'ポイフル大好き','2019-12-26 17:46:45','2019-12-26 17:46:45'),(96,12,'',4,'','2019-12-26 17:46:52','2019-12-26 17:46:52'),(97,11,'工学部情報工学科',4,'寝たい','2019-12-26 17:47:01','2019-12-26 17:47:01'),(98,8,'工学部情報工学科',4,'スキー大好き','2019-12-26 17:47:16','2019-12-26 17:47:16'),(99,26,'工学部情報工学科',3,'','2020-01-06 03:18:57','2020-01-06 03:18:57');
/*!40000 ALTER TABLE `introductions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lectures`
--

DROP TABLE IF EXISTS `lectures`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `lectures` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `file` varchar(255) DEFAULT NULL,
  `comments` varchar(255) DEFAULT NULL,
  `activation` tinyint(1) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_lectures_on_user_id` (`user_id`),
  CONSTRAINT `fk_rails_5a439a4e07` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lectures`
--

LOCK TABLES `lectures` WRITE;
/*!40000 ALTER TABLE `lectures` DISABLE KEYS */;
INSERT INTO `lectures` VALUES (1,2,'2018年度 深層学習レクチャー0','DLlecture0.pdf','PythonとTensorFlowの開発環境構築について',1,'2018-05-25 16:31:07','2018-05-25 16:31:07'),(2,2,'2018年度 深層学習レクチャー1','DLlecture1.pdf','Pythonの基礎について',1,'2018-05-25 16:31:47','2018-05-25 16:31:47'),(3,2,'2018年度 深層学習レクチャー2','DLlecture2.pdf','深層学習の基礎',1,'2018-05-25 16:32:36','2018-05-25 16:32:36'),(4,2,'2018年度 深層学習レクチャー3','DLlecture3.pdf','CNNの基礎と実装',1,'2018-05-25 16:33:23','2018-05-25 16:33:23'),(5,5,'18年度JavaScriptレクチャー資料 part.1','JSLecture1.pdf','担当:丹羽(t315065)',1,'2019-05-10 05:54:47','2019-05-23 05:57:28'),(6,5,'18年度JavaScriptレクチャー資料 part.2','JSLecture2.pdf','担当:丹羽(t315065)',1,'2019-05-10 05:55:16','2019-05-23 04:41:48'),(7,5,'19年度SQLレクチャー資料','SQLLecture.pdf','SQLについて, 簡単な操作, Javaでの操作',1,'2019-05-16 03:43:57','2019-05-23 04:42:49'),(8,26,'19年度python入門資料','PythonLecture.pdf','pythonの環境構築，基本構文',1,'2019-05-23 11:22:59','2019-05-23 11:22:59');
/*!40000 ALTER TABLE `lectures` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pages`
--

DROP TABLE IF EXISTS `pages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pages` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `contents` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pages`
--

LOCK TABLES `pages` WRITE;
/*!40000 ALTER TABLE `pages` DISABLE KEYS */;
/*!40000 ALTER TABLE `pages` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `researches`
--

DROP TABLE IF EXISTS `researches`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `researches` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `author` varchar(255) DEFAULT NULL,
  `file` varchar(255) DEFAULT NULL,
  `comments` varchar(255) DEFAULT NULL,
  `activation` tinyint(1) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `researches`
--

LOCK TABLES `researches` WRITE;
/*!40000 ALTER TABLE `researches` DISABLE KEYS */;
INSERT INTO `researches` VALUES (1,'効果的なだらけ対策アプリケーションの開発','磯谷将',NULL,'2017年度卒業研究',1,'2019-01-18 15:38:45','2019-01-18 15:38:45'),(2,'トレード管理アプリケーションの開発','吉永拓海',NULL,'2017年度卒業研究',1,'2019-01-18 15:40:22','2019-01-18 15:40:22'),(3,'Untiyを用いたゲーム開発','澤田直樹',NULL,'2017年度卒業研究',1,'2019-01-18 15:40:52','2019-01-18 15:40:52'),(4,'Kinectを用いた仮想試着ARアプリ','栗田真衣',NULL,'2017年度卒業研究',1,'2019-01-18 15:42:23','2019-01-18 15:42:23'),(5,'人検出を用いた防犯カメラアプリ','川岸慎一郎',NULL,'2017年度卒業研究',1,'2019-01-18 15:42:58','2019-01-18 15:42:58'),(6,'Bluetooth通信を用いた行動記録アプリケーション','山田隼也',NULL,'2017年度卒業研究',1,'2019-01-18 15:43:38','2019-01-18 15:43:38'),(7,'作業記録アプリケーションの開発','成瀬摩倭',NULL,'2017年度卒業研究',1,'2019-01-18 15:46:32','2019-01-18 15:46:32'),(8,'気象情報とグラフを使用したウォーキング支援アプリの開発','嶋田圭祐',NULL,'2017年度卒業研究',1,'2019-01-18 15:47:25','2019-01-18 15:47:25'),(9,'ボイストレーニングアプリ開発','芦田将輝',NULL,'2018年度卒業研究',1,'2019-01-18 15:49:10','2019-01-18 15:49:10'),(10,'株取引損益率検証アプリの開発','伊藤正大',NULL,'2018年度卒業研究',1,'2019-01-18 15:49:40','2019-01-18 15:49:40'),(11,'お料理支援アプリ開発','河合明香里',NULL,'2018年度卒業研究',1,'2019-01-18 15:50:21','2019-01-18 15:50:21'),(12,'写真撮影支援アプリ開発','近藤朱佳',NULL,'2018年度卒業研究',1,'2019-01-18 15:51:52','2019-01-18 15:51:52'),(13,'LeapMotionを用いた端末操作アプリの開発','中島悠介',NULL,'2018年度卒業研究',1,'2019-01-18 15:53:06','2019-01-18 15:53:06'),(16,'QRコードを用いたランキングシステム開発','平岩郁人',NULL,'2018年度卒業研究',1,'2019-01-18 16:01:55','2019-01-18 16:01:55'),(17,'デュアルトラック録音編集アプリ開発','若林奈々虹',NULL,'2018年度卒業研究',1,'2019-01-18 16:02:43','2019-01-18 16:02:43'),(18,'ジェスチャによるIoTコントロールシステムの開発','服部颯太','卒研発表.pptx','2018年度卒業研究',1,'2019-01-18 16:07:07','2019-01-18 16:07:07'),(20,'農業支援ロボットの開発','池上幸次朗','農業支援ロボットの開発_池上_.pdf','2019年度卒業研究',1,'2019-12-19 06:09:34','2019-12-19 06:09:34'),(21,'LeapMotionを用いた対戦ゲームの作成','今宮友輝','LeapMotionを用いた対戦ゲームの作成.pptx','2019年度卒業研究',1,'2019-12-19 06:31:13','2019-12-19 06:31:13'),(22,'ジョイコンを用いたゲーム開発','嶋﨑平','ジョイコンを用いたゲーム開発.pdf','2019年度卒業研究',1,'2019-12-19 06:31:22','2019-12-19 06:31:22'),(23,'加速度センサーを用いたレースゲームの作成','小室直斗','t316033卒論発表.pdf','2019年度卒業研究',1,'2019-12-19 06:32:56','2019-12-19 06:32:56'),(24,'GoogleAPIを用いた ルート共有アプリ開発','奥嶋碧生','プレゼンテーション2.pptx','2019年度卒業研究',1,'2019-12-19 06:47:17','2019-12-19 06:47:17'),(25,'ディープラーニングを用いたハンドジェスチャ検出及び、IoT・家電コントロールシステムの開発','矢野大喜','卒研発表.pptx','2019年度卒業研究',1,'2019-12-19 06:59:49','2019-12-19 06:59:49'),(30,'振動による方向案内アプリケーション','平野伸之介','DirectionGuidanceApplicationByVibration.pptx','2019年度卒業研究',1,'2019-12-19 08:45:37','2019-12-19 08:45:37');
/*!40000 ALTER TABLE `researches` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schema_migrations`
--

DROP TABLE IF EXISTS `schema_migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `schema_migrations` (
  `version` varchar(255) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schema_migrations`
--

LOCK TABLES `schema_migrations` WRITE;
/*!40000 ALTER TABLE `schema_migrations` DISABLE KEYS */;
INSERT INTO `schema_migrations` VALUES ('20171112033927'),('20171112035013'),('20171112040009'),('20180203042403'),('20180211074112'),('20180211081828'),('20180212082511'),('20180506123648');
/*!40000 ALTER TABLE `schema_migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tags`
--

DROP TABLE IF EXISTS `tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tags` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tags`
--

LOCK TABLES `tags` WRITE;
/*!40000 ALTER TABLE `tags` DISABLE KEYS */;
INSERT INTO `tags` VALUES (1,'パソコン','2018-05-25 16:17:09','2018-05-25 16:17:09'),(2,'タブレット','2018-05-25 16:17:17','2018-05-25 16:17:17'),(3,'ディスプレイ','2018-05-25 16:17:24','2018-05-25 16:17:24'),(4,'本','2018-05-25 16:17:31','2018-05-25 16:17:31'),(5,'コード類','2018-05-25 16:17:36','2018-05-25 16:17:36'),(6,'その他','2018-05-25 16:17:41','2018-05-25 16:17:41');
/*!40000 ALTER TABLE `tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `password_digest` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `student_id` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'吉永 拓海','$2a$10$yZelgpJAx1J95a8OZ/yCieEpqxAEVwxpQKLtdnS00TQo8KodOjoo2','owner','2018-05-23 15:36:18','2018-05-23 16:17:26','t314093'),(2,'服部 颯太','$2a$10$xODf.irGQPzRL1.uQIM4geRwQTub/O5qtXBuAJbd80Rd0uPcoLagC','owner','2018-05-23 15:41:02','2018-05-24 04:46:38','t315067'),(3,'伊藤正大','$2a$10$6oQ1V5GzB0B7riBDU368MeBszSsi4rEx2XQRI81imV7r1Hc3HnwO.','admin','2018-05-24 05:14:54','2018-05-24 05:19:21','t315006'),(4,'河合明香里','$2a$10$LtL7UfUKnuCEiS3dvoDQ0.kCmwI7imZq5s6/6zLMY/sh36K3QexgO','admin','2018-05-24 06:17:54','2018-05-24 06:17:54','t315026'),(5,'池上幸次朗','$2a$10$oOLWOPkzAnT/lye/J1BEu.7k5McYIONXYkMVnSBEacRXwJmpC4ts6','owner','2018-05-31 04:46:24','2019-05-10 04:50:56','t316003'),(6,'嶋崎平','$2a$10$YZbBkm3.mhsu7F52eJpsQ.vLIaKsVhDzhycdqvcLRdqoTw3vT9XyW','member','2018-05-31 04:47:28','2019-12-26 17:46:34','t316039'),(7,'今宮友輝','$2a$10$4rWkiCPNdRLoQygB5B/idOxuzc4hb2dT9/WsZLAtC1mytqZeWumW6','member','2018-05-31 04:48:32','2019-12-26 17:46:25','t316010'),(8,'奥嶋碧生','$2a$10$A1Bkljksaag0FwEq3KxXG.TZbTBqq/0WuBbhRpZ/YpEE0R2RWVraG','member','2018-05-31 04:54:53','2019-12-26 17:47:16','t316019'),(9,'丹下雄太','$2a$10$yApUBWlE1fMCMm1zFxgTCeheZb/naU7Z3QFoUFL7ouTOIySeQeiVG','owner','2018-05-31 04:55:44','2020-01-14 05:30:20','t316058'),(10,'小室直斗','$2a$10$PYYOJ5iMUWRhlMr/0u1QtuDxMAhLVw6si079WZsHiNdVFS1x1pQza','member','2018-05-31 04:59:21','2019-12-26 17:46:45','t316033'),(11,'平野伸之介','$2a$10$W/z1u1564sIzoGLZu7.K1eP4lbCIo/tf4DosPLgUqX8Bbt91nGFD2','member','2018-05-31 05:00:18','2019-12-26 17:47:01','t316077'),(12,'矢野大喜','$2a$10$E5uSQLi7E9kT.fJm9rlWKehT3aZiCV6UEcJhpPmm46HbgKLUBwIWK','member','2018-05-31 06:02:54','2019-12-26 17:46:52','t316091'),(13,'丹羽祐太','$2a$10$vUMsD8NODBdGH0XX7auXHe./95hXDyTRyLURMj4yuLTVsS//C7AlC','member','2018-05-31 16:12:16','2018-06-09 14:12:07','t315065'),(15,'芦田将輝','$2a$10$iuduXHz9zWEK.AXQ6HWJCe8fnZsBpNWPHcRisfPllsTx37HMIF6yW','member','2018-06-01 06:13:11','2018-06-14 05:38:56','t315002'),(16,'平岩郁人','$2a$10$VSwG.zj6SNPyQ/7erCXEeOAtEkoX0nv3vDYPjdLr.Tt74UhlzAxxa','member','2018-06-03 07:51:41','2018-06-03 07:51:41','t315073'),(17,'中島悠介','$2a$10$ni4Fb3zBQ1CC8tyxzOKD4OuDXoaCkEtrWrMG1pszjkKXqL3Q9EG6y','member','2018-06-07 14:50:58','2018-06-07 14:50:58','t315055'),(19,'若林奈々虹','$2a$10$B83uYRiRGQ6ouobASdKLZ.KZ3BHkUPOuGQicZH/.2Tc4SFolM82ZC','member','2018-06-09 14:03:28','2018-06-09 14:07:49','t315095'),(20,'嶋田　ケイスケ','$2a$10$xfsOIYJB0bjPpmM3xnLXV.ctnX5cR287GztHUkIY361iYFCnl42fG','member','2018-07-18 00:54:20','2018-07-18 00:55:28','t314001'),(21,'瀬戸 瑞樹','$2a$10$OS/leldeZMGILAp46JHhbOdcSXwmghy9aVO9hrgkorgF2fC3b8Df.','member','2019-05-10 04:54:33','2019-06-27 10:01:13','t317040'),(22,'塩塚 直人','$2a$10$Of.9isGCD2cbWw2flF72AumV3v70T/8QrwNLuSr7IKY8Gx48d7PF2','member','2019-05-10 04:58:16','2019-05-11 12:48:55','t317031'),(23,'市川 智紀','$2a$10$mcjWMicfytcV5GoQkuAdneRqvw7D9HtGUfxlqeOzeyJPfps8WDu2a','member','2019-05-10 05:00:23','2019-05-16 06:16:51','t317006'),(24,'金谷 明莉','$2a$10$XsToVrP7Y9SLB9QHoonqBO1f0RO6b5.Y4cSODmN.RFPExG6NM2Ita','member','2019-05-10 05:01:31','2019-05-10 05:26:55','t317017'),(25,'小林 優太','$2a$10$y2DxRSIe5hOlU8DA2AyEqejAELntvR.5C2U7tdIpSmNsEb11.4QB2','member','2019-05-10 05:02:07','2019-05-11 06:02:26','t317025'),(26,'大橋 萌枝','$2a$10$ZUaahVNt.SNgXmozarWQz.kMpaMO.AV0TW6aJN9JHHtyFP2iDdvPy','member','2019-05-10 05:02:57','2019-05-10 07:08:04','t317013');
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

-- Dump completed on 2020-05-13  9:16:05

-- add tables
SET CHARSET UTF8;
-- define table
DROP TABLE IF EXISTS `homepage`.`jobs`;
CREATE TABLE IF NOT EXISTS `homepage`.`jobs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `company` varchar(255) DEFAULT NULL,
  `job` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8 COMMENT='就職先企業';

DROP TABLE IF EXISTS `homepage`.`activities`;
CREATE TABLE IF NOT EXISTS `homepage`.`activities` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `show_date` varchar(255) DEFAULT NULL COMMENT '表示用',
  `last_date` datetime NOT NULL COMMENT '並び替え、年度の取得に使う',
  `activity` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='活動内容';

DROP TABLE IF EXISTS `homepage`.`societies`;
CREATE TABLE IF NOT EXISTS `homepage`.`societies` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `author` varchar(255) DEFAULT NULL,
  `society` varchar(255) DEFAULT NULL,
  `award` varchar(255) DEFAULT NULL,
  `date` date DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='学会';

-- dml
SET CHARACTER_SET_CLIENT = utf8;
SET CHARACTER_SET_CONNECTION = utf8;
INSERT INTO `homepage`.`activities`(`show_date`, `last_date`, `activity`, `created_at`, `updated_at`) VALUES ('2019/05/09', '2019/05/09', 'ゼミ説明会練習', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2019/05/16 15:15', '2019/05/16 15:15', '外部者参加可 SQLレクチャー(池上,今宮)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2019/05/23 15:15', '2019/05/23 15:15', '外部者参加可 Pythonミニ入門(大橋)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2019/05/30 18:00', '2019/05/30 18:00', '外部者参加可 深層学習 Part1:基礎(ラシキア)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2019/06/06', '2019/06/06', 'ゼミ紹介最終練習', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2019/06/06 15:15', '2019/06/06 15:15', '外部者参加可 社会人基礎力シリーズ~難しい質問への対応~(ラシキア)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2019/06/13', '2019/06/13', 'ゼミ見学', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2019/06/20', '2019/06/20', '深層学習 Part2:既存の学習済みモデル(小林)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2019/06/20 18:00~', '2019/06/20 18:00', '外部者参加可 社会人基礎力シリーズ~答えたくない質問への対応~(ラシキア)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2019/06/27', '2019/06/27', '深層学習 Part3:ファインチューニング(池上)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2019/07/04', '2019/07/04', '卒論テーマについて', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2019/07/04 18:00', '2019/07/04 18:00', '外部者参加可 プライベートコミュニケーションスキルシリーズ~相手が返事をしない・しなくなった場合の対応~(ラシキア)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2019/07/11', '2019/07/11', 'Thread入門', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2019/07/18', '2019/07/18', 'Thread作成・停止', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2019/07/25', '2019/07/25', 'Thread同期', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('', '2019/06/12', 'Threadプール', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('', '2019/06/11', 'ノンブロッキング非同期処理', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('', '2019/06/10', 'GUIとThread', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2018/06/07', '2018/06/07', 'ゼミ見学会', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2017/05/31', '2017/05/31', '質問対策講座', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2017/06/08', '2017/06/08', 'ゼミ見学会１', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2017/06/15', '2017/06/15', 'ゼミ見学会２', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2015/04/09', '2015/04/09', '春学期ゼミ開講', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2015/04/09', '2015/04/09', '春期休暇報告', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2015/06/04', '2015/06/04', 'ゼミ見学会１', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2015/06/11', '2015/06/11', 'ゼミ見学会２', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2015/07/16', '2015/07/16', '新ゼミ生歓迎会,飲み会予定', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/03/02', '2014/03/02', 'サーバー勉強会', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/04/17', '2014/04/17', 'ゼミ春期休暇報告', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/04/24', '2014/04/24', '3年4年研究発表、PHPレクチャー第1回(関谷、山脇)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/05/01', '2014/05/01', '3年4年研究発表、PHPレクチャー第2回(関谷、山脇)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/05/08', '2014/05/08', '3年4年研究発表、PHPレクチャー第3回(関谷、山脇)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/05/15', '2014/05/15', '3年4年研究発表、スレッドレクチャー', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/05/22', '2014/05/22', '3年4年研究発表、スレッドレクチャー', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/05/28', '2014/05/28', '3年4年研究発表、ゼミ見学発表練習', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/06/12', '2014/06/12', 'ゼミ見学', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/06/19', '2014/06/19', '3年4年研究発表、スレッドレクチャー', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/06/26', '2014/06/26', '3年4年研究発表、スレッドレクチャー', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/07/03', '2014/07/03', '3年4年研究発表、スレッドレクチャー', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/07/10', '2014/07/10', '3年4年研究発表、スレッドレクチャー', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/07/17', '2014/07/17', '3年4年研究発表、スレッドレクチャー', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/07/31', '2014/07/31', '飲み会', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/09/25', '2014/09/25', '夏季休暇報告', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/10/02', '2014/10/02', '2年: ITニュース、C言語レクチャー(長瀬、岩下), 3年: 研究発表、スレッドレクチャー , 4年: 中間発表練習', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/10/09', '2014/10/09', '2年: ITニュース、C言語レクチャー(長瀬、岩下), 3年: 研究発表、スレッドレクチャー , 4年: 中間発表練習', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/10/16', '2014/10/16', '2年: ITニュース、C++/javaレクチャー, 3年: 研究発表、スレッドレクチャー , 4年: 中間発表練習', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/10/23', '2014/10/23', '中間発表', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/10/30', '2014/10/30', '2年: ITニュース、C++/javaレクチャー, 3年: 研究発表、スレッドレクチャー , 4年: 卒研練習', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/11/06', '2014/11/06', '2年: ITニュース、C++/javaレクチャー, 3年: 研究発表、スレッドレクチャー , 4年: 卒研練習', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/11/13', '2014/11/13', '2年: ITニュース、iPhone(竹内、宮宅)/Android(佐野、加藤)レクチャー, 3年: 研究発表、スレッドレクチャー , 4年: 卒研練習', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/11/20', '2014/11/20', '2年: ITニュース、iPhone(竹内、宮宅)/Android(佐野、加藤)レクチャー, 3年: 研究発表、スレッドレクチャー , 4年: 卒研練習', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/11/27', '2014/11/27', '2年: ITニュース、iPhone(竹内、宮宅)/Android(佐野、加藤)レクチャー, 3年: 研究発表、スレッドレクチャー , 4年: 卒研練習', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/12/04', '2014/12/04', '2年: ITニュース、データベースレクチャー(3年宮、伊藤), 3年: 研究発表、スレッドレクチャー , 4年: 卒研練習', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/12/11', '2014/12/11', '2年: ITニュース、 3年: 研究発表、スレッドレクチャー , 4年: 卒研練習', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/12/18', '2014/12/18', '飲み会', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2013/04/11', '2013/04/11', 'ゼミ春期休暇報告', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2013/05/30', '2013/05/30', 'ゼミ紹介 時間:15:10～18:20 場所:AI棟5階', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2013/06/03', '2013/06/03', 'ゼミ紹介 時間:15:10～18:20 場所:ガーデン4階', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2013/06/06', '2013/06/06', 'ゼミ紹介 時間:15:10～18:20 場所:AI棟5階', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2013/09/26', '2013/09/26', 'ゼミ夏期休暇報告', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2013/10/03', '2013/10/03', '2年ITニュース・Cレクチャー(津曲,山脇)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2013/10/10', '2013/10/10', '2年ITニュース・Cレクチャー(津曲,山脇)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2013/10/17', '2013/10/17', '2年ITニュース・Cレクチャー(津曲,山脇)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2013/10/24', '2013/10/24', 'C++/Javaレクチャー(柏木,高田)/(関谷,安岡)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2013/10/31', '2013/10/31', '4年卒研中間発表・2年ITニュース・C++/Javaレクチャー(柏木,高田)/(関谷,安岡)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2013/11/07', '2013/11/07', '2年ITニュース・C++/Javaレクチャー(柏木,高田)/(関谷,安岡)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2013/11/14', '2013/11/14', '2年ITニュース・iPhone/Androidレクチャー(青木,前島)/(小林,高田)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2013/11/21', '2013/11/21', '2年ITニュース・iPhone/Androidレクチャー(青木,前島)/(小林,高田)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2013/11/28', '2013/11/28', 'iPhone/Androidレクチャー(青木,前島)/(小林,高田)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2013/12/05', '2013/12/05', 'JavaScriptレクチャー(池田,前島)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2013/12/12', '2013/12/12', 'JavaScriptレクチャー(池田,前島)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2013/12/19', '2013/12/19', '卒研発表・飲み会・JavaScriptレクチャー(池田,前島)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/1/9', '2014/1/9', 'C＃レクチャー(柏木,李)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2014/1/16', '2014/1/16', 'C＃レクチャー(柏木,李)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/04/05', '2012/04/05', 'ゼミ春期休暇報告', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/04/26', '2012/04/26', '3年研究発表（和田 太陽  益川 瑛大）', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/05/10', '2012/05/10', '3年研究発表（田口 竜一朗  鈴木 冬惟）', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/05/17', '2012/05/17', '3年研究発表（城田 翔平  佐々木 辰哉）', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/05/24', '2012/05/24', '3年研究発表（北野 聖弥  江川 真由）', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/05/31', '2012/05/31', '3年研究発表（石河 洋祐  井坂 尚也）', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/06/07', '2012/06/07', 'ゼミ紹介 時間:15:10～18:20 場所:15号館5階多目的演習室', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/06/11', '2012/06/11', 'ゼミ紹介 時間:15:10～18:20 場所:11号館4階ゼミ室', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/06/14', '2012/06/14', '3年研究発表（和田 太陽)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/06/14', '2012/06/14', 'ゼミ紹介 時間:15:10～18:20 場所:15号館5階多目的演習室', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/06/21', '2012/06/21', '3年研究発表（益川 瑛大  辻 裕佑)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/06/28', '2012/06/28', '3年研究発表（田口 竜一朗  鈴木 冬惟)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/07/05', '2012/07/05', '3年研究発表（城田 翔平  佐々木 辰哉)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/07/12', '2012/07/12', '2年生歓迎会', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/07/19', '2012/07/19', '3年研究発表（北野 聖弥江川 真由)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/09/27', '2012/09/27', 'ゼミ夏期休暇報告 / 2年C言語レクチャー', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/10/04', '2012/10/04', '3年研究発表（石河 洋祐井坂 尚也', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/10/11', '2012/10/11', '3年研究発表（和田 太陽益川 瑛大', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/10/18', '2012/10/18', '3年研究発表（城田 翔平田口 竜一朗', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/10/25', '2012/10/25', 'C++/Javaレクチャー', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/11/01', '2012/11/01', '3年研究発表（全員） / 2年ITニュース・C++/Javaレクチャー', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/11/08', '2012/11/08', '3年研究発表（和田 太陽辻 裕佑） / 2年ITニュース・C++/Javaレクチャー', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/11/15', '2012/11/15', '3年研究発表（鈴木 冬惟佐々木 辰哉） / 2年ITニュース・Android/iPhoneレクチャー', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/11/22', '2012/11/22', '3年研究発表（北野 聖弥江川 真由） / 2年ITニュース・Android/iPhoneレクチャー', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/11/29', '2012/11/29', '3年研究発表（石河 洋祐井坂 尚也） / 2年ITニュース・Android/iPhoneレクチャー', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/11/29', '2012/11/29', '3年中間再発表（田口 竜一朗  鈴木 冬惟  石河 洋祐  井坂 尚也  佐々木 辰哉）', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/12/06', '2012/12/06', '3年研究発表（和田 太陽  益川 瑛大） / 2年ITニュース・Rubyレクチャー', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/12/13', '2012/12/13', '3年研究発表（辻 裕佑  田口 竜一朗） / 2年ITニュース・Rubyレクチャー', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2012/12/20', '2012/12/20', '卒研発表・飲み会', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2011/02/14', '2011/02/14', '修士論文審査会', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2011/06/14', '2011/06/14', '15:10～ ゼミ訪問', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2011/06/23', '2011/06/23', '15:10～ ゼミ配属者決定', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2011/09/18', '2011/09/18', 'オープンキャンパス', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2011/10/20', '2011/10/20', '中間発表', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2011/12/17', '2011/12/17', '卒研発表', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2010/12/18', '2010/12/18', '卒業論文本発表', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2010/10/21', '2010/10/21', '中間発表', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2010/07/08', '2010/07/08', '二年生歓迎会', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2010/06/10', '2010/06/10', '新2年生対象、研究室紹介', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2010/04/05', '2010/04/05', '新入生オリエンテーション（研究室紹介）', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2009/03/04', '2009/03/04', '動的画像処理実利用化ワークショップ DIA2010', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2009/03/05', '2009/03/05', '動的画像処理実利用化ワークショップ DIA2010', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2009/01/18', '2009/01/18', '卒業論文提出期間(20日まで)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2009/12/19', '2009/12/19', '卒業研究発表打ち上げ (飲み会)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2009/12/19', '2009/12/19', '卒業研究発表', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2009/11/12', '2009/11/12', '橋本ゼミとの合同ゼミ', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2009/10/31', '2009/10/31', 'ミニオープンキャンパス', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2009/10/29', '2009/10/29', '中間発表', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2009/09/20', '2009/09/20', 'オープンキャンパス', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2009/09/14', '2009/09/14', '大学院修士課程 中間発表', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2009/09/08', '2009/09/08', 'ゼミ合宿 in蓼科', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2009/09/09', '2009/09/09', 'ゼミ合宿 in蓼科', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2009/07/09', '2009/07/09', '新ゼミ生歓迎会', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2009/05/28', '2009/05/28', '2年生を対象にゼミ紹介', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2009/04/01', '2009/04/01', '新入生オリエンテーション開催', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2008/10/24', '2008/10/24', '2008年度 卒業研究中間発表開催日程決定', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2008/03/05', '2008/03/05', '動的画像処理実利用化ワークショップ2008を中京大学で開催', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2008/01/30', '2008/01/30', '2007年度　最終ゼミ（４年生お別れ会)', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2007/12/14', '2007/12/14', '2007年度　ラシキア研究室忘年会日程決定', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2007/10/08', '2007/10/08', '2007年度　卒業研究中間発表会日程決定', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2007/09/27', '2007/09/27', '2007年度  秋学期 発表日程', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2007/08/06', '2007/08/06', '2007年夏休み勉強会＜変更＞', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2007/04/12', '2007/04/12', '2007年度 春学期 発表日程', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2007/04/12', '2007/04/12', '2007年度 春学期 ゼミ時間割り', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime)),('2007/01/15', '2007/01/15', '2007年春勉強会', cast('2020-2-10 00:00:00' as datetime), cast('2020-2-10 00:00:00' as datetime));

INSERT INTO `homepage`.`jobs` VALUES (31,'三菱電機メカトロニクス','','2020-05-15 14:17:16','2020-05-15 14:17:16'),(32,'富士ゼロックス','','2020-05-15 14:17:28','2020-05-15 14:17:28'),(33,'東海ゴム','','2020-05-15 14:17:39','2020-05-15 14:17:39'),(34,'システムリサーチ','','2020-05-15 14:17:47','2020-05-15 14:17:47'),(35,'メイケイ','','2020-05-15 14:17:57','2020-05-15 14:17:57'),(36,'中電CTI','','2020-05-15 14:18:06','2020-05-15 14:18:06'),(37,'日立エンジニアリング','','2020-05-15 14:18:17','2020-05-15 14:18:17'),(38,'SCSK','','2020-05-15 14:18:28','2020-05-15 14:18:28'),(39,'NTTデータ東海','','2020-05-15 14:19:36','2020-05-15 14:19:36'),(40,'トヨタマックス','','2020-05-15 14:19:45','2020-05-15 14:19:45'),(41,'デンソーテクノ','','2020-05-15 14:19:54','2020-05-15 14:19:54'),(42,'アイシン','','2020-05-15 14:20:02','2020-05-15 14:20:02'),(43,'リンクバル','','2020-05-15 14:20:12','2020-05-15 14:20:12'),(44,'ヤフー','','2020-05-15 14:20:19','2020-05-15 14:20:19'),(45,'すんごいゲーム会社','','2020-05-15 14:20:27','2020-05-15 14:20:27');