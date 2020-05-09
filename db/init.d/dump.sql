mysqldump: [Warning] Using a password on the command line interface can be insecure.
-- MySQL dump 10.13  Distrib 5.7.28, for Linux (x86_64)
--
-- Host: localhost    Database: homepage
-- ------------------------------------------------------
-- Server version	5.7.28

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
-- Table structure for table `activities`
--

DROP TABLE IF EXISTS `activities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `activities` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `date` date NOT NULL,
  `activity` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=284 DEFAULT CHARSET=utf8 COMMENT='活動内容';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `activities`
--

LOCK TABLES `activities` WRITE;
/*!40000 ALTER TABLE `activities` DISABLE KEYS */;
INSERT INTO `activities` VALUES (145,'2019-05-09','ゼミ説明会練習','2020-02-10 00:00:00','2020-02-10 00:00:00'),(146,'2019-05-16','外部者参加可 SQLレクチャー(池上,今宮)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(147,'2019-05-23','外部者参加可 Pythonミニ入門(大橋)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(148,'2019-05-30','外部者参加可 深層学習 Part1:基礎(ラシキア)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(149,'2019-06-06','ゼミ紹介最終練習','2020-02-10 00:00:00','2020-02-10 00:00:00'),(150,'2019-06-06','外部者参加可 社会人基礎力シリーズ~難しい質問への対応~(ラシキア)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(151,'2019-06-13','ゼミ見学','2020-02-10 00:00:00','2020-02-10 00:00:00'),(152,'2019-06-20','深層学習 Part2:既存の学習済みモデル(小林)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(153,'2019-06-20','外部者参加可 社会人基礎力シリーズ~答えたくない質問への対応~(ラシキア)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(154,'2019-06-27','深層学習 Part3:ファインチューニング(池上)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(155,'2019-07-04','卒論テーマについて','2020-02-10 00:00:00','2020-02-10 00:00:00'),(156,'2019-07-04','外部者参加可 プライベートコミュニケーションスキルシリーズ~相手が返事をしない・しなくなった場合の対応~(ラシキア)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(157,'2019-07-11','Thread入門','2020-02-10 00:00:00','2020-02-10 00:00:00'),(160,'2018-06-07','ゼミ見学会','2020-02-10 00:00:00','2020-02-10 00:00:00'),(161,'2017-05-31','質問対策講座','2020-02-10 00:00:00','2020-02-10 00:00:00'),(162,'2017-06-08','ゼミ見学会１','2020-02-10 00:00:00','2020-02-10 00:00:00'),(163,'2017-06-15','ゼミ見学会２','2020-02-10 00:00:00','2020-02-10 00:00:00'),(164,'2015-04-09','春学期ゼミ開講','2020-02-10 00:00:00','2020-02-10 00:00:00'),(165,'2015-04-09','春期休暇報告','2020-02-10 00:00:00','2020-02-10 00:00:00'),(166,'2015-06-04','ゼミ見学会１','2020-02-10 00:00:00','2020-02-10 00:00:00'),(167,'2015-06-11','ゼミ見学会２','2020-02-10 00:00:00','2020-02-10 00:00:00'),(168,'2015-07-16','新ゼミ生歓迎会,飲み会予定','2020-02-10 00:00:00','2020-02-10 00:00:00'),(169,'2014-03-02','サーバー勉強会','2020-02-10 00:00:00','2020-02-10 00:00:00'),(170,'2014-04-17','ゼミ春期休暇報告','2020-02-10 00:00:00','2020-02-10 00:00:00'),(171,'2014-04-24','3年4年研究発表、PHPレクチャー第1回(関谷、山脇)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(172,'2014-05-01','3年4年研究発表、PHPレクチャー第2回(関谷、山脇)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(173,'2014-05-08','3年4年研究発表、PHPレクチャー第3回(関谷、山脇)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(174,'2014-05-15','3年4年研究発表、スレッドレクチャー','2020-02-10 00:00:00','2020-02-10 00:00:00'),(175,'2014-05-22','3年4年研究発表、スレッドレクチャー','2020-02-10 00:00:00','2020-02-10 00:00:00'),(176,'2014-05-28','3年4年研究発表、ゼミ見学発表練習','2020-02-10 00:00:00','2020-02-10 00:00:00'),(177,'2014-06-12','ゼミ見学','2020-02-10 00:00:00','2020-02-10 00:00:00'),(178,'2014-06-19','3年4年研究発表、スレッドレクチャー','2020-02-10 00:00:00','2020-02-10 00:00:00'),(179,'2014-06-26','3年4年研究発表、スレッドレクチャー','2020-02-10 00:00:00','2020-02-10 00:00:00'),(180,'2014-07-03','3年4年研究発表、スレッドレクチャー','2020-02-10 00:00:00','2020-02-10 00:00:00'),(181,'2014-07-10','3年4年研究発表、スレッドレクチャー','2020-02-10 00:00:00','2020-02-10 00:00:00'),(182,'2014-07-17','3年4年研究発表、スレッドレクチャー','2020-02-10 00:00:00','2020-02-10 00:00:00'),(183,'2014-07-31','飲み会','2020-02-10 00:00:00','2020-02-10 00:00:00'),(184,'2014-09-25','夏季休暇報告','2020-02-10 00:00:00','2020-02-10 00:00:00'),(185,'2014-10-02','2年: ITニュース、C言語レクチャー(長瀬、岩下), 3年: 研究発表、スレッドレクチャー , 4年: 中間発表練習','2020-02-10 00:00:00','2020-02-10 00:00:00'),(186,'2014-10-09','2年: ITニュース、C言語レクチャー(長瀬、岩下), 3年: 研究発表、スレッドレクチャー , 4年: 中間発表練習','2020-02-10 00:00:00','2020-02-10 00:00:00'),(187,'2014-10-16','2年: ITニュース、C++/javaレクチャー, 3年: 研究発表、スレッドレクチャー , 4年: 中間発表練習','2020-02-10 00:00:00','2020-02-10 00:00:00'),(188,'2014-10-23','中間発表','2020-02-10 00:00:00','2020-02-10 00:00:00'),(189,'2014-10-30','2年: ITニュース、C++/javaレクチャー, 3年: 研究発表、スレッドレクチャー , 4年: 卒研練習','2020-02-10 00:00:00','2020-02-10 00:00:00'),(190,'2014-11-06','2年: ITニュース、C++/javaレクチャー, 3年: 研究発表、スレッドレクチャー , 4年: 卒研練習','2020-02-10 00:00:00','2020-02-10 00:00:00'),(191,'2014-11-13','2年: ITニュース、iPhone(竹内、宮宅)/Android(佐野、加藤)レクチャー, 3年: 研究発表、スレッドレクチャー , 4年: 卒研練習','2020-02-10 00:00:00','2020-02-10 00:00:00'),(192,'2014-11-20','2年: ITニュース、iPhone(竹内、宮宅)/Android(佐野、加藤)レクチャー, 3年: 研究発表、スレッドレクチャー , 4年: 卒研練習','2020-02-10 00:00:00','2020-02-10 00:00:00'),(193,'2014-11-27','2年: ITニュース、iPhone(竹内、宮宅)/Android(佐野、加藤)レクチャー, 3年: 研究発表、スレッドレクチャー , 4年: 卒研練習','2020-02-10 00:00:00','2020-02-10 00:00:00'),(194,'2014-12-04','2年: ITニュース、データベースレクチャー(3年宮、伊藤), 3年: 研究発表、スレッドレクチャー , 4年: 卒研練習','2020-02-10 00:00:00','2020-02-10 00:00:00'),(195,'2014-12-11','2年: ITニュース、 3年: 研究発表、スレッドレクチャー , 4年: 卒研練習','2020-02-10 00:00:00','2020-02-10 00:00:00'),(196,'2014-12-18','飲み会','2020-02-10 00:00:00','2020-02-10 00:00:00'),(197,'2013-04-11','ゼミ春期休暇報告','2020-02-10 00:00:00','2020-02-10 00:00:00'),(198,'2013-05-30','ゼミ紹介 時間:15:10～18:20 場所:AI棟5階','2020-02-10 00:00:00','2020-02-10 00:00:00'),(199,'2013-06-03','ゼミ紹介 時間:15:10～18:20 場所:ガーデン4階','2020-02-10 00:00:00','2020-02-10 00:00:00'),(200,'2013-06-06','ゼミ紹介 時間:15:10～18:20 場所:AI棟5階','2020-02-10 00:00:00','2020-02-10 00:00:00'),(201,'2013-09-26','ゼミ夏期休暇報告','2020-02-10 00:00:00','2020-02-10 00:00:00'),(202,'2013-10-03','2年ITニュース・Cレクチャー(津曲,山脇)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(203,'2013-10-10','2年ITニュース・Cレクチャー(津曲,山脇)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(204,'2013-10-17','2年ITニュース・Cレクチャー(津曲,山脇)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(205,'2013-10-24','C++/Javaレクチャー(柏木,高田)/(関谷,安岡)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(206,'2013-10-31','4年卒研中間発表・2年ITニュース・C++/Javaレクチャー(柏木,高田)/(関谷,安岡)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(207,'2013-11-07','2年ITニュース・C++/Javaレクチャー(柏木,高田)/(関谷,安岡)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(208,'2013-11-14','2年ITニュース・iPhone/Androidレクチャー(青木,前島)/(小林,高田)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(209,'2013-11-21','2年ITニュース・iPhone/Androidレクチャー(青木,前島)/(小林,高田)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(210,'2013-11-28','iPhone/Androidレクチャー(青木,前島)/(小林,高田)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(211,'2013-12-05','JavaScriptレクチャー(池田,前島)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(212,'2013-12-12','JavaScriptレクチャー(池田,前島)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(213,'2013-12-19','卒研発表・飲み会・JavaScriptレクチャー(池田,前島)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(214,'2014-01-09','C＃レクチャー(柏木,李)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(215,'2014-01-16','C＃レクチャー(柏木,李)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(216,'2012-04-05','ゼミ春期休暇報告','2020-02-10 00:00:00','2020-02-10 00:00:00'),(217,'2012-04-26','3年研究発表（和田 太陽  益川 瑛大）','2020-02-10 00:00:00','2020-02-10 00:00:00'),(218,'2012-05-10','3年研究発表（田口 竜一朗  鈴木 冬惟）','2020-02-10 00:00:00','2020-02-10 00:00:00'),(219,'2012-05-17','3年研究発表（城田 翔平  佐々木 辰哉）','2020-02-10 00:00:00','2020-02-10 00:00:00'),(220,'2012-05-24','3年研究発表（北野 聖弥  江川 真由）','2020-02-10 00:00:00','2020-02-10 00:00:00'),(221,'2012-05-31','3年研究発表（石河 洋祐  井坂 尚也）','2020-02-10 00:00:00','2020-02-10 00:00:00'),(222,'2012-06-07','ゼミ紹介 時間:15:10～18:20 場所:15号館5階多目的演習室','2020-02-10 00:00:00','2020-02-10 00:00:00'),(223,'2012-06-11','ゼミ紹介 時間:15:10～18:20 場所:11号館4階ゼミ室','2020-02-10 00:00:00','2020-02-10 00:00:00'),(224,'2012-06-14','3年研究発表（和田 太陽)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(225,'2012-06-14','ゼミ紹介 時間:15:10～18:20 場所:15号館5階多目的演習室','2020-02-10 00:00:00','2020-02-10 00:00:00'),(226,'2012-06-21','3年研究発表（益川 瑛大  辻 裕佑)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(227,'2012-06-28','3年研究発表（田口 竜一朗  鈴木 冬惟)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(228,'2012-07-05','3年研究発表（城田 翔平  佐々木 辰哉)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(229,'2012-07-12','2年生歓迎会','2020-02-10 00:00:00','2020-02-10 00:00:00'),(230,'2012-07-19','3年研究発表（北野 聖弥江川 真由)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(231,'2012-09-27','ゼミ夏期休暇報告 / 2年C言語レクチャー','2020-02-10 00:00:00','2020-02-10 00:00:00'),(232,'2012-10-04','3年研究発表（石河 洋祐井坂 尚也','2020-02-10 00:00:00','2020-02-10 00:00:00'),(233,'2012-10-11','3年研究発表（和田 太陽益川 瑛大','2020-02-10 00:00:00','2020-02-10 00:00:00'),(234,'2012-10-18','3年研究発表（城田 翔平田口 竜一朗','2020-02-10 00:00:00','2020-02-10 00:00:00'),(235,'2012-10-25','C++/Javaレクチャー','2020-02-10 00:00:00','2020-02-10 00:00:00'),(236,'2012-11-01','3年研究発表（全員） / 2年ITニュース・C++/Javaレクチャー','2020-02-10 00:00:00','2020-02-10 00:00:00'),(237,'2012-11-08','3年研究発表（和田 太陽辻 裕佑） / 2年ITニュース・C++/Javaレクチャー','2020-02-10 00:00:00','2020-02-10 00:00:00'),(238,'2012-11-15','3年研究発表（鈴木 冬惟佐々木 辰哉） / 2年ITニュース・Android/iPhoneレクチャー','2020-02-10 00:00:00','2020-02-10 00:00:00'),(239,'2012-11-22','3年研究発表（北野 聖弥江川 真由） / 2年ITニュース・Android/iPhoneレクチャー','2020-02-10 00:00:00','2020-02-10 00:00:00'),(240,'2012-11-29','3年研究発表（石河 洋祐井坂 尚也） / 2年ITニュース・Android/iPhoneレクチャー','2020-02-10 00:00:00','2020-02-10 00:00:00'),(241,'2012-11-29','3年中間再発表（田口 竜一朗  鈴木 冬惟  石河 洋祐  井坂 尚也  佐々木 辰哉）','2020-02-10 00:00:00','2020-02-10 00:00:00'),(242,'2012-12-06','3年研究発表（和田 太陽  益川 瑛大） / 2年ITニュース・Rubyレクチャー','2020-02-10 00:00:00','2020-02-10 00:00:00'),(243,'2012-12-13','3年研究発表（辻 裕佑  田口 竜一朗） / 2年ITニュース・Rubyレクチャー','2020-02-10 00:00:00','2020-02-10 00:00:00'),(244,'2012-12-20','卒研発表・飲み会','2020-02-10 00:00:00','2020-02-10 00:00:00'),(245,'2011-02-14','修士論文審査会','2020-02-10 00:00:00','2020-02-10 00:00:00'),(246,'2011-06-14','15:10～ ゼミ訪問','2020-02-10 00:00:00','2020-02-10 00:00:00'),(247,'2011-06-23','15:10～ ゼミ配属者決定','2020-02-10 00:00:00','2020-02-10 00:00:00'),(248,'2011-09-18','オープンキャンパス','2020-02-10 00:00:00','2020-02-10 00:00:00'),(249,'2011-10-20','中間発表','2020-02-10 00:00:00','2020-02-10 00:00:00'),(250,'2011-12-17','卒研発表','2020-02-10 00:00:00','2020-02-10 00:00:00'),(251,'2010-12-18','卒業論文本発表','2020-02-10 00:00:00','2020-02-10 00:00:00'),(252,'2010-10-21','中間発表','2020-02-10 00:00:00','2020-02-10 00:00:00'),(253,'2010-07-08','二年生歓迎会','2020-02-10 00:00:00','2020-02-10 00:00:00'),(254,'2010-06-10','新2年生対象、研究室紹介','2020-02-10 00:00:00','2020-02-10 00:00:00'),(255,'2010-04-05','新入生オリエンテーション（研究室紹介）','2020-02-10 00:00:00','2020-02-10 00:00:00'),(256,'2009-03-04','動的画像処理実利用化ワークショップ DIA2010','2020-02-10 00:00:00','2020-02-10 00:00:00'),(257,'2009-03-05','動的画像処理実利用化ワークショップ DIA2010','2020-02-10 00:00:00','2020-02-10 00:00:00'),(258,'2009-01-18','卒業論文提出期間(20日まで)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(259,'2009-12-19','卒業研究発表打ち上げ (飲み会)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(260,'2009-12-19','卒業研究発表','2020-02-10 00:00:00','2020-02-10 00:00:00'),(261,'2009-11-12','橋本ゼミとの合同ゼミ','2020-02-10 00:00:00','2020-02-10 00:00:00'),(262,'2009-10-31','ミニオープンキャンパス','2020-02-10 00:00:00','2020-02-10 00:00:00'),(263,'2009-10-29','中間発表','2020-02-10 00:00:00','2020-02-10 00:00:00'),(264,'2009-09-20','オープンキャンパス','2020-02-10 00:00:00','2020-02-10 00:00:00'),(265,'2009-09-14','大学院修士課程 中間発表','2020-02-10 00:00:00','2020-02-10 00:00:00'),(266,'2009-09-08','ゼミ合宿 in蓼科','2020-02-10 00:00:00','2020-02-10 00:00:00'),(267,'2009-09-09','ゼミ合宿 in蓼科','2020-02-10 00:00:00','2020-02-10 00:00:00'),(268,'2009-07-09','新ゼミ生歓迎会','2020-02-10 00:00:00','2020-02-10 00:00:00'),(269,'2009-05-28','2年生を対象にゼミ紹介','2020-02-10 00:00:00','2020-02-10 00:00:00'),(270,'2009-04-01','新入生オリエンテーション開催','2020-02-10 00:00:00','2020-02-10 00:00:00'),(271,'2008-10-24','2008年度 卒業研究中間発表開催日程決定','2020-02-10 00:00:00','2020-02-10 00:00:00'),(272,'2008-03-05','動的画像処理実利用化ワークショップ2008を中京大学で開催','2020-02-10 00:00:00','2020-02-10 00:00:00'),(273,'2008-01-30','2007年度　最終ゼミ（４年生お別れ会)','2020-02-10 00:00:00','2020-02-10 00:00:00'),(274,'2007-12-14','2007年度　ラシキア研究室忘年会日程決定','2020-02-10 00:00:00','2020-02-10 00:00:00'),(275,'2007-10-08','2007年度　卒業研究中間発表会日程決定','2020-02-10 00:00:00','2020-02-10 00:00:00'),(276,'2007-09-27','2007年度  秋学期 発表日程','2020-02-10 00:00:00','2020-02-10 00:00:00'),(277,'2007-08-06','2007年夏休み勉強会＜変更＞','2020-02-10 00:00:00','2020-02-10 00:00:00'),(278,'2007-04-12','2007年度 春学期 発表日程','2020-02-10 00:00:00','2020-02-10 00:00:00'),(279,'2007-04-12','2007年度 春学期 ゼミ時間割り','2020-02-10 00:00:00','2020-02-10 00:00:00'),(280,'2007-01-15','2007年春勉強会','2020-02-10 00:00:00','2020-02-10 00:00:00'),(281,'2020-02-08','Thread同期','2020-02-13 01:27:37','2020-02-13 01:27:37'),(282,'2020-02-13','ThreadPool','2020-02-13 01:43:25','2020-02-19 07:08:11'),(283,'2020-02-07','Thread非同期','2020-02-13 02:13:53','2020-02-13 02:13:53');
/*!40000 ALTER TABLE `activities` ENABLE KEYS */;
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
  `num` int(11) DEFAULT NULL COMMENT '所持数',
  `note` varchar(255) DEFAULT NULL COMMENT 'コメントみたいな',
  `tag_id` bigint(20) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_equipments_on_tag_id` (`tag_id`)
) ENGINE=InnoDB AUTO_INCREMENT=60 DEFAULT CHARSET=utf8 COMMENT='備品';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `equipments`
--

