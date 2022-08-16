/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50730
 Source Host           : localhost:3306
 Source Schema         : mall_go

 Target Server Type    : MySQL
 Target Server Version : 50730
 File Encoding         : 65001

 Date: 16/08/2022 15:39:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for mall_balance
-- ----------------------------
DROP TABLE IF EXISTS `mall_balance`;
CREATE TABLE `mall_balance`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `type` tinyint(1) NOT NULL COMMENT '类型\r\n1 用户\r\n2 商户',
  `available` bigint(22) NOT NULL DEFAULT 0 COMMENT '可用',
  `frozen` bigint(22) NOT NULL DEFAULT 0 COMMENT '冻结',
  `status` int(1) NOT NULL DEFAULT 1 COMMENT '状态 \r\n1 正常\r\n2 冻结',
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '余额表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
