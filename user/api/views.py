from rest_framework import viewsets, mixins
from django.http import HttpResponse
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

    basename = 'client api'
    queryset = Client.objects.all()
    serializer_class = ClientSerializer

    def get_serializer_class(self):
        serializer_class = self.serializer_class

        if self.request.method == 'PUT':
            serializer_class = ClientSerializerLocation
        return serializer_class

    def get_permissions(self):
        self.permission_classes = [IsAdminOrIsSelf]
        return super(ClientViewSet, self).get_permissions()

    def post(self, request, *args, **kwargs):
        return self.create(request, *args, **kwargs)

def check_user(request):
    user = request.user.username
    if user is '':
        return HttpResponse(status=401)
    else:
        return HttpResponse(status=200)
