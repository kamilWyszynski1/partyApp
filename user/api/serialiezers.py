from django.contrib.auth.models import User
from rest_framework import serializers
from .models import Client


class ClientSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = Client
        fields = ( 'id', 'email', 'password', 'location')
        extra_kwargs = {'password': {
            'write_only': True
        }}

    def create(self, validated_data):
        email = validated_data.pop('email')
        password = validated_data.pop('password')
        try:
            location = validated_data.pop('location')
        except KeyError as e:
            location = ""
        client = Client(email=email, location=location)
        client.username = email
        client.set_password(password)
        client.save()
        return client


class ClientSerializerLocation(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = Client
        fields = ('location',)


