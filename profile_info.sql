/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : spider_profile

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2020-02-19 18:24:34
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for profile_info
-- ----------------------------
DROP TABLE IF EXISTS `profile_info`;
CREATE TABLE `profile_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `nickname` varchar(50) DEFAULT '' COMMENT '用户名称',
  `workcity` varchar(100) DEFAULT '' COMMENT '所在城市',
  `marriage` varchar(20) DEFAULT '0' COMMENT '婚否',
  `age` tinyint(4) DEFAULT '0' COMMENT '年龄',
  `gender` tinyint(1) DEFAULT NULL,
  `height` decimal(6,2) DEFAULT '0.00' COMMENT '身高',
  `salary` decimal(10,2) DEFAULT '0.00' COMMENT '薪酬',
  `education` varchar(50) DEFAULT '' COMMENT '教育水平',
  `introduce` varchar(50) DEFAULT '' COMMENT '备注',
  `avatar` varchar(255) DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT '2020-01-01 00:00:00' COMMENT '创建时间',
  `last_modtime` timestamp NULL DEFAULT NULL,
  `status` tinyint(1) DEFAULT '1' COMMENT '0：禁用；1：启用',
  PRIMARY KEY (`id`),
  KEY `idx_pro_name` (`nickname`),
  KEY `idx_create_time` (`create_time`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8 COMMENT='电影信息表';