LOCK TABLES `equipments` WRITE;
/*!40000 ALTER TABLE `equipments` DISABLE KEYS */;
INSERT INTO `equipments` VALUES (59,'golangでサーバを作る',1,'web技術',9,'2020-01-22 11:38:42','2020-02-18 01:30:50');
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
  `department` varchar(255) DEFAULT NULL COMMENT '所属学科',
  `grade` int(11) DEFAULT NULL COMMENT '0が卒業生っぽい',
  `comments` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_introductions_on_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=111 DEFAULT CHARSET=utf8 COMMENT='メンバー紹介';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `introductions`
--

LOCK TABLES `introductions` WRITE;
/*!40000 ALTER TABLE `introductions` DISABLE KEYS */;
INSERT INTO `introductions` VALUES (103,32,'情報',3,'こめんお！','2020-01-20 05:45:39','2020-02-18 01:19:49'),(105,34,'工学部情報工学科',4,'こんにちは。','2020-01-20 12:40:02','2020-02-18 01:46:56'),(107,36,'情報工学',2,'','2020-02-02 10:28:37','2020-02-03 08:23:42'),(110,39,'',0,'','2020-02-16 08:16:09','2020-02-16 08:16:40');
/*!40000 ALTER TABLE `introductions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jobs`
--

DROP TABLE IF EXISTS `jobs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `jobs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `company` varchar(255) DEFAULT NULL,
  `job` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8 COMMENT='就職先企業';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jobs`
--

LOCK TABLES `jobs` WRITE;
/*!40000 ALTER TABLE `jobs` DISABLE KEYS */;
INSERT INTO `jobs` VALUES (33,'cyber agent','server side engineer','2020-02-02 05:06:03','2020-02-18 00:55:34'),(34,'yahoo','エンジニア','2020-02-16 08:50:40','2020-02-16 08:50:51'),(35,'capcon','エンジニア','2020-02-18 00:55:51','2020-02-18 00:55:51');
/*!40000 ALTER TABLE `jobs` ENABLE KEYS */;
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
  `activation` tinyint(1) DEFAULT NULL COMMENT '謎1',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_lectures_on_user_id` (`user_id`),
  CONSTRAINT `fk_rails_5a439a4e07` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8 COMMENT='レクチャーの資料';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lectures`
