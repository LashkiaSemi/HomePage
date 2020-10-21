SET CHARSET UTF8;
SET foreign_key_checks=0;

-- define database
-- DROP DATABASE `homepage`;
CREATE DATABASE IF NOT EXISTS `homepage`;
USE `homepage`;

-- define table

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `password_digest` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL COMMENT 'owner/admin/member',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `student_id` varchar(255) DEFAULT NULL COMMENT '学籍番号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='メンバー';

DROP TABLE IF EXISTS `lectures`;
CREATE TABLE `lectures` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `file` varchar(255) DEFAULT NULL,
  `comments` varchar(255) DEFAULT NULL,
  `activation` tinyint(1) DEFAULT NULL COMMENT '公開/非公開',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_lectures_on_user_id` (`user_id`),
  CONSTRAINT `fk_rails_5a439a4e07` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='レクチャー';

DROP TABLE IF EXISTS `equipments`;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='備品';

DROP TABLE IF EXISTS `introductions`;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='メンバープロフィール';

DROP TABLE IF EXISTS `pages`;
CREATE TABLE `pages` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `contents` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `researches`;
CREATE TABLE `researches` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `author` varchar(255) DEFAULT NULL,
  `file` varchar(255) DEFAULT NULL,
  `comments` varchar(255) DEFAULT NULL,
  `activation` tinyint(1) DEFAULT NULL COMMENT '公開/非公開',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='卒業研究';

DROP TABLE IF EXISTS `jobs`;
CREATE TABLE `jobs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `company` varchar(255) DEFAULT NULL,
  `job` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='就職先企業';

DROP TABLE IF EXISTS `activities`;
CREATE TABLE `activities` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `show_date` varchar(255) DEFAULT NULL COMMENT '表示用',
  `date` datetime NOT NULL,
  `activity` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `annotation` varchar(255) NOT NULL DEFAULT '',
  `is_important` tinyint(1) NOT NULL DEFAULT '0',
  `is_notify` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='活動内容';

DROP TABLE IF EXISTS `societies`;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='学会';

DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='タグ';
