from flask_sqlalchemy import SQLAlchemy

db = SQLAlchemy()


class Party(db.Model):
    __tablename__ = 'party'

    id = db.Column(db.Integer, primary_key=True)
    user = db.Column(db.Integer)
    location = db.Column(db.String())
    order = db.Column(db.String())
    hostesses = db.Column(db.Integer)

    def __init__(self, user, location, order, hostesses):
        self.hostesses = hostesses
        self.order = order
        self.location = location
        self.user = user

    def __repr__(self):
        return '<id {}>'.format(self.id)

    def serialize(self):
        return {
            'id': self.id,
            'user': self.user,
            'location': self.location,
            'order': self.order,
            'hostess': self.hostesses
        }

