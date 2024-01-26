-- 创建数据库
DROP DATABASE IF EXISTS materials_sort;
CREATE DATABASE materials_sort;
-- 使用数据库
USE materials_sort;
-- 用户表
CREATE TABLE users (
       user_id INT PRIMARY KEY AUTO_INCREMENT,
       username VARCHAR(255) UNIQUE NOT NULL,
       password VARCHAR(255) NOT NULL,
       email VARCHAR(255) UNIQUE NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 资料类型表
CREATE TABLE document_type (
      type_id INT PRIMARY KEY AUTO_INCREMENT,
      type_name VARCHAR(255) UNIQUE NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO document_type (type_name) VALUES ('text');
INSERT INTO document_type (type_name) VALUES ('link');
INSERT INTO document_type (type_name) VALUES ('file');

-- 资料表
CREATE TABLE document (
       document_id INT PRIMARY KEY AUTO_INCREMENT,
       type_id INT NOT NULL,                -- 资料类型 0为简单文本，1为链接，2为链接，3为文件 （后期可能考虑将链接分为纯文本和视频，视频将利用爬虫进行爬取，文本也将进行一定的分析）
       title VARCHAR(255) NOT NULL,         -- 资料标题，简单文本后期会利用api进行分析，或者训练一个自己的机器人来使用，前期这个部分就跟取用户输入的前20个字
       description TEXT,                    -- 资料描述 如果类型是一个简单文本，那么这个部分就是文本内容，如果是链接，那么这个部分就是链接的描述，需要爬取，如果是文件，哪儿买这个部分就是文件的描述，如果不填入的话默认是标题
       file_path VARCHAR(255),              -- 这个是在七牛云上的文件路径，如果是链接的会员，
       file_md5 VARCHAR(32),
       upload_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       update_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 更新时间
       user_id INT,
       FOREIGN KEY (user_id) REFERENCES users(user_id),
       FOREIGN KEY (type_id) REFERENCES document_type(type_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 资料分类表
CREATE TABLE category (
        category_id INT PRIMARY KEY AUTO_INCREMENT,
        category_name VARCHAR(255) UNIQUE NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 资料-分类关系表
-- 这里是由于这里是 many-many 的关系类型，所以要定义一个中间表来存储关系
CREATE TABLE document_category (
       document_id INT,
       category_id INT,
       PRIMARY KEY (document_id, category_id),
       FOREIGN KEY (document_id) REFERENCES document(document_id),
       FOREIGN KEY (category_id) REFERENCES category(category_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 管理员表
DROP TABLE IF EXISTS `admin`;
CREATE TABLE admin (
        id INT PRIMARY KEY AUTO_INCREMENT,
        username VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
