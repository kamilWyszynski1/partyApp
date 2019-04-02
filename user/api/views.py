from rest_framework import viewsets, mixins

from .models import Client
from .permissions import IsAdminOrIsSelf
from .serialiezers import (
    ClientSerializer,
    ClientSerializerLocation
)


# class ClientViewSet(viewsets.ModelViewSet):
#     queryset = Client.objects.all()
#     serializer_class = ClientSerializer
#
#     def get_serializer_class(self):
#         serializer_class = self.serializer_class
#
#         if self.request.method == 'PUT':
#             serializer_class = ClientSerializerLocation
#         return serializer_class
#
#     def get_permissions(self):
#         self.permission_classes = [IsAdminOrIsSelf]
#         return super(ClientViewSet, self).get_permissions()
#

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



# class SingleClientGet(generics.RetrieveAPIView):
#     queryset = Client.objects.all()
#     serializer_class = ClientSerializer
#     permission_classes = [IsAuthenticated, TokenHasReadWriteScope]
#
#     def retrieve(self, request, *args, **kwargs):
#         if request.user.pk == kwargs['pk']:
#             instance = self.get_object()
#             serializer = self.get_serializer(instance)
#             return Response(serializer.data)
#         else:
#             raise PermissionDenied({'message': 'You dont have perrmision to change others data'})
