CREATE USER :user WITH PASSWORD ':passwrd';

GRANT 
SELECT 
  ON ALL TABLES IN SCHEMA public TO :user;

GRANT INSERT, 
UPDATE 
  ON address, 
  comments, 
  orderitems, 
  orders, 
  favorites, 
  usedpromocodes, 
  users, 
  payment TO :user;
