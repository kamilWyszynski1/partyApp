#!bin/sh

python manage.py db init
python manage.py db migrate
python manage.py db update

exec python app.py