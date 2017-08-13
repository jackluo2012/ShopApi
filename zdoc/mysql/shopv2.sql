CREATE DATABASE IF NOT EXISTS `shop`;
USE `shop`;


CREATE TABLE IF NOT EXISTS `admin` (
  `id` int(16) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `password` varchar(128) NOT NULL COMMENT '密码',
  `email` varchar(64) NOT NULL COMMENT '邮箱',
  `phone` int(16) NOT NULL COMMENT '手机号',
  `name` varchar(64) NOT NULL COMMENT '真实姓名',
  `status` int(8) DEFAULT '0' COMMENT '状态',
  `created` datetime NOT NULL DEFAULT current_timestamp,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


CREATE TABLE IF NOT EXISTS `user` (
  `id` int(16) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) UNIQUE DEFAULT NULL,
  `password` varchar(128) NOT NULL DEFAULT '',
  `status` int(8) DEFAULT NULL,
  `created` datetime NOT NULL DEFAULT current_timestamp,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


CREATE TABLE IF NOT EXISTS `userinfo` (
  `userid`   int(16),
  `nickname` varchar(100) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `phone` varchar(20) UNIQUE NOT NULL DEFAULT '',
  `sex` TINYINT(1) DEFAULT NULL COMMENT '0:男;1:女',
  PRIMARY KEY (`userid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


CREATE TABLE IF NOT EXISTS `address` (
  `id` varchar(256) NOT NULL DEFAULT '',
  `name` varchar(64) NOT NULL,
  `userid` int(16) NOT NULL,
  `phone` varchar(16) NOT NULL,
  `area` varchar(256) NOT NULL,
  `address` varchar(256) NOT NULL,
  `created` datetime NOT NULL DEFAULT current_timestamp,
  `updated` datetime DEFAULT NULL,
  `isdefault` int(8) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS `orderproduct` (
  `id` int(16) unsigned NOT NULL AUTO_INCREMENT,
  `productid` int(16) NOT NULL,
  `orderid` int(16) DEFAULT '0',
  `discount`  int(8)  NOT NULL ,
  `name`  varchar(64) NOT NULL COMMENT '商品名称',
  `size`  varchar(64) NOT NULL ,
  `count` int(64) NOT NULL ,
  `color` varchar(64)NOT NULL ,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS `orders` (
  `id` int(16) unsigned NOT NULL AUTO_INCREMENT,
  `userid` int(16) NOT NULL,
  `addressid` varchar(64) NOT NULL,
  `totalprice` double NOT NULL COMMENT '商品总价',
  `freight` double DEFAULT '0' COMMENT '运费',
  `remark` text COMMENT '备注',
  `status` int(8) NOT NULL,
  `payway` int NOT NULL ,
  `created` datetime NOT NULL DEFAULT current_timestamp,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)  
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


CREATE TABLE IF NOT EXISTS `cart` (
  `id` int(16) unsigned NOT NULL AUTO_INCREMENT,
  `productid` int(16) unsigned NOT NULL,
  `orderid` int(16) DEFAULT '0',
  `userid` int(16) NOT NULL,  
  `name` varchar(256) NOT NULL COMMENT '商品名称',
  `count` int(16) unsigned NOT NULL,  
  `size`    varchar(64) DEFAULT '',
  `color`   varchar(64) DEFAULT '',
  `price`   double NOT NULL,
  `status`  int(8) NOT NULL DEFAULT '233' COMMENT'是否在购物车  0: 在, 1: 不在',
  `paystatus`  int(8) NOT NULL DEFAULT '236' COMMENT'是否购买  0: 购买, 1: 不购买',
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


CREATE TABLE IF NOT EXISTS `category` (
  `id` int(16) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL DEFAULT '',
  `pid` int(16) NOT NULL DEFAULT '0',
  `status` int(16) NOT NULL,
  `created` datetime NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


CREATE TABLE IF NOT EXISTS `product` (
  `id` int(16) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(256) NOT NULL DEFAULT '',
  `totalsale` int(16) NOT NULL DEFAULT '0' COMMENT'销售量',
  `category` int(16) NOT NULL,
  `price` double NOT NULL,
  `detail` varchar(1024) DEFAULT '',
  `status` int(8) NOT NULL,
  `created` datetime NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
