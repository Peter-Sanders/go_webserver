SELECT 
  id, 
  fname, 
  lname, 
  phone, 
  email, 
  password, 
  username 
FROM users
WHERE email = ?;
