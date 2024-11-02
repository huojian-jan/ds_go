-- 创建gin_demo数据库
CREATE DATABASE `gin_blog` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 创建gin_demo数据库的用户
CREATE USER 'gin_user'@'%' IDENTIFIED WITH mysql_native_password BY '123456' PASSWORD EXPIRE NEVER;
-- 授权gin_demo数据库的用户
GRANT ALL PRIVILEGES ON gin_blog.* TO 'gin_user'@'%';
-- 刷新权限
FLUSH PRIVILEGES;