DROP TRIGGER IF EXISTS trigger_delete_company ON companies;
DROP FUNCTION IF EXISTS cascade_company_delete;
DROP TABLE IF EXISTS robots;
DROP INDEX IF EXISTS index_unique_driver;
DROP TABLE IF EXISTS drivers;
DROP INDEX IF EXISTS index_unique_delivery;
DROP TABLE IF EXISTS deliveries;
DROP INDEX IF EXISTS index_unique_user;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS companies;
DROP TABLE IF EXISTS admins;
