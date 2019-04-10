import os
from flask import Flask, jsonify, request, abort
from flask_restful import Api, Resource
from models import db, Party
from flask_cors import CORS
import requests

DBUSER = 'postgres'
DBPASS = 'postgres'
DBHOST = '0.0.0.0'
DBPORT = '5432'
DBNAME = 'postgres'



app = Flask(__name__)
CORS(app)

app.config['SQLALCHEMY_DATABASE_URI'] = \
    'postgresql+psycopg2://{user}:{passwd}@{host}:{port}/{db}'.format(
        user=DBUSER,
        passwd=DBPASS,
        host=DBHOST,
        port=DBPORT,
        db=DBNAME)
app.config.from_object(os.environ['APP_SETTINGS'])
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

db.init_app(app)
api = Api(app)


@app.route("/")
def hello():
    return "Hello World!"


class PartyResource(Resource):

    def authenticate(self):
        '''
            Method that authenticate user's token in users' service
        '''
        try:
            token = request.headers['Authorization'].split(' ')[1]
            print(token)
            url = 'http://127.0.0.1:8000/check'
            headers = {'Authorization': f'Bearer {token}'}
            
            content = requests.get(
                url,
                headers=headers
            )

            if content.status_code == 401:
                abort(401)
            return content.status_code
        except:
            abort(401)

    def get(self, party_id=None):
        self.authenticate()
        try:
            if party_id is not None:
                return jsonify(
                    Party.query.filter_by(id=party_id).first().serialize()
                )
            else:
                return jsonify(
                    [x.serialize() for x in Party.query.all()]
                )
        except:
            return '', 204

    def post(self):
        self.authenticate()
        content = request.json

        try:
            party = Party(
                user=content['user'],
                location=content['location'],
                order=content['order'],
                hostesses=content['hostesses']
            )

            db.session.add(party)
            db.session.commit()
            return jsonify(party.serialize())
        except:
            abort(404)

    def delete(self, party_id=None):
        self.authenticate()

        if party_id is None:
            [db.session.delete(x) for x in Party.query.all()]
            db.session.commit()
        else:
            try:
                party = Party.query.filter_by(id=party_id).first()
                db.session.delete(party)
                db.session.commit()
                return 'deleted', 200
            except Exception as e:
               abort(404)

    def put(self, party_id):
        self.authenticate()

        try:
            party = Party.query.filter_by(id=party_id).first()
            content = request.json
            party.order = content['order']
            db.session.commit()
        except Exception as e:
            abort(404)

api.add_resource(PartyResource,
                 '/party',
                 '/party/<int:party_id>')

if __name__ == '__main__':
    db.create_all()
    app.run(host="0.0.0.0")
