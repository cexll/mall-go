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

 Date: 24/08/2022 11:31:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for mall_balance_change_log
-- ----------------------------
DROP TABLE IF EXISTS `mall_balance_change_log`;
CREATE TABLE `mall_balance_change_log`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `balance_id` int(11) NOT NULL COMMENT '余额ID',
  `amount` int(22) NOT NULL COMMENT '变动金额',
  `before_amount` int(22) NOT NULL COMMENT '变动前余额',
  `after_amount` int(22) NOT NULL COMMENT '变动后余额',
  `type` tinyint(1) NOT NULL COMMENT '变动类型 \r\n1 增加\r\n2 减少',
  `type_amount` tinyint(1) NOT NULL COMMENT '余额类型\r\n1 可用余额\r\n2 冻结余额',
  `is_delete` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除 ',
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '余额变动记录表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
