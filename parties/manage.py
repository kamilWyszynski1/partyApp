from flask_script import Manager
from flask_migrate import Migrate, MigrateCommand

from app import app, db

migrate = Migrate(app, db)
manager = Manager(app)

manager.add_command('db', MigrateCommand)


def runserver():
    create_db()
    app.run(host="0.0.0.0", port=5000)


## alows you to run this function from 'python manage.py <function_name>'
@manager.command
def create_db():
    """
    Create database migrations and upgrade it
    """
    db.create_all()


if __name__ == '__main__':
    manager.run()
