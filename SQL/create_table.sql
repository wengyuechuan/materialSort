-- 创建数据库
CREATE DATABASE IF NOT EXISTS materials_sort;
-- 使用数据库
USE materials_sort;
-- 用户表
CREATE TABLE IF NOT EXISTS users (
       id INT PRIMARY KEY AUTO_INCREMENT,
       username VARCHAR(255) UNIQUE NOT NULL,
       password VARCHAR(255) NOT NULL,
       email VARCHAR(255) UNIQUE NOT NULL,
       full_name VARCHAR(255)
)CHARSET=utf8mb4 ;

-- 资料表
CREATE TABLE IF NOT EXISTS documents (
       id INT PRIMARY KEY AUTO_INCREMENT,
       title VARCHAR(255) NOT NULL,
       description TEXT,
       file_path VARCHAR(255) NOT NULL,
       file_md5 VARCHAR(32) NOT NULL,
       upload_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       user_id INT,
       FOREIGN KEY (user_id) REFERENCES users(id)
);

-- 资料分类表
CREATE TABLE IF NOT EXISTS categories (
        id INT PRIMARY KEY AUTO_INCREMENT,
        category_name VARCHAR(255) UNIQUE NOT NULL
);

-- 资料-分类关系表
CREATE TABLE IF NOT EXISTS document_category (
       document_id INT,
       category_id INT,
       PRIMARY KEY (document_id, category_id),
       FOREIGN KEY (document_id) REFERENCES documents(id),
       FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- 管理员表
CREATE TABLE IF NOT EXISTS admins (
        id INT PRIMARY KEY AUTO_INCREMENT,
        username VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        full_name VARCHAR(255)
);
