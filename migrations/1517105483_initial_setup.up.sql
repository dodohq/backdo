CREATE TABLE IF NOT EXISTS admins (
  id SERIAL PRIMARY KEY,
  email VARCHAR(100) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS companies (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  contact_number VARCHAR(20) NOT NULL,
  deleted BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(100) NOT NULL,
  password VARCHAR(255) NOT NULL,
  company_id INT NOT NULL,
  FOREIGN KEY (company_id) REFERENCES companies(id),
  deleted BOOLEAN NOT NULL DEFAULT FALSE
);
CREATE UNIQUE INDEX index_unique_user ON users (email, company_id) WHERE NOT deleted;

CREATE TABLE IF NOT EXISTS deliveries (
  id SERIAL PRIMARY KEY,
  customer_name VARCHAR(255) NOT NULL,
  contact_number VARCHAR(20) NOT NULL,
  passcode VARCHAR(50),
  qr_code_url VARCHAR(100),
  company_id INT NOT NULL,
  FOREIGN KEY (company_id) REFERENCES companies(id),
  deleted BOOLEAN NOT NULL DEFAULT FALSE
);
CREATE UNIQUE INDEX index_unique_delivery ON deliveries (contact_number, company_id) WHERE NOT deleted;

CREATE TABLE IF NOT EXISTS drivers (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  phone_number VARCHAR(20) NOT NULL,
  company_id INT NOT NULL,
  FOREIGN KEY (company_id) REFERENCES companies(id),
  deleted BOOLEAN NOT NULL DEFAULT FALSE
);
CREATE UNIQUE INDEX index_unique_driver ON drivers (phone_number, company_id) WHERE NOT deleted;

CREATE TABLE IF NOT EXISTS robots (
  id SERIAL PRIMARY KEY,
  uuid VARCHAR(100) UNIQUE NOT NULL,
  company_id INT NOT NULL,
  FOREIGN KEY (company_id) REFERENCES companies(id),
  deleted BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE OR REPLACE FUNCTION cascade_company_delete()
RETURNS trigger AS
$BODY$
BEGIN
  UPDATE users
  SET deleted = NEW.deleted
  WHERE company_id = NEW.id;
  UPDATE deliveries
  SET deleted = NEW.deleted
  WHERE company_id = NEW.id;
  UPDATE drivers
  SET deleted = NEW.deleted
  WHERE company_id = NEW.id;
  UPDATE robots
  SET deleted = NEW.deleted
  WHERE company_id = NEW.id;
  RETURN NEW;
END;
$BODY$
LANGUAGE plpgsql;
CREATE TRIGGER trigger_delete_company
BEFORE UPDATE ON companies
FOR EACH ROW 
WHEN (OLD.deleted IS DISTINCT FROM NEW.deleted)
EXECUTE PROCEDURE cascade_company_delete();
