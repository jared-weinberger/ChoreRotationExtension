insert into users (id, name, email)
values (?, ?, ?)
on conflict(id)
do update set name = excluded.name