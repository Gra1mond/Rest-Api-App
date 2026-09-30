CREATE TABLE  IF NOT EXISTS users(
    id serial primary key,
    email VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS projects(
    id serial primary key,
    user_id integer not null,
    name varchar(255) not null,

    foreign key (user_id)
        references users(id)
        on delete cascade
        on update cascade
);

CREATE TABLE IF NOT EXISTS tasks(
    id SERIAL PRIMARY KEY,
    project_id integer not null,
    title varchar(255) not null,
    description text,
    done BOOLEAN NOT NULL DEFAULT FALSE,

    foreign key(project_id)
        references projects(id)
        on delete cascade
        on update cascade
)