insert into user (fname, lname, phone, email, username, password) 
with data as (
  select  
  ? as fname, 
  ? as lname, 
  ? as phone, 
  ? as email, 
  ? as username,
  ? as password
)
select 
  d.*
from data d 
left join user u on u.email = d.email and u.username = d.username 
where u.id is null;
