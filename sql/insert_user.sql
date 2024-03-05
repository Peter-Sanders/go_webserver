insert into users 
with data as (
  select 
  NULL as id, 
  ? as fname, 
  ? as lname, 
  ? as phone, 
  ? as email, 
  ? as time 
)
select 
  d.*
from data d 
left join users u on u.email = d.email and lower(u.fname) = lower(d.fname) and lower(u.lname) = lower(d.lname)
where u.id is null;
