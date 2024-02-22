/*
 Navicat Premium Data Transfer

 Source Server         : 127
 Source Server Type    : MySQL
 Source Server Version : 80033
 Source Host           : localhost:3306
 Source Schema         : fund

 Target Server Type    : MySQL
 Target Server Version : 80033
 File Encoding         : 65001

 Date: 21/02/2024 22:03:02
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for df_func_purchase
-- ----------------------------
DROP TABLE IF EXISTS `df_func_purchase`;
CREATE TABLE `df_func_purchase`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `purchase_amount` decimal(10, 4) NULL DEFAULT NULL COMMENT '购买金额',
  `holding_cost_price` decimal(10, 4) NULL DEFAULT NULL COMMENT '持仓成本价',
  `holding_quantity` decimal(10, 4) NULL DEFAULT NULL COMMENT '持有份额',
  `distribution` decimal(10, 4) NULL DEFAULT NULL COMMENT '分红',
  `add_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `last_update_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后更新时间',
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for df_fund_earnings
-- ----------------------------
DROP TABLE IF EXISTS `df_fund_earnings`;
CREATE TABLE `df_fund_earnings`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '基金代码',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '基金简称',
  `date` date NULL DEFAULT NULL,
  `nav_per_unit` decimal(10, 4) NULL DEFAULT NULL COMMENT '单位净值',
  `cumulative_nav` decimal(10, 4) NULL DEFAULT NULL COMMENT '累计净值',
  `daily_growth_rate` decimal(6, 4) NULL DEFAULT NULL COMMENT '日增长率',
  `past_1_week` decimal(8, 4) NULL DEFAULT NULL COMMENT '近1周增长率',
  `past_1_month` decimal(8, 4) NULL DEFAULT NULL COMMENT '近1个月增长率',
  `past_3_months` decimal(8, 4) NULL DEFAULT NULL COMMENT '近3个月增长率',
  `past_6_months` decimal(8, 4) NULL DEFAULT NULL COMMENT '近6个月增长率',
  `past_1_year` decimal(8, 4) NULL DEFAULT NULL COMMENT '近1年增长率',
  `past_2_years` decimal(8, 4) NULL DEFAULT NULL COMMENT '近2年增长率',
  `past_3_years` decimal(8, 4) NULL DEFAULT NULL COMMENT '近3年增长率',
  `this_year` decimal(8, 4) NULL DEFAULT NULL COMMENT '今年来增长率',
  `since_inception` decimal(8, 4) NULL DEFAULT NULL COMMENT '成立来增长率',
  `add_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `last_update_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后更新时间',
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3184 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for df_fund_earnings_rank
-- ----------------------------
DROP TABLE IF EXISTS `df_fund_earnings_rank`;
CREATE TABLE `df_fund_earnings_rank`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `date` date NULL DEFAULT NULL,
  `rank` int NULL DEFAULT NULL,
  `rank_precent` decimal(5, 2) NULL DEFAULT NULL,
  `total_rate` decimal(12, 2) NULL DEFAULT NULL,
  `kind_avg_rate` decimal(12, 2) NULL DEFAULT NULL,
  `unit_NV` decimal(8, 4) NULL DEFAULT NULL COMMENT '单位净值',
  `total_NV` decimal(8, 4) NULL DEFAULT NULL COMMENT '累计净值',
  `day_incre_val` decimal(6, 4) NULL DEFAULT NULL COMMENT '日增长值',
  `day_incre_rate` decimal(6, 2) NULL DEFAULT NULL COMMENT '日增长率',
  `add_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `last_update_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后更新时间',
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `code`(`code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4751291 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for df_fund_list
-- ----------------------------
DROP TABLE IF EXISTS `df_fund_list`;
CREATE TABLE `df_fund_list`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '代码',
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '名字',
  `pinyin` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '拼音',
  `date` datetime(0) NULL DEFAULT NULL COMMENT '日期',
  `abbr_pinyin` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '拼音简写',
  `type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '基金类型',
  `Inc_date` datetime(0) NULL DEFAULT NULL COMMENT '成立日',
  `buy` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '购买',
  `sell` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '赎回',
  `add_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `last_update_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后更新时间',
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18983 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for df_fund_star
-- ----------------------------
DROP TABLE IF EXISTS `df_fund_star`;
CREATE TABLE `df_fund_star`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `ZhaoShang_Securities_star` int NULL DEFAULT NULL COMMENT '招商证券-星',
  `ZhaoShang_Securities_trend` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '招商证券-趋势 up down',
  `Shanghai_Securities_star` int NULL DEFAULT NULL COMMENT '上海证券-星',
  `Shanghai_Securities_trend` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '上海证券-趋势 up down',
  `Jianan_Jinxin_star` int NULL DEFAULT NULL COMMENT '济安金信-星',
  `Jianan_Jinxin_trend` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '济安金信-趋势 up down',
  `ZhaoShang_Securities_date` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '招商证券-更新时间',
  `Shanghai_Securities_date` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '上海证券-更新时间',
  `Jianan_Jinxin_Securities_date` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '济安金信-更新时间',
  `add_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `last_update_time` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后更新时间',
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2491 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for trade_day
-- ----------------------------
DROP TABLE IF EXISTS `trade_day`;
CREATE TABLE `trade_day`  (
  `date` int NOT NULL,
  PRIMARY KEY (`date`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