--

LOCK TABLES `lectures` WRITE;
/*!40000 ALTER TABLE `lectures` DISABLE KEYS */;
INSERT INTO `lectures` VALUES (25,34,'adminサンプル','IMG_20191018_143029.jpg','コメント',1,'2020-02-17 01:50:19','2020-02-17 07:05:29'),(26,39,'lecturetiiii','IMG_20190827_213826 (1)のコピー.jpg','',1,'2020-02-17 07:04:31','2020-02-17 07:05:36'),(41,39,'あくび。たいとる','images.png','あくびちゃんのそつけん',1,'2020-02-20 05:17:18','2020-02-20 05:17:38');
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='なにこれ謎';
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
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8 COMMENT='研究';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `researches`
--

LOCK TABLES `researches` WRITE;
/*!40000 ALTER TABLE `researches` DISABLE KEYS */;
INSERT INTO `researches` VALUES (34,'sample title','','images.png','sample comment',NULL,'2020-02-14 08:30:01','2020-02-14 08:30:01'),(35,'たいとる！','あくび','OUT_6527643220685032268.pdf','',1,'2020-02-14 08:52:24','2020-02-18 01:21:44'),(37,'卒研新規','アナベル','image_2019_s.jpg','あいがも！！！！',NULL,'2020-02-16 06:35:37','2020-02-16 06:52:05');
/*!40000 ALTER TABLE `researches` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `societies`
--

DROP TABLE IF EXISTS `societies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `societies` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `author` varchar(255) DEFAULT NULL,
  `society` varchar(255) DEFAULT NULL,
  `award` varchar(255) DEFAULT NULL,
  `date` date DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COMMENT='学会';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `societies`
--

LOCK TABLES `societies` WRITE;
/*!40000 ALTER TABLE `societies` DISABLE KEYS */;
INSERT INTO `societies` VALUES (5,'ひとからたのしい','ぼく','sample','なし','2020-01-21','2020-02-14 05:40:41','2020-02-14 05:40:41'),(6,'けんきゅう','','sample','','2020-01-21','2020-02-14 05:42:13','2020-02-14 05:42:13'),(8,'けんきゅう001','アナベル','CTF','学生賞','2020-02-06','2020-02-16 06:35:04','2020-02-16 06:35:04');
/*!40000 ALTER TABLE `societies` ENABLE KEYS */;
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
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COMMENT='備品かなんか？';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tags`
--

LOCK TABLES `tags` WRITE;
/*!40000 ALTER TABLE `tags` DISABLE KEYS */;
INSERT INTO `tags` VALUES (6,'その他','2020-01-20 12:14:30','2020-01-20 12:14:30'),(10,'ソフトウェア','2020-02-17 05:58:35','2020-02-17 05:58:35');
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
  `role` varchar(255) DEFAULT NULL COMMENT 'owner/admin/member',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `student_id` varchar(255) DEFAULT NULL COMMENT '学籍番号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8 COMMENT='メンバー';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (32,'yawn','$2a$10$anqKnlC9NSlpnrpSj5Ij1OjVztcfVd10nlXBWYfxBb7hYNytf1GBK','owner','2020-01-20 05:45:39','2020-02-18 01:19:49','t317013'),(34,'18番','$2a$10$jiq1.YJTfa4tfakQX2qDyeS9EwDDF0TXr9qJpvuzr9rjhl0KX2qh6','owner','2020-01-20 12:40:02','2020-02-18 01:48:20','t317018'),(36,'20番','$2a$10$h2/VTHGPNIRXQP8Qiwpky.JSdXSCt4yZWoYhfNhojrwbxakWpeu9S','member','2020-02-02 10:28:37','2020-02-03 08:23:42','t317020'),(39,'あくびちゃん','$2a$10$OUwwulzBL6JVLvT15soCf.OWaykfnuWQDbHO4kjsDLjMt/akGJniq','member','2020-02-16 08:16:09','2020-02-16 08:26:10','t317014');
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

-- Dump completed on 2020-05-09 15:03:35
