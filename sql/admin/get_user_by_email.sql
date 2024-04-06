SELECT 
  id, 
  fname, 
  lname, 
  phone, 
  email, 
  password, 
  username 
FROM user
WHERE email = ?;
