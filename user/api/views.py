from django.shortcuts import render

# Create your views here.

from rest_framework import viewsets, generics
from rest_framework.decorators import permission_classes, api_view
from rest_framework.permissions import IsAuthenticated
from oauth2_provider.contrib.rest_framework import TokenHasReadWriteScope

from .models import Client
from .serialiezers import (
    ClientSerializer,
    ClientSerializerLocation
)


# class ClientViewSet(viewsets.ModelViewSet):
#     queryset = Client.objects.all()
#     serializer_class = ClientSerializer
#     permission_classes = [IsAuthenticated, TokenHasReadWriteScope]
#
#     def get_serializer_class(self):
#         serializer_class = self.serializer_class
#
#         if self.request.method == 'PUT':
#             serializer_class = ClientSerializerLocation
#         return serializer_class


class SingleClientView(generics.RetrieveAPIView):
    queryset = Client.objects.all()
    serializer_class = ClientSerializer
    permission_classes = [IsAuthenticated, TokenHasReadWriteScope]

