-- Check if the database exists
SELECT SCHEMA_NAME
FROM INFORMATION_SCHEMA.SCHEMATA
WHERE SCHEMA_NAME = 'authen_system';

-- If the database doesn't exist, create it
CREATE DATABASE IF NOT EXISTS authen_system;

-- Grant privileges to the user
GRANT ALL PRIVILEGES ON authen_system.* TO 'admin'@'%' IDENTIFIED BY 'admin';
