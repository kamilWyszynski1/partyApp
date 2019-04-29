import requests

from rest_framework import serializers
from .models import Client


class ClientSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = Client
        fields = ('id', 'email', 'password', 'location')
        extra_kwargs = {'password': {
            'write_only': True
        }}

    @classmethod
    def send_mail(cls, email):
        mail_url = "http://localhost:8080/verification/"
        post_data = {'email': email}
        response = requests.post(mail_url, json=post_data, timeout=5)

        return response

    def create(self, validated_data):
        email = validated_data.pop('email')
        password = validated_data.pop('password')
        try:
            location = validated_data.pop('location')
        except KeyError:
            location = ""
        client = Client(email=email, location=location)
        client.username = email
        client.set_password(password)

        if not client.is_admin:
            client.is_active = False
        response = self.send_mail(email)

        if response.status_code == 200:
            client.save()
            return client
        else:
            return None


class ClientSerializerLocation(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = Client
        fields = ('location',)
