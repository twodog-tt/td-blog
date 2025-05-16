## 个人博客系统后端

### 概述：
#### 本次作业要求你使用 Go 语言结合 Gin 框架和 GORM 库开发一个个人博客系统的后端，实现博客文章的基本管理功能，包括文章的创建、读取、更新和删除（CRUD）操作，同时支持用户认证和简单的评论功能。

### 1.项目初始化：
    - 使用go mod init 初始化项目依赖管理
    - Gin + GORM
    - 使用go-zero完成部分功能

### 2.数据库设计与模型定义
    - 按照产品功能定义表结构，实现部分功能
    CREATE TABLE `users` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(50) NOT NULL COMMENT '用户名',
    `password` varchar(255) NOT NULL COMMENT '密码(加密存储)',
    `email` varchar(100) NOT NULL COMMENT '邮箱',
    `avatar` varchar(255) DEFAULT NULL COMMENT '头像URL',
    `bio` varchar(255) DEFAULT NULL COMMENT '个人简介',
    `status` tinyint(1) DEFAULT '1' COMMENT '状态(0-禁用,1-正常)',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`),
    UNIQUE KEY `idx_email` (`email`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

    CREATE TABLE `articles` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_id` int(11) NOT NULL COMMENT '作者ID',
    `title` varchar(255) NOT NULL COMMENT '文章标题',
    `slug` varchar(255) NOT NULL COMMENT '文章URL别名',
    `content` longtext NOT NULL COMMENT '文章内容',
    `summary` varchar(500) DEFAULT NULL COMMENT '文章摘要',
    `cover_image` varchar(255) DEFAULT NULL COMMENT '封面图URL',
    `view_count` int(11) DEFAULT '0' COMMENT '浏览次数',
    `status` tinyint(1) DEFAULT '1' COMMENT '状态(0-草稿,1-已发布,2-已删除)',
    `published_at` datetime DEFAULT NULL COMMENT '发布时间',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_slug` (`slug`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_status` (`status`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文章表';

    CREATE TABLE `categories` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL COMMENT '分类名称',
    `slug` varchar(50) NOT NULL COMMENT '分类URL别名',
    `description` varchar(255) DEFAULT NULL COMMENT '分类描述',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_slug` (`slug`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='分类表';

    CREATE TABLE `article_category` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `article_id` int(11) NOT NULL COMMENT '文章ID',
    `category_id` int(11) NOT NULL COMMENT '分类ID',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_article_category` (`article_id`,`category_id`),
    KEY `idx_category_id` (`category_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文章分类关联表';

    CREATE TABLE `comments` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `article_id` int(11) NOT NULL COMMENT '文章ID',
    `user_id` int(11) DEFAULT NULL COMMENT '用户ID(未登录用户为NULL)',
    `parent_id` int(11) DEFAULT NULL COMMENT '父评论ID',
    `content` text NOT NULL COMMENT '评论内容',
    `username` varchar(50) DEFAULT NULL COMMENT '评论者名称(未登录用户)',
    `email` varchar(100) DEFAULT NULL COMMENT '评论者邮箱(未登录用户)',
    `status` tinyint(1) DEFAULT '1' COMMENT '状态(0-待审核,1-已发布,2-已删除)',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_article_id` (`article_id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_parent_id` (`parent_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='评论表';

#### 2.1 关键关系说明：
    用户 → 文章 (1:N)
    一个用户可以写多篇文章 (users.id → articles.user_id)
    用户 → 评论 (1:N)
    一个用户可以发表多条评论 (users.id → comments.user_id)
    文章 → 评论 (1:N)
    一篇文章可以有多条评论 (articles.id → comments.article_id)
    文章 ↔ 分类 (M:N)
    通过关联表 article_category 实现多对多关系
    (articles.id → article_category.article_id，categories.id → article_category.category_id)
    评论自关联
    评论可以回复其他评论 (comments.id → comments.parent_id)

#### 2.2 命令
    goctl model mysql datasource -url="root:123456789@tcp(127.0.0.1:3306)/td-blog" -table="users,articles,categories,article_category,comments" -dir="/Users/wangxintong/GolandProjects/td-blog/model" -cache=false   # 可选：启用 Redis 缓存
### 3.鉴权
    - 实现用户登陆与注册
    - JWT的使用

### 4.文章管理
    - CRUD

### 5.评论
    - CRUD

### 6.异常处理与日志记录
    - trace的使用

### 提交要求：
    - 代码完整
    - 展示测试用例与结果