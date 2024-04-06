insert into user (fname, lname, phone, email, username, password) 
with data as (
   select  
     :fname as fname, 
     :lname as lname, 
     :phone as phone, 
     :email as email, 
     :username as username,
     :password as password
)
select 
  d.*
from data d 
left join user u on u.email = d.email and u.username = d.username 
where u.id is null;
