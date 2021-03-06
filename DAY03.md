# Взаимодействие с базами данных. Хранение секретов в переменных окружения.

Цель: научиться работать с базами данных и обрабатывать ошибки при взаимодействии с ними.

База (postgres)
Host:  #######
Порт 5432
DbName postgresdb
User postgresadmin
Password admin123

## Задание 1.

Написать программу, которая умеет выполнять запрос в БД

```pg
 test.do_something(
	_employee_id integer,
           _test_data jsonb,
	_date timestamp
)
```

## Задание 2.

Научиться обрабатывать коды ошибок при запросах БД. В зависимости от типа ошибки возвращать соответствующий статус код.

Процедура test.get_db_error(_id integer), где
Id = 1 вернет пользовательскую ошибку
Id = 2 вернет ошибку бд,
Любой другой id вернет bool (то есть отработает успешно, без ошибок)

Ошибка с кодом больше 50 000 (либо уровень ошибки 50) – пользовательская, иначе – внутренняя, в бд в зависимости от уровня ошибки необходимо отдавать сервисом 400 или 500 код. При пользовательской ошибке необходимо в ответ также отдавать текст ошибки из БД.

## Задание 3. 

Изучить переменные окружения и основные инструменты для работы с ними в Go. Переписать программу таким образом, чтобы секретные данные, такие как User и Password считывались из переменных окружения.

Вспомогательная литература:

* The_Ultimate_Guide_To_Building_Database-Driven_Apps_with_Go.pdf
