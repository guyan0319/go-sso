-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2020-05-06 15:19:24
-- 服务器版本： 8.0.12
-- PHP 版本： 7.3.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";



/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `ssodb`
--

-- --------------------------------------------------------

--
-- 表的结构 `device`
--

CREATE TABLE `device` (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键',
  `uid` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户主键',
  `client` varchar(50) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '客户端',
  `model` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '设备型号',
  `ip` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT 'ip地址',
  `ext` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '扩展信息',
  `ctime` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '注册时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- 表的结构 `trace`
--

CREATE TABLE `trace` (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键',
  `uid` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户主键',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型(0:注册1::登录2:退出3:修改4:删除)',
  `ip` int(10) UNSIGNED NOT NULL COMMENT 'ip',
  `ext` varchar(1000) COLLATE utf8mb4_general_ci NOT NULL COMMENT '扩展字段',
  `ctime` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '注册时间'
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- 表的结构 `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键',
  `name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `email` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `mobile` varchar(20) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `passwd` varchar(40) COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `salt` char(4) COLLATE utf8mb4_general_ci NOT NULL COMMENT '盐值',
  `ext` text COLLATE utf8mb4_general_ci NOT NULL COMMENT '扩展字段',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态（0：未审核,1:通过 10删除）',
  `ctime` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- 转储表的索引
--

--
-- 表的索引 `device`
--
ALTER TABLE `device`
  ADD PRIMARY KEY (`id`),
  ADD KEY `uid` (`uid`);

--
-- 表的索引 `trace`
--
ALTER TABLE `trace`
  ADD PRIMARY KEY (`id`),
  ADD KEY `UT` (`uid`,`type`) USING BTREE;

--
-- 表的索引 `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `ctime` (`ctime`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `device`
--
ALTER TABLE `device`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键';

--
-- 使用表AUTO_INCREMENT `trace`
--
ALTER TABLE `trace`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键';
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
