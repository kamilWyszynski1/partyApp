import json

import requests
from rest_framework import viewsets, mixins
from rest_framework.decorators import api_view, permission_classes
from rest_framework.permissions import IsAuthenticated, IsAdminUser
from django.http import HttpResponse
from django.core.exceptions import ObjectDoesNotExist
from .models import Client
from .permissions import IsAdminOrIsSelf
from .serialiezers import (
    ClientSerializer,
    ClientSerializerLocation
)


class ClientViewSet(mixins.CreateModelMixin,
                    mixins.RetrieveModelMixin,
                    mixins.UpdateModelMixin,
                    mixins.DestroyModelMixin,
                    viewsets.GenericViewSet):
    queryset = Client.objects.all()
    serializer_class = ClientSerializer

    def get_serializer_class(self):
        serializer_class = self.serializer_class

        if self.request.method == 'PUT':
            serializer_class = ClientSerializerLocation
        return serializer_class

    def get_permissions(self):
        if self.request.method == 'GET':
            self.permission_classes = [IsAdminOrIsSelf]
        elif self.request.method == 'POST':
            self.permission_classes = [IsAdminUser]
        return super(ClientViewSet, self).get_permissions()


def check_user(request):
    """
        Function which takes request with Access Token and 
        defines if user is authenticated or not
    """
    user = request.user.username
    if user is '':
        return HttpResponse(status=401)
    else:
        return HttpResponse(status=200)


@api_view(['POST'])
@permission_classes((IsAdminUser,))
def activate_user(request):
    """
        Function takes request from email service - it has constant permission 
        provided on admin account. Email services store infinite access token which allows
        service to provide this action.
    """
    email = request.data['email']
    try:
        user = Client.objects.get(email=email)
        user.is_active = True
        user.save()
        return HttpResponse(status=200)
    except ObjectDoesNotExist:
        return HttpResponse(status=404)


def send_mail(request):
    mail_url = "http://localhost:8080/verification/"
    url = "http://127.0.0.1:5000"
    email = 'kamil-wyszynski@wp.pl'
    post_data = {'email': email}
    # response = requests.post(mail_url, json=post_data, timeout=5)
    response = requests.get(url)

    return HttpResponse(response.status_code)
