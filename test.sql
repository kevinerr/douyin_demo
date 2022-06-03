/*
SQLyog Ultimate v10.00 Beta1
MySQL - 8.0.16 : Database - douyin_db
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`douyin_db` /*!40100 DEFAULT CHARACTER SET utf8 */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `douyin_db`;

/*Table structure for table `comment` */

DROP TABLE IF EXISTS `comment`;

CREATE TABLE `comment` (
  `id` bigint(20) NOT NULL COMMENT '评论ID',
  `video_id` bigint(20) NOT NULL COMMENT '视频ID',
  `user_id` bigint(20) NOT NULL COMMENT '评论者ID',
  `content` text NOT NULL COMMENT '评论内容',
  `create_time` datetime NOT NULL COMMENT '评论时间',
  PRIMARY KEY (`id`),
  KEY `idx_comment_video_id` (`video_id`) USING BTREE COMMENT '查询视频评论列表'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `comment` */

insert  into `comment`(`id`,`video_id`,`user_id`,`content`,`create_time`) values (716492269342949376,716484862625710080,716483629923958784,'热巴真好看','2022-06-01 03:24:29'),(717195863331438592,716484822712713216,716483629923958784,'111','2022-06-03 02:00:19'),(717198824258404352,716484862625710080,716483691932549120,'真好看','2022-06-03 02:12:05');

/*Table structure for table `favorite` */

DROP TABLE IF EXISTS `favorite`;

CREATE TABLE `favorite` (
  `id` bigint(20) NOT NULL COMMENT '点赞记录ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户的ID',
  `video_id` bigint(20) NOT NULL COMMENT '点赞视频的ID',
  `create_time` datetime NOT NULL COMMENT '点赞时间',
  PRIMARY KEY (`id`),
  KEY `idx_favorite_user_id` (`user_id`) USING BTREE COMMENT '查询用户点赞的视频列表'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `favorite` */

insert  into `favorite`(`id`,`user_id`,`video_id`,`create_time`) values (716534500737155072,716483731774242816,716484822712713216,'2022-06-01 06:12:18');

/*Table structure for table `follow` */

DROP TABLE IF EXISTS `follow`;

CREATE TABLE `follow` (
  `id` bigint(20) NOT NULL COMMENT '关注记录ID',
  `follower_id` bigint(20) NOT NULL COMMENT '粉丝用户ID',
  `follow_id` bigint(20) NOT NULL COMMENT '被关注用户ID',
  `create_time` datetime NOT NULL COMMENT '关注记录时间',
  PRIMARY KEY (`id`),
  KEY `idx_follow_follower_id` (`follower_id`) USING BTREE COMMENT '查询关注列表',
  KEY `idx_follow_follow_id` (`follow_id`) USING BTREE COMMENT '查询粉丝列表'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `follow` */

insert  into `follow`(`id`,`follower_id`,`follow_id`,`create_time`) values (716494042132643840,716483629923958784,716483731774242816,'2022-06-01 03:31:31'),(717197661576364032,716483666192105472,716483731774242816,'2022-06-03 02:07:27'),(717198712866078720,716483691932549120,716483731774242816,'2022-06-03 02:11:38');

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` bigint(20) NOT NULL COMMENT '用户ID',
  `username` varchar(32) NOT NULL COMMENT '登录账号',
  `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '登录密码',
  `nickname` varchar(20) NOT NULL COMMENT '昵称',
  `follow_count` bigint(20) DEFAULT NULL COMMENT '关注总数',
  `follower_count` bigint(20) DEFAULT NULL COMMENT '粉丝总数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_user_username` (`username`) USING BTREE COMMENT '用户名唯一'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `user` */

insert  into `user`(`id`,`username`,`password`,`nickname`,`follow_count`,`follower_count`) values (716483629923958784,'张三','$2a$12$V3oRo/5g3ZaouOhBLewo2uCB7H5P56c93PKdz6mc2o5BcWW9xXFk6','张三',1,0),(716483666192105472,'李四','$2a$12$BLj9w/R1azr7y5xjltrP0.j1.y2.Ze66iHMrOaL/foSM7OhLScCRG','李四',1,0),(716483691932549120,'王五','$2a$12$8plkrRG6XMre5nLEs7ACt.9BSIDEycIR2XpbRCf7ugMILC0j9iaJm','王五',1,0),(716483731774242816,'迪丽热巴','$2a$12$Xsk/lp7TYqagBd16DmZSmuJIdHTFpsXF5coIZrSZhAW73tJGB0/Xm','迪丽热巴',0,3);

/*Table structure for table `video` */

DROP TABLE IF EXISTS `video`;

CREATE TABLE `video` (
  `id` bigint(20) NOT NULL COMMENT '视频ID',
  `author_id` bigint(20) NOT NULL COMMENT '作者的ID',
  `title` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '视频描述',
  `play_url` varchar(1024) NOT NULL COMMENT '视频播放地址',
  `cover_url` varchar(1024) NOT NULL COMMENT '视频封面地址',
  `favorite_count` bigint(20) DEFAULT NULL COMMENT '视频的点赞总数',
  `comment_count` bigint(20) DEFAULT NULL COMMENT '视频的评论总数',
  `create_time` datetime NOT NULL COMMENT '视频创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_video_author_id` (`author_id`) USING BTREE COMMENT '查询发布视频列表',
  KEY `idx_video_create_time` (`create_time`) USING BTREE COMMENT '查询视频流'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `video` */

insert  into `video`(`id`,`author_id`,`title`,`play_url`,`cover_url`,`favorite_count`,`comment_count`,`create_time`) values (716484822712713216,716483731774242816,'testTitle','http://10.115.112.67:8080/static/716483731774242816_1.MP4','https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg',10,0,'2022-06-01 02:54:53'),(716484862625710080,716483731774242816,'testTitle','http://10.115.112.67:8080/static/716483731774242816_2.MP4','https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg',8,0,'2022-06-01 02:55:03');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
