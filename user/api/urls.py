from django.urls import path, include
from rest_framework import routers
from . import views
# router = routers.DefaultRouter()
# router.register('clients', views.ClientViewSet)


urlpatterns = [
    #path('', include(router.urls)),
    path('clients/<int:pk>', views.SingleClientView.as_view(), name='single_client'),
]