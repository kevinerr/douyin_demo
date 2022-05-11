/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80022
 Source Host           : localhost:3306
 Source Schema         : douyin_db

 Target Server Type    : MySQL
 Target Server Version : 80022
 File Encoding         : 65001

 Date: 11/05/2022 10:56:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` int NOT NULL COMMENT '评论ID',
  `video_id` int NOT NULL COMMENT '视频ID',
  `user_id` int NOT NULL COMMENT '评论者ID',
  `content` text NOT NULL COMMENT '评论内容',
  `create_time` datetime NOT NULL COMMENT '评论时间',
  PRIMARY KEY (`id`),
  INDEX `idx_comment_video_id` (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite` (
  `id` int NOT NULL COMMENT '点赞记录ID',
  `user_id` int NOT NULL COMMENT '用户的ID',
  `video_id` int NOT NULL COMMENT '点赞视频的ID',
  `create_time` datetime NOT NULL COMMENT '点赞时间',
  PRIMARY KEY (`id`),
  INDEX `idx_favorite_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for follow
-- ----------------------------
DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow` (
  `id` int NOT NULL COMMENT '关注记录ID',
  `follower_id` int NOT NULL COMMENT '粉丝用户ID',
  `follow_id` int NOT NULL COMMENT '被关注用户ID',
  `create_time` datetime NOT NULL COMMENT '关注记录时间',
  PRIMARY KEY (`id`),
  INDEX `idx_follow_follow_id` (`follow_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(32) NOT NULL COMMENT '登录账号',
  `password` varchar(32) NOT NULL COMMENT '登录密码',
  `nickname` varchar(20) NOT NULL COMMENT '昵称',
  `follow_count` int DEFAULT NULL COMMENT '关注总数',
  `follower_count` int DEFAULT NULL COMMENT '粉丝总数',
  `salt` varchar(255) NOT NULL COMMENT '加密盐值',
  PRIMARY KEY (`id`),
  UNIQUE `uni_user_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video` (
  `id` int NOT NULL COMMENT '视频ID',
  `author_id` int NOT NULL COMMENT '作者的ID',
  `description` text COMMENT '视频描述',
  `play_url` varchar(1024) NOT NULL COMMENT '视频播放地址',
  `cover_url` varchar(1024) NOT NULL COMMENT '视频封面地址',
  `favorite_count` int DEFAULT NULL COMMENT '视频的点赞总数',
  `comment_count` int DEFAULT NULL COMMENT '视频的评论总数',
  `create_time` datetime NOT NULL COMMENT '视频创建时间',
  PRIMARY KEY (`id`),
  INDEX `idx_video_author_id` (`author_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
