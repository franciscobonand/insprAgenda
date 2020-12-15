docker pull postgres:alpine
docker run --name pginspr -e POSTGRES_PASSWORD=postgres123 -d -p 5432:5432 postgres:alpine
sleep 5
docker exec pginspr bash -c "PGPASSWORD=postgres123 psql -h localhost -U postgres -c 'create table tasks (id serial primary key,priority integer,status integer,title varchar,description varchar,dependency varchar,deadline date,workstart timestamp without time zone,workend timestamp without time zone,creationdate date,lastupdate date,timeestimate integer);'"