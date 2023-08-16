drop table if exists todos;

drop sequence if exists todos_seq;

create sequence todos_seq;

create table todos(
    id int primary key default nextval('todos_seq'),
    description text not null,
    status_id int not null,
    created_on timestamp default now() not null
);

-- name: GetTodo :one
select
    t.*,
    s.description as status
from todos t
join status s on
t.status_id = s.id
where t.id = $1 limit 1;

-- name: GetAllTodos :many
select
    t.*,
    s.description as status
from todos t
join status s on
    t.status_id = s.id;

-- name: CreateTodo :one
insert into todos (description, status_id) values ($1, '1') returning *;

-- name: DeleteTodo :exec
delete from todos where id = $1;

-- name: UpdateTodo :exec
update todos set description = $2, status_id = $3 where id = $1;