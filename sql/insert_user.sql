insert into users (fname, lname, phone, email, username, password) 
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
left join users u on u.email = d.email and u.username = d.username 
where u.id is null;
