
# dtm官方文档
  部署需要的库表 https://dtm.pub/deploy/base.html#%E9%97%AE%E9%A2%98%E8%AF%8A%E6%96%AD
# 关于docker安装dtm，代码dtm都是桥接模式，
    直连模式修改 http://host.docker.internal  为http://localhost
# 以下数据库都在一数据库实例下面，如果需要改动需要调整代码
# 准备 RM 数据表
### 创建数据库：dtm_barrier
### 创建表：dtm_barrier
    create database if not exists dtm_barrier;
    drop table if exists dtm_barrier.barrier;
    create table if not exists dtm_barrier.barrier(
    id bigint(22) PRIMARY KEY AUTO_INCREMENT,
    trans_type varchar(45) default '',
    gid varchar(128) default '',
    branch_id varchar(128) default '',
    op varchar(45) default '',
    barrier_id varchar(45) default '',
    reason varchar(45) default '' comment 'the branch type who insert this record',
    create_time datetime DEFAULT now(),
    update_time datetime DEFAULT now(),
    key(create_time),
    key(update_time),
    UNIQUE key(gid, branch_id, op, barrier_id)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

# 转账微服务依赖库表
### 创建数据库：busi
### 创建表：user_account
     SET NAMES utf8mb4;
     SET FOREIGN_KEY_CHECKS = 0;
     
     -- ----------------------------
     -- Table structure for user_account
     -- ----------------------------
     DROP TABLE IF EXISTS `user_account`;
     CREATE TABLE `user_account` (
     `id` int NOT NULL AUTO_INCREMENT,
     `user_id` int NOT NULL,
     `balance` decimal(10,2) NOT NULL DEFAULT '0.00',
     `trading_balance` decimal(10,2) NOT NULL DEFAULT '0.00',
     `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
     `update_time` datetime DEFAULT CURRENT_TIMESTAMP,
     PRIMARY KEY (`id`),
     UNIQUE KEY `user_id` (`user_id`)
     ) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
     
     -- ----------------------------
     -- Records of user_account
     -- ----------------------------
     BEGIN;
     INSERT INTO `user_account` (`id`, `user_id`, `balance`, `trading_balance`, `create_time`, `update_time`) VALUES (1, 1000, 11051.00, 0.00, '2024-08-16 09:19:44', '2024-08-16 13:52:11');
     INSERT INTO `user_account` (`id`, `user_id`, `balance`, `trading_balance`, `create_time`, `update_time`) VALUES (2, 1001, 11051.00, 0.00, '2024-08-16 09:20:13', '2024-08-16 13:52:11');
     COMMIT;
     
     SET FOREIGN_KEY_CHECKS = 1;